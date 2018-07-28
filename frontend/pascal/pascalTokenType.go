package pascal

import "github.com/object88/writing-compilers-and-interpreters/frontend"

const (
	// AND is a Pascal keyword token type
	AND PascalTokenType = "AND"

	ERROR PascalTokenType = "ERROR"
)

// PascalTokenType is a Pascal-specific token type
type PascalTokenType frontend.TokenType
