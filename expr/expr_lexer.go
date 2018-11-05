// Code generated from Expr.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 54, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	6, 9, 39, 10, 9, 13, 9, 14, 9, 40, 3, 10, 6, 10, 44, 10, 10, 13, 10, 14,
	10, 45, 3, 11, 6, 11, 49, 10, 11, 13, 11, 14, 11, 50, 3, 11, 3, 11, 2,
	2, 12, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21,
	12, 3, 2, 5, 4, 2, 67, 92, 99, 124, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15,
	34, 34, 2, 56, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 3, 23, 3, 2, 2,
	2, 5, 25, 3, 2, 2, 2, 7, 27, 3, 2, 2, 2, 9, 29, 3, 2, 2, 2, 11, 31, 3,
	2, 2, 2, 13, 33, 3, 2, 2, 2, 15, 35, 3, 2, 2, 2, 17, 38, 3, 2, 2, 2, 19,
	43, 3, 2, 2, 2, 21, 48, 3, 2, 2, 2, 23, 24, 7, 63, 2, 2, 24, 4, 3, 2, 2,
	2, 25, 26, 7, 42, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 43, 2, 2, 28, 8,
	3, 2, 2, 2, 29, 30, 7, 44, 2, 2, 30, 10, 3, 2, 2, 2, 31, 32, 7, 49, 2,
	2, 32, 12, 3, 2, 2, 2, 33, 34, 7, 45, 2, 2, 34, 14, 3, 2, 2, 2, 35, 36,
	7, 47, 2, 2, 36, 16, 3, 2, 2, 2, 37, 39, 9, 2, 2, 2, 38, 37, 3, 2, 2, 2,
	39, 40, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 18, 3,
	2, 2, 2, 42, 44, 9, 3, 2, 2, 43, 42, 3, 2, 2, 2, 44, 45, 3, 2, 2, 2, 45,
	43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 20, 3, 2, 2, 2, 47, 49, 9, 4, 2,
	2, 48, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 48, 3, 2, 2, 2, 50, 51,
	3, 2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 53, 8, 11, 2, 2, 53, 22, 3, 2, 2, 2,
	6, 2, 40, 45, 50, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'='", "'('", "')'", "'*'", "'/'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "MUL", "DIV", "ADD", "SUB", "ID", "INT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "MUL", "DIV", "ADD", "SUB", "ID", "INT", "WS",
}

type ExprLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewExprLexer(input antlr.CharStream) *ExprLexer {

	l := new(ExprLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Expr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// ExprLexer tokens.
const (
	ExprLexerT__0 = 1
	ExprLexerT__1 = 2
	ExprLexerT__2 = 3
	ExprLexerMUL  = 4
	ExprLexerDIV  = 5
	ExprLexerADD  = 6
	ExprLexerSUB  = 7
	ExprLexerID   = 8
	ExprLexerINT  = 9
	ExprLexerWS   = 10
)