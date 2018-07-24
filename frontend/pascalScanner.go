package frontend

import (
	"github.com/pkg/errors"
)

// PascalScanner is a Pascal language specific scanner
type PascalScanner struct {
	BaseScanner
}

// NewPascalScanner instantiates and returns a new PascalScanner
func NewPascalScanner(s *Source) *PascalScanner {
	return &PascalScanner{
		BaseScanner: *NewBaseScanner(s),
	}
}

func (ps *PascalScanner) extractToken() (Token, error) {
	r, err := ps.BaseScanner.CurrentChar()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to retrieve current char from base scanner")
	}

	if r == EOF {
		return NewEOFToken(ps.source), nil
	}

	t := NewBaseToken(ps.source)
	return t, nil
}

func (ps *PascalScanner) NextToken() (Token, error) {
	return ps.extractToken()
}
