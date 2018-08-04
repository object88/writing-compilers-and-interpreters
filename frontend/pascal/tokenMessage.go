package pascal

import (
	"fmt"
	"strings"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

// TokenMessage message is the results from parsing a file
type TokenMessage struct {
	message.BaseMessage
	token frontend.Token
}

// NewTokenMessage returns a new instance of the TokenMessage struct
func NewTokenMessage(token frontend.Token) *TokenMessage {
	return &TokenMessage{
		BaseMessage: message.NewBaseMessage(4),
		token:       token,
	}
}

func (tm *TokenMessage) String() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf(">>> % 15s line=%03d, pos=%2d, text=\"%s\"\n", tm.token.GetType().String(), tm.token.GetLineNumber(), tm.token.GetPosition(), tm.token.GetText()))
	v := tm.token.GetValue()
	if v != nil {
		sb.WriteString(">>>                 value=\"")
		switch v0 := v.(type) {
		case string:
			sb.WriteString(v0)
		default:
			sb.WriteString(fmt.Sprintf("%v", v0))
		}
		sb.WriteString("\"\n")
	}

	return sb.String()
}
