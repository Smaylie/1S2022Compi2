// Code generated from Gramatica.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGramaticaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGramaticaVisitor) VisitStat(ctx *StatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitExpNum(ctx *ExpNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitExpMulDiv(ctx *ExpMulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitExpStr(ctx *ExpStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGramaticaVisitor) VisitExpSumRes(ctx *ExpSumResContext) interface{} {
	return v.VisitChildren(ctx)
}
