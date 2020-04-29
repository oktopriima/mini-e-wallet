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

type UserBalance struct {
	ID                   int                 `json:"id"`
	UserID               int                 `json:"user_id"`
	Balance              float64             `json:"balance"`
	BalanceAchieve       float64             `json:"balance_achieve"`
	UserBalanceHistories *UserBalanceHistory `json:"user_balance_histories" gorm:"ForeignKey:UserBalanceID"`
}

// TableName sets the insert table name for this struct type
func (u *UserBalance) TableName() string {
	return "user_balances"
}
