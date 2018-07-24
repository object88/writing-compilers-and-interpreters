package frontend

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

type PascalParserTD struct {
	BaseParser
}

func NewPascalParserTD(scanner Scanner, messageHandler *message.MessageHandler) *PascalParserTD {
	p := &PascalParserTD{
		BaseParser: *NewBaseParser(scanner, messageHandler),
	}
	return p
}

func (p *PascalParserTD) Parse() error {
	startTime := time.Now()

	for {
		t, err := p.NextToken()
		if err != nil {
			return err
		}
		if _, ok := t.(*EOFToken); ok {
			break
		}
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	statementCount := 0
	errorCount := 0

	m := NewParserSummary(statementCount, errorCount, elapsedTime)
	p.mh.SendMessage(m)

	return nil
}

func (p *PascalParserTD) CurrentToken() (Token, error) {
	return p.BaseParser.CurrentToken()
}

func (p *PascalParserTD) GetErrorCount() int {
	return 0
}
