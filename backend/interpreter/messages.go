package interpreter

import (
	"fmt"
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

func (is *InterpreterSummary) String() string {
	return fmt.Sprintf("\n%20d statements executed\n%20d runtime errors\n%20.2f seconds total execution\n", is.executionCount, is.runtimeErrors, is.elapsedTime.Seconds())
}
