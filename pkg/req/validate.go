package req

import "effective_mobile/internal/validator"

func IsValid[T any](payload T) error {
	return validator.Validate(payload)
}