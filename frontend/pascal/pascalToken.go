package pascal

import (
	"github.com/object88/writing-compilers-and-interpreters/frontend"
)

type PascalToken struct {
	frontend.BaseToken
}

func NewPascalToken(s *frontend.Source) *PascalToken {
	return &PascalToken{
		BaseToken: frontend.NewBaseToken(s),
	}
}
