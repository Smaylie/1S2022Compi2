// Code generated from c:\Users\eily2\Documents\Ingenieria\Compi 2\1S2022Compi2\Ejemplo_visitor_interpreter\Gramatica.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 26, 4,
	2, 9, 2, 4, 3, 9, 3, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 5, 3, 13, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 21, 10, 3, 12, 3, 14, 3, 24,
	11, 3, 3, 3, 2, 3, 4, 4, 2, 4, 2, 4, 3, 2, 5, 6, 3, 2, 3, 4, 2, 26, 2,
	6, 3, 2, 2, 2, 4, 12, 3, 2, 2, 2, 6, 7, 5, 4, 3, 2, 7, 8, 7, 2, 2, 3, 8,
	3, 3, 2, 2, 2, 9, 10, 8, 3, 1, 2, 10, 13, 7, 8, 2, 2, 11, 13, 7, 7, 2,
	2, 12, 9, 3, 2, 2, 2, 12, 11, 3, 2, 2, 2, 13, 22, 3, 2, 2, 2, 14, 15, 12,
	4, 2, 2, 15, 16, 9, 2, 2, 2, 16, 21, 5, 4, 3, 5, 17, 18, 12, 3, 2, 2, 18,
	19, 9, 3, 2, 2, 19, 21, 5, 4, 3, 4, 20, 14, 3, 2, 2, 2, 20, 17, 3, 2, 2,
	2, 21, 24, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 22, 23, 3, 2, 2, 2, 23, 5, 3,
	2, 2, 2, 24, 22, 3, 2, 2, 2, 5, 12, 20, 22,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'",
}
var symbolicNames = []string{
	"", "MAS", "MEN", "POR", "DIV", "STR", "NUM", "WHITESPACE",
}

var ruleNames = []string{
	"start", "exp",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GramaticaParser struct {
	*antlr.BaseParser
}

func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

// GramaticaParser tokens.
const (
	GramaticaParserEOF        = antlr.TokenEOF
	GramaticaParserMAS        = 1
	GramaticaParserMEN        = 2
	GramaticaParserPOR        = 3
	GramaticaParserDIV        = 4
	GramaticaParserSTR        = 5
	GramaticaParserNUM        = 6
	GramaticaParserWHITESPACE = 7
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start = 0
	GramaticaParserRULE_exp   = 1
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) CopyFrom(ctx *StartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StatContext struct {
	*StartContext
}

func NewStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StatContext {
	var p = new(StatContext)

	p.StartContext = NewEmptyStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StartContext))

	return p
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *StatContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewStatContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.exp(0)
	}
	{
		p.SetState(5)
		p.Match(GramaticaParserEOF)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) CopyFrom(ctx *ExpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpNumContext struct {
	*ExpContext
}

func NewExpNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpNumContext {
	var p = new(ExpNumContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *ExpNumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpNumContext) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

type ExpMulDivContext struct {
	*ExpContext
	left  IExpContext
	op    antlr.Token
	right IExpContext
}

func NewExpMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpMulDivContext {
	var p = new(ExpMulDivContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *ExpMulDivContext) GetOp() antlr.Token { return s.op }

func (s *ExpMulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpMulDivContext) GetLeft() IExpContext { return s.left }

func (s *ExpMulDivContext) GetRight() IExpContext { return s.right }

func (s *ExpMulDivContext) SetLeft(v IExpContext) { s.left = v }

func (s *ExpMulDivContext) SetRight(v IExpContext) { s.right = v }

func (s *ExpMulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpMulDivContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpMulDivContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpMulDivContext) POR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserPOR, 0)
}

func (s *ExpMulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(GramaticaParserDIV, 0)
}

type ExpStrContext struct {
	*ExpContext
}

func NewExpStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpStrContext {
	var p = new(ExpStrContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *ExpStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpStrContext) STR() antlr.TerminalNode {
	return s.GetToken(GramaticaParserSTR, 0)
}

type ExpSumResContext struct {
	*ExpContext
	left  IExpContext
	op    antlr.Token
	right IExpContext
}

func NewExpSumResContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpSumResContext {
	var p = new(ExpSumResContext)

	p.ExpContext = NewEmptyExpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpContext))

	return p
}

func (s *ExpSumResContext) GetOp() antlr.Token { return s.op }

func (s *ExpSumResContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpSumResContext) GetLeft() IExpContext { return s.left }

func (s *ExpSumResContext) GetRight() IExpContext { return s.right }

func (s *ExpSumResContext) SetLeft(v IExpContext) { s.left = v }

func (s *ExpSumResContext) SetRight(v IExpContext) { s.right = v }

func (s *ExpSumResContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpSumResContext) AllExp() []IExpContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpContext)(nil)).Elem())
	var tst = make([]IExpContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpContext)
		}
	}

	return tst
}

func (s *ExpSumResContext) Exp(i int) IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *ExpSumResContext) MAS() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMAS, 0)
}

func (s *ExpSumResContext) MEN() antlr.TerminalNode {
	return s.GetToken(GramaticaParserMEN, 0)
}

func (p *GramaticaParser) Exp() (localctx IExpContext) {
	return p.exp(0)
}

func (p *GramaticaParser) exp(_p int) (localctx IExpContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, GramaticaParserRULE_exp, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(10)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserNUM:
		localctx = NewExpNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(8)
			p.Match(GramaticaParserNUM)
		}

	case GramaticaParserSTR:
		localctx = NewExpStrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(9)
			p.Match(GramaticaParserSTR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(20)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(18)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpMulDivContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*ExpMulDivContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp)
				p.SetState(12)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(13)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpMulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserPOR || _la == GramaticaParserDIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpMulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(14)

					var _x = p.exp(3)

					localctx.(*ExpMulDivContext).right = _x
				}

			case 2:
				localctx = NewExpSumResContext(p, NewExpContext(p, _parentctx, _parentState))
				localctx.(*ExpSumResContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_exp)
				p.SetState(15)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(16)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpSumResContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserMAS || _la == GramaticaParserMEN) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpSumResContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(17)

					var _x = p.exp(2)

					localctx.(*ExpSumResContext).right = _x
				}

			}

		}
		p.SetState(22)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpContext = nil
		if localctx != nil {
			t = localctx.(*ExpContext)
		}
		return p.Exp_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Exp_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
