package main

import (
	"learnAnltr/expr"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
)

type EvalVisitor struct {
	parser.ExprVisitor
	memory map[string]interface{}
}

func NewEvaVisitor()*EvalVisitor {
	visitor := new(EvalVisitor)
	visitor.memory = make(map[string]interface{})
	return visitor
}

/*
*func (prc *BaseParserRuleContext) Accept(visitor ParseTreeVisitor) interface{} {
	return visitor.VisitChildren(prc)
}
*/
//
func (this *EvalVisitor)Visit(node antlr.ParseTree) interface{} {
	res := node.Accept(this)
	return res
}

func (this *EvalVisitor)VisitPrintExpr(ctx *parser.PrintExprContext) interface{} {
	value := this.Visit(ctx.Expr())
	return value
}

func (this *EvalVisitor)VisitAssign(ctx *parser.AssignContext) interface{} {
	id := ctx.ID().GetText()
	value := this.Visit(ctx.Expr())
	this.memory[id] = value
	return value
}

func (this *EvalVisitor)VisitParens(ctx *parser.ParensContext) interface{} {
	return this.Visit(ctx.Expr())
}

func (this *EvalVisitor)VisitId(ctx *parser.IdContext) interface{} {
	id := ctx.ID().GetText()
	value, ok := this.memory[id]
	if ok {
		return value
	}
	return 0
}

func (this *EvalVisitor)VisitInt(ctx *parser.IntContext) interface{} {
	value, _ := strconv.Atoi( ctx.INT().GetText())
	return value
}

func (this *EvalVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	left := this.Visit(ctx.Expr(0)).(int)
	right := this.Visit(ctx.Expr(1)).(int)
	if ctx.GetOp().GetTokenType() == parser.ExprLexerADD {
		return left + right
	}
	return left - right
}

func (this *EvalVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	left := this.Visit(ctx.Expr(0)).(int)
	right := this.Visit(ctx.Expr(1)).(int)
	if ctx.GetOp().GetTokenType() == parser.ExprLexerMUL {
		return left * right
	}
	return left / right
}


func Calc(input string) int {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewExprLexer(inputStream)
	tokens := antlr.NewCommonTokenStream(lexer, 0)
	parse := parser.NewExprParser(tokens)

	tree := parse.Stat()
	//fmt.Println(tree.GetText())
	return tree.Accept(NewEvaVisitor()).(int)
}