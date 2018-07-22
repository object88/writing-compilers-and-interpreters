package frontend

import (
	"github.com/object88/writing-compilers-and-interpreters/intermediate"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

type Parser interface {
	Parse() error
	GetErrorCount() int
	GetICode() intermediate.ICode
	GetSymTab() intermediate.SymTab
	CurrentToken() (Token, error)
	GetMessageHandler() *message.MessageHandler
}

type BaseParser struct {
	scanner Scanner
	icode   intermediate.ICode
	symTab  intermediate.SymTab
	mh      *message.MessageHandler
}

func NewBaseParser(scanner Scanner, messageHandler *message.MessageHandler) *BaseParser {
	return &BaseParser{
		scanner: scanner,
		mh:      messageHandler,
	}
}

func (bp *BaseParser) CurrentToken() (Token, error) {
	return bp.scanner.CurrentToken()
}

func (bp *BaseParser) NextToken() (Token, error) {
	return bp.scanner.NextToken()
}

func (bp *BaseParser) GetICode() intermediate.ICode {
	return bp.icode
}

func (bp *BaseParser) GetSymTab() intermediate.SymTab {
	return bp.symTab
}

func (bp *BaseParser) GetMessageHandler() *message.MessageHandler {
	return bp.mh
}
