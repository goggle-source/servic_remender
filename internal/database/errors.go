package database

import "errors"

var (
	ErrDatabase = errors.New("error database")
	ErrForeignKey = errors.New("foreign key violation")
	ErrNotNull = errors.New("NOT NULL violation")
	ErrNoRows = errors.New("No rows found")
)
