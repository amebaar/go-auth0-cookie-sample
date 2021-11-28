package model

type Role struct {
	name        roleName
	permissions []*Permission
}

type roleName string

func RoleName(value string) (*roleName, error) {
	// TODO("value validation")
	name := roleName(value)
	return &name, nil
}
