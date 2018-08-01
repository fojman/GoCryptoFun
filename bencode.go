package main

import "errors"

/*
	- i4e
	- 3XYZ
	- l<content>e
	- d<content>e
*/

type Decoder struct {
	stream []byte
	pos    int
}

func (self *Decoder) IsEos() bool {
	if self.pos >= len(self.stream) {
		return true
	}

	return false
}

func (self *Decoder) Next() (res interface{}, err error) {

	if self.IsEos() {
		return nil, errors.New("Cannot Next() on EOS")
	}

	switch c := self.stream[self.pos]; {
	case c == 'i':
		{

		}

	case c >= '0' && c <= '9':
		{

		}
	case c == 'l':
		{

		}
	case c == 'd':
		{

		}
	} //swith
}

func (self *Decoder) parseInt() (res interface{}, err error) {

}

// isDigit
func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (self *Decoder) parseString() (res interface{}, err error) {

	// <len>:<stringgggg.....>
	// 3:xyz
	strLen := 0
	value := 0
	for {
		d := self.stream[self.pos]
		if d == ':' {
			break
		}
		digit := (int(d) - int('0'))
		value := value*10 + digit

		// check len
		self.pos++
	}
	return string(self.stream[self.pos:value]), nil
}
