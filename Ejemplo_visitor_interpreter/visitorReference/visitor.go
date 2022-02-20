package visitorReference

import (
	"Ejemplo_visitor_interpreter/interpreter/expresion"
	"Ejemplo_visitor_interpreter/interpreter/tipos"
	"Ejemplo_visitor_interpreter/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Visitor struct {
	*parser.BaseGramaticaVisitor
}

func (v *Visitor) Visit(tree antlr.ParseTree) expresion.ExpresionAbstracta {
	switch val := tree.(type) {
	case *parser.StatContext:
		return v.VisitStat(val)
	case *parser.ExpMulDivContext:
		return v.VisitExpMulDiv(val)
	case *parser.ExpSumResContext:
		return v.VisitExpSumRes(val)
	case *parser.ExpNumContext:
		return v.VisitExpNum(val)
	case *parser.ExpStrContext:
		return v.VisitExpStr(val)
	default:
		panic("Unknown context")
	}
}

func (v *Visitor) VisitStat(ctx *parser.StatContext) expresion.ExpresionAbstracta {
	return v.Visit(ctx.Exp())
}

func (v *Visitor) VisitExpMulDiv(ctx *parser.ExpMulDivContext) expresion.ExpresionAbstracta {
	left := v.Visit(ctx.GetLeft())
	right := v.Visit(ctx.GetRight())

	var op int

	if ctx.GetOp().GetTokenType() == parser.GramaticaParserPOR {
		op = expresion.POR
	} else {
		op = expresion.DIV
	}

	exp := expresion.NewExpresionAritmetica(left, op, right)
	return &exp
}

func (v *Visitor) VisitExpSumRes(ctx *parser.ExpSumResContext) expresion.ExpresionAbstracta {
	left := v.Visit(ctx.GetLeft())
	right := v.Visit(ctx.GetRight())

	var op int
	if ctx.GetOp().GetTokenType() == parser.GramaticaParserMAS {
		op = expresion.MAS
	} else {
		op = expresion.MEN
	}

	exp := expresion.NewExpresionAritmetica(left, op, right)
	return &exp
}

func (v *Visitor) VisitExpNum(ctx *parser.ExpNumContext) expresion.ExpresionAbstracta {
	tipo := tipos.NewTipo(tipos.NUMBER)
	exp := expresion.ExpresionLiteral{Tipo: &tipo, Valor: ctx.NUM().GetText()}
	return &exp
}

func (v *Visitor) VisitExpStr(ctx *parser.ExpStrContext) expresion.ExpresionAbstracta {
	tipo := tipos.NewTipo(tipos.STRING)
	exp := expresion.ExpresionLiteral{Tipo: &tipo, Valor: ctx.STR().GetText()}
	return &exp
}
