// Code generated from SimpleExpr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gen // SimpleExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SimpleExprParser.
type SimpleExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SimpleExprParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by SimpleExprParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by SimpleExprParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by SimpleExprParser#unaryExpr.
	VisitUnaryExpr(ctx *UnaryExprContext) interface{}

	// Visit a parse tree produced by SimpleExprParser#additiveExpr.
	VisitAdditiveExpr(ctx *AdditiveExprContext) interface{}

	// Visit a parse tree produced by SimpleExprParser#multiplicativeExpr.
	VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{}
}
