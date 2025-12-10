package domain

import "errors"

var (
	ErrInvalidTimeStamp      = errors.New("invalid timestamp")
	ErrEqualName             = errors.New("name is nil")
	ErrEqualNotificationType = errors.New("notificationType is nill")
	ErrMaxName               = errors.New("The name is very more")
	ErrMaxDescription        = errors.New("The description is very more")
)
