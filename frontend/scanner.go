package frontend

type Scanner interface {
	CurrentToken() (Token, error)
	NextToken() (Token, error)
}

type BaseScanner struct {
	source *Source
}

func NewBaseScanner(s *Source) *BaseScanner {
	return &BaseScanner{
		source: s,
	}
}

func (bs *BaseScanner) CurrentChar() (rune, error) {
	return bs.source.CurrentChar()
}

func (bs *BaseScanner) NextChar() (rune, error) {
	return bs.source.NextChar()
}
