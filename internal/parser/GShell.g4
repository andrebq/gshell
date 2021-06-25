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

// This section was copied from ponylang grammar
// https://github.com/ponylang/ponyc/blob/main/pony.g
//
// Adjust it later such that it fits gshell better,
// but it works for now.
fragment HEX : DIGIT | 'a'..'f' | 'A'..'F';
fragment HEX_ESC : '\\' 'x' HEX HEX ;
fragment UNICODE_ESC : '\\' 'u' HEX HEX HEX HEX ;
fragment UNICODE2_ESC : '\\' 'U' HEX HEX HEX HEX HEX HEX;
fragment STRING_CHAR : '\\' '"' | ESC | ~('"' | '\\');
fragment CHAR_ESC
  : '\\' ('a' | 'b' | 'e' | 'f' | 'n' | 'r' | 't' | 'v' | '\\' | '0')
  | HEX_ESC
  ;
fragment ESC : CHAR_ESC | UNICODE_ESC | UNICODE2_ESC;
STRING
  : '"' STRING_CHAR* '"'
  | '"""' (('"' | '""') ? ~'"')* '"""' '"'*
  ;
// end of the section copied from ponylang

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
   | listArgument arguments*
   | textArgument arguments*;

namedArgument : IDENTIFIER ;
numericArgument : NUMBER ;
variableArgument : '$' IDENTIFIER ;
scriptArgument: commandBlock ;
listArgument: '[' listArgumentItems ']' ;
listArgumentItems:
   NL* arguments NL* listArgumentItems
   | ;
textArgument: STRING;

// expression
//    : expression op=('*'|'/') expression # MulDiv
//    | expression op=('+'|'-') expression # AddSub
//    | NUMBER                             # Number
//    ;
