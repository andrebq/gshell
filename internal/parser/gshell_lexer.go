// Code generated from GShell.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 9, 94, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 
	13, 9, 13, 4, 14, 9, 14, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 
	3, 6, 3, 6, 3, 7, 3, 7, 6, 7, 42, 10, 7, 13, 7, 14, 7, 43, 3, 7, 6, 7, 
	47, 10, 7, 13, 7, 14, 7, 48, 5, 7, 51, 10, 7, 3, 8, 3, 8, 3, 8, 6, 8, 56, 
	10, 8, 13, 8, 14, 8, 57, 3, 9, 6, 9, 61, 10, 9, 13, 9, 14, 9, 62, 3, 9, 
	6, 9, 66, 10, 9, 13, 9, 14, 9, 67, 5, 9, 70, 10, 9, 3, 10, 3, 10, 3, 10, 
	7, 10, 75, 10, 10, 12, 10, 14, 10, 78, 11, 10, 3, 11, 3, 11, 5, 11, 82, 
	10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 6, 14, 89, 10, 14, 13, 14, 14, 
	14, 90, 3, 14, 3, 14, 2, 2, 15, 3, 3, 5, 4, 7, 2, 9, 2, 11, 2, 13, 2, 15, 
	2, 17, 2, 19, 5, 21, 6, 23, 7, 25, 8, 27, 9, 3, 2, 6, 9, 2, 35, 35, 37, 
	40, 44, 45, 47, 48, 65, 66, 96, 96, 128, 128, 3, 2, 61, 61, 3, 2, 12, 12, 
	5, 2, 11, 11, 15, 15, 34, 34, 4, 136, 2, 67, 2, 92, 2, 99, 2, 124, 2, 126, 
	2, 126, 2, 183, 2, 183, 2, 194, 2, 216, 2, 218, 2, 248, 2, 250, 2, 444, 
	2, 446, 2, 449, 2, 454, 2, 454, 2, 456, 2, 457, 2, 459, 2, 460, 2, 462, 
	2, 499, 2, 501, 2, 661, 2, 663, 2, 689, 2, 882, 2, 885, 2, 888, 2, 889, 
	2, 893, 2, 895, 2, 897, 2, 897, 2, 904, 2, 904, 2, 906, 2, 908, 2, 910, 
	2, 910, 2, 912, 2, 931, 2, 933, 2, 1015, 2, 1017, 2, 1155, 2, 1164, 2, 
	1329, 2, 1331, 2, 1368, 2, 1379, 2, 1417, 2, 4258, 2, 4295, 2, 4297, 2, 
	4297, 2, 4303, 2, 4303, 2, 5026, 2, 5111, 2, 5114, 2, 5119, 2, 7298, 2, 
	7306, 2, 7426, 2, 7469, 2, 7533, 2, 7545, 2, 7547, 2, 7580, 2, 7682, 2, 
	7959, 2, 7962, 2, 7967, 2, 7970, 2, 8007, 2, 8010, 2, 8015, 2, 8018, 2, 
	8025, 2, 8027, 2, 8027, 2, 8029, 2, 8029, 2, 8031, 2, 8031, 2, 8033, 2, 
	8063, 2, 8066, 2, 8073, 2, 8082, 2, 8089, 2, 8098, 2, 8105, 2, 8114, 2, 
	8118, 2, 8120, 2, 8125, 2, 8128, 2, 8128, 2, 8132, 2, 8134, 2, 8136, 2, 
	8141, 2, 8146, 2, 8149, 2, 8152, 2, 8157, 2, 8162, 2, 8174, 2, 8180, 2, 
	8182, 2, 8184, 2, 8189, 2, 8452, 2, 8452, 2, 8457, 2, 8457, 2, 8460, 2, 
	8469, 2, 8471, 2, 8471, 2, 8475, 2, 8479, 2, 8486, 2, 8486, 2, 8488, 2, 
	8488, 2, 8490, 2, 8490, 2, 8492, 2, 8495, 2, 8497, 2, 8502, 2, 8507, 2, 
	8507, 2, 8510, 2, 8513, 2, 8519, 2, 8523, 2, 8528, 2, 8528, 2, 8581, 2, 
	8582, 2, 11266, 2, 11312, 2, 11314, 2, 11360, 2, 11362, 2, 11389, 2, 11392, 
	2, 11494, 2, 11501, 2, 11504, 2, 11508, 2, 11509, 2, 11522, 2, 11559, 2, 
	11561, 2, 11561, 2, 11567, 2, 11567, 2, 42562, 2, 42607, 2, 42626, 2, 42653, 
	2, 42788, 2, 42865, 2, 42867, 2, 42889, 2, 42893, 2, 42896, 2, 42898, 2, 
	42928, 2, 42930, 2, 42937, 2, 43004, 2, 43004, 2, 43826, 2, 43868, 2, 43874, 
	2, 43879, 2, 43890, 2, 43969, 2, 64258, 2, 64264, 2, 64277, 2, 64281, 2, 
	65315, 2, 65340, 2, 65347, 2, 65372, 2, 1026, 3, 1105, 3, 1202, 3, 1237, 
	3, 1242, 3, 1277, 3, 3202, 3, 3252, 3, 3266, 3, 3316, 3, 6306, 3, 6369, 
	3, 54274, 3, 54358, 3, 54360, 3, 54430, 3, 54432, 3, 54433, 3, 54436, 3, 
	54436, 3, 54439, 3, 54440, 3, 54443, 3, 54446, 3, 54448, 3, 54459, 3, 54461, 
	3, 54461, 3, 54463, 3, 54469, 3, 54471, 3, 54535, 3, 54537, 3, 54540, 3, 
	54543, 3, 54550, 3, 54552, 3, 54558, 3, 54560, 3, 54587, 3, 54589, 3, 54592, 
	3, 54594, 3, 54598, 3, 54600, 3, 54600, 3, 54604, 3, 54610, 3, 54612, 3, 
	54951, 3, 54954, 3, 54978, 3, 54980, 3, 55004, 3, 55006, 3, 55036, 3, 55038, 
	3, 55062, 3, 55064, 3, 55094, 3, 55096, 3, 55120, 3, 55122, 3, 55152, 3, 
	55154, 3, 55178, 3, 55180, 3, 55210, 3, 55212, 3, 55236, 3, 55238, 3, 55245, 
	3, 59650, 3, 59717, 3, 57, 2, 50, 2, 59, 2, 1634, 2, 1643, 2, 1778, 2, 
	1787, 2, 1986, 2, 1995, 2, 2408, 2, 2417, 2, 2536, 2, 2545, 2, 2664, 2, 
	2673, 2, 2792, 2, 2801, 2, 2920, 2, 2929, 2, 3048, 2, 3057, 2, 3176, 2, 
	3185, 2, 3304, 2, 3313, 2, 3432, 2, 3441, 2, 3560, 2, 3569, 2, 3666, 2, 
	3675, 2, 3794, 2, 3803, 2, 3874, 2, 3883, 2, 4162, 2, 4171, 2, 4242, 2, 
	4251, 2, 6114, 2, 6123, 2, 6162, 2, 6171, 2, 6472, 2, 6481, 2, 6610, 2, 
	6619, 2, 6786, 2, 6795, 2, 6802, 2, 6811, 2, 6994, 2, 7003, 2, 7090, 2, 
	7099, 2, 7234, 2, 7243, 2, 7250, 2, 7259, 2, 42530, 2, 42539, 2, 43218, 
	2, 43227, 2, 43266, 2, 43275, 2, 43474, 2, 43483, 2, 43506, 2, 43515, 2, 
	43602, 2, 43611, 2, 44018, 2, 44027, 2, 65298, 2, 65307, 2, 1186, 3, 1195, 
	3, 4200, 3, 4209, 3, 4338, 3, 4347, 3, 4408, 3, 4417, 3, 4562, 3, 4571, 
	3, 4850, 3, 4859, 3, 5202, 3, 5211, 3, 5330, 3, 5339, 3, 5714, 3, 5723, 
	3, 5826, 3, 5835, 3, 5938, 3, 5947, 3, 6370, 3, 6379, 3, 7250, 3, 7259, 
	3, 7506, 3, 7515, 3, 27234, 3, 27243, 3, 27474, 3, 27483, 3, 55248, 3, 
	55297, 3, 59730, 3, 59739, 3, 98, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 
	19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 
	2, 27, 3, 2, 2, 2, 3, 29, 3, 2, 2, 2, 5, 31, 3, 2, 2, 2, 7, 33, 3, 2, 2, 
	2, 9, 35, 3, 2, 2, 2, 11, 37, 3, 2, 2, 2, 13, 50, 3, 2, 2, 2, 15, 52, 3, 
	2, 2, 2, 17, 69, 3, 2, 2, 2, 19, 71, 3, 2, 2, 2, 21, 81, 3, 2, 2, 2, 23, 
	83, 3, 2, 2, 2, 25, 85, 3, 2, 2, 2, 27, 88, 3, 2, 2, 2, 29, 30, 7, 125, 
	2, 2, 30, 4, 3, 2, 2, 2, 31, 32, 7, 127, 2, 2, 32, 6, 3, 2, 2, 2, 33, 34, 
	9, 6, 2, 2, 34, 8, 3, 2, 2, 2, 35, 36, 9, 7, 2, 2, 36, 10, 3, 2, 2, 2, 
	37, 38, 9, 2, 2, 2, 38, 12, 3, 2, 2, 2, 39, 41, 7, 47, 2, 2, 40, 42, 5, 
	9, 5, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 
	44, 3, 2, 2, 2, 44, 51, 3, 2, 2, 2, 45, 47, 5, 9, 5, 2, 46, 45, 3, 2, 2, 
	2, 47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 51, 
	3, 2, 2, 2, 50, 39, 3, 2, 2, 2, 50, 46, 3, 2, 2, 2, 51, 14, 3, 2, 2, 2, 
	52, 53, 5, 13, 7, 2, 53, 55, 7, 48, 2, 2, 54, 56, 5, 9, 5, 2, 55, 54, 3, 
	2, 2, 2, 56, 57, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 
	16, 3, 2, 2, 2, 59, 61, 5, 7, 4, 2, 60, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 
	2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 70, 3, 2, 2, 2, 64, 66, 
	5, 11, 6, 2, 65, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 
	67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 60, 3, 2, 2, 2, 69, 65, 3, 
	2, 2, 2, 70, 18, 3, 2, 2, 2, 71, 76, 5, 17, 9, 2, 72, 75, 5, 9, 5, 2, 73, 
	75, 5, 17, 9, 2, 74, 72, 3, 2, 2, 2, 74, 73, 3, 2, 2, 2, 75, 78, 3, 2, 
	2, 2, 76, 74, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 20, 3, 2, 2, 2, 78, 76, 
	3, 2, 2, 2, 79, 82, 5, 13, 7, 2, 80, 82, 5, 15, 8, 2, 81, 79, 3, 2, 2, 
	2, 81, 80, 3, 2, 2, 2, 82, 22, 3, 2, 2, 2, 83, 84, 9, 3, 2, 2, 84, 24, 
	3, 2, 2, 2, 85, 86, 9, 4, 2, 2, 86, 26, 3, 2, 2, 2, 87, 89, 9, 5, 2, 2, 
	88, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 
	2, 2, 2, 91, 92, 3, 2, 2, 2, 92, 93, 8, 14, 2, 2, 93, 28, 3, 2, 2, 2, 14, 
	2, 43, 48, 50, 57, 62, 67, 69, 74, 76, 81, 90, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'{'", "'}'",
}

var lexerSymbolicNames = []string{
	"", "", "", "IDENTIFIER", "NUMBER", "TERMINATOR", "NL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "LETTER", "DIGIT", "PUCTUATION", "INT", "FLOAT", "IDENTIFER_START", 
	"IDENTIFIER", "NUMBER", "TERMINATOR", "NL", "WS",
}

type GShellLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewGShellLexer(input antlr.CharStream) *GShellLexer {

	l := new(GShellLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "GShell.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GShellLexer tokens.
const (
	GShellLexerT__0 = 1
	GShellLexerT__1 = 2
	GShellLexerIDENTIFIER = 3
	GShellLexerNUMBER = 4
	GShellLexerTERMINATOR = 5
	GShellLexerNL = 6
	GShellLexerWS = 7
)

