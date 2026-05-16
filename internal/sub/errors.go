package sub

import "errors"

var (
	ErrSubExists = errors.New("subscription user_service exists")
	ErrSubNotFound = errors.New("subscription not found")
)