package interpreter

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

type InterpreterSummary struct {
	message.BaseMessage
	executionCount int
	runtimeErrors  int
	elapsedTime    time.Duration
}

func NewInterpreterSummary(executionCount, runtimeErrors int, elapsedTime time.Duration) *InterpreterSummary {
	return &InterpreterSummary{
		BaseMessage:    message.NewBaseMessage(0),
		executionCount: executionCount,
		runtimeErrors:  runtimeErrors,
		elapsedTime:    elapsedTime,
	}
}
