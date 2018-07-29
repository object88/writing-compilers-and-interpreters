package cmd

import (
	"fmt"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/frontend/pascal"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

// ParserListener listens for messages from the parser
type ParserListener struct{}

// NewParserListener creates a new ParserListener instance
func NewParserListener() *ParserListener {
	return &ParserListener{}
}

// MessageReceived injests a message
func (bl *ParserListener) MessageReceived(m message.Message) {
	switch m0 := m.(type) {
	case *frontend.ParserSummaryMessage:
		fmt.Printf(m0.String())
	case *pascal.SyntaxErrorMessage:
		fmt.Printf(m0.String())
	case *pascal.TokenMessage:
		fmt.Printf(m0.String())
	}
}
