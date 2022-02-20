// Code generated from Gramatica.g4 by ANTLR 4.7.2. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 46, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3,
	6, 7, 6, 28, 10, 6, 12, 6, 14, 6, 31, 11, 6, 3, 6, 3, 6, 3, 7, 6, 7, 36,
	10, 7, 13, 7, 14, 7, 37, 3, 8, 6, 8, 41, 10, 8, 13, 8, 14, 8, 42, 3, 8,
	3, 8, 3, 29, 2, 9, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 3, 2, 4,
	3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 48, 2, 3, 3, 2, 2, 2, 2,
	5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2,
	13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 3, 17, 3, 2, 2, 2, 5, 19, 3, 2, 2, 2,
	7, 21, 3, 2, 2, 2, 9, 23, 3, 2, 2, 2, 11, 25, 3, 2, 2, 2, 13, 35, 3, 2,
	2, 2, 15, 40, 3, 2, 2, 2, 17, 18, 7, 45, 2, 2, 18, 4, 3, 2, 2, 2, 19, 20,
	7, 47, 2, 2, 20, 6, 3, 2, 2, 2, 21, 22, 7, 44, 2, 2, 22, 8, 3, 2, 2, 2,
	23, 24, 7, 49, 2, 2, 24, 10, 3, 2, 2, 2, 25, 29, 7, 36, 2, 2, 26, 28, 11,
	2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 31, 3, 2, 2, 2, 29, 30, 3, 2, 2, 2, 29,
	27, 3, 2, 2, 2, 30, 32, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 32, 33, 7, 36,
	2, 2, 33, 12, 3, 2, 2, 2, 34, 36, 9, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 37,
	3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 14, 3, 2, 2, 2,
	39, 41, 9, 3, 2, 2, 40, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 40, 3,
	2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45, 8, 8, 2, 2, 45,
	16, 3, 2, 2, 2, 6, 2, 29, 37, 42, 3, 8, 2, 2,
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
	"", "'+'", "'-'", "'*'", "'/'",
}

var lexerSymbolicNames = []string{
	"", "MAS", "MEN", "POR", "DIV", "STR", "NUM", "WHITESPACE",
}

var lexerRuleNames = []string{
	"MAS", "MEN", "POR", "DIV", "STR", "NUM", "WHITESPACE",
}

type GramaticaLexer struct {
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

func NewGramaticaLexer(input antlr.CharStream) *GramaticaLexer {

	l := new(GramaticaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Gramatica.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GramaticaLexer tokens.
const (
	GramaticaLexerMAS        = 1
	GramaticaLexerMEN        = 2
	GramaticaLexerPOR        = 3
	GramaticaLexerDIV        = 4
	GramaticaLexerSTR        = 5
	GramaticaLexerNUM        = 6
	GramaticaLexerWHITESPACE = 7
)
