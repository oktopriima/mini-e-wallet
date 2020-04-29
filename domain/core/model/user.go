package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type User struct {
	ID           int            `json:"id"`
	Username     string         `json:"username"`
	Email        string         `json:"email"`
	Password     string         `json:"-"`
	LastLogin    time.Time      `json:"last_login"`
	IsVerified   bool           `json:"is_verified"`
	IsActive     bool           `json:"is_active"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	UserBalances []*UserBalance `json:"user_balances" gorm:"ForeignKey:UserID"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "users"
}
