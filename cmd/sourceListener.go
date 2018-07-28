package cmd

import (
	"fmt"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

type SourceListener struct{}

// NewSourceListener creates a new SourceListener instance
func NewSourceListener() *SourceListener {
	return &SourceListener{}
}

// MessageReceived injests a message
func (sl *SourceListener) MessageReceived(m message.Message) {
	switch m0 := m.(type) {
	case *frontend.SourceLineMessage:
		fmt.Printf(m0.String())
	}
}
