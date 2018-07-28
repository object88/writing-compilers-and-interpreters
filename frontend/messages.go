package frontend

import (
	"fmt"
	"time"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

// ParserSummary message is the results from parsing a file
type ParserSummary struct {
	message.BaseMessage
	statementCount int
	errorCount     int
	elapsedTime    time.Duration
}

// NewParserSummary returns a new instance of the ParserSummary struct
func NewParserSummary(statementCount, errorCount int, elapsedTime time.Duration) *ParserSummary {
	return &ParserSummary{
		BaseMessage:    message.NewBaseMessage(3),
		statementCount: statementCount,
		errorCount:     errorCount,
		elapsedTime:    elapsedTime,
	}
}

func (ps *ParserSummary) String() string {
	return fmt.Sprintf("\n%20d source lines\n%20d syntax errors\n%20.2f seconds total execution\n", ps.statementCount, ps.errorCount, ps.elapsedTime.Seconds())
}

// SourceLineMessage is a Message containing a line number and that line's text
type SourceLineMessage struct {
	message.BaseMessage
	lineNumber int
	lineText   string
}

// NewSourceLineMessage creates a new SourceLineMessage
func NewSourceLineMessage(lineNumber int, lineText string) *SourceLineMessage {
	return &SourceLineMessage{
		BaseMessage: message.NewBaseMessage(4),
		lineNumber:  lineNumber,
		lineText:    lineText,
	}
}

func (sl *SourceLineMessage) String() string {
	return fmt.Sprintf("%03d %s", sl.lineNumber, sl.lineText)
}
