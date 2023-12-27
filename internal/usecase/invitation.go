package usecase

import (
	"context"
	"prel/internal/gateway/repository"
	"prel/internal/model"
	"prel/pkg/custom_error"

	"github.com/cockroachdb/errors"
	"github.com/jackc/pgx/v5"
)

func (uc *Usecase) Invite(ctx context.Context, email string, role model.UserRole) error {
	user := model.GetUser(ctx)
	if !user.IsAdmin() {
		return errors.WithDetail(errors.New("user is not admin"), string(custom_error.OnlyAdmin))
	}

	now := model.GetClock(ctx).Now()
	return repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {

		err := uc.checkIsUserPresent(ctx, email)
		if err != nil {
			return err
		}

		invitation, err := uc.checkIsValidInvitationPresent(ctx, email)
		if err != nil {
			return err
		}

		// if invitation is not found, it means that the user is not invited
		if invitation == nil {
			user := model.GetUser(ctx)
			invitation = model.NewInvitation(user.ID(), email, role, now)
		} else {
			// if invitation is found, it means that the user is invited but expired
			invitation.UpdateExpiredAt(now)
		}

		return uc.invitationRepo.Save(ctx, invitation)
	})
}

func (uc *Usecase) checkIsUserPresent(ctx context.Context, email string) error {
	u, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return errors.Wrap(err, "failed to find user by email")
	}

	if u != nil {
		return errors.WithDetail(errors.Newf("user already exists by email(%s)", email), string(custom_error.AlreadyRegistered))
	}

	return nil
}

func (uc *Usecase) checkIsValidInvitationPresent(ctx context.Context, email string) (*model.Invitation, error) {
	i, err := uc.invitationRepo.FindByInviteeMail(ctx, email)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.Wrap(err, "failed to find invitation by email")
	}

	// if invitation is not found, it means that the user is not invited
	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	// if invitation is found and it is expired, it means that the user is not invited
	if i.IsExpired(model.GetClock(ctx).Now()) {
		return i, nil
	}

	return nil, errors.WithDetail(errors.Newf("invitation already exists by email(%s)", email), string(custom_error.AlreadyInvited))
}
