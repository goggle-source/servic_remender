package servicelayer

import (
	"fmt"
	"servic_remender/internal/database"
	"servic_remender/internal/domain"
)

type MessageError struct {
	Op  string
	Err error
}

func ValidationError(op string, err error) error {
	errDomain := ValidateDomainErrors(op, err)
	if errDomain != nil {
		return errDomain
	}

	errDatabase := ValidateDatabaseErrors(op, err)
	if errDatabase != nil {
		return errDatabase
	}

	return fmt.Errorf("%s: %w", op, ErrInternal)
}

func ValidateDomainErrors(op string, err error) error {

	ErrArr := map[error]error{
		domain.ErrEqualName:             ErrEqualName,
		domain.ErrMaxDescription:        ErrMaxDescription,
		domain.ErrEqualNotificationType: ErrEqualNotificationType,
		domain.ErrMaxName:               ErrMaxName,
		domain.ErrInvalidTimeStamp:      ErrInvalidTimeStamp,
	}

	err, ok := ErrArr[err]
	if ok {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func ValidateReminderSent(status string) error {
	if status == ErrSending {
		return ErrStatus
	}

	return nil
}

func ValidateDatabaseErrors(op string, err error) error {
	ErrArr := map[error]error{
		database.ErrDatabase:   ErrDatabase,
		database.ErrForeignKey: ErrClientForeignKey,
		database.ErrNoRows:     ErrClientNoRows,
		database.ErrNotNull:    ErrClientNoRows,
	}

	err, ok := ErrArr[err]
	if ok {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
