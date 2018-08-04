package frontend

// TokenTypeCode is a numeric indicator for a token type
type TokenTypeCode int

// TokenType is a TokenTypeCode and a matching name
type TokenType struct {
	Code TokenTypeCode
	Name string
}

const (
	EOFTokenTypeCode TokenTypeCode = iota
)

// EOFTokenType indicates an EOFToken
var EOFTokenType = NewTokenType(EOFTokenTypeCode, "EOF")

// NewTokenType returns a new instance of a TokenType
func NewTokenType(code TokenTypeCode, name string) TokenType {
	return TokenType{
		Code: code,
		Name: name,
	}
}

func (tt TokenType) String() string {
	return tt.Name
}
