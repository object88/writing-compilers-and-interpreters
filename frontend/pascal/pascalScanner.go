package pascal

import (
	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/pkg/errors"
)

// PascalScanner is a Pascal language specific scanner
type PascalScanner struct {
	frontend.BaseScanner
}

// NewPascalScanner instantiates and returns a new PascalScanner
func NewPascalScanner(s *frontend.Source) *PascalScanner {
	return &PascalScanner{
		BaseScanner: *frontend.NewBaseScanner(s),
	}
}

func (ps *PascalScanner) extractToken() (frontend.Token, error) {
	r, err := ps.BaseScanner.CurrentChar()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to retrieve current char from base scanner")
	}

	if r == frontend.EOF {
		return frontend.NewEOFToken(ps.Source), nil
	}

	t := frontend.NewBaseToken(ps.Source)
	return &t, nil
}

func (ps *PascalScanner) NextToken() (frontend.Token, error) {
	return ps.extractToken()
}
