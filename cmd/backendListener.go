package cmd

import (
	"fmt"

	"github.com/object88/writing-compilers-and-interpreters/backend/compiler"
	"github.com/object88/writing-compilers-and-interpreters/backend/interpreter"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

type BackendListener struct{}

// NewBackendListener creates a new BackendListener instance
func NewBackendListener() *BackendListener {
	return &BackendListener{}
}

// MessageReceived injests a message
func (bl *BackendListener) MessageReceived(m message.Message) {
	switch m0 := m.(type) {
	case *compiler.CompilerSummary:
		fmt.Printf(m0.String())
	case *interpreter.InterpreterSummary:
		fmt.Printf(m0.String())
	}
}
