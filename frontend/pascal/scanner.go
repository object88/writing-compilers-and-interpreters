package pascal

import (
	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/pkg/errors"
)

// Scanner is a Pascal language specific scanner
type Scanner struct {
	frontend.BaseScanner
}

// NewScanner instantiates and returns a new Scanner
func NewScanner(s *frontend.Source) *Scanner {
	return &Scanner{
		BaseScanner: *frontend.NewBaseScanner(s),
	}
}

func (ps *Scanner) extractToken() (frontend.Token, error) {
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

func (ps *Scanner) NextToken() (frontend.Token, error) {
	return ps.extractToken()
}
