package errors

import (
	"context"
	"fmt"
	"reflect"

	pb "github.com/lcnascimento/event-sourcing-atm/proto/impl"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

// New ...
func New(args ...interface{}) *infra.Error {
	err := infra.Error{
		CustomData: infra.CustomData{},
	}

	for _, arg := range args {
		switch arg := arg.(type) {
		case context.Context:
			err.Ctx = arg
		case error:
			err.Err = arg
		case infra.OpName:
			err.OpName = arg
		case infra.ErrorKind:
			err.Kind = arg
		case infra.Severity:
			err.Severity = arg
		case infra.CustomData:
			err.CustomData = err.CustomData.Merge(&arg)
		}
	}

	return &err
}

// Context ...
func Context(err error) context.Context {
	e, ok := err.(*infra.Error)
	if !ok {
		return context.Background()
	}

	if e.Ctx != nil {
		return e.Ctx
	}

	return Context(e.Err)
}

// Trace ...
func Trace(err *infra.Error) []infra.OpName {
	trace := []infra.OpName{err.OpName}

	nextError, ok := err.Err.(*infra.Error)
	if !ok {
		return trace
	}

	trace = append(trace, Trace(nextError)...)

	return trace
}

// Kind ...
func Kind(err error) infra.ErrorKind {
	e, ok := err.(*infra.Error)
	if !ok {
		return infra.KindUnexpected
	}

	if e.Kind != 0 {
		return e.Kind
	}

	return Kind(e.Err)
}

// Severity ...
func Severity(err error) infra.Severity {
	e, ok := err.(*infra.Error)
	if !ok {
		return infra.SeverityError
	}

	if e.Severity != "" {
		return e.Severity
	}

	return Severity(e.Err)
}

// CustomData ...
func CustomData(err *infra.Error) infra.CustomData {
	nextErr, ok := err.Err.(*infra.Error)
	if !ok {
		return err.CustomData
	}

	return CustomData(nextErr).Merge(&err.CustomData)
}

// OpName ...
func OpName(err *infra.Error) infra.OpName {
	nextError, ok := err.Err.(*infra.Error)
	if !ok {
		return err.OpName
	}

	return OpName(nextError)
}

// Error ...
func Error(err *infra.Error) error {
	nextError, ok := err.Err.(*infra.Error)
	if !ok {
		return err.Err
	}

	return Error(nextError)
}

// Log ...
func Log(log infra.LogProvider, err *infra.Error) {
	method := fmt.Sprintf("%sCustomData", Severity(err))

	values := []reflect.Value{
		reflect.ValueOf(Context(err)),
		reflect.ValueOf(OpName(err)),
		reflect.ValueOf(err.Err.Error()),
		reflect.ValueOf(CustomData(err).Merge(&infra.CustomData{
			"trace": Trace(err),
			"kind":  Kind(err),
		})),
	}

	reflect.ValueOf(log).MethodByName(method).Call(values)
}

// ToProtoError ...
func ToProtoError(err *infra.Error) *pb.Error {
	return &pb.Error{
		Message: Error(err).Error(),
		Kind:    pb	.ErrorKind(Kind(err)),
	}
}
