package model

type User struct {
	id     userId
	name   userName
	tenant Tenant
	role   Role
}

type Users []*User

type userId int

func UserId(value int) (*userId, error) {
	// TODO("value validation")
	id := userId(value)
	return &id, nil
}

type userName string

func UserName(value string) (*userName, error) {
	// TODO("value validation")
	name := userName(value)
	return &name, nil
}
