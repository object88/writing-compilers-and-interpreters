package frontend

// Token defines the interface for a token in the source code
type Token interface {
	GetLineNumber() int
	GetPosition() int
	GetText() string
	GetType() TokenType
	GetValue() interface{}
	// extract() error
}

// BaseToken contains essential token information; line number, position, and
// type
type BaseToken struct {
	typ      TokenType
	text     string
	value    interface{}
	source   *Source
	lineNum  int
	position int
}

// NewBaseToken creates a new instance of a BaseToken.  This should be used by
// other Tokens.
func NewBaseToken(s *Source) BaseToken {
	t := BaseToken{
		source:   s,
		lineNum:  s.GetLineNumber(),
		position: s.GetPosition(),
	}

	// // TODO: It doesn't seem like it's idiomatic to return an error from a
	// // "NewFoo" method; but what should we do if we encounter one here?
	// _ = t.extract()

	return t
}

// GetLineNumber returns the line number for this token
func (t *BaseToken) GetLineNumber() int {
	return t.lineNum
}

// GetPosition returns the position in the token's line
func (t *BaseToken) GetPosition() int {
	return t.position
}

// GetText returns the token text
func (t *BaseToken) GetText() string {
	return t.text
}

// GetType retuns the token type (EOFTokenType, etc.)
func (t *BaseToken) GetType() TokenType {
	return t.typ
}

// GetValue returns the token value
func (t *BaseToken) GetValue() interface{} {
	return t.value
}

// func (t *BaseToken) extract() error {
// 	c, err := t.CurrentChar()
// 	if err != nil {
// 		return err
// 	}
// 	t.text = string(c)
// 	t.value = nil

// 	_, err = t.NextChar()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (t *BaseToken) AssignTypeAndText(typ TokenType, text string) {
	t.text = text
	t.typ = typ
}

func (t *BaseToken) AssignValue(val interface{}) {
	t.value = val
}

func (t *BaseToken) CurrentChar() (rune, error) {
	return t.source.CurrentChar()
}

func (t *BaseToken) NextChar() (rune, error) {
	return t.source.NextChar()
}

func (t *BaseToken) PeekChar() (rune, error) {
	return t.source.PeekChar()
}
