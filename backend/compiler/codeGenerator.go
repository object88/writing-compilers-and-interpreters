package compiler

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/backend"
	"github.com/object88/writing-compilers-and-interpreters/intermediate"
)

type CodeGenerator struct {
	backend.BaseBackend
}

func (cg *CodeGenerator) Process(iCode intermediate.ICode, symTab intermediate.SymTab) error {
	startTime := time.Now()
	endTime := time.Now()
	elapsedTime :	= endTime.Sub(startTime)
	instructionCount := 0

	m := NewCompilerSummary(instructionCount, elapsedTime)
	cg.BaseBackend.MessageHandler.SendMessage(m)

	return nil
}
