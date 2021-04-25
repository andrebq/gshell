grammar GShell;

// Tokens
fragment LETTER: [\p{Ll}|\p{Lu}];
fragment DIGIT: [\p{Nd}];
fragment PUNCTUATION_HEAD: [!?.\-+*&^%#@~];
fragment PUCTUATION_TAIL: PUNCTUATION_HEAD | [$];
fragment INT: '-' DIGIT+ | DIGIT+;
fragment FLOAT: INT '.' DIGIT+;
fragment IDENTIFER_START: LETTER|PUNCTUATION_HEAD;
fragment IDENTIFIER_TAIL: (DIGIT|LETTER|PUCTUATION_TAIL);

IDENTIFIER: IDENTIFER_START IDENTIFIER_TAIL*;
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
   : NL* singleCommand terminator NL*;

openBlock
   : '{';

closeBlock
   : '}';

commandBlock
   //: openBlock commandListItem* closeBlock
   : openBlock commandBlockTail;

commandBlockTail
   : NL* singleCommand terminator NL* commandBlockTail
   | NL* singleCommand NL* closeBlock
   | closeBlock;

script
   : commandBlock EOF;

singleCommand
   : commandLine ;

commandLine
   : commandName
   | commandName arguments ;

commandName : IDENTIFIER ;

arguments
   : namedArgument arguments*
   | numericArgument arguments*
   | variableArgument arguments*
   | scriptArgument arguments*
   | listArgument arguments*;

namedArgument : IDENTIFIER ;
numericArgument : NUMBER ;
variableArgument : '$' IDENTIFIER ;
scriptArgument: commandBlock ;
listArgument: '[' arguments ']' ;

// expression
//    : expression op=('*'|'/') expression # MulDiv
//    | expression op=('+'|'-') expression # AddSub
//    | NUMBER                             # Number
//    ;
