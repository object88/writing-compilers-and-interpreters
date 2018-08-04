package pascal

import "github.com/object88/writing-compilers-and-interpreters/frontend"

const (
	// AndTokenType is a Pascal keyword token type for "AND"
	AndTokenType TokenType = iota + 100

	PlusTokenType
	MinusTokenType
	AsteriskTokenType
	SlashTokenType
	CommaTokenType
	SemiColonTokenType
	SingleQuoteTokenType
	EqualsTokenType
	OpenParenthesisTokenType
	CloseParenthesisTokenType
	OpenHardBraceTokenType
	CloseHardBraceTokenType
	OpenCurlyBraceTokenType
	CloseCurlyBraceTokenType
	CarotTokenType
	ColonEqualsTokenType
	ColonTokenType
	LessThanOrEqualsTokenType
	LessThanOrLessThanTokenType
	LessThanTokenType
	GreaterThanOrEqualsTokenType
	GreaterThanTokenType
	DotTokenType
	DotDotTokenType

	IdentifierTokenType
	IntegerTokenType
	RealTokenType
	StringTokenType
	ErrorTokenType
)

// TokenType is a Pascal-specific token type
type TokenType frontend.TokenType

var reservedTokenTypes = map[string]TokenType{
	"AND": AndTokenType,
}

var specialSymbolTokenTypes = map[string]TokenType{
	"+":  PlusTokenType,
	"-":  MinusTokenType,
	"*":  AsteriskTokenType,
	"/":  SlashTokenType,
	",":  CommaTokenType,
	";":  SemiColonTokenType,
	"'":  SingleQuoteTokenType,
	"=":  EqualsTokenType,
	"(":  OpenParenthesisTokenType,
	")":  CloseParenthesisTokenType,
	"[":  OpenHardBraceTokenType,
	"]":  CloseHardBraceTokenType,
	"{":  OpenCurlyBraceTokenType,
	"}":  CloseCurlyBraceTokenType,
	"^":  CarotTokenType,
	":=": ColonEqualsTokenType,
	":":  ColonTokenType,
	"<=": LessThanOrEqualsTokenType,
	"<>": LessThanOrLessThanTokenType,
	"<":  LessThanTokenType,
	">=": GreaterThanOrEqualsTokenType,
	">":  GreaterThanTokenType,
	".":  DotTokenType,
	"..": DotDotTokenType,
}
