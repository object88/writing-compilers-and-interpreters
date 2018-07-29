package pascal

import (
	"github.com/object88/writing-compilers-and-interpreters/frontend"
)

// Token is a pascal token
type Token struct {
	frontend.BaseToken
}

// NewToken returns a new pascal token from the source
func NewToken(s *frontend.Source) *Token {
	return &Token{
		BaseToken: frontend.NewBaseToken(s),
	}
}
