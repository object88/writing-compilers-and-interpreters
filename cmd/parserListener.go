package cmd

import (
	"fmt"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

type ParserListener struct{}

// NewParserListener creates a new ParserListener instance
func NewParserListener() *ParserListener {
	return &ParserListener{}
}

// MessageReceived injests a message
func (bl *ParserListener) MessageReceived(m message.Message) {
	switch m0 := m.(type) {
	case *frontend.ParserSummary:
		fmt.Printf(m0.String())
	}
}
