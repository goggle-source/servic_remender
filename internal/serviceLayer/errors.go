package servicelayer

import "errors"

var (
	ErrInternal              = errors.New("error servic")
	ErrMaxName               = errors.New("The name is too big")
	ErrMaxDescription        = errors.New("The description is too big")
	ErrInvalidTimeStamp      = errors.New("The timestamp is invalid format")
	ErrEqualNotificationType = errors.New("The notification_type is equal")
	ErrEqualName             = errors.New("The name is nil")
)
