package handler

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("parameter %s of type %s is required", name, typ)
}

type CreateOpeningRequest struct {
	Title    string          `json:"title"`
	Role     string          `json:"role"`
	Company  string          `json:"company"`
	Location string          `json:"location"`
	Remote   *bool           `json:"remote"`
	Salary   decimal.Decimal `json:"salary"`
	Link     string          `json:"link"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Salary.LessThanOrEqual(decimal.Zero) {
		return errParamIsRequired("salary", "decimal")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	return nil
}
