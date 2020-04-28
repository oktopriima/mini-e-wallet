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

type Migration struct {
	ID        string    `json:"id"`
	AppliedAt time.Time `json:"applied_at"`
}

// TableName sets the insert table name for this struct type
func (m *Migration) TableName() string {
	return "migrations"
}
