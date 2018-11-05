// Code generated from Expr.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Expr

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseExprListener is a complete listener for a parse tree produced by ExprParser.
type BaseExprListener struct{}

var _ ExprListener = &BaseExprListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseExprListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseExprListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseExprListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseExprListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrintExpr is called when production printExpr is entered.
func (s *BaseExprListener) EnterPrintExpr(ctx *PrintExprContext) {}

// ExitPrintExpr is called when production printExpr is exited.
func (s *BaseExprListener) ExitPrintExpr(ctx *PrintExprContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseExprListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseExprListener) ExitAssign(ctx *AssignContext) {}

// EnterBlank is called when production blank is entered.
func (s *BaseExprListener) EnterBlank(ctx *BlankContext) {}

// ExitBlank is called when production blank is exited.
func (s *BaseExprListener) ExitBlank(ctx *BlankContext) {}

// EnterParens is called when production parens is entered.
func (s *BaseExprListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production parens is exited.
func (s *BaseExprListener) ExitParens(ctx *ParensContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseExprListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseExprListener) ExitAddSub(ctx *AddSubContext) {}

// EnterId is called when production id is entered.
func (s *BaseExprListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseExprListener) ExitId(ctx *IdContext) {}

// EnterInt is called when production int is entered.
func (s *BaseExprListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseExprListener) ExitInt(ctx *IntContext) {}

// EnterMulDiv is called when production mulDiv is entered.
func (s *BaseExprListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production mulDiv is exited.
func (s *BaseExprListener) ExitMulDiv(ctx *MulDivContext) {}
