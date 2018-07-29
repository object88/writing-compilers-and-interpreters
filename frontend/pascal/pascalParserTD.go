package pascal

import (
	"errors"
	"time"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

// PascalParserTD is a top-down Pascal language parser
type PascalParserTD struct {
	frontend.BaseParser
	errorHandler errorHandler
}

func NewPascalParserTD(scanner frontend.Scanner, messageHandler *message.MessageHandler) *PascalParserTD {
	p := &PascalParserTD{
		BaseParser:   *frontend.NewBaseParser(scanner, messageHandler),
		errorHandler: errorHandler{},
	}
	return p
}

func (p *PascalParserTD) Parse() error {
	mh := p.GetMessageHandler()
	startTime := time.Now()
	var t frontend.Token

	for {
		var err error
		t, err = p.NextToken()
		if err != nil {
			return err
		}

		if _, ok := t.(*frontend.EOFToken); ok {
			break
		}

		if PascalTokenType(t.GetType()) == ERROR {
			p.errorHandler.flag(t, errors.New("WTF"), p)
		} else {
			tm := NewTokenMessage(t)
			mh.SendMessage(tm)
		}

	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	m := frontend.NewParserSummaryMessage(t.GetLineNumber(), p.errorHandler.errorCount, elapsedTime)
	p.GetMessageHandler().SendMessage(m)

	return nil
}

func (p *PascalParserTD) CurrentToken() (frontend.Token, error) {
	return p.BaseParser.CurrentToken()
}

// GetErrorCount returns the number of errors encountered by the parser
func (p *PascalParserTD) GetErrorCount() int {
	return 0
}
