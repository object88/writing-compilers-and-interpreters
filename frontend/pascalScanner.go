package frontend

import (
	"io"

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
	_, err := ps.BaseScanner.CurrentChar()
	if err != nil {
		if err == io.EOF {
			return NewEofToken(ps.source), nil
		}
		return nil, errors.Wrapf(err, "Failed to retrieve current char from base scanner")
	}

	t := NewBaseToken(ps.source)

	return t, nil
}
