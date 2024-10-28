package parser

import (
	"fmt"
	"strconv"

	"github.com/SwanHtetAungPhyo/interpreter/internal/lexer"
)

type NodeType int

const (
	NODE_INT NodeType = iota
	NODE_BINOP
)

type Node interface {
	Type() NodeType
}

type IntNode struct {
	Value int
}

//	+
///		\
//1		2

func (n IntNode) Type() NodeType {
	return NODE_INT
}

type BinOpNode struct {
	Left  Node
	Op    lexer.TokenType
	Right Node
}

func (n BinOpNode) Type() NodeType {
	return NODE_BINOP
}

type Parser struct {
	lexer  *lexer.Lexer
	curTok lexer.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	p := &Parser{lexer: lexer}
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curTok = p.lexer.NextToken()
}

func (p *Parser) Parse() (Node, error) {
	return p.expr()
}

func (p *Parser) expr() (Node, error) {
	left, err := p.term()
	if err != nil {
		return nil, err
	}

	for p.curTok.Type == lexer.TOKEN_PLUS || p.curTok.Type == lexer.TOKEN_MINUS {
		op := p.curTok.Type
		p.nextToken()
		right, err := p.term()
		if err != nil {
			return nil, err
		}
		left = BinOpNode{Left: left, Op: op, Right: right}
	}

	return left, nil
}

func (p *Parser) term() (Node, error) {
	left, err := p.factor()
	if err != nil {
		return nil, err
	}

	for p.curTok.Type == lexer.TOKEN_MULTIPLY || p.curTok.Type == lexer.TOKEN_DIVIDE {
		op := p.curTok.Type
		p.nextToken()
		right, err := p.factor()
		if err != nil {
			return nil, err
		}
		left = BinOpNode{Left: left, Op: op, Right: right}
	}

	return left, nil
}

// ( 12 ) }
func (p *Parser) factor() (Node, error) {
	if p.curTok.Type == lexer.TOKEN_INT {
		value, err := strconv.Atoi(p.curTok.Value)
		if err != nil {
			return nil, err
		}
		p.nextToken()
		return IntNode{Value: value}, nil
	} else if p.curTok.Type == lexer.TOKEN_LPAREN {
		p.nextToken()
		node, err := p.expr()
		if err != nil {
			return nil, err
		}
		if p.curTok.Type != lexer.TOKEN_RPAREN {
			return nil, fmt.Errorf("expected ')', got %v", p.curTok)
		}
		p.nextToken()
		return node, nil
	}
	return nil, fmt.Errorf("unexpected token: %v", p.curTok)
}
