package pascal

import (
	"strings"
	"unicode"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/pkg/errors"
)

// WordToken is a word in the pascal source code
type WordToken struct {
	Token
}

// NewWordToken returns a new instance of a WordToken
func NewWordToken(s *frontend.Source) (*WordToken, error) {
	wt := &WordToken{
		Token: *NewToken(s),
	}

	err := wt.extract()

	return wt, err
}

// Extract will process the source to populate the token
func (wt *WordToken) extract() error {
	var sb strings.Builder

	r, err := wt.BaseToken.CurrentChar()
	if err != nil {
		return errors.Wrap(err, "pascal.WordToken::Extract: while extracting first char")
	}
	for {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			break
		}
		sb.WriteRune(r)
		r, err = wt.BaseToken.NextChar()
		if err != nil {
			return errors.Wrap(err, "pascal.WordToken::Extract: while extracting char")
		}
	}

	text := sb.String()
	var typ frontend.TokenType
	if t, ok := reservedTokenTypes[strings.ToUpper(text)]; ok {
		// text is a reserved word
		typ = t
	} else {
		typ = IdentifierTokenType
	}

	// Apply text & typ to the token..., somehow.
	wt.BaseToken.AssignTypeAndText(frontend.TokenType(typ), text)

	return nil
}
