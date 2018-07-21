package frontend

import (
	"github.com/object88/writing-compilers-and-interpreters/intermediate"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

type Parser interface {
	Parse() error
	GetErrorCount() int
	CurrentToken() Token
}

type BaseParser struct {
	scanner Scanner
	icode   intermediate.ICode
	mh      *message.MessageHandler
}

func NewBaseParser(scanner Scanner, messageHandler *message.MessageHandler) *BaseParser {
	return &BaseParser{
		scanner: scanner,
		mh:      messageHandler,
	}
}

func (bp *BaseParser) CurrentToken() (*Token, error) {
	return bp.scanner.CurrentToken()
}

func (bp *BaseParser) NextToken() (*Token, error) {
	return bp.scanner.NextToken()
}
