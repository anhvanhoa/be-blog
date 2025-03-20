package rbac

import "strings"

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

func Allow(roles ...string) FunPermission {
	return func(r ...string) bool {
		for _, role := range roles {
			for _, r := range r {
				if strings.EqualFold(role, r) {
					return true
				}
			}
		}
		return false
	}
}

func AllowAll() FunPermission {
	return func(r ...string) bool {
		return true
	}
}
