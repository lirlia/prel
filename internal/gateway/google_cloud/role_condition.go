package google_cloud

import (
	"regexp"
	"slices"

	"google.golang.org/api/iam/v1"
)

var basicRoles = []string{
	"roles/owner",
	"roles/editor",
	"roles/viewer",
}

type PrincipalType string

var ignoreMap = map[PrincipalType][]*regexp.Regexp{
	"user": {
		// https://cloud.google.com/iam/docs/service-account-permissions
		regexp.MustCompile(`[sS]erviceAgent`),
	},
}

func ExcludeRoleByPrincipalType(roles []*iam.Role, principalType PrincipalType) []*iam.Role {
	return filterRoles(roles, func(role *iam.Role) bool {
		for _, ignore := range ignoreMap[principalType] {
			if ignore.MatchString(role.Name) {
				return false
			}
		}
		return true
	})
}

func ExcludeBasicRole(roles []*iam.Role) []*iam.Role {
	return filterRoles(roles, func(role *iam.Role) bool {
		return !slices.Contains(basicRoles, role.Name)
	})
}

func filterRoles(roles []*iam.Role, filter func(*iam.Role) bool) []*iam.Role {
	filteredRoles := make([]*iam.Role, 0, len(roles))
	for _, role := range roles {
		if filter(role) {
			filteredRoles = append(filteredRoles, role)
		}
	}
	return filteredRoles
}
