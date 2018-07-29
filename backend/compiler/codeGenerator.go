package compiler

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/backend"
	"github.com/object88/writing-compilers-and-interpreters/intermediate"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

// CodeGenerator is a compiler back end
type CodeGenerator struct {
	backend.BaseBackend
}

// NewCodeGenerator creates a new CodeGenerator instance
func NewCodeGenerator(mh *message.MessageHandler) *CodeGenerator {
	return &CodeGenerator{
		BaseBackend: backend.NewBaseBackend(mh),
	}
}

// Process transforms icode and symbol tables into machine language
func (cg *CodeGenerator) Process(iCode intermediate.ICode, symTab intermediate.SymTab) error {
	startTime := time.Now()
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	instructionCount := 0

	m := NewSummaryMessage(instructionCount, elapsedTime)
	cg.BaseBackend.MessageHandler.SendMessage(m)

	return nil
}
