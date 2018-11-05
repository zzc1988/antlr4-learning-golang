// Code generated from Expr.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ExprParser.
type ExprVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ExprParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by ExprParser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by ExprParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by ExprParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by ExprParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by ExprParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by ExprParser#int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by ExprParser#mulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}
}
