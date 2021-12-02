package model

type Company struct {
	id *companyId
}

type companyId int

func CompanyId(value int) (*companyId, error) {
	// TODO("value validation")
	id := companyId(value)
	return &id, nil
}

func NewCompany(id *companyId) *Company {
	return &Company{
		id,
	}
}
