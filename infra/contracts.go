package infra

import (
	"context"
)

// LogProvider ...
type LogProvider interface {
	Critical(context.Context, OpName, string)
	Criticalf(context.Context, OpName, string, ...interface{})
	CriticalCustomData(context.Context, OpName, string, CustomData)
	Info(context.Context, OpName, string)
	Infof(context.Context, OpName, string, ...interface{})
	InfoCustomData(context.Context, OpName, string, CustomData)
	Error(context.Context, OpName, string)
	Errorf(context.Context, OpName, string, ...interface{})
	ErrorCustomData(context.Context, OpName, string, CustomData)
	Warning(context.Context, OpName, string)
	Warningf(context.Context, OpName, string, ...interface{})
	WarningCustomData(context.Context, OpName, string, CustomData)
	Debug(context.Context, OpName, string)
	Debugf(context.Context, OpName, string, ...interface{})
	DebugCustomData(context.Context, OpName, string, CustomData)
}