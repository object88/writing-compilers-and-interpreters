package pascal

import (
	"github.com/object88/writing-compilers-and-interpreters/frontend"
)

type Token struct {
	frontend.BaseToken
}

func NewToken(s *frontend.Source) *Token {
	return &Token{
		BaseToken: frontend.NewBaseToken(s),
	}
}
