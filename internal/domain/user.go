package domain

import (
	"context"
	"database/sql"
	"time"
)

type AdminUser struct {
	ID        int64        `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Gender    string       `json:"gender"`
	Age       int8         `json:"age"`
	Address   AdminAddress `json:"address"`
	Email     string       `json:"email"`
	Telephone string       `json:"telephone"`
	AdminInfo AdminInfo    `json:"admin_info"`
	Password  string       `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type AdminAddress struct {
	StreetLine1 string         `json:"street_line_1"`
	StreetLine2 sql.NullString `json:"street_line_2"`
	City        string         `json:"city"`
	State       string         `json:"state"`
	Country     string         `json:"country"`
	Continent   string         `json:"continent"`
}

type AdminInfo struct {
	Role            string `json:"role"`
	Access          string `json:"access"`
	HasMasterAccess string `json:"has_master_access"`
}

type AdminUserRepository interface {
	Login(context.Context, string, string) (AdminUser, error)
}
