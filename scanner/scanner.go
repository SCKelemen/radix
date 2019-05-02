package scanner

import (
	// "unicode/utf8"
)


func New(input string) *Scanner {
	return &Scanner{input: input, readHead: 0, previewHead: 0}
}

type Scanner struct {
	input string 
	readHead int 
	previewHead int 
	ch rune 
}

func (s *Scanner) scan() {

	for isLetter(s.ch) || isDigit(s.ch) { 
		s.readChar() // advance through letters or digits
	}
}


func (l *Lexer) readNumber() string {
	position := l.position
	for util.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (s *Scanner) readChar() {

	if s.previewHead >= len(s.input) {
		s.ch = 0
	} else {
		s.ch = rune(s.input[s.previewHead])
	}

	s.readHead = s.previewHead
	s.previewHead++
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' //|| ch >= utf8.RuneSelf && unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9' //|| ch >= utf8.RuneSelf && unicode.IsDigit(ch)
}

func digitVal(ch rune) int {
	switch {
	case '0' <= ch && ch <= '9':
		return int(ch - '0')
	case 'a' <= ch && ch <= 'j':
		return int(ch - 'a' + 10)
	case 'A' <= ch && ch <= 'J':
		return int(ch - 'A' + 10)
	case 'k' <= ch && ch <= 't':
		return int(ch - 'k' + 20)
	case 'K' <= ch && ch <= 'T':
		return int(ch - 'K' + 20)
	case 'u' <= ch && ch <= 'z':
		return int(ch - 'u' + 30)
	case 'U' <= ch && ch <= 'Z':
		return int(ch - 'U' + 30)
	}
	return 36 // larger than any legal digit val
}

// 0 A K U
// 1 B L V
// 2 C M W
// 3 D N X
// 4 E O Y
// 5 F P Z
// 6 G Q
// 7 H R
// 8 I S
// 9 J T

type Token int 

const (
	INVALID Token = iota 
	Number  
)

// numbers 

// decimal numbers
// 8000
// 8_000
// 10r8_000  

// hexadecimal numbers
// 16r1F40
// 16r1F_40
// 16r_1F_40

// binary numbers 
// 2r1111101000000
// 1_1111_0100_0000
// 2r_1_1111_0100_0000

// octal 
// 8r17500
// 8r17_500
// 8r_17_500

type Number struct {
	base int 
	exponent int 
}

type Floating struct {
	sign bool
	mantissa int 
	exponent int 
}