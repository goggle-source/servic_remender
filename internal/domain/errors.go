package domain

import "errors"

var (
	InvalidTimeStamp      = errors.New("invalid timestamp")
	EqualName             = errors.New("name is nil")
	EqualNotificationType = errors.New("notificationType is nill")
	ErrMaxParameter                = errors.New("parameter is very more")
)
