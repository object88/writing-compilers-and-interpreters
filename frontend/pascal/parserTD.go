package pascal

import (
	"errors"
	"time"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

// ParserTD is a top-down Pascal language parser
type ParserTD struct {
	frontend.BaseParser
	errorHandler errorHandler
}

// NewParserTD returns a new ParserTD instance
func NewParserTD(scanner frontend.Scanner, messageHandler *message.MessageHandler) *ParserTD {
	p := &ParserTD{
		BaseParser:   *frontend.NewBaseParser(scanner, messageHandler),
		errorHandler: errorHandler{},
	}
	return p
}

// Parse consumes from the scanner and produces tokens
func (p *ParserTD) Parse() error {
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

		if t.GetType() == ErrorTokenType {
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

func (p *ParserTD) CurrentToken() (frontend.Token, error) {
	return p.BaseParser.CurrentToken()
}

// GetErrorCount returns the number of errors encountered by the parser
func (p *ParserTD) GetErrorCount() int {
	return 0
}
