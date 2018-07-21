package frontend

type TokenType interface{}

type Token struct {
	typ      TokenType
	text     string
	value    interface{}
	source   *Source
	lineNum  int
	position int
}

func NewToken(s *Source) *Token {
	t := &Token{
		source:   s,
		lineNum:  s.GetLineNumber(),
		position: s.GetPosition(),
	}

	// TODO: It doesn't seem like it's idiomatic to return an error from a
	// "NewFoo" method; but what should we do if we encounter one here?
	_ = t.extract()

	return t
}

func (t *Token) extract() error {
	c, err := t.currentChar()
	if err != nil {
		return err
	}
	t.text = string(c)
	t.value = nil

	t.nextChar()
	return nil
}

func (t *Token) currentChar() (rune, error) {
	return t.source.CurrentChar()
}

func (t *Token) nextChar() (rune, error) {
	return t.source.NextChar()
}

func (t *Token) peekChar() (rune, error) {
	return t.source.PeekChar()
}
