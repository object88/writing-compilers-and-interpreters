package frontend

// EOFToken represents the end of a file
type EOFToken struct{}

// NewEOFToken creates a new token for the end of the file
func NewEOFToken(s *Source) *EOFToken {
	return &EOFToken{}
}

func (ef *EOFToken) extract() error {
	return nil
}
