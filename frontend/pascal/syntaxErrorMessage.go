package pascal

import (
	"strings"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
)

// SyntaxErrorMessage message is the results of a syntactic error in a pascal
// source code
type SyntaxErrorMessage struct {
	TokenMessage
	err error
}

// NewSyntaxErrorMessage returns a new instance of the SyntaxErrorMessage
// struct
func NewSyntaxErrorMessage(token frontend.Token, err error) *SyntaxErrorMessage {
	return &SyntaxErrorMessage{
		TokenMessage: *NewTokenMessage(token),
		err:          err,
	}
}

func (tm *SyntaxErrorMessage) String() string {
	var sb strings.Builder

	spaceCount := 5 + tm.token.GetPosition()
	sb.WriteString(strings.Repeat(" ", spaceCount))
	sb.WriteString("^\n*** ")
	sb.WriteString(tm.err.Error())

	t := tm.token.GetText()
	if len(t) > 0 {
		sb.WriteString(" [at \"")
		sb.WriteString(t)
		sb.WriteString("\"]")
	}

	return sb.String()
}
