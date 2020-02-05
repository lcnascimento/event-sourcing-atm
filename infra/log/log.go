package log

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/lcnascimento/event-sourcing-atm/infra"
)

type logValues struct {
	ContextID  string            `json:"contextID,omitempty"`
	Severity   infra.Severity    `json:"severity"`
	CustomData *infra.CustomData `json:"customData,omitempty"`
	Message    string            `json:"message"`
}

var logLevels = map[infra.Severity]int{
	infra.SeverityCritical: 1,
	infra.SeverityError:    2,
	infra.SeverityWarning:  3,
	infra.SeverityInfo:     4,
	infra.SeverityDebug:    5,
}

// ClientInput ...
type ClientInput struct {
	Level infra.Severity
	GoEnv infra.Environment
}

// Client ...
type Client struct {
	in        ClientInput
	levelCode int
}

// NewClient ...
func NewClient(in ClientInput) *Client {
	levelCode, ok := logLevels[in.Level]
	if !ok {
		levelCode = logLevels[infra.SeverityInfo]
	}

	return &Client{in: in, levelCode: levelCode}
}

// Critical ...
func (c Client) Critical(ctx context.Context, opName infra.OpName, msg string) {
	if c.levelCode >= 1 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityCritical, msg, nil))
	}
}

// Criticalf ...
func (c Client) Criticalf(ctx context.Context, opName infra.OpName, msg string, infos ...interface{}) {
	if c.levelCode >= 1 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityCritical, fmt.Sprintf(msg, infos...), nil))
	}
}

// CriticalCustomData ...
func (c Client) CriticalCustomData(ctx context.Context, opName infra.OpName, msg string, customData infra.CustomData) {
	if c.levelCode >= 1 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityCritical, msg, customData))
	}
}

// Error ...
func (c Client) Error(ctx context.Context, opName infra.OpName, msg string) {
	if c.levelCode >= 2 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityError, msg, nil))
	}
}

// Errorf ...
func (c Client) Errorf(ctx context.Context, opName infra.OpName, msg string, infos ...interface{}) {
	if c.levelCode >= 2 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityError, fmt.Sprintf(msg, infos...), nil))
	}
}

// ErrorCustomData ...
func (c Client) ErrorCustomData(ctx context.Context, opName infra.OpName, msg string, customData infra.CustomData) {
	if c.levelCode >= 2 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityError, msg, customData))
	}
}

// Warning ...
func (c Client) Warning(ctx context.Context, opName infra.OpName, msg string) {
	if c.levelCode >= 3 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityWarning, msg, nil))
	}
}

// Warningf ...
func (c Client) Warningf(ctx context.Context, opName infra.OpName, msg string, infos ...interface{}) {
	if c.levelCode >= 3 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityWarning, fmt.Sprintf(msg, infos...), nil))
	}
}

// WarningCustomData ...
func (c Client) WarningCustomData(ctx context.Context, opName infra.OpName, msg string, customData infra.CustomData) {
	if c.levelCode >= 3 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityWarning, msg, customData))
	}
}

// Info ...
func (c Client) Info(ctx context.Context, opName infra.OpName, msg string) {
	if c.levelCode >= 4 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityInfo, msg, nil))
	}
}

// Infof ...
func (c Client) Infof(ctx context.Context, opName infra.OpName, msg string, infos ...interface{}) {
	if c.levelCode >= 4 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityInfo, fmt.Sprintf(msg, infos...), nil))
	}
}

// InfoCustomData ...
func (c Client) InfoCustomData(ctx context.Context, opName infra.OpName, msg string, customData infra.CustomData) {
	if c.levelCode >= 4 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityInfo, msg, customData))
	}
}

// Debug ...
func (c Client) Debug(ctx context.Context, opName infra.OpName, msg string) {
	if c.levelCode >= 5 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityDebug, msg, nil))
	}
}

// Debugf ...
func (c Client) Debugf(ctx context.Context, opName infra.OpName, msg string, infos ...interface{}) {
	if c.levelCode >= 5 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityDebug, fmt.Sprintf(msg, infos...), nil))
	}
}

// DebugCustomData ...
func (c Client) DebugCustomData(ctx context.Context, opName infra.OpName, msg string, customData infra.CustomData) {
	if c.levelCode >= 5 {
		fmt.Println(c.buildLogValues(ctx, opName, infra.SeverityDebug, msg, customData))
	}
}

func (c Client) buildLogValues(ctx context.Context, opName infra.OpName, level infra.Severity, message string, customData infra.CustomData) string {
	values := logValues{
		Severity: level,
		Message:  fmt.Sprintf("[%s]: %s", opName, message),
	}

	if contextID := ctx.Value(infra.IDContextValueKey); contextID != nil {
		values.ContextID = fmt.Sprintf("%v", contextID)
	}

	if customData != nil {
		values.CustomData = &customData
	}

	logMessage, err := json.Marshal(values)

	if c.in.GoEnv == infra.EnvironmentDevelopment {
		logMessage, err = json.MarshalIndent(values, "", " ")
	}

	if err != nil {
		logMessage = []byte(message)
	}

	return string(logMessage)
}
