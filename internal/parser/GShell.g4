grammar GShell;

// Tokens
fragment LETTERS: [\p{Ll}|\p{Lu}]+;
fragment DIGITS: [\p{Nd}]+;
fragment INT: '-' DIGITS | DIGITS;
fragment FLOAT: INT '.' DIGITS;

IDENTIFIER: LETTERS DIGITS|LETTERS;
NUMBER: INT | FLOAT;

TERMINATOR: [;];
NL: [\n];

WHITESPACE: [ \r\t]+ -> skip;

// Rules
start
   : script EOF;

terminator
   : TERMINATOR
   | NL ;

commandListItem
   : NL* singleCommand NL*;
  
script
   : '{' commandListItem* '}' EOF;

singleCommand
   : commandLine terminator ;

commandLine
   : commandName
   | commandName arguments ;

commandName : IDENTIFIER ;

arguments
   : namedArgument arguments*
   | numericArgument arguments*;

namedArgument : IDENTIFIER ;
numericArgument : NUMBER ;

// expression
//    : expression op=('*'|'/') expression # MulDiv
//    | expression op=('+'|'-') expression # AddSub
//    | NUMBER                             # Number
//    ;
