package pascal

import (
	"github.com/object88/writing-compilers-and-interpreters/frontend"
)

type AbortionReason int

const (
	maxErrors int = 25

	TooManyErrors AbortionReason = iota
)

type errorHandler struct {
	errorCount int
}

func (eh *errorHandler) flag(token frontend.Token, err error, parser frontend.Parser) {
	parser.GetMessageHandler().SendMessage(nil)

	eh.errorCount++
	if eh.errorCount > maxErrors {
		eh.abortTranslation(TooManyErrors, parser)
	}
}

func (eh *errorHandler) abortTranslation(reason AbortionReason, parser frontend.Parser) {}
