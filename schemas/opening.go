package schemas

import (
	"time"

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

type OpeningResponse struct {
	ID        uint            `json:"id"`
	Title     string          `json:"title"`
	Role      string          `json:"role"`
	Company   string          `json:"company"`
	Location  string          `json:"location"`
	Remote    bool            `json:"remote"`
	Salary    decimal.Decimal `json:"salary"`
	Link      string          `json:"link"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	DeletedAt time.Time       `json:"deleted_at,omitempty"`
}
