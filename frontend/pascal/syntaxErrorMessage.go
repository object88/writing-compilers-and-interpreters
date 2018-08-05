package pascal

import (
	"strings"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
)

// SyntaxErrorMessage message is the results of a syntactic error in a pascal
// source code
type SyntaxErrorMessage struct {
	TokenMessage
	errCode errorCode
}

// NewSyntaxErrorMessage returns a new instance of the SyntaxErrorMessage
// struct
func NewSyntaxErrorMessage(token frontend.Token, errCode errorCode) *SyntaxErrorMessage {
	return &SyntaxErrorMessage{
		TokenMessage: *NewTokenMessage(token),
		errCode:      errCode,
	}
}

func (tm *SyntaxErrorMessage) String() string {
	var sb strings.Builder

	spaceCount := 5 + tm.token.GetPosition()
	sb.WriteString(strings.Repeat(" ", spaceCount))
	sb.WriteString("^\n*** ")
	sb.WriteString(tm.errCode.message)

	t := tm.token.GetText()
	if len(t) > 0 {
		sb.WriteString(" [at \"")
		sb.WriteString(t)
		sb.WriteString("\"]")
	}
	sb.WriteRune('\n')

	return sb.String()
}
