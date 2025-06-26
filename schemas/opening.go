package schemas

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Title    string          `json:"title"`
	Role     string          `json:"role"`
	Company  string          `json:"company"`
	Location string          `json:"location"`
	Remote   bool            `json:"remote"`
	Salary   decimal.Decimal `json:"salary"`
	Link     string          `json:"link"`
}
