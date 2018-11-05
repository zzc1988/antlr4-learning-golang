
## 一、anlter4的安装
   [具体安装参考](https://github.com/antlr/antlr4/blob/master/doc/getting-started.md)
## 二、anlter4的具体使用请看官方文档
### 官方文档是使用java语言描述的，golang的例子网上真的太少了。
## 三、使用golang与antlr4实现词法解析器
###  anter4支持两种遍历*parse-tree*的方式*visitor*与*listener*， *listener*是基于事件触发的，只需要定义好相应的Entry()与Exit()函数，执行Walk()函数遍历的时候就会自行触发相应的函数
### *visitor*则需要自行进行对*parse-tree*遍历，但是控制性更强。
### *listener*的方式有点像回溯的方法来对各个结点进行遍历
```
//伪代码
void walk_node(context) {
    context.Entry()
    loop:
        walk_node(conntext.child())
    end loop;
    context.Exit()
}
```

**举一个四则运算的例子(visitor方式)**

#### 首先定义语法
antlr4使用的是LL(*)文法
```
grammar Expr;
options
{
    language = Go;
    output=AST;
}

stat:   expr EOF               #printExpr
    |   ID '=' expr EOF        #assign
    |   EOF                    #blank
    ;

expr:   expr op=('*' | '/') expr    #mulDiv
    |   expr op=('+' | '-') expr    #AddSub
    |   INT                         #int
    |   ID                          #id
    |   '(' expr ')'                #parens
    ;

MUL :   '*';
DIV :   '/';
ADD :   '+';
SUB :   '-';

ID  : [a-zA-Z]+;
INT : [0-9]+;
WS  : [ \r\n\t]+ -> skip;
```
使用 *#symbol* label对rule进行修饰可以生成唯一的visitor方法，通过实现具体的visitor的 *interface{}* 来遍历*parse-tree*

实现的visitor代码
```
package mian

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
```
以上代码其实就是对不同action实现的具体的遍历方法。

最后执行Test函数测试一下
```
package mian

import (
	"testing"
	"fmt"
)

func TestCalc(t *testing.T) {
	fmt.Printf("d = 2     : %d\n", Calc("d = 2"))
	fmt.Printf("3         : %d\n", Calc("3"))
	fmt.Printf("3 * 4     : %d\n", Calc("3 * 4"))
	fmt.Printf("b = 5 * 6 : %d\n", Calc("b = 5 * 6"))
	fmt.Printf("2 / 2     : %d\n", Calc("2 / 2"))
	fmt.Printf("c = 4 / 2 : %d\n", Calc("c = 4 / 2"))
	fmt.Printf("4 / 2 + 3 : %d\n", Calc("4 / 2 + 3"))
	fmt.Printf("3 + 4 / 2 : %d\n", Calc("3 + 4 / 2"))
	fmt.Printf("4 / 2 * 2 : %d\n", Calc("4 / 2 * 2"))
}
/*
=== RUN   TestCalc
d = 2     : 2
3         : 3
3 * 4     : 12
b = 5 * 6 : 30
2 / 2     : 1
c = 4 / 2 : 2s
4 / 2 + 3 : 5
3 + 4 / 2 : 5
4 / 2 * 2 : 4
--- PASS: TestCalc (0.00s)
PASS
*/
```

## 最后再放一个使用listener模式的 [链接](https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/)

