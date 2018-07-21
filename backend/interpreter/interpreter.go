package interpreter

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/backend"
	"github.com/object88/writing-compilers-and-interpreters/intermediate"
)

type Interpreter struct {
	backend.BaseBackend
}

func (i *Interpreter) Process(iCode intermediate.ICode, symTab intermediate.SymTab) error {
	startTime := time.Now()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	executionCount := 0
	runtimeErrors := 0

	m := NewInterpreterSummary(executionCount, runtimeErrors, elapsedTime)
	i.BaseBackend.MessageHandler.SendMessage(m)

	return nil
}
