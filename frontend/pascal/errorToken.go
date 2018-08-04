package pascal

import "github.com/object88/writing-compilers-and-interpreters/frontend"

type ErrorToken struct {
	Token
}

func NewErrorToken(s *frontend.Source, ec errorCode, r rune) *ErrorToken {
	et := &ErrorToken{
		Token: *NewToken(s),
	}

	et.AssignTypeAndText(frontend.TokenType(ErrorTokenType), string(r))
	et.AssignValue(ec)

	return et
}
