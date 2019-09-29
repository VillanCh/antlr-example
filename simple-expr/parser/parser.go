package parser

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"se/parser/gen"
	"strconv"
)

func ExecSe(raw string) (int, error) {
	return 0, nil
}

type Variable struct {
	IsTmp bool
	Type  string
	Value int
}

func NewTmpVariable(v int) *Variable {
	return &Variable{
		IsTmp: true,
		Type:  "int",
		Value: v,
	}
}

func NewVariable(v int) *Variable {
	return &Variable{
		IsTmp: false,
		Type:  "int",
		Value: v,
	}
}

func ParseToVariable(raw string) (*Variable, error) {
	i, err := strconv.ParseInt(raw, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse [%s] to int failed: %s", raw, err)
	}
	return NewVariable(int(i)), nil
}

type Opcode struct {
	Result Variable
}

type seParser struct {
	gen.SimpleExprVisitor

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
	return s.Execute(parser.Start())
}

func (s *seParser) Execute(tree antlr.ParseTree) (int, error) {
	ret, ok := s.Visit(tree).(int)
	if !ok {
		return 0, errors.New("failed to convert result to int")
	}
	return ret, nil
}

func (s *seParser) Visit(tree antlr.ParseTree) interface{} {
	switch node := tree.(type) {
	case *gen.StartContext:
		return s.VisitStart(node)
	case *gen.ExprContext:
		return s.VisitExpr(node)
	case *gen.PrimaryExprContext:
		return s.VisitPrimaryExpr(node)
	case *gen.UnaryExprContext:
		return s.VisitUnaryExpr(node)
	case *gen.AdditiveExprContext:
		return s.VisitAdditiveExpr(node)
	case *gen.MultiplicativeExprContext:
		return s.VisitMultiplicativeExpr(node)
	}

	panic(fmt.Errorf("unknown parsetree type: %+v: %+v (%v)", reflect.TypeOf(tree), tree, tree.GetText()))
}

func (s *seParser) VisitStart(n *gen.StartContext) interface{} {
	return s.Visit(n.Expr())
}

func (s *seParser) VisitExpr(n *gen.ExprContext) interface{} {
	return s.Visit(n.AdditiveExpr())
}

func (s *seParser) VisitAdditiveExpr(n *gen.AdditiveExprContext) interface{} {
	if m := n.MultiplicativeExpr(); m != nil {
		return s.VisitMultiplicativeExpr(m.(*gen.MultiplicativeExprContext))
	}

	if m := n.Plus(); m != nil {
		aExpr, mExpr := n.AdditiveExpr(), n.MultiplicativeExpr()
		value1 := s.VisitAdditiveExpr(aExpr.(*gen.AdditiveExprContext))
		value2 := s.VisitMultiplicativeExpr(mExpr.(*gen.MultiplicativeExprContext))
		return NewVariable(value1.(*Variable).Value + value2.(*Variable).Value)
	} else if m := n.Minus(); m != nil {
		aExpr, mExpr := n.AdditiveExpr(), n.MultiplicativeExpr()
		value1 := s.VisitAdditiveExpr(aExpr.(*gen.AdditiveExprContext))
		value2 := s.VisitMultiplicativeExpr(mExpr.(*gen.MultiplicativeExprContext))
		return NewVariable(value1.(*Variable).Value - value2.(*Variable).Value)
	}

	return errors.New("BUG for parse additiveExpr")
}

func (s *seParser) VisitMultiplicativeExpr(n *gen.MultiplicativeExprContext) interface{} {
	if m := n.PrimaryExpr(); m != nil {
		return s.VisitPrimaryExpr(m.(*gen.PrimaryExprContext))
	}

	if m := n.Multiple(); m != nil {
		aExpr, mExpr := n.PrimaryExpr(), n.MultiplicativeExpr()
		value1 := s.VisitPrimaryExpr(aExpr.(*gen.PrimaryExprContext))
		value2 := s.VisitMultiplicativeExpr(mExpr.(*gen.MultiplicativeExprContext))
		return NewVariable(value1.(*Variable).Value * value2.(*Variable).Value)
	} else if m := n.Divide(); m != nil {
		aExpr, mExpr := n.PrimaryExpr(), n.MultiplicativeExpr()
		value1 := s.VisitPrimaryExpr(aExpr.(*gen.PrimaryExprContext))
		value2 := s.VisitMultiplicativeExpr(mExpr.(*gen.MultiplicativeExprContext))
		return NewVariable(value1.(*Variable).Value / value2.(*Variable).Value)
	}

	return errors.New("BUG for parse additiveExpr")
}

func (s *seParser) VisitPrimaryExpr(n *gen.PrimaryExprContext) interface{} {
	if m := n.Number(); m != nil {
	}
	return nil
}
