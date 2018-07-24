package frontend

import (
	"bufio"
	"io"

	"github.com/object88/writing-compilers-and-interpreters/message"
)

// Source consumes a buffered reader to produce runes
type Source struct {
	bufr *bufio.Reader
	r    io.Reader
	line []rune
	c    rune

	lineNum    int
	currentPos int

	mh *message.MessageHandler
}

// NewSource returns a new Source instance
func NewSource(reader io.Reader, messageHandler *message.MessageHandler) *Source {
	s := &Source{
		bufr:       bufio.NewReader(reader),
		r:          reader,
		lineNum:    0,
		currentPos: -2,
		mh:         messageHandler,
	}

	return s
}

func (s *Source) CurrentChar() (rune, error) {
	if s.currentPos == -2 {
		// TODO: not sure what to do with this err yet.
		_ = s.readLine()
		return s.NextChar()
	} else if s.line == nil {
		return EOF, nil
	} else if s.currentPos == -1 || s.currentPos == len(s.line) {
		return EOL, nil
	} else if s.currentPos > len(s.line) {
		err := s.readLine()
		if err == io.EOF {
			return EOF, nil
		} else if err != nil {
			return rune(0), err
		}
		return s.NextChar()
	}

	r := s.line[s.currentPos]
	return r, nil
}

func (s *Source) NextChar() (rune, error) {
	s.currentPos++
	return s.CurrentChar()
}

func (s *Source) PeekChar() (rune, error) {
	s.CurrentChar()
	if s.line == nil {
		return EOF, nil
	}

	nextPos := s.currentPos + 1
	if nextPos < len(s.line) {
		return s.line[nextPos], nil
	}

	return EOL, nil
}

func (s *Source) GetLineNumber() int {
	return s.lineNum
}

func (s *Source) GetPosition() int {
	return s.currentPos
}

// Close terminates the reader
func (s *Source) Close() error {
	if c, ok := s.r.(io.Closer); ok {
		c.Close()
	}

	return nil
}

func (s *Source) readLine() error {
	line, err := s.bufr.ReadString('\n')
	if err == io.EOF {
		s.currentPos = -1
		s.line = nil
		s.lineNum++
	} else if err != nil {
		return err
	} else {
		s.currentPos = -1
		s.line = []rune(line)
		s.lineNum++
		s.mh.SendMessage(NewSourceLine(s.lineNum, line))
	}

	return err
}

func (s *Source) GetMessageHandler() *message.MessageHandler {
	return s.mh
}
