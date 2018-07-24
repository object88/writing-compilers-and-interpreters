package frontend

import (
	"fmt"
	"time"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

type ParserSummary struct {
	message.BaseMessage
	statementCount int
	errorCount     int
	elapsedTime    time.Duration
}

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

type SourceLine struct {
	message.BaseMessage
	lineNumber int
	lineText   string
}

func NewSourceLine(lineNumber int, lineText string) *SourceLine {
	return &SourceLine{
		BaseMessage: message.NewBaseMessage(4),
		lineNumber:  lineNumber,
		lineText:    lineText,
	}
}

func (sl *SourceLine) String() string {
	return fmt.Sprintf("%03d %s", sl.lineNumber, sl.lineText)
}
