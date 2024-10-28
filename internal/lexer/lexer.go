package lexer

import "unicode"

type TokenType int

const (
	TOKEN_INT TokenType = iota
	TOKEN_PLUS
	TOKEN_MINUS
	TOKEN_MULTIPLY
	TOKEN_DIVIDE
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_EOF
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input string
	pos   int
}

// 10 + 12 , [10,"", +, 12]
func NewLexer(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	if l.pos >= len(l.input) {
		return Token{Type: TOKEN_EOF}
	}

	switch l.input[l.pos] {
	case '+':
		l.pos++
		return Token{Type: TOKEN_PLUS, Value: "+"}
	case '-':
		l.pos++
		return Token{Type: TOKEN_MINUS, Value: "-"}
	case '*':
		l.pos++
		return Token{Type: TOKEN_MULTIPLY, Value: "*"}
	case '/':
		l.pos++
		return Token{Type: TOKEN_DIVIDE, Value: "/"}
	case '(':
		l.pos++
		return Token{Type: TOKEN_LPAREN, Value: "("}
	case ')':
		l.pos++
		return Token{Type: TOKEN_RPAREN, Value: ")"}
	default:
		if unicode.IsDigit(rune(l.input[l.pos])) {
			start := l.pos
			for l.pos < len(l.input) && unicode.IsDigit(rune(l.input[l.pos])) {
				l.pos++
			}
			return Token{Type: TOKEN_INT, Value: l.input[start:l.pos]}
		}
	}

	// Return EOF if the character is unknown
	l.pos++
	return Token{Type: TOKEN_EOF, Value: string(l.input[l.pos-1])}
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) && unicode.IsSpace(rune(l.input[l.pos])) {
		l.pos++
	}
}
