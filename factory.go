package wcai

import (
	"strings"

	"github.com/object88/writing-compilers-and-interpreters/backend"
	"github.com/object88/writing-compilers-and-interpreters/backend/compiler"
	"github.com/object88/writing-compilers-and-interpreters/backend/interpreter"
	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/frontend/pascal"
	"github.com/object88/writing-compilers-and-interpreters/message"
	"github.com/pkg/errors"
)

func NewBackend(operation string, messageHandler *message.MessageHandler) (backend.Backend, error) {
	switch strings.ToLower(operation) {
	case "compile":
		return compiler.NewCodeGenerator(messageHandler), nil
	case "execute":
		return interpreter.NewInterpreter(messageHandler), nil
	default:
		return nil, errors.Errorf("Operation '%s' is not supported", operation)

	}
}

func NewParser(language, parserType string, source *frontend.Source, messageHandler *message.MessageHandler) (frontend.Parser, error) {
	switch strings.ToLower(language) {
	case "pascal":
		switch strings.ToLower(parserType) {
		case "top-down":
			s := pascal.NewScanner(source)
			return pascal.NewParserTD(s, messageHandler), nil
		default:
			return nil, errors.Errorf("For language '%s'; parser type '%s' is not supported", language, parserType)
		}
	default:
		return nil, errors.Errorf("Language '%s' is not supported", language)
	}
}
