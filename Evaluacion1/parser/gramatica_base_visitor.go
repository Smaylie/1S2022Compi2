// Code generated from Gramatica.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseGramaticaVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGramaticaVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}
