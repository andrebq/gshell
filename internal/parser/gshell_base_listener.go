// Code generated from GShell.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GShell

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGShellListener is a complete listener for a parse tree produced by GShellParser.
type BaseGShellListener struct{}

var _ GShellListener = &BaseGShellListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGShellListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGShellListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGShellListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGShellListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseGShellListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseGShellListener) ExitStart(ctx *StartContext) {}

// EnterTerminator is called when production terminator is entered.
func (s *BaseGShellListener) EnterTerminator(ctx *TerminatorContext) {}

// ExitTerminator is called when production terminator is exited.
func (s *BaseGShellListener) ExitTerminator(ctx *TerminatorContext) {}

// EnterCommandListItem is called when production commandListItem is entered.
func (s *BaseGShellListener) EnterCommandListItem(ctx *CommandListItemContext) {}

// ExitCommandListItem is called when production commandListItem is exited.
func (s *BaseGShellListener) ExitCommandListItem(ctx *CommandListItemContext) {}

// EnterOpenBlock is called when production openBlock is entered.
func (s *BaseGShellListener) EnterOpenBlock(ctx *OpenBlockContext) {}

// ExitOpenBlock is called when production openBlock is exited.
func (s *BaseGShellListener) ExitOpenBlock(ctx *OpenBlockContext) {}

// EnterCloseBlock is called when production closeBlock is entered.
func (s *BaseGShellListener) EnterCloseBlock(ctx *CloseBlockContext) {}

// ExitCloseBlock is called when production closeBlock is exited.
func (s *BaseGShellListener) ExitCloseBlock(ctx *CloseBlockContext) {}

// EnterCommandBlock is called when production commandBlock is entered.
func (s *BaseGShellListener) EnterCommandBlock(ctx *CommandBlockContext) {}

// ExitCommandBlock is called when production commandBlock is exited.
func (s *BaseGShellListener) ExitCommandBlock(ctx *CommandBlockContext) {}

// EnterCommandBlockTail is called when production commandBlockTail is entered.
func (s *BaseGShellListener) EnterCommandBlockTail(ctx *CommandBlockTailContext) {}

// ExitCommandBlockTail is called when production commandBlockTail is exited.
func (s *BaseGShellListener) ExitCommandBlockTail(ctx *CommandBlockTailContext) {}

// EnterScript is called when production script is entered.
func (s *BaseGShellListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseGShellListener) ExitScript(ctx *ScriptContext) {}

// EnterSingleCommand is called when production singleCommand is entered.
func (s *BaseGShellListener) EnterSingleCommand(ctx *SingleCommandContext) {}

// ExitSingleCommand is called when production singleCommand is exited.
func (s *BaseGShellListener) ExitSingleCommand(ctx *SingleCommandContext) {}

// EnterCommandLine is called when production commandLine is entered.
func (s *BaseGShellListener) EnterCommandLine(ctx *CommandLineContext) {}

// ExitCommandLine is called when production commandLine is exited.
func (s *BaseGShellListener) ExitCommandLine(ctx *CommandLineContext) {}

// EnterCommandName is called when production commandName is entered.
func (s *BaseGShellListener) EnterCommandName(ctx *CommandNameContext) {}

// ExitCommandName is called when production commandName is exited.
func (s *BaseGShellListener) ExitCommandName(ctx *CommandNameContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGShellListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGShellListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterNamedArgument is called when production namedArgument is entered.
func (s *BaseGShellListener) EnterNamedArgument(ctx *NamedArgumentContext) {}

// ExitNamedArgument is called when production namedArgument is exited.
func (s *BaseGShellListener) ExitNamedArgument(ctx *NamedArgumentContext) {}

// EnterNumericArgument is called when production numericArgument is entered.
func (s *BaseGShellListener) EnterNumericArgument(ctx *NumericArgumentContext) {}

// ExitNumericArgument is called when production numericArgument is exited.
func (s *BaseGShellListener) ExitNumericArgument(ctx *NumericArgumentContext) {}

// EnterVariableArgument is called when production variableArgument is entered.
func (s *BaseGShellListener) EnterVariableArgument(ctx *VariableArgumentContext) {}

// ExitVariableArgument is called when production variableArgument is exited.
func (s *BaseGShellListener) ExitVariableArgument(ctx *VariableArgumentContext) {}

// EnterScriptArgument is called when production scriptArgument is entered.
func (s *BaseGShellListener) EnterScriptArgument(ctx *ScriptArgumentContext) {}

// ExitScriptArgument is called when production scriptArgument is exited.
func (s *BaseGShellListener) ExitScriptArgument(ctx *ScriptArgumentContext) {}

// EnterListArgument is called when production listArgument is entered.
func (s *BaseGShellListener) EnterListArgument(ctx *ListArgumentContext) {}

// ExitListArgument is called when production listArgument is exited.
func (s *BaseGShellListener) ExitListArgument(ctx *ListArgumentContext) {}
