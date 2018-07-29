package interpreter

import (
	"fmt"
	"time"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

// SummaryMessage reveals interpreter execution metadata
type SummaryMessage struct {
	message.BaseMessage
	executionCount int
	runtimeErrors  int
	elapsedTime    time.Duration
}

// NewSummaryMessage returns a new instance of a SummaryMessage
func NewSummaryMessage(executionCount, runtimeErrors int, elapsedTime time.Duration) *SummaryMessage {
	return &SummaryMessage{
		BaseMessage:    message.NewBaseMessage(0),
		executionCount: executionCount,
		runtimeErrors:  runtimeErrors,
		elapsedTime:    elapsedTime,
	}
}

func (is *SummaryMessage) String() string {
	return fmt.Sprintf("\n%20d statements executed\n%20d runtime errors\n%20.2f seconds total execution\n", is.executionCount, is.runtimeErrors, is.elapsedTime.Seconds())
}
