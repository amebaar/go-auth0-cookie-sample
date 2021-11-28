package model

type Permission struct {
	operation Operation
}

type Operation int

const (
	ReadOperation Operation = iota + 1
	CreateOperation
	UpdateOperation
	DeleteOperation
)

func (o Operation) String() string {
	switch o {
	case ReadOperation:
		return "Read"
	case CreateOperation:
		return "Create"
	case UpdateOperation:
		return "Update"
	case DeleteOperation:
		return "Delete"
	default:
		return "Invalid_Operation"
	}
}
