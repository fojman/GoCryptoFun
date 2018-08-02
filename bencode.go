package main

import (
	"bufio"
	"errors"
)

/*
	- i4e
	- 3XYZ
	- l<content>e
	- d<content>e
*/

type Decoder struct {
	bufio.Reader
}

type Stream struct {
	pos int    // current position
	buf []byte // bencode data
}

// NewStream .ctor(benc string)
func NewStream(bs string) *Stream {
	return &Stream{0, []byte(bs)}
}

// IsEOS end of stream?
func (s *Stream) IsEOS() bool {
	if s.pos >= len(s.buf) {
		return true
	}

	return false
}

func (s *Stream) readByte() (byte, error) {
	if s.IsEOS() {
		return 0, errors.New("cannot read from empty stream")
	}

	b := s.buf[s.pos]
	s.pos++ // move to next byte
	return b, nil
}

func (s *Stream) read(inBuf []byte) error {

	for index := 0; index < len(inBuf); index++ {
		b, err := s.readByte()
		if err != nil {
			return err
		}
		inBuf[index] = b
	}
	return nil
}

func (s *Stream) readUntil(until byte) ([]byte, error) {
	if s.IsEOS() {
		return nil, errors.New("readUntil: end of stream")
	}

	start := s.pos
	var i int
	for {
		if s.IsEOS() {
			return nil, errors.New("readUntil: end of stream")
		}
		// d 2:xy 2:zx e
		c := s.buf[s.pos]
		if c == until {
			break
		}
		i++
		s.pos++

	}
	end := start + i
	return s.buf[start:end], nil
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func toInt(buf []byte) int {

	var inValue int
	for _, element := range buf {
		d := element - '0'
		inValue = inValue * 10
		inValue = inValue + int(d)
	}
	//_ = inValue
	return inValue
}

func (s *Stream) readInt() (interface{}, error) {
	digits, err := s.readUntil('e')
	if err != nil {
		return nil, err
	}
	// parse digit
	value := toInt(digits)

	// move 1 byte forward (pass over i...E)
	_, err = s.readByte()

	return value, nil
}

func (s *Stream) unread() error {
	if s.pos == 0 {
		return errors.New("cannot unread at the beginning")
	}

	s.pos--

	return nil
}

func (s *Stream) readString() (string, error) {
	// 5:abcyy

	digits, err := s.readUntil(':')
	if err != nil {
		return "", err
	}

	// skip ':'
	if _, err := s.readByte(); err != nil {
		return "", err
	}

	len := toInt(digits)
	strBuf := make([]byte, len)
	if err := s.read(strBuf); err != nil {
		return "", err
	}

	return string(strBuf), nil
}

func (s *Stream) readList() (interface{}, error) {

	var list []interface{}

	return list, nil
}

func (s *Stream) current() byte {
	return s.buf[s.pos]
}

func (s *Stream) readDictionary() (interface{}, error) {

	dict := make(map[string]interface{})

	c, err := s.readByte()
	if err != nil {
		return nil, err
	}

	if c != 'd' {
		return nil, errors.New("'d' -dictionary expected")
	}

	for {

		// d 3:key e
		key, err := s.readString()
		if err != nil {
			return nil, err
		}

		ch, err := s.readByte()
		if err != nil {
			return nil, err
		}

		item, err := s.parseNext(ch)
		if err != nil {
			return nil, err
		}
		dict[key] = item

		if s.current() == 'e' {
			break
		}
	}

	return dict, nil
}

func (s *Stream) parseNext(ch byte) (item interface{}, err error) {

	switch {
	case isDigit(ch):
		{
			if err := s.unread(); err != nil {
				return nil, err
			}
			return s.readString()
		}

	case ch == 'i':
		return s.readInt()
	case ch == 'l':
		return s.readList()
	case ch == 'd':
		return s.readDictionary()
	default:
		return nil, errors.New("ch")
	}
}

func decode(str string) (interface{}, error) {
	if len(str) <= 0 {
		return nil, errors.New("Empty Bencode string passed in")
	}

	b := []byte(str)
	if b[0] != 'd' {
		return nil, errors.New("first byte must be 'd'")
	}

	stream := NewStream(str)

	dict, err := stream.readDictionary()
	if err != nil {
		return nil, err
	}

	return dict, nil
}
