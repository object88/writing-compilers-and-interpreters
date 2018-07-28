package frontend

// EOFToken represents the end of a file
type EOFToken struct{}

// NewEOFToken creates a new token for the end of the file
func NewEOFToken(s *Source) *EOFToken {
	return &EOFToken{}
}

// GetLineNumber is always -1 for the EOFToken
func (*EOFToken) GetLineNumber() int {
	return -1
}

// GetPosition is always -1 for the EOFToken
func (*EOFToken) GetPosition() int {
	return -1
}

// GetText returns the empty string for EOFToken
func (*EOFToken) GetText() string {
	return ""
}

// GetType is always EOFTokenType for EOFToken
func (*EOFToken) GetType() TokenType {
	return EOFTokenType
}

// GetValue returns nil for EOFToken
func (*EOFToken) GetValue() interface{} {
	return nil
}

func (*EOFToken) extract() error {
	return nil
}
