package frontend

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

type ParserSummary struct {
	message.BaseMessage
	errorCount  int
	elapsedTime time.Duration
}

func NewParserSummary(errorCount int, elapsedTime time.Duration) *ParserSummary {
	return &ParserSummary{
		BaseMessage: message.NewBaseMessage(3),
		errorCount:  errorCount,
		elapsedTime: elapsedTime,
	}
}
