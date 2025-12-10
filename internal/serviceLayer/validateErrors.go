package servicelayer

import (
	"fmt"
	"servic_remender/internal/domain"
)

type MessageError struct {
	Op  string
	Err error
}

func ValidateDomainErrors(op string, err error) error {

	if err == domain.ErrEqualName {
		return fmt.Errorf("%s: %w", op, ErrEqualName)
	}

	if err == domain.ErrMaxDescription {
		return fmt.Errorf("%s: %w", op, ErrMaxDescription)
	}

	if err == domain.ErrEqualNotificationType {
		return fmt.Errorf("%s: %w", op, ErrEqualNotificationType)
	}

	if err == domain.ErrMaxName {
		return fmt.Errorf("%s: %w", op, ErrMaxName)
	}

	if err == domain.ErrInvalidTimeStamp {
		return fmt.Errorf("%s: %w", op, ErrInvalidTimeStamp)
	}

	return fmt.Errorf("%s: %w", op, ErrInternal)
}
