package pascal

import "github.com/object88/writing-compilers-and-interpreters/frontend"

// ErrorToken represents a token that is where it should not be
type ErrorToken struct {
	Token
}

// NewErrorToken returns a new instance of an ErrorToken
func NewErrorToken(s *frontend.Source, ec errorCode, r rune) *ErrorToken {
	et := &ErrorToken{
		Token: *NewToken(s),
	}

	et.AssignTypeAndText(frontend.TokenType(ErrorTokenType), string(r))
	et.AssignValue(ec)

	return et
}
