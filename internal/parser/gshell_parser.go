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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 13, 191, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 7, 
	4, 47, 10, 4, 12, 4, 14, 4, 50, 11, 4, 3, 4, 3, 4, 3, 4, 7, 4, 55, 10, 
	4, 12, 4, 14, 4, 58, 11, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 
	8, 7, 8, 68, 10, 8, 12, 8, 14, 8, 71, 11, 8, 3, 8, 3, 8, 3, 8, 7, 8, 76, 
	10, 8, 12, 8, 14, 8, 79, 11, 8, 3, 8, 3, 8, 3, 8, 7, 8, 84, 10, 8, 12, 
	8, 14, 8, 87, 11, 8, 3, 8, 3, 8, 7, 8, 91, 10, 8, 12, 8, 14, 8, 94, 11, 
	8, 3, 8, 3, 8, 3, 8, 5, 8, 99, 10, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 
	11, 3, 11, 3, 11, 3, 11, 5, 11, 110, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 
	7, 13, 116, 10, 13, 12, 13, 14, 13, 119, 11, 13, 3, 13, 3, 13, 7, 13, 123, 
	10, 13, 12, 13, 14, 13, 126, 11, 13, 3, 13, 3, 13, 7, 13, 130, 10, 13, 
	12, 13, 14, 13, 133, 11, 13, 3, 13, 3, 13, 7, 13, 137, 10, 13, 12, 13, 
	14, 13, 140, 11, 13, 3, 13, 3, 13, 7, 13, 144, 10, 13, 12, 13, 14, 13, 
	147, 11, 13, 3, 13, 3, 13, 7, 13, 151, 10, 13, 12, 13, 14, 13, 154, 11, 
	13, 5, 13, 156, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 7, 19, 172, 10, 19, 12, 
	19, 14, 19, 175, 11, 19, 3, 19, 3, 19, 7, 19, 179, 10, 19, 12, 19, 14, 
	19, 182, 11, 19, 3, 19, 3, 19, 3, 19, 5, 19, 187, 10, 19, 3, 20, 3, 20, 
	3, 20, 2, 2, 21, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 
	32, 34, 36, 38, 2, 3, 3, 2, 11, 12, 2, 194, 2, 40, 3, 2, 2, 2, 4, 43, 3, 
	2, 2, 2, 6, 48, 3, 2, 2, 2, 8, 59, 3, 2, 2, 2, 10, 61, 3, 2, 2, 2, 12, 
	63, 3, 2, 2, 2, 14, 98, 3, 2, 2, 2, 16, 100, 3, 2, 2, 2, 18, 103, 3, 2, 
	2, 2, 20, 109, 3, 2, 2, 2, 22, 111, 3, 2, 2, 2, 24, 155, 3, 2, 2, 2, 26, 
	157, 3, 2, 2, 2, 28, 159, 3, 2, 2, 2, 30, 161, 3, 2, 2, 2, 32, 164, 3, 
	2, 2, 2, 34, 166, 3, 2, 2, 2, 36, 186, 3, 2, 2, 2, 38, 188, 3, 2, 2, 2, 
	40, 41, 5, 16, 9, 2, 41, 42, 7, 2, 2, 3, 42, 3, 3, 2, 2, 2, 43, 44, 9, 
	2, 2, 2, 44, 5, 3, 2, 2, 2, 45, 47, 7, 12, 2, 2, 46, 45, 3, 2, 2, 2, 47, 
	50, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 3, 2, 2, 
	2, 50, 48, 3, 2, 2, 2, 51, 52, 5, 18, 10, 2, 52, 56, 5, 4, 3, 2, 53, 55, 
	7, 12, 2, 2, 54, 53, 3, 2, 2, 2, 55, 58, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 
	56, 57, 3, 2, 2, 2, 57, 7, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 59, 60, 7, 3, 
	2, 2, 60, 9, 3, 2, 2, 2, 61, 62, 7, 4, 2, 2, 62, 11, 3, 2, 2, 2, 63, 64, 
	5, 8, 5, 2, 64, 65, 5, 14, 8, 2, 65, 13, 3, 2, 2, 2, 66, 68, 7, 12, 2, 
	2, 67, 66, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69, 70, 
	3, 2, 2, 2, 70, 72, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 73, 5, 18, 10, 
	2, 73, 77, 5, 4, 3, 2, 74, 76, 7, 12, 2, 2, 75, 74, 3, 2, 2, 2, 76, 79, 
	3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 80, 3, 2, 2, 2, 
	79, 77, 3, 2, 2, 2, 80, 81, 5, 14, 8, 2, 81, 99, 3, 2, 2, 2, 82, 84, 7, 
	12, 2, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85, 
	86, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 92, 5, 18, 
	10, 2, 89, 91, 7, 12, 2, 2, 90, 89, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 
	90, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 95, 3, 2, 2, 2, 94, 92, 3, 2, 2, 
	2, 95, 96, 5, 10, 6, 2, 96, 99, 3, 2, 2, 2, 97, 99, 5, 10, 6, 2, 98, 69, 
	3, 2, 2, 2, 98, 85, 3, 2, 2, 2, 98, 97, 3, 2, 2, 2, 99, 15, 3, 2, 2, 2, 
	100, 101, 5, 12, 7, 2, 101, 102, 7, 2, 2, 3, 102, 17, 3, 2, 2, 2, 103, 
	104, 5, 20, 11, 2, 104, 19, 3, 2, 2, 2, 105, 110, 5, 22, 12, 2, 106, 107, 
	5, 22, 12, 2, 107, 108, 5, 24, 13, 2, 108, 110, 3, 2, 2, 2, 109, 105, 3, 
	2, 2, 2, 109, 106, 3, 2, 2, 2, 110, 21, 3, 2, 2, 2, 111, 112, 7, 9, 2, 
	2, 112, 23, 3, 2, 2, 2, 113, 117, 5, 26, 14, 2, 114, 116, 5, 24, 13, 2, 
	115, 114, 3, 2, 2, 2, 116, 119, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 
	118, 3, 2, 2, 2, 118, 156, 3, 2, 2, 2, 119, 117, 3, 2, 2, 2, 120, 124, 
	5, 28, 15, 2, 121, 123, 5, 24, 13, 2, 122, 121, 3, 2, 2, 2, 123, 126, 3, 
	2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 156, 3, 2, 2, 
	2, 126, 124, 3, 2, 2, 2, 127, 131, 5, 30, 16, 2, 128, 130, 5, 24, 13, 2, 
	129, 128, 3, 2, 2, 2, 130, 133, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 
	132, 3, 2, 2, 2, 132, 156, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 134, 138, 
	5, 32, 17, 2, 135, 137, 5, 24, 13, 2, 136, 135, 3, 2, 2, 2, 137, 140, 3, 
	2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 156, 3, 2, 2, 
	2, 140, 138, 3, 2, 2, 2, 141, 145, 5, 34, 18, 2, 142, 144, 5, 24, 13, 2, 
	143, 142, 3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 
	146, 3, 2, 2, 2, 146, 156, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 152, 
	5, 38, 20, 2, 149, 151, 5, 24, 13, 2, 150, 149, 3, 2, 2, 2, 151, 154, 3, 
	2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 156, 3, 2, 2, 
	2, 154, 152, 3, 2, 2, 2, 155, 113, 3, 2, 2, 2, 155, 120, 3, 2, 2, 2, 155, 
	127, 3, 2, 2, 2, 155, 134, 3, 2, 2, 2, 155, 141, 3, 2, 2, 2, 155, 148, 
	3, 2, 2, 2, 156, 25, 3, 2, 2, 2, 157, 158, 7, 9, 2, 2, 158, 27, 3, 2, 2, 
	2, 159, 160, 7, 10, 2, 2, 160, 29, 3, 2, 2, 2, 161, 162, 7, 5, 2, 2, 162, 
	163, 7, 9, 2, 2, 163, 31, 3, 2, 2, 2, 164, 165, 5, 12, 7, 2, 165, 33, 3, 
	2, 2, 2, 166, 167, 7, 6, 2, 2, 167, 168, 5, 36, 19, 2, 168, 169, 7, 7, 
	2, 2, 169, 35, 3, 2, 2, 2, 170, 172, 7, 12, 2, 2, 171, 170, 3, 2, 2, 2, 
	172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 
	176, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176, 180, 5, 24, 13, 2, 177, 179, 
	7, 12, 2, 2, 178, 177, 3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 178, 3, 2, 
	2, 2, 180, 181, 3, 2, 2, 2, 181, 183, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 
	183, 184, 5, 36, 19, 2, 184, 187, 3, 2, 2, 2, 185, 187, 3, 2, 2, 2, 186, 
	173, 3, 2, 2, 2, 186, 185, 3, 2, 2, 2, 187, 37, 3, 2, 2, 2, 188, 189, 7, 
	8, 2, 2, 189, 39, 3, 2, 2, 2, 20, 48, 56, 69, 77, 85, 92, 98, 109, 117, 
	124, 131, 138, 145, 152, 155, 173, 180, 186,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'{'", "'}'", "'$'", "'['", "']'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "STRING", "IDENTIFIER", "NUMBER", "TERMINATOR", 
	"NL", "WS",
}

var ruleNames = []string{
	"start", "terminator", "commandListItem", "openBlock", "closeBlock", "commandBlock", 
	"commandBlockTail", "script", "singleCommand", "commandLine", "commandName", 
	"arguments", "namedArgument", "numericArgument", "variableArgument", "scriptArgument", 
	"listArgument", "listArgumentItems", "textArgument",
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
	GShellParserSTRING = 6
	GShellParserIDENTIFIER = 7
	GShellParserNUMBER = 8
	GShellParserTERMINATOR = 9
	GShellParserNL = 10
	GShellParserWS = 11
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
	GShellParserRULE_listArgumentItems = 17
	GShellParserRULE_textArgument = 18
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
		p.SetState(38)
		p.Script()
	}
	{
		p.SetState(39)
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
		p.SetState(41)
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
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == GShellParserNL {
		{
			p.SetState(43)
			p.Match(GShellParserNL)
		}


		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(49)
		p.SingleCommand()
	}
	{
		p.SetState(50)
		p.Terminator()
	}
	p.SetState(54)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for _la == GShellParserNL {
		{
			p.SetState(51)
			p.Match(GShellParserNL)
		}


		p.SetState(56)
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
		p.SetState(57)
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
		p.SetState(59)
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
		p.SetState(61)
		p.OpenBlock()
	}
	{
		p.SetState(62)
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

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == GShellParserNL {
			{
				p.SetState(64)
				p.Match(GShellParserNL)
			}


			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(70)
			p.SingleCommand()
		}
		{
			p.SetState(71)
			p.Terminator()
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(72)
					p.Match(GShellParserNL)
				}


			}
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		{
			p.SetState(78)
			p.CommandBlockTail()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == GShellParserNL {
			{
				p.SetState(80)
				p.Match(GShellParserNL)
			}


			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(86)
			p.SingleCommand()
		}
		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == GShellParserNL {
			{
				p.SetState(87)
				p.Match(GShellParserNL)
			}


			p.SetState(92)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(93)
			p.CloseBlock()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
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
		p.SetState(98)
		p.CommandBlock()
	}
	{
		p.SetState(99)
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
		p.SetState(101)
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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.CommandName()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.CommandName()
		}
		{
			p.SetState(105)
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
		p.SetState(109)
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

func (s *ArgumentsContext) TextArgument() ITextArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITextArgumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITextArgumentContext)
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

	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GShellParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)
			p.NamedArgument()
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(112)
					p.Arguments()
				}


			}
			p.SetState(117)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
		}


	case GShellParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.NumericArgument()
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(119)
					p.Arguments()
				}


			}
			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
		}


	case GShellParserT__2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.VariableArgument()
		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(126)
					p.Arguments()
				}


			}
			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
		}


	case GShellParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(132)
			p.ScriptArgument()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(133)
					p.Arguments()
				}


			}
			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}


	case GShellParserT__3:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(139)
			p.ListArgument()
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(140)
					p.Arguments()
				}


			}
			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
		}


	case GShellParserSTRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(146)
			p.TextArgument()
		}
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(147)
					p.Arguments()
				}


			}
			p.SetState(152)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
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
		p.SetState(155)
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
		p.SetState(157)
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
		p.SetState(159)
		p.Match(GShellParserT__2)
	}
	{
		p.SetState(160)
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
		p.SetState(162)
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

func (s *ListArgumentContext) ListArgumentItems() IListArgumentItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArgumentItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListArgumentItemsContext)
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
		p.SetState(164)
		p.Match(GShellParserT__3)
	}
	{
		p.SetState(165)
		p.ListArgumentItems()
	}
	{
		p.SetState(166)
		p.Match(GShellParserT__4)
	}



	return localctx
}


// IListArgumentItemsContext is an interface to support dynamic dispatch.
type IListArgumentItemsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListArgumentItemsContext differentiates from other interfaces.
	IsListArgumentItemsContext()
}

type ListArgumentItemsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListArgumentItemsContext() *ListArgumentItemsContext {
	var p = new(ListArgumentItemsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_listArgumentItems
	return p
}

func (*ListArgumentItemsContext) IsListArgumentItemsContext() {}

func NewListArgumentItemsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArgumentItemsContext {
	var p = new(ListArgumentItemsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_listArgumentItems

	return p
}

func (s *ListArgumentItemsContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArgumentItemsContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ListArgumentItemsContext) ListArgumentItems() IListArgumentItemsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArgumentItemsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListArgumentItemsContext)
}

func (s *ListArgumentItemsContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(GShellParserNL)
}

func (s *ListArgumentItemsContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(GShellParserNL, i)
}

func (s *ListArgumentItemsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArgumentItemsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ListArgumentItemsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterListArgumentItems(s)
	}
}

func (s *ListArgumentItemsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitListArgumentItems(s)
	}
}




func (p *GShellParser) ListArgumentItems() (localctx IListArgumentItemsContext) {
	localctx = NewListArgumentItemsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GShellParserRULE_listArgumentItems)
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

	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GShellParserT__0, GShellParserT__2, GShellParserT__3, GShellParserSTRING, GShellParserIDENTIFIER, GShellParserNUMBER, GShellParserNL:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for _la == GShellParserNL {
			{
				p.SetState(168)
				p.Match(GShellParserNL)
			}


			p.SetState(173)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(174)
			p.Arguments()
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(175)
					p.Match(GShellParserNL)
				}


			}
			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}
		{
			p.SetState(181)
			p.ListArgumentItems()
		}


	case GShellParserT__4:
		p.EnterOuterAlt(localctx, 2)



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ITextArgumentContext is an interface to support dynamic dispatch.
type ITextArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTextArgumentContext differentiates from other interfaces.
	IsTextArgumentContext()
}

type TextArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTextArgumentContext() *TextArgumentContext {
	var p = new(TextArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GShellParserRULE_textArgument
	return p
}

func (*TextArgumentContext) IsTextArgumentContext() {}

func NewTextArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TextArgumentContext {
	var p = new(TextArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GShellParserRULE_textArgument

	return p
}

func (s *TextArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *TextArgumentContext) STRING() antlr.TerminalNode {
	return s.GetToken(GShellParserSTRING, 0)
}

func (s *TextArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TextArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TextArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.EnterTextArgument(s)
	}
}

func (s *TextArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GShellListener); ok {
		listenerT.ExitTextArgument(s)
	}
}




func (p *GShellParser) TextArgument() (localctx ITextArgumentContext) {
	localctx = NewTextArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GShellParserRULE_textArgument)

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
		p.SetState(186)
		p.Match(GShellParserSTRING)
	}



	return localctx
}


