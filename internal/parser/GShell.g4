grammar GShell;

// Tokens
fragment LETTER: [\p{Ll}|\p{Lu}];
fragment DIGIT: [\p{Nd}];
fragment PUCTUATION: [!?.\-+*&^%$#@~];
fragment INT: '-' DIGIT+ | DIGIT+;
fragment FLOAT: INT '.' DIGIT+;
fragment IDENTIFER_START: LETTER+|PUCTUATION+;

IDENTIFIER: IDENTIFER_START (DIGIT|IDENTIFER_START)*;
NUMBER: INT | FLOAT;

TERMINATOR: [;];
NL: [\n];

WS: [ \r\t]+ -> skip;

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
