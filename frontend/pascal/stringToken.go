package pascal

import (
	"strings"
	"unicode"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/pkg/errors"
)

// StringToken contains a string
type StringToken struct {
	Token
	value string
}

// NewStringToken returns a new instance of a StringToken
func NewStringToken(s *frontend.Source) (*StringToken, error) {
	st := &StringToken{
		Token: *NewToken(s),
	}

	err := st.extract()

	return st, err
}

func (st *StringToken) extract() error {
	var tb strings.Builder
	var vb strings.Builder

	currentChar, err := st.NextChar()
	if err != nil {
		return errors.Wrap(err, "pascal.StringToken::extract: failed to consume initial single-quote")
	}
	tb.WriteRune('\'')

	for {
		// TODO: actually consume this string.
		if unicode.Is(unicode.Pattern_White_Space, currentChar) {
			currentChar = ' '
		}

		if currentChar != '\'' && currentChar != frontend.EOF {
			tb.WriteRune(currentChar)
			vb.WriteRune(currentChar)
			currentChar, err = st.NextChar()
		}

		if currentChar == '\'' {
			peekedChar, err := st.PeekChar()
			if err != nil {
				return errors.Wrap(err, "Error peeking")
			}
			for currentChar == '\'' && peekedChar == '\'' {
				tb.WriteString("''")
				vb.WriteRune('\'')
				_, err = st.NextChar()
				currentChar, err = st.NextChar()
			}
		}

		if currentChar == '\'' || currentChar == frontend.EOF {
			break
		}
	}

	if currentChar == '\'' {
		_, err = st.NextChar()
		if err != nil {
			return errors.Wrap(err, "pascal.StringToken::extract: failed to consume final single-quote")
		}
		tb.WriteRune('\'')
		st.AssignTypeAndText(frontend.TokenType(StringTokenType), tb.String())
		st.AssignValue(vb.String())
	} else {
		st.AssignTypeAndText(frontend.TokenType(ErrorTokenType), tb.String())
		st.AssignValue("UNEXPECTED_EOF")
	}

	return nil
}
