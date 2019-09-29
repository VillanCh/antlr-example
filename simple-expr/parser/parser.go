package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"se/parser/gen"
)

func ExecSe(raw string) (int, error) {
	return 0, nil
}

type operator int

const (
	add operator = iota
	sub
	mul
	div
)

type opcode struct {
	op1 varIf
	op2 varIf
	op  operator
}

type varIf interface {
	Value() int
}

type variable struct {
	from opcode
}

func (v *variable) Value() int {
	return 0
}

type seParser struct {
	gen.SimpleExprListener

	// 表达式解析结果
	currentResult int

	// 表达式解析失败
	err error

	// Calc Stack
	stacks []int
}

func (s *seParser) parse(expr string) (int, error) {
	stream := antlr.NewInputStream(expr)
	lexer := gen.NewSimpleExprLexer(stream)
	parser := gen.NewSimpleExprParser(antlr.NewCommonTokenStream(lexer, 0))
	_ = parser
	return 0, nil
}

func (s *seParser) ExitUnaryExpr(c *gen.UnaryExprContext) {

}
