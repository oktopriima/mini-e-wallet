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

type UserBalanceHistory struct {
	ID            int     `json:"id"`
	UserBalanceID int     `json:"user_balance_id"`
	BalanceBefore float64 `json:"balance_before"`
	BalanceAfter  float64 `json:"balance_after"`
	Activity      string  `json:"activity"`
	Type          string  `json:"type"`
	IP            string  `json:"ip"`
	Location      string  `json:"location"`
	UserAgent     string  `json:"user_agent"`
	Author        string  `json:"author"`
}

// TableName sets the insert table name for this struct type
func (u *UserBalanceHistory) TableName() string {
	return "user_balance_histories"
}
