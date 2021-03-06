package user

import (
	"github.com/jmoiron/sqlx"
)

type userDomain struct {
	DB *sqlx.DB
}

func New(DB *sqlx.DB) *userDomain {
	return &userDomain{
		DB: DB,
	}
}
