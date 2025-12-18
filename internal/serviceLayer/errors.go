package servicelayer

import "errors"

var (
	ErrInternal              = errors.New("error servic")
	ErrMaxName               = errors.New("The name is too big")
	ErrMaxDescription        = errors.New("The description is too big")
	ErrInvalidTimeStamp      = errors.New("The timestamp is invalid format")
	ErrEqualNotificationType = errors.New("The notification_type is equal")
	ErrEqualName             = errors.New("The name is nil")
	ErrStatus                = errors.New("an error occurred when sending")
)

var (
	ErrDatabase         = errors.New("Err database")
	ErrClientForeignKey = errors.New("there is no reminder destination")
	ErrClientNoRows     = errors.New("there is no such value")
	ErrClientNotNull = errors.New("there are no necessary values")
)
