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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 12, 160, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 7, 4, 43, 10, 4, 12, 4, 14, 4, 
	46, 11, 4, 3, 4, 3, 4, 3, 4, 7, 4, 51, 10, 4, 12, 4, 14, 4, 54, 11, 4, 
	3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 7, 8, 64, 10, 8, 12, 8, 
	14, 8, 67, 11, 8, 3, 8, 3, 8, 3, 8, 7, 8, 72, 10, 8, 12, 8, 14, 8, 75, 
	11, 8, 3, 8, 3, 8, 3, 8, 7, 8, 80, 10, 8, 12, 8, 14, 8, 83, 11, 8, 3, 8, 
	3, 8, 7, 8, 87, 10, 8, 12, 8, 14, 8, 90, 11, 8, 3, 8, 3, 8, 3, 8, 5, 8, 
	95, 10, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 
	5, 11, 106, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 7, 13, 112, 10, 13, 12, 
	13, 14, 13, 115, 11, 13, 3, 13, 3, 13, 7, 13, 119, 10, 13, 12, 13, 14, 
	13, 122, 11, 13, 3, 13, 3, 13, 7, 13, 126, 10, 13, 12, 13, 14, 13, 129, 
	11, 13, 3, 13, 3, 13, 7, 13, 133, 10, 13, 12, 13, 14, 13, 136, 11, 13, 
	3, 13, 3, 13, 7, 13, 140, 10, 13, 12, 13, 14, 13, 143, 11, 13, 5, 13, 145, 
	10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 2, 2, 19, 2, 4, 6, 8, 10, 12, 14, 16, 
	18, 20, 22, 24, 26, 28, 30, 32, 34, 2, 3, 3, 2, 10, 11, 2, 160, 2, 36, 
	3, 2, 2, 2, 4, 39, 3, 2, 2, 2, 6, 44, 3, 2, 2, 2, 8, 55, 3, 2, 2, 2, 10, 
	57, 3, 2, 2, 2, 12, 59, 3, 2, 2, 2, 14, 94, 3, 2, 2, 2, 16, 96, 3, 2, 2, 
	2, 18, 99, 3, 2, 2, 2, 20, 105, 3, 2, 2, 2, 22, 107, 3, 2, 2, 2, 24, 144, 
	3, 2, 2, 2, 26, 146, 3, 2, 2, 2, 28, 148, 3, 2, 2, 2, 30, 150, 3, 2, 2, 
	2, 32, 153, 3, 2, 2, 2, 34, 155, 3, 2, 2, 2, 36, 37, 5, 16, 9, 2, 37, 38, 
	7, 2, 2, 3, 38, 3, 3, 2, 2, 2, 39, 40, 9, 2, 2, 2, 40, 5, 3, 2, 2, 2, 41, 
	43, 7, 11, 2, 2, 42, 41, 3, 2, 2, 2, 43, 46, 3, 2, 2, 2, 44, 42, 3, 2, 
	2, 2, 44, 45, 3, 2, 2, 2, 45, 47, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 47, 48, 
	5, 18, 10, 2, 48, 52, 5, 4, 3, 2, 49, 51, 7, 11, 2, 2, 50, 49, 3, 2, 2, 
	2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 7, 3, 
	2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 56, 7, 3, 2, 2, 56, 9, 3, 2, 2, 2, 57, 
	58, 7, 4, 2, 2, 58, 11, 3, 2, 2, 2, 59, 60, 5, 8, 5, 2, 60, 61, 5, 14, 
	8, 2, 61, 13, 3, 2, 2, 2, 62, 64, 7, 11, 2, 2, 63, 62, 3, 2, 2, 2, 64, 
	67, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 
	2, 67, 65, 3, 2, 2, 2, 68, 69, 5, 18, 10, 2, 69, 73, 5, 4, 3, 2, 70, 72, 
	7, 11, 2, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 
	73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 77, 5, 
	14, 8, 2, 77, 95, 3, 2, 2, 2, 78, 80, 7, 11, 2, 2, 79, 78, 3, 2, 2, 2, 
	80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 84, 3, 
	2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 88, 5, 18, 10, 2, 85, 87, 7, 11, 2, 2, 
	86, 85, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 
	2, 2, 2, 89, 91, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 92, 5, 10, 6, 2, 92, 
	95, 3, 2, 2, 2, 93, 95, 5, 10, 6, 2, 94, 65, 3, 2, 2, 2, 94, 81, 3, 2, 
	2, 2, 94, 93, 3, 2, 2, 2, 95, 15, 3, 2, 2, 2, 96, 97, 5, 12, 7, 2, 97, 
	98, 7, 2, 2, 3, 98, 17, 3, 2, 2, 2, 99, 100, 5, 20, 11, 2, 100, 19, 3, 
	2, 2, 2, 101, 106, 5, 22, 12, 2, 102, 103, 5, 22, 12, 2, 103, 104, 5, 24, 
	13, 2, 104, 106, 3, 2, 2, 2, 105, 101, 3, 2, 2, 2, 105, 102, 3, 2, 2, 2, 
	106, 21, 3, 2, 2, 2, 107, 108, 7, 8, 2, 2, 108, 23, 3, 2, 2, 2, 109, 113, 
	5, 26, 14, 2, 110, 112, 5, 24, 13, 2, 111, 110, 3, 2, 2, 2, 112, 115, 3, 
	2, 2, 2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 145, 3, 2, 2, 
	2, 115, 113, 3, 2, 2, 2, 116, 120, 5, 28, 15, 2, 117, 119, 5, 24, 13, 2, 
	118, 117, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 
	121, 3, 2, 2, 2, 121, 145, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 127, 
	5, 30, 16, 2, 124, 126, 5, 24, 13, 2, 125, 124, 3, 2, 2, 2, 126, 129, 3, 
	2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 145, 3, 2, 2, 
	2, 129, 127, 3, 2, 2, 2, 130, 134, 5, 32, 17, 2, 131, 133, 5, 24, 13, 2, 
	132, 131, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 
	135, 3, 2, 2, 2, 135, 145, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137, 141, 
	5, 34, 18, 2, 138, 140, 5, 24, 13, 2, 139, 138, 3, 2, 2, 2, 140, 143, 3, 
	2, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 145, 3, 2, 2, 
	2, 143, 141, 3, 2, 2, 2, 144, 109, 3, 2, 2, 2, 144, 116, 3, 2, 2, 2, 144, 
	123, 3, 2, 2, 2, 144, 130, 3, 2, 2, 2, 144, 137, 3, 2, 2, 2, 145, 25, 3, 
	2, 2, 2, 146, 147, 7, 8, 2, 2, 147, 27, 3, 2, 2, 2, 148, 149, 7, 9, 2, 
	2, 149, 29, 3, 2, 2, 2, 150, 151, 7, 5, 2, 2, 151, 152, 7, 8, 2, 2, 152, 
	31, 3, 2, 2, 2, 153, 154, 5, 12, 7, 2, 154, 33, 3, 2, 2, 2, 155, 156, 7, 
	6, 2, 2, 156, 157, 5, 24, 13, 2, 157, 158, 7, 7, 2, 2, 158, 35, 3, 2, 2, 
	2, 16, 44, 52, 65, 73, 81, 88, 94, 105, 113, 120, 127, 134, 141, 144,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'$'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "IDENTIFIER", "NUMBER", "TERMINATOR", "NL", "WS",
}

var ruleNames = []string{
	"start", "terminator", "commandListItem", "openBlock", "closeBlock", "commandBlock", 
	"commandBlockTail", "script", "singleCommand", "commandLine", "commandName", 
	"arguments", "namedArgument", "numericArgument", "variableArgument", "scriptArgument", 
	"listArgument",
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
	GShellParserEOF = antlr.TokenEOF
	GShellParserT__0 = 1
	GShellParserT__1 = 2
	GShellParserT__2 = 3
	GShellParserT__3 = 4
	GShellParserT__4 = 5
	GShellParserIDENTIFIER = 6
	GShellParserNUMBER = 7
	GShellParserTERMINATOR = 8
	GShellParserNL = 9
	GShellParserWS = 10
)

// GShellParser rules.
const (
	GShellParserRULE_start = 0
	GShellParserRULE_terminator = 1
	GShellParserRULE_commandListItem = 2
	GShellParserRULE_openBlock = 3
	GShellParserRULE_closeBlock = 4
	GShellParserRULE_commandBlock = 5
	GShellParserRULE_commandBlockTail = 6
	GShellParserRULE_script = 7
	GShellParserRULE_singleCommand = 8
	GShellParserRULE_commandLine = 9
	GShellParserRULE_commandName = 10
	GShellParserRULE_arguments = 11
	GShellParserRULE_namedArgument = 12
	GShellParserRULE_numericArgument = 13
	GShellParserRULE_variableArgument = 14
	GShellParserRULE_scriptArgument = 15
	GShellParserRULE_listArgument = 16
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
		p.SetState(34)
		p.Script()
	}
	{
		p.SetState(35)
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
		p.SetState(37)
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

func (s *CommandListItemContext) Terminator() ITerminatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerminatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == GShellParserNL {
		{
			p.SetState(39)
			p.Match(GShellParserNL)
		}


		p.SetState(44)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(45)
		p.SingleCommand()
	}
	{
		p.SetState(46)
		p.Terminator()
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == GShellParserNL {
		{
			p.SetState(47)
			p.Match(GShellParserNL)
		}


		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IOpenBlockContext is an interface to support dynamic dispatch.
type IOpenBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpenBlockContext differentiates from other interfaces.
	IsOpenBlockContext()
}

type OpenBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpenBlockContext() *OpenBlockContext {
	var p = new(OpenBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_openBlock
	return p
}

func (*OpenBlockContext) IsOpenBlockContext() {}

func NewOpenBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpenBlockContext {
	var p = new(OpenBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_openBlock

	return p
}

func (s *OpenBlockContext) GetParser() antlr.Parser { return s.parser }
func (s *OpenBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpenBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *OpenBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterOpenBlock(s)
	}
}

func (s *OpenBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitOpenBlock(s)
	}
}




func (p *GShellParser) OpenBlock() (localctx IOpenBlockContext) {
	localctx = NewOpenBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GShellParserRULE_openBlock)

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
		p.SetState(53)
		p.Match(GShellParserT__0)
	}



	return localctx
}


// ICloseBlockContext is an interface to support dynamic dispatch.
type ICloseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCloseBlockContext differentiates from other interfaces.
	IsCloseBlockContext()
}

type CloseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCloseBlockContext() *CloseBlockContext {
	var p = new(CloseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_closeBlock
	return p
}

func (*CloseBlockContext) IsCloseBlockContext() {}

func NewCloseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CloseBlockContext {
	var p = new(CloseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_closeBlock

	return p
}

func (s *CloseBlockContext) GetParser() antlr.Parser { return s.parser }
func (s *CloseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CloseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CloseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterCloseBlock(s)
	}
}

func (s *CloseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitCloseBlock(s)
	}
}




func (p *GShellParser) CloseBlock() (localctx ICloseBlockContext) {
	localctx = NewCloseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GShellParserRULE_closeBlock)

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
		p.SetState(55)
		p.Match(GShellParserT__1)
	}



	return localctx
}


// ICommandBlockContext is an interface to support dynamic dispatch.
type ICommandBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandBlockContext differentiates from other interfaces.
	IsCommandBlockContext()
}

type CommandBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandBlockContext() *CommandBlockContext {
	var p = new(CommandBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_commandBlock
	return p
}

func (*CommandBlockContext) IsCommandBlockContext() {}

func NewCommandBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandBlockContext {
	var p = new(CommandBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_commandBlock

	return p
}

func (s *CommandBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandBlockContext) OpenBlock() IOpenBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpenBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpenBlockContext)
}

func (s *CommandBlockContext) CommandBlockTail() ICommandBlockTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandBlockTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandBlockTailContext)
}

func (s *CommandBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CommandBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterCommandBlock(s)
	}
}

func (s *CommandBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitCommandBlock(s)
	}
}




func (p *GShellParser) CommandBlock() (localctx ICommandBlockContext) {
	localctx = NewCommandBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GShellParserRULE_commandBlock)

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
		p.SetState(57)
		p.OpenBlock()
	}
	{
		p.SetState(58)
		p.CommandBlockTail()
	}



	return localctx
}


// ICommandBlockTailContext is an interface to support dynamic dispatch.
type ICommandBlockTailContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandBlockTailContext differentiates from other interfaces.
	IsCommandBlockTailContext()
}

type CommandBlockTailContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandBlockTailContext() *CommandBlockTailContext {
	var p = new(CommandBlockTailContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_commandBlockTail
	return p
}

func (*CommandBlockTailContext) IsCommandBlockTailContext() {}

func NewCommandBlockTailContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandBlockTailContext {
	var p = new(CommandBlockTailContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_commandBlockTail

	return p
}

func (s *CommandBlockTailContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandBlockTailContext) SingleCommand() ISingleCommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleCommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleCommandContext)
}

func (s *CommandBlockTailContext) Terminator() ITerminatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerminatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *CommandBlockTailContext) CommandBlockTail() ICommandBlockTailContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandBlockTailContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandBlockTailContext)
}

func (s *CommandBlockTailContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(GShellParserNL)
}

func (s *CommandBlockTailContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(GShellParserNL, i)
}

func (s *CommandBlockTailContext) CloseBlock() ICloseBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICloseBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICloseBlockContext)
}

func (s *CommandBlockTailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandBlockTailContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *CommandBlockTailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterCommandBlockTail(s)
	}
}

func (s *CommandBlockTailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitCommandBlockTail(s)
	}
}




func (p *GShellParser) CommandBlockTail() (localctx ICommandBlockTailContext) {
	localctx = NewCommandBlockTailContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GShellParserRULE_commandBlockTail)
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

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == GShellParserNL {
			{
				p.SetState(60)
				p.Match(GShellParserNL)
			}


			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(66)
			p.SingleCommand()
		}
		{
			p.SetState(67)
			p.Terminator()
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(68)
					p.Match(GShellParserNL)
				}


			}
			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		{
			p.SetState(74)
			p.CommandBlockTail()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == GShellParserNL {
			{
				p.SetState(76)
				p.Match(GShellParserNL)
			}


			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(82)
			p.SingleCommand()
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == GShellParserNL {
			{
				p.SetState(83)
				p.Match(GShellParserNL)
			}


			p.SetState(88)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(89)
			p.CloseBlock()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.CloseBlock()
		}

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

func (s *ScriptContext) CommandBlock() ICommandBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandBlockContext)
}

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(GShellParserEOF, 0)
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
	p.EnterRule(localctx, 14, GShellParserRULE_script)

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
		p.SetState(94)
		p.CommandBlock()
	}
	{
		p.SetState(95)
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
	p.EnterRule(localctx, 16, GShellParserRULE_singleCommand)

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
		p.CommandLine()
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
	p.EnterRule(localctx, 18, GShellParserRULE_commandLine)

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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.CommandName()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(100)
			p.CommandName()
		}
		{
			p.SetState(101)
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
	p.EnterRule(localctx, 20, GShellParserRULE_commandName)

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
		p.SetState(105)
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

func (s *ArgumentsContext) ListArgument() IListArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListArgumentContext)
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
	p.EnterRule(localctx, 22, GShellParserRULE_arguments)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GShellParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.NamedArgument()
		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(108)
					p.Arguments()
				}


			}
			p.SetState(113)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}


	case GShellParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(114)
			p.NumericArgument()
		}
		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(115)
					p.Arguments()
				}


			}
			p.SetState(120)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}


	case GShellParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(121)
			p.VariableArgument()
		}
		p.SetState(125)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(122)
					p.Arguments()
				}


			}
			p.SetState(127)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}


	case GShellParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(128)
			p.ScriptArgument()
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(129)
					p.Arguments()
				}


			}
			p.SetState(134)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}


	case GShellParserT__3:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(135)
			p.ListArgument()
		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(136)
					p.Arguments()
				}


			}
			p.SetState(141)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 24, GShellParserRULE_namedArgument)

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
		p.SetState(144)
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
	p.EnterRule(localctx, 26, GShellParserRULE_numericArgument)

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
		p.SetState(146)
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
	p.EnterRule(localctx, 28, GShellParserRULE_variableArgument)

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
		p.SetState(148)
		p.Match(GShellParserT__2)
	}
	{
		p.SetState(149)
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

func (s *ScriptArgumentContext) CommandBlock() ICommandBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandBlockContext)
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
	p.EnterRule(localctx, 30, GShellParserRULE_scriptArgument)

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
		p.SetState(151)
		p.CommandBlock()
	}



	return localctx
}


// IListArgumentContext is an interface to support dynamic dispatch.
type IListArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListArgumentContext differentiates from other interfaces.
	IsListArgumentContext()
}

type ListArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListArgumentContext() *ListArgumentContext {
	var p = new(ListArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_listArgument
	return p
}

func (*ListArgumentContext) IsListArgumentContext() {}

func NewListArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArgumentContext {
	var p = new(ListArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_listArgument

	return p
}

func (s *ListArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArgumentContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ListArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ListArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterListArgument(s)
	}
}

func (s *ListArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitListArgument(s)
	}
}




func (p *GShellParser) ListArgument() (localctx IListArgumentContext) {
	localctx = NewListArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GShellParserRULE_listArgument)

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
		p.SetState(153)
		p.Match(GShellParserT__3)
	}
	{
		p.SetState(154)
		p.Arguments()
	}
	{
		p.SetState(155)
		p.Match(GShellParserT__4)
	}



	return localctx
}


