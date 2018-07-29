package interpreter

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/backend"
	"github.com/object88/writing-compilers-and-interpreters/intermediate"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

// Interpreter executes code without compiling
type Interpreter struct {
	backend.BaseBackend
}

// NewInterpreter returns a new instance of Interpreter
func NewInterpreter(messageHandler *message.MessageHandler) *Interpreter {
	return &Interpreter{
		BaseBackend: backend.NewBaseBackend(messageHandler),
	}
}

// Process executes icode and symbol tables
func (i *Interpreter) Process(iCode intermediate.ICode, symTab intermediate.SymTab) error {
	startTime := time.Now()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	executionCount := 0
	runtimeErrors := 0

	m := NewSummaryMessage(executionCount, runtimeErrors, elapsedTime)
	i.BaseBackend.MessageHandler.SendMessage(m)

	return nil
}
