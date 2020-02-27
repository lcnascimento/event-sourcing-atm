package domain

import "errors"

var (
	// ErrApplyEventIntoAggregate ...
	ErrApplyEventIntoAggregate error = errors.New("could not apply event into aggregate")
)
