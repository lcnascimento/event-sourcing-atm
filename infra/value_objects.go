package infra

import (
	"context"
	"errors"
	"net/http"
	"time"
)

// ObjectID represents a document identifier
type ObjectID string

var (
	// ErrNotImplemented ...
	ErrNotImplemented = errors.New("not implemented yet")
	// ErrCastEntity ...
	ErrCastEntity = errors.New("could not cast entity")
	// ErrInvalidDataType ...
	ErrInvalidDataType = errors.New("invalid data type error")
	// ErrEntityNotFound ...
	ErrEntityNotFound = errors.New("entity not found")
)

// Environment ...
type Environment string

const (
	// EnvironmentDevelopment ...
	EnvironmentDevelopment Environment = "development"
	// EnvironmentProduction ...
	EnvironmentProduction Environment = "production"
)

// Severity ...
type Severity string

const (
	// SeverityCritical ...
	SeverityCritical Severity = "Critical"
	// SeverityError ...
	SeverityError Severity = "Error"
	// SeverityWarning ...
	SeverityWarning Severity = "Warning"
	// SeverityInfo ...
	SeverityInfo Severity = "Info"
	// SeverityDebug ...
	SeverityDebug Severity = "Debug"
)

// Error ...
type Error struct {
	Ctx        context.Context `json:"-"`
	Err        error           `json:"-"`
	Severity   Severity        `json:"severity"`
	OpName     OpName          `json:"opName"`
	Kind       ErrorKind       `json:"kind"`
	CustomData CustomData      `json:"customData"`
}

// Error ...
func (e Error) Error() string {
	return e.Err.Error()
}

// OpName ...
type OpName string

// Op ...
type Op struct {
	Name      OpName
	StartTime time.Time
}

// NewOp ...
func NewOp(name OpName) *Op {
	return &Op{
		Name:      name,
		StartTime: time.Now(),
	}
}

// Duration ...
func (op Op) Duration() float64 {
	return time.Since(op.StartTime).Seconds()
}

// CustomData ...
type CustomData map[string]interface{}

// Merge ...
func (cd CustomData) Merge(newCd *CustomData) CustomData {
	if newCd == nil {
		return cd
	}

	for k, v := range *newCd {
		cd[k] = v
	}

	return cd
}

// ErrorKind ...
type ErrorKind int

const (
	// KindIncompleteResponse ...
	KindIncompleteResponse ErrorKind = http.StatusPartialContent
	// KindBadRequest ...
	KindBadRequest ErrorKind = http.StatusBadRequest
	// KindNotFound ...
	KindNotFound ErrorKind = http.StatusNotFound
	// KindUnexpected ...
	KindUnexpected ErrorKind = http.StatusInternalServerError
)

const (
	// AppKeyContextValueKey ...
	AppKeyContextValueKey string = "appKey"
	// IDContextValueKey ...
	IDContextValueKey string = "contextID"
	// ReferenceContextValueKey ...
	ReferenceContextValueKey string = "reference"
)

// EventStreamTopic ...
type EventStreamTopic string

// EventName ...
type EventName string

// EventRowID ...
type EventRowID string

// AggregateID ...
type AggregateID string

// ListEventsInput ...
type ListEventsInput struct {
	RowID       *EventRowID
	AggregateID *AggregateID
}
