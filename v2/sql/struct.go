package sql

import (
	"auth-service/whatsauth"
	"database/sql"
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
