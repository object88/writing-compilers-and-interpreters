package compiler

import (
	"fmt"
	"time"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

// SummaryMessage reveals compiler execution metadata
type SummaryMessage struct {
	message.BaseMessage
	instructionCount int
	elapsedTime      time.Duration
}

// NewSummaryMessage returns a new instance of a SummaryMessage
func NewSummaryMessage(instructionCount int, elapsedTime time.Duration) *SummaryMessage {
	return &SummaryMessage{
		BaseMessage:      message.NewBaseMessage(0),
		instructionCount: instructionCount,
		elapsedTime:      elapsedTime,
	}
}

func (cs *SummaryMessage) String() string {
	return fmt.Sprintf("\n%20d instructions generated\n%20.2f seconds total execution", cs.instructionCount, cs.elapsedTime.Seconds())
}
