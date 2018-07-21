package frontend

import (
	"io"
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
		_, err := p.NextToken()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	errorCount := 0

	m := NewParserSummary(errorCount, elapsedTime)
	p.mh.SendMessage(m)

	return nil
}

func (p *PascalParserTD) CurrentToken() (Token, error) {
	return p.BaseParser.CurrentToken()
}

func (p *PascalParserTD) GetErrorCount() int {
	return 0
}
