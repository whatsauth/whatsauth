package sql

import (
	"database/sql"
	"github.com/whatsauth/whatsauth/v2"
	"sync"
)

type Queriers struct {
	mut    sync.Mutex
	config whatsauth.LoginInfo
	db     *sql.DB
}

func NewQuerier(config whatsauth.LoginInfo, db *sql.DB) *Queriers {
	return &Queriers{
		config: config,
		db:     db,
	}
}
