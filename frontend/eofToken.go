package frontend

type EofToken struct{}

func NewEofToken(s *Source) *EofToken {
	return &EofToken{}
}

func (ef *EofToken) extract() error {
	return nil
}
