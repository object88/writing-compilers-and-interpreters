package compiler

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

type CompilerSummary struct {
	message.BaseMessage
	instructionCount int
	elapsedTime      time.Duration
}

func NewCompilerSummary(instructionCount int, elapsedTime time.Duration) *CompilerSummary {
	return &CompilerSummary{
		BaseMessage:      message.NewBaseMessage(0),
		instructionCount: instructionCount,
		elapsedTime:      elapsedTime,
	}
}
