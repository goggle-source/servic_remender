package database

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

var arrErr = map[error]error{}

func ValidateErrors(err error) error {
	if pq, ok := err.(*pq.Error); ok {
		switch pq.Code {
		case "23503":
			return ErrForeignKey
		case "23502":
			return ErrNotNull
		}
	}

	if errors.Is(err, sql.ErrNoRows) {
		return ErrNoRows
	}

	return ErrDatabase
}
