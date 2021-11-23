package policy

import (
	"log"
	"strings"

	"go-auth0-cookie-sample/auth"
)

var instance *Manager

type Manager struct {
	client     *auth.Auth0Client
	identifier string
	roles      []*Role
}

type Role struct {
	id          string
	name        string
	permissions []*Permission
}

type Permission struct {
	operation  string
	resource   string
	identifier string
}

func GetManager() *Manager {
	return instance
}

func (m *Manager) Refresh() {
	roles, err := m.client.GetRoles()
	if err != nil {
		log.Fatalf("Failed to get role list: %+v", err)
	}

	result := make([]*Role, 0, len(roles))

	for _, r := range roles {
		permissions := make([]*Permission, 0, len(r.Permissions))
		for _, p := range r.Permissions {
			ops := strings.Split(p.Name, ":")
			if len(ops) != 2 {
				continue
			}
			permissions = append(permissions, &Permission{
				ops[0], ops[1], p.Identifier,
			})
		}
		result = append(result, &Role{
			id:          r.Id,
			name:        r.Name,
			permissions: permissions,
		})
	}
	instance.roles = result
}

func (m *Manager) HasPermission(role string, operation string, resource string) bool {
	if m.roles == nil {
		m.Refresh()
	}

	for _, r := range m.roles {
		if r.name == role {
			for _, p := range r.permissions {
				if p.operation == operation && p.resource == resource && p.identifier == m.identifier {
					return true
				}
			}
		}
	}
	return false
}

func InitPolicyManager(cl *auth.Auth0Client, identifier string) {
	instance = &Manager{
		cl, identifier, nil,
	}
	instance.Refresh()
}
