// Code generated from SimpleExpr.g4 by ANTLR 4.7.2. DO NOT EDIT.

package gen

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 80, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4,
	13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12,
	59, 10, 12, 5, 12, 61, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 5, 14, 67, 10,
	14, 3, 15, 6, 15, 70, 10, 15, 13, 15, 14, 15, 71, 3, 16, 6, 16, 75, 10,
	16, 13, 16, 14, 16, 76, 3, 16, 3, 16, 2, 2, 17, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 2, 23, 2, 25, 2, 27, 2, 29, 2,
	31, 12, 3, 2, 4, 3, 2, 51, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 79, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 3, 33, 3, 2, 2, 2, 5, 35, 3, 2, 2,
	2, 7, 37, 3, 2, 2, 2, 9, 39, 3, 2, 2, 2, 11, 41, 3, 2, 2, 2, 13, 43, 3,
	2, 2, 2, 15, 45, 3, 2, 2, 2, 17, 48, 3, 2, 2, 2, 19, 51, 3, 2, 2, 2, 21,
	53, 3, 2, 2, 2, 23, 60, 3, 2, 2, 2, 25, 62, 3, 2, 2, 2, 27, 66, 3, 2, 2,
	2, 29, 69, 3, 2, 2, 2, 31, 74, 3, 2, 2, 2, 33, 34, 7, 45, 2, 2, 34, 4,
	3, 2, 2, 2, 35, 36, 7, 47, 2, 2, 36, 6, 3, 2, 2, 2, 37, 38, 7, 44, 2, 2,
	38, 8, 3, 2, 2, 2, 39, 40, 7, 49, 2, 2, 40, 10, 3, 2, 2, 2, 41, 42, 7,
	42, 2, 2, 42, 12, 3, 2, 2, 2, 43, 44, 7, 43, 2, 2, 44, 14, 3, 2, 2, 2,
	45, 46, 7, 45, 2, 2, 46, 47, 7, 45, 2, 2, 47, 16, 3, 2, 2, 2, 48, 49, 7,
	47, 2, 2, 49, 50, 7, 47, 2, 2, 50, 18, 3, 2, 2, 2, 51, 52, 5, 21, 11, 2,
	52, 20, 3, 2, 2, 2, 53, 54, 5, 23, 12, 2, 54, 22, 3, 2, 2, 2, 55, 61, 7,
	50, 2, 2, 56, 58, 5, 25, 13, 2, 57, 59, 5, 29, 15, 2, 58, 57, 3, 2, 2,
	2, 58, 59, 3, 2, 2, 2, 59, 61, 3, 2, 2, 2, 60, 55, 3, 2, 2, 2, 60, 56,
	3, 2, 2, 2, 61, 24, 3, 2, 2, 2, 62, 63, 9, 2, 2, 2, 63, 26, 3, 2, 2, 2,
	64, 67, 7, 50, 2, 2, 65, 67, 5, 25, 13, 2, 66, 64, 3, 2, 2, 2, 66, 65,
	3, 2, 2, 2, 67, 28, 3, 2, 2, 2, 68, 70, 5, 27, 14, 2, 69, 68, 3, 2, 2,
	2, 70, 71, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 30,
	3, 2, 2, 2, 73, 75, 9, 3, 2, 2, 74, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2,
	76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 79, 8,
	16, 2, 2, 79, 32, 3, 2, 2, 2, 8, 2, 58, 60, 66, 71, 76, 3, 8, 2, 2,
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
	"", "'+'", "'-'", "'*'", "'/'", "'('", "')'", "'++'", "'--'",
}

var lexerSymbolicNames = []string{
	"", "Plus", "Minus", "Multiple", "Divide", "LeftParen", "RightParen", "Inc",
	"Dec", "Number", "WhiteSpace",
}

var lexerRuleNames = []string{
	"Plus", "Minus", "Multiple", "Divide", "LeftParen", "RightParen", "Inc",
	"Dec", "Number", "DecimalIntegerLiteral", "DecimalNumeral", "NonZeroDigit",
	"Digit", "Digits", "WhiteSpace",
}

type SimpleExprLexer struct {
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

func NewSimpleExprLexer(input antlr.CharStream) *SimpleExprLexer {

	l := new(SimpleExprLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "SimpleExpr.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SimpleExprLexer tokens.
const (
	SimpleExprLexerPlus       = 1
	SimpleExprLexerMinus      = 2
	SimpleExprLexerMultiple   = 3
	SimpleExprLexerDivide     = 4
	SimpleExprLexerLeftParen  = 5
	SimpleExprLexerRightParen = 6
	SimpleExprLexerInc        = 7
	SimpleExprLexerDec        = 8
	SimpleExprLexerNumber     = 9
	SimpleExprLexerWhiteSpace = 10
)
