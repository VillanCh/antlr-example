// Code generated from SimpleExpr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gen // SimpleExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSimpleExprListener is a complete listener for a parse tree produced by SimpleExprParser.
type BaseSimpleExprListener struct{}

var _ SimpleExprListener = &BaseSimpleExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSimpleExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSimpleExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSimpleExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSimpleExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSimpleExprListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSimpleExprListener) ExitStart(ctx *StartContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseSimpleExprListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseSimpleExprListener) ExitExpr(ctx *ExprContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BaseSimpleExprListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BaseSimpleExprListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterUnaryExpr is called when production unaryExpr is entered.
func (s *BaseSimpleExprListener) EnterUnaryExpr(ctx *UnaryExprContext) {}

// ExitUnaryExpr is called when production unaryExpr is exited.
func (s *BaseSimpleExprListener) ExitUnaryExpr(ctx *UnaryExprContext) {}

// EnterAdditiveExpr is called when production additiveExpr is entered.
func (s *BaseSimpleExprListener) EnterAdditiveExpr(ctx *AdditiveExprContext) {}

// ExitAdditiveExpr is called when production additiveExpr is exited.
func (s *BaseSimpleExprListener) ExitAdditiveExpr(ctx *AdditiveExprContext) {}

// EnterMultiplicativeExpr is called when production multiplicativeExpr is entered.
func (s *BaseSimpleExprListener) EnterMultiplicativeExpr(ctx *MultiplicativeExprContext) {}

// ExitMultiplicativeExpr is called when production multiplicativeExpr is exited.
func (s *BaseSimpleExprListener) ExitMultiplicativeExpr(ctx *MultiplicativeExprContext) {}
