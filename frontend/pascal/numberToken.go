package pascal

import (
	"math"
	"strings"
	"unicode"

	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/pkg/errors"
)

const (
	// TODO: This needs to be researched; picked a number out of thin air
	maxExponent int32 = 32
)

// NumberToken represents a hard-coded number in pascal source
type NumberToken struct {
	Token
}

// NewNumberToken returns a new instance of a NumberToken struct
func NewNumberToken(s *frontend.Source) (*NumberToken, error) {
	nt := &NumberToken{
		Token: *NewToken(s),
	}

	var tb strings.Builder
	err := nt.extract(&tb)
	nt.AssignTypeAndText(nt.GetType(), tb.String())

	return nt, err
}

// Extract will process the source to populate the token
func (nt *NumberToken) extract(tb *strings.Builder) error {

	nt.AssignTypeAndText(frontend.TokenType(IntegerTokenType), "")

	wholeDigits, err := nt.unsignedIntegerDigits(tb)
	if err != nil {
		// TODO
	}
	if nt.GetType() == ErrorTokenType {
		return nil
	}

	currentChar, err := nt.CurrentChar()
	if err != nil {
		return errors.Wrap(err, "pascal.NumberToken::extract: ...")
	}

	sawDotDot := false
	fractionalDigits := ""

	if currentChar == '.' {
		peekedChar, err := nt.PeekChar()
		if err != nil {
			return errors.Wrap(err, "pascal.NumberToken::extract: !!!")
		}
		if peekedChar == '.' {
			sawDotDot = true
		} else {
			nt.AssignTypeAndText(frontend.TokenType(RealTokenType), "")
			tb.WriteRune('.')
			_, err = nt.NextChar()
			if err != nil {
				return errors.Wrap(err, "pascal.NumberToken::extract: ???")
			}
			fractionalDigits, err = nt.unsignedIntegerDigits(tb)
			if err != nil {
				return errors.Wrap(err, "pascal.NumberToken::extract: ***")
			}
			if nt.GetType() == ErrorTokenType {
				return nil
			}
		}
	}

	var exponentDigits string
	var exponentSign rune

	currentChar, err = nt.CurrentChar()
	if !sawDotDot && (currentChar == 'E' || currentChar == 'e') {
		nt.AssignTypeAndText(frontend.TokenType(RealTokenType), "")
		tb.WriteRune(currentChar)
		currentChar, err = nt.NextChar()
		if err != nil {
			return errors.Wrap(err, "pascal.NumberToken::extract: ---")
		}

		if currentChar == '+' || currentChar == '-' {
			tb.WriteRune(currentChar)
			exponentSign = currentChar
			currentChar, err = nt.NextChar()
			if err != nil {
				return errors.Wrap(err, "pascal.NumberToken::extract: +++")
			}
		}

		exponentDigits, err = nt.unsignedIntegerDigits(tb)
		if err != nil {
			return errors.Wrap(err, "pascal.NumberToken::extract: ///")
		}
	}

	if nt.GetType() == IntegerTokenType {
		integerValue := nt.computeIntegerValue(wholeDigits)
		if nt.GetType() != ErrorTokenType {
			nt.AssignValue(integerValue)
		}
	} else if nt.GetType() == RealTokenType {
		floatValue := nt.computeFloatValue(wholeDigits, fractionalDigits, exponentDigits, exponentSign)
		if nt.GetType() != ErrorTokenType {
			nt.AssignValue(floatValue)
		}
	}

	return nil
}

func (nt *NumberToken) unsignedIntegerDigits(tb *strings.Builder) (string, error) {
	currentChar, err := nt.CurrentChar()
	if err != nil {
		return "", errors.Wrap(err, "pascal.NumberToken::unsignedIntegerDigits: failed to get current char")
	}
	var digits strings.Builder

	if !unicode.IsDigit(currentChar) {
		nt.AssignTypeAndText(frontend.TokenType(ErrorTokenType), "")
		nt.AssignValue(InvalidNumber)
		return "", nil
	}

	for unicode.IsDigit(currentChar) {
		tb.WriteRune(currentChar)
		digits.WriteRune(currentChar)
		currentChar, err = nt.NextChar()
		if err != nil {
			return "", errors.Wrap(err, "pascal.NumberToken::unsignedIntegerDigits: error encountered reading next character")
		}
	}

	return digits.String(), nil
}

func (nt *NumberToken) computeIntegerValue(digits string) int32 {
	if digits == "" {
		return 0
	}

	var integerValue int32
	var prevValue int32 = -1
	index := 0

	for index < len(digits) && integerValue >= prevValue {
		prevValue = integerValue
		integerValue *= 10
		integerValue += int32(digits[index] - byte('0'))
		index++
	}

	if integerValue >= prevValue {
		// We did not overflow, so that's nice.
		return integerValue
	}

	nt.AssignTypeAndText(frontend.TokenType(ErrorTokenType), "")
	nt.AssignValue(RangeInteger)
	return 0
}

func (nt *NumberToken) computeFloatValue(wholeDigits, fractionalDigits, exponentDigits string, exponentSign rune) float32 {
	exponentValue := nt.computeIntegerValue(exponentDigits)
	digits := wholeDigits

	if exponentSign == '-' {
		exponentValue = -exponentValue
	}

	if fractionalDigits != "" {
		exponentValue -= int32(len(fractionalDigits))
		digits += fractionalDigits
	}

	exp := exponentValue + int32(len(wholeDigits))
	if exp < 0 {
		exp = -exp
	}

	if exp > maxExponent {
		nt.AssignTypeAndText(frontend.TokenType(ErrorTokenType), "")
		nt.AssignValue(RangeReal)
		return 0.0
	}

	var floatValue float64

	for index := 0; index < len(digits); index++ {
		floatValue *= 10
		floatValue += float64(digits[index] - byte('0'))
	}

	if exponentValue != 0 {
		floatValue *= math.Pow10(int(exponentValue))
	}

	return float32(floatValue)
}
