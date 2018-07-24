package frontend

type Scanner interface {
	CurrentToken() (Token, error)
	NextToken() (Token, error)
}

type BaseScanner struct {
	source       *Source
	currentToken Token
}

// EOF represents the end-of-file
var EOF = rune(0)

// EOL represents the end-of-line
var EOL = rune('\n')

func NewBaseScanner(s *Source) *BaseScanner {
	return &BaseScanner{
		source: s,
	}
}

// CurrentChar returns the current rune from the source
func (bs *BaseScanner) CurrentChar() (rune, error) {
	return bs.source.CurrentChar()
}

// NextChar retreives the next rune from the source
func (bs *BaseScanner) NextChar() (rune, error) {
	return bs.source.NextChar()
}

func (bs *BaseScanner) CurrentToken() (Token, error) {
	return bs.currentToken, nil
}
