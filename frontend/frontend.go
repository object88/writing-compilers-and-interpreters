package frontend

import (
	"strings"

	"github.com/pkg/errors"
)

func NewParser(language, parserType string, source *Source) (Parser, error) {

	switch strings.ToLower(language) {
	case "pascal":
		switch strings.ToLower(parserType) {
		case "top-down":
			s := NewPascalScanner(source)
			return NewPascalParserTD(s, nil), nil
		default:
			return nil, errors.Errorf("For language '%s'; parser type '%s' is not supported", language, parserType)
		}
	default:
		return nil, errors.Errorf("Language '%s' is not supported", language)
	}
}
