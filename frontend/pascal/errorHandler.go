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

func (eh *errorHandler) flag(token frontend.Token, errCode errorCode, parser frontend.Parser) {
	m := NewSyntaxErrorMessage(token, errCode)
	parser.GetMessageHandler().SendMessage(m)

	eh.errorCount++
	if eh.errorCount > maxErrors {
		eh.abortTranslation(TooManyErrors, parser)
	}
}

func (eh *errorHandler) abortTranslation(reason AbortionReason, parser frontend.Parser) {

	// m := NewSyntaxErrorMessage(reason.)
	// parser.GetMessageHandler().SendMessage(m)
}
