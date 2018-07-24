package pascal

import (
	"time"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/message"
)

type PascalParserTD struct {
	frontend.BaseParser
}

func NewPascalParserTD(scanner frontend.Scanner, messageHandler *message.MessageHandler) *PascalParserTD {
	p := &PascalParserTD{
		BaseParser: *frontend.NewBaseParser(scanner, messageHandler),
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
		if _, ok := t.(*frontend.EOFToken); ok {
			break
		}
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	statementCount := 0
	errorCount := 0

	m := frontend.NewParserSummary(statementCount, errorCount, elapsedTime)
	p.GetMessageHandler().SendMessage(m)

	return nil
}

func (p *PascalParserTD) CurrentToken() (frontend.Token, error) {
	return p.BaseParser.CurrentToken()
}

func (p *PascalParserTD) GetErrorCount() int {
	return 0
}
