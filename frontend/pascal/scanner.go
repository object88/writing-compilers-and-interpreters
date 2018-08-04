package pascal

import (
	"unicode"

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
	ps.skipWhiteSpace()

	var token frontend.Token

	r, err := ps.BaseScanner.CurrentChar()
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to retrieve current char from base scanner")
	}

	if r == frontend.EOF {
		return frontend.NewEOFToken(ps.Source), nil
	}

	if unicode.IsLetter(r) {
		token, err = NewWordToken(ps.Source)
	} else if unicode.IsNumber(r) {
		token, err = NewNumberToken(ps.Source)
	} else if r == '\'' {
		token, err = NewStringToken(ps.Source)
	} else if false {
		// Special character case
		token, err = NewSpecialSymbolToken(ps.Source)
	} else {
		// Invalid character case
		token = NewErrorToken(ps.Source, InvalidCharacter, r)
		_, err = ps.NextChar()
	}
	// t := frontend.NewBaseToken(ps.Source)

	if err != nil {
		// Handle the error from constructing and extracting the token
	}

	return token, nil
}

func (ps *Scanner) NextToken() (frontend.Token, error) {
	return ps.extractToken()
}

// skipWhiteSpace consumes white space and comments
func (ps *Scanner) skipWhiteSpace() error {
	for {
		currentChar, err := ps.CurrentChar()
		if err != nil {
			return errors.Wrap(err, "pascal.Scanner::skipWhiteSpace: While skipping white space")
		}

		if !unicode.Is(unicode.Pattern_White_Space, currentChar) && currentChar != '{' {
			break
		}

		if currentChar == '{' {
			var nextChar rune
			for {
				nextChar, err = ps.NextChar()
				if err != nil {
					return errors.Wrap(err, "pascal.Scanner::skipWhiteSpace: While getting next char inside comment block")
				}
				if nextChar == '}' || nextChar == frontend.EOF {
					break
				}
			}

			if nextChar == '}' {
				_, err = ps.NextChar()
				if err != nil {
					return errors.Wrap(err, "pascal.Scanner::skipWhiteSpace: Whie getting next char at end of comment block")
				}
			}
		} else {
			_, err = ps.NextChar()
			if err != nil {
				return errors.Wrap(err, "pascal.Scanner::skipWhiteSpace: While consuming white space")
			}
		}

	}

	return nil
}
