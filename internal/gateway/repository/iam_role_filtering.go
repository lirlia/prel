package repository

import (
	"context"
	"prel/internal/gateway/postgresql"
	"prel/internal/model"

	"github.com/cockroachdb/errors"
)

type iamRoleFilteringRuleRepository struct{}

func NewIamRoleFilteringRuleRepository() *iamRoleFilteringRuleRepository {
	return &iamRoleFilteringRuleRepository{}
}

func (i *iamRoleFilteringRuleRepository) FindByID(ctx context.Context, id string) (*model.IamRoleFilteringRule, error) {
	req, err := postgresql.GetQuery(ctx).FindIamRoleFilteringRuleByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find iam role filtering rule by id")
	}

	return model.ReconstructIamRoleFilteringRule(
		req.ID,
		req.Pattern,
		req.UserID)
}

func (i *iamRoleFilteringRuleRepository) FindAll(ctx context.Context) (model.IamRoleFilteringRules, error) {
	req, err := postgresql.GetQuery(ctx).FindIamRoleFilteringRule(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find all iam role filtering rule")
	}

	ret := make(model.IamRoleFilteringRules, 0, len(req))
	for _, r := range req {
		rule, err := model.ReconstructIamRoleFilteringRule(
			r.ID,
			r.Pattern,
			r.UserID)

		if err != nil {
			return nil, errors.Wrap(err, "failed to reconstruct iam role filtering rule")
		}

		ret = append(ret, rule)
	}

	return ret, nil
}

func (i *iamRoleFilteringRuleRepository) Save(ctx context.Context, rule *model.IamRoleFilteringRule) error {
	err := postgresql.GetQuery(ctx).UpsertIamRoleFilteringRule(ctx, postgresql.UpsertIamRoleFilteringRuleParams{
		ID:      rule.ID(),
		Pattern: rule.Pattern(),
		UserID:  string(rule.UserID()),
	})
	if err != nil {
		return errors.Wrap(err, "failed to save iam role filtering rule")
	}

	return nil
}

func (i *iamRoleFilteringRuleRepository) Delete(ctx context.Context, rule *model.IamRoleFilteringRule) error {
	err := postgresql.GetQuery(ctx).DeleteIamRoleFilteringRule(ctx, rule.ID())
	if err != nil {
		return errors.Wrap(err, "failed to delete iam role filtering rule")
	}

	return nil
}
