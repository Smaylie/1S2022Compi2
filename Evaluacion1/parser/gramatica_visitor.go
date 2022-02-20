// Code generated from Gramatica.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GramaticaParser.
type GramaticaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GramaticaParser#start.
	VisitStart(ctx *StartContext) interface{}
}
