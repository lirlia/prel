package model

import (
	"prel/pkg/custom_error"
	"unicode/utf8"

	"github.com/cockroachdb/errors"
	"github.com/google/uuid"
)

const (
	minPatternLength = 3
	maxPatternLength = 20
)

type IamRoleFilteringRule struct {
	id      string
	pattern string

	// userID is the user who added the pattern.
	userID UserID
}

type IamRoleFilteringRules []*IamRoleFilteringRule

func (r *IamRoleFilteringRule) ID() string {
	return r.id
}

func (r *IamRoleFilteringRule) Pattern() string {
	return r.pattern
}

func (r *IamRoleFilteringRule) UserID() UserID {
	return r.userID
}

func newIamRoleFilteringRule(id, pattern string, userID UserID) (*IamRoleFilteringRule, error) {
	if utf8.RuneCountInString(pattern) < minPatternLength || utf8.RuneCountInString(pattern) > maxPatternLength {
		return nil, errors.WithDetail(errors.New("pattern length must be between 3 and 20"), string(custom_error.InvalidArgument))
	}

	return &IamRoleFilteringRule{
		id:      id,
		pattern: pattern,
		userID:  userID,
	}, nil
}

func NewIamRoleFilteringRule(pattern string, userID UserID) (*IamRoleFilteringRule, error) {
	return newIamRoleFilteringRule(uuid.New().String(), pattern, userID)
}

func ReconstructIamRoleFilteringRule(id, pattern, userID string) (*IamRoleFilteringRule, error) {
	rule, err := newIamRoleFilteringRule(id, pattern, UserID(userID))
	if err != nil {
		return nil, err
	}

	return rule, nil
}

func (r IamRoleFilteringRules) Patterns() []string {
	patterns := make([]string, 0, len(r))
	for _, rule := range r {
		patterns = append(patterns, rule.Pattern())
	}
	return patterns
}
