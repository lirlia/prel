package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"prel/internal/gateway/repository"
	"prel/internal/model"
	"prel/pkg/custom_error"
	"sort"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
)

func (uc *Usecase) UserWithPaging(ctx context.Context, page int, pageSize int) (model.Users, model.Invitations, Paging, error) {
	user := model.GetUser(ctx)
	if !user.IsAdmin() {
		return nil, nil, Paging{}, errors.WithDetail(errors.New("user is not admin"), string(custom_error.OnlyAdmin))
	}

	users, paging, err := uc.calcUsers(ctx, page, pageSize)
	if err != nil {
		return nil, nil, Paging{}, err
	}

	emails := make([]string, 0, len(users))
	for _, u := range users {
		emails = append(emails, u.Email())
	}

	invitations, err := uc.invitationRepo.FindByInviteeMailsAndExpiredAt(ctx, emails, model.GetClock(ctx).Now())
	if err != nil {
		return nil, nil, Paging{}, err
	}

	return users, invitations, paging, nil
}

func (uc *Usecase) calcUsers(ctx context.Context, page, pageSize int) (model.Users, Paging, error) {

	totalPage, currentPage, start, end, err := calcPager(ctx, uc.userAndInvitationRepo, page, pageSize)
	if err != nil {
		return nil, Paging{}, err
	}

	now := model.GetClock(ctx).Now()
	users, err := uc.userAndInvitationRepo.FindUserAndInvitationPagedByExpiredAt(ctx, start, end-start, now)
	if err != nil {
		return nil, Paging{}, err
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Email() < users[j].Email()
	})

	return users, Paging{totalPage, currentPage}, nil
}

func (uc *Usecase) UpdateUser(ctx context.Context, targetUserID string, isAvailable bool, role model.UserRole) error {
	user := model.GetUser(ctx)
	if !user.IsAdmin() {
		return errors.WithDetail(errors.New("user is not admin"), string(custom_error.OnlyAdmin))
	}

	return repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {
		user, err := uc.userRepo.FindByID(ctx, model.UserID(targetUserID))
		if err != nil {
			return err
		}

		if user == nil {
			return errors.Newf("user not found")
		}

		user.SetAvailable(isAvailable)
		user.SetRole(role)

		return uc.userRepo.Save(ctx, user)
	})
}

func (uc *Usecase) GoogleSignIn(ctx context.Context, oauth *oauth2.Config, code string) (*model.User, error) {

	token, err := oauth.Exchange(ctx, code)
	if err != nil {
		return nil, errors.WithDetail(
			errors.Wrap(err, "failed to exchange token"),
			string(custom_error.InvalidArgument))
	}

	resp, err := uc.h.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user info")
	}

	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}

	info := struct {
		GoogleID string `json:"id"`
		Email    string `json:"email"`
	}{}
	err = json.Unmarshal(content, &info)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to unmarshal json %v", content))
	}

	return uc.signin(ctx, info.GoogleID, info.Email)
}

func (uc *Usecase) IapSignIn(ctx context.Context, idToken string) (*model.User, error) {
	aud := uc.config.IapAudience
	if aud == "" {
		return nil, errors.WithDetail(errors.New("invalid iap audience"), string(custom_error.InvalidArgument))
	}

	payload, err := idtoken.Validate(ctx, idToken, aud)
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate id token")
	}

	email, ok := payload.Claims["email"].(string)
	if !ok {
		return nil, errors.Newf("email is not found in id token %+v", payload)
	}

	googleID, ok := payload.Claims["sub"].(string)
	if !ok {
		return nil, errors.Newf("googleID is not found in id token %+v", payload)
	}

	// remove prefix "accounts.google.com:"
	googleID = strings.Split(googleID, ":")[1]

	return uc.signin(ctx, googleID, email)
}

func (uc *Usecase) signin(ctx context.Context, googleID, email string) (*model.User, error) {
	var user *model.User
	var err error
	err = repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {
		user, err = uc.userRepo.FindByGoogleID(ctx, googleID)

		// if user not found, register user
		if errors.Is(err, pgx.ErrNoRows) {
			user, err = uc.register(ctx, googleID, email)
			if err != nil {
				return err
			}

			return nil
		}

		if err != nil {
			return err
		}

		now := model.GetClock(ctx).Now()
		user.UpdateLastSignInAt(now)

		expiredAt := now.Add(time.Duration(uc.config.SessionExpireSeconds) * time.Second)
		user.UpdateSession(expiredAt)

		return uc.userRepo.Save(ctx, user)
	})

	return user, err
}

func (uc *Usecase) register(ctx context.Context, googleID string, email string) (*model.User, error) {

	now := model.GetClock(ctx).Now()
	count, err := uc.userRepo.Count(ctx)
	if err != nil {
		return nil, err
	}

	var role model.UserRole

	// if no user, first user is admin
	if count == 0 {
		role = model.UserRoleAdmin
	} else {

		// if user count is not 0, new user must be invited by admin
		invitations, err := uc.invitationRepo.FindByInviteeMailsAndExpiredAt(ctx, []string{email}, now)
		if errors.Is(err, pgx.ErrNoRows) || len(invitations) == 0 {
			return nil, errors.WithDetail(
				errors.Newf("invitation not found by email(%s)", email),
				string(custom_error.NotInvited))
		}

		if err != nil {
			return nil, errors.Wrap(err, "failed to find invitation by email")
		}

		if len(invitations) > 1 {
			return nil, errors.Newf("invalid invitation count(%d)", len(invitations))
		}

		invitation := invitations[0]
		role = invitation.InviteeRole()

		err = uc.invitationRepo.Delete(ctx, invitation)
		if err != nil {
			return nil, errors.Wrap(err, "failed to delete invitation")
		}
	}

	expiredAt := now.Add(time.Duration(uc.config.SessionExpireSeconds) * time.Second)
	user := model.NewUser(googleID, email, role, expiredAt, now)
	return user, uc.userRepo.Save(ctx, user)
}
