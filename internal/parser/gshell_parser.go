// Code generated from GShell.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GShell

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 112,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 7, 4, 33, 10, 4, 12, 4, 14,
	4, 36, 11, 4, 3, 4, 3, 4, 7, 4, 40, 10, 4, 12, 4, 14, 4, 43, 11, 4, 3,
	5, 3, 5, 7, 5, 47, 10, 5, 12, 5, 14, 5, 50, 11, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 62, 10, 7, 3, 8, 3, 8, 3,
	9, 3, 9, 7, 9, 68, 10, 9, 12, 9, 14, 9, 71, 11, 9, 3, 9, 3, 9, 7, 9, 75,
	10, 9, 12, 9, 14, 9, 78, 11, 9, 3, 9, 3, 9, 7, 9, 82, 10, 9, 12, 9, 14,
	9, 85, 11, 9, 3, 9, 3, 9, 7, 9, 89, 10, 9, 12, 9, 14, 9, 92, 11, 9, 5,
	9, 94, 10, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3,
	13, 7, 13, 105, 10, 13, 12, 13, 14, 13, 108, 11, 13, 3, 13, 3, 13, 3, 13,
	2, 2, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 3, 3, 2, 8, 9,
	2, 111, 2, 26, 3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6, 34, 3, 2, 2, 2, 8, 44,
	3, 2, 2, 2, 10, 54, 3, 2, 2, 2, 12, 61, 3, 2, 2, 2, 14, 63, 3, 2, 2, 2,
	16, 93, 3, 2, 2, 2, 18, 95, 3, 2, 2, 2, 20, 97, 3, 2, 2, 2, 22, 99, 3,
	2, 2, 2, 24, 102, 3, 2, 2, 2, 26, 27, 5, 8, 5, 2, 27, 28, 7, 2, 2, 3, 28,
	3, 3, 2, 2, 2, 29, 30, 9, 2, 2, 2, 30, 5, 3, 2, 2, 2, 31, 33, 7, 9, 2,
	2, 32, 31, 3, 2, 2, 2, 33, 36, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 34, 35,
	3, 2, 2, 2, 35, 37, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 37, 41, 5, 10, 6, 2,
	38, 40, 7, 9, 2, 2, 39, 38, 3, 2, 2, 2, 40, 43, 3, 2, 2, 2, 41, 39, 3,
	2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 7, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 44,
	48, 7, 3, 2, 2, 45, 47, 5, 6, 4, 2, 46, 45, 3, 2, 2, 2, 47, 50, 3, 2, 2,
	2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 3, 2, 2, 2, 50, 48,
	3, 2, 2, 2, 51, 52, 7, 4, 2, 2, 52, 53, 7, 2, 2, 3, 53, 9, 3, 2, 2, 2,
	54, 55, 5, 12, 7, 2, 55, 56, 5, 4, 3, 2, 56, 11, 3, 2, 2, 2, 57, 62, 5,
	14, 8, 2, 58, 59, 5, 14, 8, 2, 59, 60, 5, 16, 9, 2, 60, 62, 3, 2, 2, 2,
	61, 57, 3, 2, 2, 2, 61, 58, 3, 2, 2, 2, 62, 13, 3, 2, 2, 2, 63, 64, 7,
	6, 2, 2, 64, 15, 3, 2, 2, 2, 65, 69, 5, 18, 10, 2, 66, 68, 5, 16, 9, 2,
	67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 3,
	2, 2, 2, 70, 94, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 76, 5, 20, 11, 2,
	73, 75, 5, 16, 9, 2, 74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3,
	2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 94, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 79,
	83, 5, 22, 12, 2, 80, 82, 5, 16, 9, 2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2,
	2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 94, 3, 2, 2, 2, 85, 83,
	3, 2, 2, 2, 86, 90, 5, 24, 13, 2, 87, 89, 5, 16, 9, 2, 88, 87, 3, 2, 2,
	2, 89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 94,
	3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 65, 3, 2, 2, 2, 93, 72, 3, 2, 2, 2,
	93, 79, 3, 2, 2, 2, 93, 86, 3, 2, 2, 2, 94, 17, 3, 2, 2, 2, 95, 96, 7,
	6, 2, 2, 96, 19, 3, 2, 2, 2, 97, 98, 7, 7, 2, 2, 98, 21, 3, 2, 2, 2, 99,
	100, 7, 5, 2, 2, 100, 101, 7, 6, 2, 2, 101, 23, 3, 2, 2, 2, 102, 106, 7,
	3, 2, 2, 103, 105, 5, 6, 4, 2, 104, 103, 3, 2, 2, 2, 105, 108, 3, 2, 2,
	2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108,
	106, 3, 2, 2, 2, 109, 110, 7, 4, 2, 2, 110, 25, 3, 2, 2, 2, 12, 34, 41,
	48, 61, 69, 76, 83, 90, 93, 106,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'$'",
}
var symbolicNames = []string{
	"", "", "", "", "IDENTIFIER", "NUMBER", "TERMINATOR", "NL", "WS",
}

var ruleNames = []string{
	"start", "terminator", "commandListItem", "script", "singleCommand", "commandLine",
	"commandName", "arguments", "namedArgument", "numericArgument", "variableArgument",
	"scriptArgument",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GShellParser struct {
	*antlr.BaseParser
}

func NewGShellParser(input antlr.TokenStream) *GShellParser {
	this := new(GShellParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GShell.g4"

	return this
}

// GShellParser tokens.
const (
	GShellParserEOF        = antlr.TokenEOF
	GShellParserT__0       = 1
	GShellParserT__1       = 2
	GShellParserT__2       = 3
	GShellParserIDENTIFIER = 4
	GShellParserNUMBER     = 5
	GShellParserTERMINATOR = 6
	GShellParserNL         = 7
	GShellParserWS         = 8
)

// GShellParser rules.
const (
	GShellParserRULE_start            = 0
	GShellParserRULE_terminator       = 1
	GShellParserRULE_commandListItem  = 2
	GShellParserRULE_script           = 3
	GShellParserRULE_singleCommand    = 4
	GShellParserRULE_commandLine      = 5
	GShellParserRULE_commandName      = 6
	GShellParserRULE_arguments        = 7
	GShellParserRULE_namedArgument    = 8
	GShellParserRULE_numericArgument  = 9
	GShellParserRULE_variableArgument = 10
	GShellParserRULE_scriptArgument   = 11
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
	p.RuleIndex = GShellParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Script() IScriptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScriptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScriptContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GShellParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *GShellParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GShellParserRULE_start)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Script()
	}
	{
		p.SetState(25)
		p.Match(GShellParserEOF)
	}

	return localctx
}

// ITerminatorContext is an interface to support dynamic dispatch.
type ITerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerminatorContext differentiates from other interfaces.
	IsTerminatorContext()
}

type TerminatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminatorContext() *TerminatorContext {
	var p = new(TerminatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_terminator
	return p
}

func (*TerminatorContext) IsTerminatorContext() {}

func NewTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminatorContext {
	var p = new(TerminatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_terminator

	return p
}

func (s *TerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminatorContext) TERMINATOR() antlr.TerminalNode {
	return s.GetToken(GShellParserTERMINATOR, 0)
}

func (s *TerminatorContext) NL() antlr.TerminalNode {
	return s.GetToken(GShellParserNL, 0)
}

func (s *TerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterTerminator(s)
	}
}

func (s *TerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitTerminator(s)
	}
}

func (p *GShellParser) Terminator() (localctx ITerminatorContext) {
	localctx = NewTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GShellParserRULE_terminator)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GShellParserTERMINATOR || _la == GShellParserNL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICommandListItemContext is an interface to support dynamic dispatch.
type ICommandListItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandListItemContext differentiates from other interfaces.
	IsCommandListItemContext()
}

type CommandListItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandListItemContext() *CommandListItemContext {
	var p = new(CommandListItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_commandListItem
	return p
}

func (*CommandListItemContext) IsCommandListItemContext() {}

func NewCommandListItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandListItemContext {
	var p = new(CommandListItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_commandListItem

	return p
}

func (s *CommandListItemContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandListItemContext) SingleCommand() ISingleCommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleCommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleCommandContext)
}

func (s *CommandListItemContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(GShellParserNL)
}

func (s *CommandListItemContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(GShellParserNL, i)
}

func (s *CommandListItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandListItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandListItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterCommandListItem(s)
	}
}

func (s *CommandListItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitCommandListItem(s)
	}
}

func (p *GShellParser) CommandListItem() (localctx ICommandListItemContext) {
	localctx = NewCommandListItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GShellParserRULE_commandListItem)
	var _la int

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GShellParserNL {
		{
			p.SetState(29)
			p.Match(GShellParserNL)
		}

		p.SetState(34)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.SingleCommand()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(36)
				p.Match(GShellParserNL)
			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(GShellParserEOF, 0)
}

func (s *ScriptContext) AllCommandListItem() []ICommandListItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandListItemContext)(nil)).Elem())
	var tst = make([]ICommandListItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandListItemContext)
		}
	}

	return tst
}

func (s *ScriptContext) CommandListItem(i int) ICommandListItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandListItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandListItemContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *GShellParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GShellParserRULE_script)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.Match(GShellParserT__0)
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GShellParserIDENTIFIER || _la == GShellParserNL {
		{
			p.SetState(43)
			p.CommandListItem()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.Match(GShellParserT__1)
	}
	{
		p.SetState(50)
		p.Match(GShellParserEOF)
	}

	return localctx
}

// ISingleCommandContext is an interface to support dynamic dispatch.
type ISingleCommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleCommandContext differentiates from other interfaces.
	IsSingleCommandContext()
}

type SingleCommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleCommandContext() *SingleCommandContext {
	var p = new(SingleCommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_singleCommand
	return p
}

func (*SingleCommandContext) IsSingleCommandContext() {}

func NewSingleCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleCommandContext {
	var p = new(SingleCommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_singleCommand

	return p
}

func (s *SingleCommandContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleCommandContext) CommandLine() ICommandLineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandLineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandLineContext)
}

func (s *SingleCommandContext) Terminator() ITerminatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerminatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *SingleCommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleCommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleCommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterSingleCommand(s)
	}
}

func (s *SingleCommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitSingleCommand(s)
	}
}

func (p *GShellParser) SingleCommand() (localctx ISingleCommandContext) {
	localctx = NewSingleCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GShellParserRULE_singleCommand)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.CommandLine()
	}
	{
		p.SetState(53)
		p.Terminator()
	}

	return localctx
}

// ICommandLineContext is an interface to support dynamic dispatch.
type ICommandLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandLineContext differentiates from other interfaces.
	IsCommandLineContext()
}

type CommandLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandLineContext() *CommandLineContext {
	var p = new(CommandLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_commandLine
	return p
}

func (*CommandLineContext) IsCommandLineContext() {}

func NewCommandLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandLineContext {
	var p = new(CommandLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_commandLine

	return p
}

func (s *CommandLineContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandLineContext) CommandName() ICommandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandNameContext)
}

func (s *CommandLineContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *CommandLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterCommandLine(s)
	}
}

func (s *CommandLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitCommandLine(s)
	}
}

func (p *GShellParser) CommandLine() (localctx ICommandLineContext) {
	localctx = NewCommandLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GShellParserRULE_commandLine)

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

	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.CommandName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.CommandName()
		}
		{
			p.SetState(57)
			p.Arguments()
		}

	}

	return localctx
}

// ICommandNameContext is an interface to support dynamic dispatch.
type ICommandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandNameContext differentiates from other interfaces.
	IsCommandNameContext()
}

type CommandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandNameContext() *CommandNameContext {
	var p = new(CommandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_commandName
	return p
}

func (*CommandNameContext) IsCommandNameContext() {}

func NewCommandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandNameContext {
	var p = new(CommandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_commandName

	return p
}

func (s *CommandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandNameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GShellParserIDENTIFIER, 0)
}

func (s *CommandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterCommandName(s)
	}
}

func (s *CommandNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitCommandName(s)
	}
}

func (p *GShellParser) CommandName() (localctx ICommandNameContext) {
	localctx = NewCommandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GShellParserRULE_commandName)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(61)
		p.Match(GShellParserIDENTIFIER)
	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) NamedArgument() INamedArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedArgumentContext)
}

func (s *ArgumentsContext) AllArguments() []IArgumentsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentsContext)(nil)).Elem())
	var tst = make([]IArgumentsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentsContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Arguments(i int) IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ArgumentsContext) NumericArgument() INumericArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericArgumentContext)
}

func (s *ArgumentsContext) VariableArgument() IVariableArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableArgumentContext)
}

func (s *ArgumentsContext) ScriptArgument() IScriptArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScriptArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScriptArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *GShellParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GShellParserRULE_arguments)

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

	var _alt int

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GShellParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(63)
			p.NamedArgument()
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(64)
					p.Arguments()
				}

			}
			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	case GShellParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.NumericArgument()
		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(71)
					p.Arguments()
				}

			}
			p.SetState(76)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}

	case GShellParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(77)
			p.VariableArgument()
		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(78)
					p.Arguments()
				}

			}
			p.SetState(83)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
		}

	case GShellParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(84)
			p.ScriptArgument()
		}
		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(85)
					p.Arguments()
				}

			}
			p.SetState(90)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INamedArgumentContext is an interface to support dynamic dispatch.
type INamedArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedArgumentContext differentiates from other interfaces.
	IsNamedArgumentContext()
}

type NamedArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedArgumentContext() *NamedArgumentContext {
	var p = new(NamedArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_namedArgument
	return p
}

func (*NamedArgumentContext) IsNamedArgumentContext() {}

func NewNamedArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedArgumentContext {
	var p = new(NamedArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_namedArgument

	return p
}

func (s *NamedArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedArgumentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GShellParserIDENTIFIER, 0)
}

func (s *NamedArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterNamedArgument(s)
	}
}

func (s *NamedArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitNamedArgument(s)
	}
}

func (p *GShellParser) NamedArgument() (localctx INamedArgumentContext) {
	localctx = NewNamedArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GShellParserRULE_namedArgument)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(93)
		p.Match(GShellParserIDENTIFIER)
	}

	return localctx
}

// INumericArgumentContext is an interface to support dynamic dispatch.
type INumericArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericArgumentContext differentiates from other interfaces.
	IsNumericArgumentContext()
}

type NumericArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericArgumentContext() *NumericArgumentContext {
	var p = new(NumericArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_numericArgument
	return p
}

func (*NumericArgumentContext) IsNumericArgumentContext() {}

func NewNumericArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericArgumentContext {
	var p = new(NumericArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_numericArgument

	return p
}

func (s *NumericArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericArgumentContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(GShellParserNUMBER, 0)
}

func (s *NumericArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterNumericArgument(s)
	}
}

func (s *NumericArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitNumericArgument(s)
	}
}

func (p *GShellParser) NumericArgument() (localctx INumericArgumentContext) {
	localctx = NewNumericArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GShellParserRULE_numericArgument)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(95)
		p.Match(GShellParserNUMBER)
	}

	return localctx
}

// IVariableArgumentContext is an interface to support dynamic dispatch.
type IVariableArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableArgumentContext differentiates from other interfaces.
	IsVariableArgumentContext()
}

type VariableArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableArgumentContext() *VariableArgumentContext {
	var p = new(VariableArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_variableArgument
	return p
}

func (*VariableArgumentContext) IsVariableArgumentContext() {}

func NewVariableArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableArgumentContext {
	var p = new(VariableArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_variableArgument

	return p
}

func (s *VariableArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableArgumentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(GShellParserIDENTIFIER, 0)
}

func (s *VariableArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterVariableArgument(s)
	}
}

func (s *VariableArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitVariableArgument(s)
	}
}

func (p *GShellParser) VariableArgument() (localctx IVariableArgumentContext) {
	localctx = NewVariableArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GShellParserRULE_variableArgument)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(GShellParserT__2)
	}
	{
		p.SetState(98)
		p.Match(GShellParserIDENTIFIER)
	}

	return localctx
}

// IScriptArgumentContext is an interface to support dynamic dispatch.
type IScriptArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScriptArgumentContext differentiates from other interfaces.
	IsScriptArgumentContext()
}

type ScriptArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptArgumentContext() *ScriptArgumentContext {
	var p = new(ScriptArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_scriptArgument
	return p
}

func (*ScriptArgumentContext) IsScriptArgumentContext() {}

func NewScriptArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptArgumentContext {
	var p = new(ScriptArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_scriptArgument

	return p
}

func (s *ScriptArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptArgumentContext) AllCommandListItem() []ICommandListItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandListItemContext)(nil)).Elem())
	var tst = make([]ICommandListItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandListItemContext)
		}
	}

	return tst
}

func (s *ScriptArgumentContext) CommandListItem(i int) ICommandListItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandListItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandListItemContext)
}

func (s *ScriptArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterScriptArgument(s)
	}
}

func (s *ScriptArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitScriptArgument(s)
	}
}

func (p *GShellParser) ScriptArgument() (localctx IScriptArgumentContext) {
	localctx = NewScriptArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GShellParserRULE_scriptArgument)
	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Match(GShellParserT__0)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == GShellParserIDENTIFIER || _la == GShellParserNL {
		{
			p.SetState(101)
			p.CommandListItem()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(107)
		p.Match(GShellParserT__1)
	}

	return localctx
}
