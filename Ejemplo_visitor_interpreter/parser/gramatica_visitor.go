// Code generated from Gramatica.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by GramaticaParser.
type GramaticaVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GramaticaParser#stat.
	VisitStat(ctx *StatContext) interface{}

	// Visit a parse tree produced by GramaticaParser#expNum.
	VisitExpNum(ctx *ExpNumContext) interface{}

	// Visit a parse tree produced by GramaticaParser#expMulDiv.
	VisitExpMulDiv(ctx *ExpMulDivContext) interface{}

	// Visit a parse tree produced by GramaticaParser#expStr.
	VisitExpStr(ctx *ExpStrContext) interface{}

	// Visit a parse tree produced by GramaticaParser#expSumRes.
	VisitExpSumRes(ctx *ExpSumResContext) interface{}
}
