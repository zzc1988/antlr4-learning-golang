// Code generated from Expr.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseExprVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseExprVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitBlank(ctx *BlankContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseExprVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}
