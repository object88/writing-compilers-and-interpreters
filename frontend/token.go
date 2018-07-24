package frontend

type TokenType interface{}

type Token interface {
	extract() error
}

type BaseToken struct {
	typ      TokenType
	text     string
	value    interface{}
	source   *Source
	lineNum  int
	position int
}

func NewBaseToken(s *Source) *BaseToken {
	t := &BaseToken{
		source:   s,
		lineNum:  s.GetLineNumber(),
		position: s.GetPosition(),
	}

	// TODO: It doesn't seem like it's idiomatic to return an error from a
	// "NewFoo" method; but what should we do if we encounter one here?
	_ = t.extract()

	return t
}

func (t *BaseToken) extract() error {
	c, err := t.currentChar()
	if err != nil {
		return err
	}
	t.text = string(c)
	t.value = nil

	_, err = t.nextChar()
	if err != nil {
		return err
	}
	return nil
}

func (t *BaseToken) currentChar() (rune, error) {
	return t.source.CurrentChar()
}

func (t *BaseToken) nextChar() (rune, error) {
	return t.source.NextChar()
}

func (t *BaseToken) peekChar() (rune, error) {
	return t.source.PeekChar()
}
