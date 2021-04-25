// Code generated from GShell.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // GShell

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GShellListener is a complete listener for a parse tree produced by GShellParser.
type GShellListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterTerminator is called when entering the terminator production.
	EnterTerminator(c *TerminatorContext)

	// EnterCommandListItem is called when entering the commandListItem production.
	EnterCommandListItem(c *CommandListItemContext)

	// EnterOpenBlock is called when entering the openBlock production.
	EnterOpenBlock(c *OpenBlockContext)

	// EnterCloseBlock is called when entering the closeBlock production.
	EnterCloseBlock(c *CloseBlockContext)

	// EnterCommandBlock is called when entering the commandBlock production.
	EnterCommandBlock(c *CommandBlockContext)

	// EnterCommandBlockTail is called when entering the commandBlockTail production.
	EnterCommandBlockTail(c *CommandBlockTailContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterSingleCommand is called when entering the singleCommand production.
	EnterSingleCommand(c *SingleCommandContext)

	// EnterCommandLine is called when entering the commandLine production.
	EnterCommandLine(c *CommandLineContext)

	// EnterCommandName is called when entering the commandName production.
	EnterCommandName(c *CommandNameContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterNamedArgument is called when entering the namedArgument production.
	EnterNamedArgument(c *NamedArgumentContext)

	// EnterNumericArgument is called when entering the numericArgument production.
	EnterNumericArgument(c *NumericArgumentContext)

	// EnterVariableArgument is called when entering the variableArgument production.
	EnterVariableArgument(c *VariableArgumentContext)

	// EnterScriptArgument is called when entering the scriptArgument production.
	EnterScriptArgument(c *ScriptArgumentContext)

	// EnterListArgument is called when entering the listArgument production.
	EnterListArgument(c *ListArgumentContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTerminator is called when exiting the terminator production.
	ExitTerminator(c *TerminatorContext)

	// ExitCommandListItem is called when exiting the commandListItem production.
	ExitCommandListItem(c *CommandListItemContext)

	// ExitOpenBlock is called when exiting the openBlock production.
	ExitOpenBlock(c *OpenBlockContext)

	// ExitCloseBlock is called when exiting the closeBlock production.
	ExitCloseBlock(c *CloseBlockContext)

	// ExitCommandBlock is called when exiting the commandBlock production.
	ExitCommandBlock(c *CommandBlockContext)

	// ExitCommandBlockTail is called when exiting the commandBlockTail production.
	ExitCommandBlockTail(c *CommandBlockTailContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitSingleCommand is called when exiting the singleCommand production.
	ExitSingleCommand(c *SingleCommandContext)

	// ExitCommandLine is called when exiting the commandLine production.
	ExitCommandLine(c *CommandLineContext)

	// ExitCommandName is called when exiting the commandName production.
	ExitCommandName(c *CommandNameContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitNamedArgument is called when exiting the namedArgument production.
	ExitNamedArgument(c *NamedArgumentContext)

	// ExitNumericArgument is called when exiting the numericArgument production.
	ExitNumericArgument(c *NumericArgumentContext)

	// ExitVariableArgument is called when exiting the variableArgument production.
	ExitVariableArgument(c *VariableArgumentContext)

	// ExitScriptArgument is called when exiting the scriptArgument production.
	ExitScriptArgument(c *ScriptArgumentContext)

	// ExitListArgument is called when exiting the listArgument production.
	ExitListArgument(c *ListArgumentContext)
}
