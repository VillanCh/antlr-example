// Code generated from SimpleExpr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gen // SimpleExpr
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSimpleExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSimpleExprVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleExprVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleExprVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleExprVisitor) VisitUnaryExpr(ctx *UnaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleExprVisitor) VisitAdditiveExpr(ctx *AdditiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSimpleExprVisitor) VisitMultiplicativeExpr(ctx *MultiplicativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}
