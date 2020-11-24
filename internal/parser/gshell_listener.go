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

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitTerminator is called when exiting the terminator production.
	ExitTerminator(c *TerminatorContext)

	// ExitCommandListItem is called when exiting the commandListItem production.
	ExitCommandListItem(c *CommandListItemContext)

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
}
