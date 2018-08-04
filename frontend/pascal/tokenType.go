package pascal

import "github.com/object88/writing-compilers-and-interpreters/frontend"

const (
	AndTokenTypeCode frontend.TokenTypeCode = iota + 100
	BeginTokenTypeCode

	PlusTokenTypeCode frontend.TokenTypeCode = iota + 200
	MinusTokenTypeCode
	AsteriskTokenTypeCode
	SlashTokenTypeCode
	CommaTokenTypeCode
	SemiColonTokenTypeCode
	SingleQuoteTokenTypeCode
	EqualsTokenTypeCode
	OpenParenthesisTokenTypeCode
	CloseParenthesisTokenTypeCode
	OpenHardBraceTokenTypeCode
	CloseHardBraceTokenTypeCode
	OpenCurlyBraceTokenTypeCode
	CloseCurlyBraceTokenTypeCode
	CarotTokenTypeCode
	ColonEqualsTokenTypeCode
	ColonTokenTypeCode
	LessThanOrEqualsTokenTypeCode
	LessThanOrLessThanTokenTypeCode
	LessThanTokenTypeCode
	GreaterThanOrEqualsTokenTypeCode
	GreaterThanTokenTypeCode
	DotTokenTypeCode
	DotDotTokenTypeCode

	IdentifierTokenTypeCode frontend.TokenTypeCode = iota + 200
	IntegerTokenTypeCode
	RealTokenTypeCode
	StringTokenTypeCode
	ErrorTokenTypeCode
)

// AndTokenType is a Pascal keyword token type for "AND"
var AndTokenType = frontend.NewTokenType(AndTokenTypeCode, "And")
var BeginTokenType = frontend.NewTokenType(BeginTokenTypeCode, "Begin")

var PlusTokenType = frontend.NewTokenType(PlusTokenTypeCode, "Plus")
var MinusTokenType = frontend.NewTokenType(MinusTokenTypeCode, "Minus")
var AsteriskTokenType = frontend.NewTokenType(AsteriskTokenTypeCode, "Asterisk")
var SlashTokenType = frontend.NewTokenType(SlashTokenTypeCode, "Slash")
var CommaTokenType = frontend.NewTokenType(CommaTokenTypeCode, "Comma")
var SemiColonTokenType = frontend.NewTokenType(SemiColonTokenTypeCode, "SemiColon")
var SingleQuoteTokenType = frontend.NewTokenType(SingleQuoteTokenTypeCode, "SingleQuote")
var EqualsTokenType = frontend.NewTokenType(EqualsTokenTypeCode, "Equals")
var OpenParenthesisTokenType = frontend.NewTokenType(OpenParenthesisTokenTypeCode, "OpenParenthesis")
var CloseParenthesisTokenType = frontend.NewTokenType(CloseParenthesisTokenTypeCode, "CloseParenthesis")
var OpenHardBraceTokenType = frontend.NewTokenType(OpenHardBraceTokenTypeCode, "OpenHardBrace")
var CloseHardBraceTokenType = frontend.NewTokenType(CloseHardBraceTokenTypeCode, "CloseHardBrace")
var OpenCurlyBraceTokenType = frontend.NewTokenType(OpenCurlyBraceTokenTypeCode, "OpenCurlyBrace")
var CloseCurlyBraceTokenType = frontend.NewTokenType(CloseCurlyBraceTokenTypeCode, "CloseCurlyBrace")
var CarotTokenType = frontend.NewTokenType(CarotTokenTypeCode, "Carot")
var ColonEqualsTokenType = frontend.NewTokenType(ColonEqualsTokenTypeCode, "ColonEquals")
var ColonTokenType = frontend.NewTokenType(ColonTokenTypeCode, "Colon")
var LessThanOrEqualsTokenType = frontend.NewTokenType(LessThanOrEqualsTokenTypeCode, "LessThanOrEquals")
var LessThanOrLessThanTokenType = frontend.NewTokenType(LessThanOrLessThanTokenTypeCode, "LessThanOrLessThan")
var LessThanTokenType = frontend.NewTokenType(LessThanTokenTypeCode, "LessThan")
var GreaterThanOrEqualsTokenType = frontend.NewTokenType(GreaterThanOrEqualsTokenTypeCode, "GreaterThanOrEquals")
var GreaterThanTokenType = frontend.NewTokenType(GreaterThanTokenTypeCode, "GreaterThan")
var DotTokenType = frontend.NewTokenType(DotTokenTypeCode, "Dot")
var DotDotTokenType = frontend.NewTokenType(DotDotTokenTypeCode, "DotDot")

var IdentifierTokenType = frontend.NewTokenType(IdentifierTokenTypeCode, "Identifier")
var IntegerTokenType = frontend.NewTokenType(IntegerTokenTypeCode, "Integer")
var RealTokenType = frontend.NewTokenType(RealTokenTypeCode, "Real")
var StringTokenType = frontend.NewTokenType(StringTokenTypeCode, "String")
var ErrorTokenType = frontend.NewTokenType(ErrorTokenTypeCode, "Error")

var reservedTokenTypes = map[string]frontend.TokenType{
	"AND":   AndTokenType,
	"BEGIN": BeginTokenType,
}

var specialSymbolTokenTypes = map[string]frontend.TokenType{
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
