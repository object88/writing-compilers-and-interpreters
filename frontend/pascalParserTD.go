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
	_ = endTime.Sub(startTime)

	var m interface{}
	p.mh.SendMessage(m)

	return nil
}

func (p *PascalParserTD) GetErrorCount() int {
	return 0
}
