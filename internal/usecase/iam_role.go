package usecase

import (
	"context"
	"prel/internal/gateway/repository"
	"prel/internal/model"
	"sort"
	"strings"
)

func (uc *Usecase) GetIamRoles(ctx context.Context, projectID string) ([]string, error) {
	user := model.GetUser(ctx)
	now := model.GetClock(ctx).Now()

	// Currently, only roles that can be applied to user principal are returned.
	// so we set user as principal.
	roles, err := uc.c.GetIamRoles(now, projectID, user)
	if err != nil {
		return nil, err
	}

	roleIDs := make([]string, 0, len(roles))
	for _, role := range roles {
		roleIDs = append(roleIDs, role.Name)
	}

	sort.Slice(roleIDs, func(i, j int) bool {
		return roleIDs[i] < roleIDs[j]
	})

	var rulePatterns []string
	err = repository.NewTransactionManager().Transaction(ctx, func(ctx context.Context) error {
		rules, err := uc.iamRoleFilteringRuleRepo.FindAll(ctx)
		if err != nil {
			return err
		}
		rulePatterns = rules.Patterns()

		return nil
	})
	if err != nil {
		return nil, err
	}

	if len(rulePatterns) == 0 {
		return roleIDs, nil
	}
	return FilterRoleIDs(roleIDs, rulePatterns), nil
}

func FilterRoleIDs(roleIDs []string, patterns []string) []string {
	var filtered []string
	for _, id := range roleIDs {
		for _, pattern := range patterns {
			if strings.Contains(id, pattern) {
				filtered = append(filtered, id)
				break
			}
		}
	}
	return filtered
}
