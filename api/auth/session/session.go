package session

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-auth0-cookie-sample/api/app/domain/model"
)

func GetUserCompany(ctx echo.Context) (*model.Company, error) {
	companyId, err := model.CompanyId(ctx.Get("cid").(int))
	if err != nil {
		return nil, fmt.Errorf("faield to get cid from session: %+v", err)
	}
	return model.NewCompany(companyId), nil
}
