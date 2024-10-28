package environment

import (
	"fmt"

	"github.com/SwanHtetAungPhyo/interpreter/internal/lexer"
	"github.com/SwanHtetAungPhyo/interpreter/internal/parser"
)

type Interpreter struct{}

func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Interpret(input string) (int, error) {
	lexer := lexer.NewLexer(input)
	parser := parser.NewParser(lexer)
	ast, err := parser.Parse()
	if err != nil {
		return 0, err
	}
	return i.visit(ast)
}

func (i *Interpreter) visit(node parser.Node) (int, error) {
	switch n := node.(type) {
	case parser.IntNode:
		return n.Value, nil
	case parser.BinOpNode:
		left, err := i.visit(n.Left)
		if err != nil {
			return 0, err
		}
		right, err := i.visit(n.Right)
		if err != nil {
			return 0, err
		}
		switch n.Op {
		case lexer.TOKEN_PLUS:
			return left + right, nil
		case lexer.TOKEN_MINUS:
			return left - right, nil
		case lexer.TOKEN_MULTIPLY:
			return left * right, nil
		case lexer.TOKEN_DIVIDE:
			if right == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			return left / right, nil
		default:
			return 0, fmt.Errorf("unknown operator: %v", n.Op)
		}
	default:
		return 0, fmt.Errorf("unknown node type: %T", node)
	}
}
