// Generated from c:\Users\andre\Source\github.com\andrebq\gshell\internal\parser\gshell.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class GShellLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, IDENTIFIER=3, NUMBER=4, TERMINATOR=5, NL=6, WS=7;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "LETTER", "DIGIT", "PUCTUATION", "INT", "FLOAT", "IDENTIFER_START", 
			"IDENTIFIER", "NUMBER", "TERMINATOR", "NL", "WS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'{'", "'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, "IDENTIFIER", "NUMBER", "TERMINATOR", "NL", "WS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public GShellLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "gshell.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\t^\b\1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6"+
		"\3\7\3\7\6\7*\n\7\r\7\16\7+\3\7\6\7/\n\7\r\7\16\7\60\5\7\63\n\7\3\b\3"+
		"\b\3\b\6\b8\n\b\r\b\16\b9\3\t\6\t=\n\t\r\t\16\t>\3\t\6\tB\n\t\r\t\16\t"+
		"C\5\tF\n\t\3\n\3\n\3\n\7\nK\n\n\f\n\16\nN\13\n\3\13\3\13\5\13R\n\13\3"+
		"\f\3\f\3\r\3\r\3\16\6\16Y\n\16\r\16\16\16Z\3\16\3\16\2\2\17\3\3\5\4\7"+
		"\2\t\2\13\2\r\2\17\2\21\2\23\5\25\6\27\7\31\b\33\t\3\2\6\t\2##%(,-/\60"+
		"AB``\u0080\u0080\3\2==\3\2\f\f\5\2\13\13\17\17\"\"\4\u0088\2C\2\\\2c\2"+
		"|\2~\2~\2\u00b7\2\u00b7\2\u00c2\2\u00d8\2\u00da\2\u00f8\2\u00fa\2\u01bc"+
		"\2\u01be\2\u01c1\2\u01c6\2\u01c6\2\u01c8\2\u01c9\2\u01cb\2\u01cc\2\u01ce"+
		"\2\u01f3\2\u01f5\2\u0295\2\u0297\2\u02b1\2\u0372\2\u0375\2\u0378\2\u0379"+
		"\2\u037d\2\u037f\2\u0381\2\u0381\2\u0388\2\u0388\2\u038a\2\u038c\2\u038e"+
		"\2\u038e\2\u0390\2\u03a3\2\u03a5\2\u03f7\2\u03f9\2\u0483\2\u048c\2\u0531"+
		"\2\u0533\2\u0558\2\u0563\2\u0589\2\u10a2\2\u10c7\2\u10c9\2\u10c9\2\u10cf"+
		"\2\u10cf\2\u13a2\2\u13f7\2\u13fa\2\u13ff\2\u1c82\2\u1c8a\2\u1d02\2\u1d2d"+
		"\2\u1d6d\2\u1d79\2\u1d7b\2\u1d9c\2\u1e02\2\u1f17\2\u1f1a\2\u1f1f\2\u1f22"+
		"\2\u1f47\2\u1f4a\2\u1f4f\2\u1f52\2\u1f59\2\u1f5b\2\u1f5b\2\u1f5d\2\u1f5d"+
		"\2\u1f5f\2\u1f5f\2\u1f61\2\u1f7f\2\u1f82\2\u1f89\2\u1f92\2\u1f99\2\u1fa2"+
		"\2\u1fa9\2\u1fb2\2\u1fb6\2\u1fb8\2\u1fbd\2\u1fc0\2\u1fc0\2\u1fc4\2\u1fc6"+
		"\2\u1fc8\2\u1fcd\2\u1fd2\2\u1fd5\2\u1fd8\2\u1fdd\2\u1fe2\2\u1fee\2\u1ff4"+
		"\2\u1ff6\2\u1ff8\2\u1ffd\2\u2104\2\u2104\2\u2109\2\u2109\2\u210c\2\u2115"+
		"\2\u2117\2\u2117\2\u211b\2\u211f\2\u2126\2\u2126\2\u2128\2\u2128\2\u212a"+
		"\2\u212a\2\u212c\2\u212f\2\u2131\2\u2136\2\u213b\2\u213b\2\u213e\2\u2141"+
		"\2\u2147\2\u214b\2\u2150\2\u2150\2\u2185\2\u2186\2\u2c02\2\u2c30\2\u2c32"+
		"\2\u2c60\2\u2c62\2\u2c7d\2\u2c80\2\u2ce6\2\u2ced\2\u2cf0\2\u2cf4\2\u2cf5"+
		"\2\u2d02\2\u2d27\2\u2d29\2\u2d29\2\u2d2f\2\u2d2f\2\ua642\2\ua66f\2\ua682"+
		"\2\ua69d\2\ua724\2\ua771\2\ua773\2\ua789\2\ua78d\2\ua790\2\ua792\2\ua7b0"+
		"\2\ua7b2\2\ua7b9\2\ua7fc\2\ua7fc\2\uab32\2\uab5c\2\uab62\2\uab67\2\uab72"+
		"\2\uabc1\2\ufb02\2\ufb08\2\ufb15\2\ufb19\2\uff23\2\uff3c\2\uff43\2\uff5c"+
		"\2\u0402\3\u0451\3\u04b2\3\u04d5\3\u04da\3\u04fd\3\u0c82\3\u0cb4\3\u0cc2"+
		"\3\u0cf4\3\u18a2\3\u18e1\3\ud402\3\ud456\3\ud458\3\ud49e\3\ud4a0\3\ud4a1"+
		"\3\ud4a4\3\ud4a4\3\ud4a7\3\ud4a8\3\ud4ab\3\ud4ae\3\ud4b0\3\ud4bb\3\ud4bd"+
		"\3\ud4bd\3\ud4bf\3\ud4c5\3\ud4c7\3\ud507\3\ud509\3\ud50c\3\ud50f\3\ud516"+
		"\3\ud518\3\ud51e\3\ud520\3\ud53b\3\ud53d\3\ud540\3\ud542\3\ud546\3\ud548"+
		"\3\ud548\3\ud54c\3\ud552\3\ud554\3\ud6a7\3\ud6aa\3\ud6c2\3\ud6c4\3\ud6dc"+
		"\3\ud6de\3\ud6fc\3\ud6fe\3\ud716\3\ud718\3\ud736\3\ud738\3\ud750\3\ud752"+
		"\3\ud770\3\ud772\3\ud78a\3\ud78c\3\ud7aa\3\ud7ac\3\ud7c4\3\ud7c6\3\ud7cd"+
		"\3\ue902\3\ue945\39\2\62\2;\2\u0662\2\u066b\2\u06f2\2\u06fb\2\u07c2\2"+
		"\u07cb\2\u0968\2\u0971\2\u09e8\2\u09f1\2\u0a68\2\u0a71\2\u0ae8\2\u0af1"+
		"\2\u0b68\2\u0b71\2\u0be8\2\u0bf1\2\u0c68\2\u0c71\2\u0ce8\2\u0cf1\2\u0d68"+
		"\2\u0d71\2\u0de8\2\u0df1\2\u0e52\2\u0e5b\2\u0ed2\2\u0edb\2\u0f22\2\u0f2b"+
		"\2\u1042\2\u104b\2\u1092\2\u109b\2\u17e2\2\u17eb\2\u1812\2\u181b\2\u1948"+
		"\2\u1951\2\u19d2\2\u19db\2\u1a82\2\u1a8b\2\u1a92\2\u1a9b\2\u1b52\2\u1b5b"+
		"\2\u1bb2\2\u1bbb\2\u1c42\2\u1c4b\2\u1c52\2\u1c5b\2\ua622\2\ua62b\2\ua8d2"+
		"\2\ua8db\2\ua902\2\ua90b\2\ua9d2\2\ua9db\2\ua9f2\2\ua9fb\2\uaa52\2\uaa5b"+
		"\2\uabf2\2\uabfb\2\uff12\2\uff1b\2\u04a2\3\u04ab\3\u1068\3\u1071\3\u10f2"+
		"\3\u10fb\3\u1138\3\u1141\3\u11d2\3\u11db\3\u12f2\3\u12fb\3\u1452\3\u145b"+
		"\3\u14d2\3\u14db\3\u1652\3\u165b\3\u16c2\3\u16cb\3\u1732\3\u173b\3\u18e2"+
		"\3\u18eb\3\u1c52\3\u1c5b\3\u1d52\3\u1d5b\3\u6a62\3\u6a6b\3\u6b52\3\u6b5b"+
		"\3\ud7d0\3\ud801\3\ue952\3\ue95b\3b\2\3\3\2\2\2\2\5\3\2\2\2\2\23\3\2\2"+
		"\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\3\35\3\2\2\2\5"+
		"\37\3\2\2\2\7!\3\2\2\2\t#\3\2\2\2\13%\3\2\2\2\r\62\3\2\2\2\17\64\3\2\2"+
		"\2\21E\3\2\2\2\23G\3\2\2\2\25Q\3\2\2\2\27S\3\2\2\2\31U\3\2\2\2\33X\3\2"+
		"\2\2\35\36\7}\2\2\36\4\3\2\2\2\37 \7\177\2\2 \6\3\2\2\2!\"\t\6\2\2\"\b"+
		"\3\2\2\2#$\t\7\2\2$\n\3\2\2\2%&\t\2\2\2&\f\3\2\2\2\')\7/\2\2(*\5\t\5\2"+
		")(\3\2\2\2*+\3\2\2\2+)\3\2\2\2+,\3\2\2\2,\63\3\2\2\2-/\5\t\5\2.-\3\2\2"+
		"\2/\60\3\2\2\2\60.\3\2\2\2\60\61\3\2\2\2\61\63\3\2\2\2\62\'\3\2\2\2\62"+
		".\3\2\2\2\63\16\3\2\2\2\64\65\5\r\7\2\65\67\7\60\2\2\668\5\t\5\2\67\66"+
		"\3\2\2\289\3\2\2\29\67\3\2\2\29:\3\2\2\2:\20\3\2\2\2;=\5\7\4\2<;\3\2\2"+
		"\2=>\3\2\2\2><\3\2\2\2>?\3\2\2\2?F\3\2\2\2@B\5\13\6\2A@\3\2\2\2BC\3\2"+
		"\2\2CA\3\2\2\2CD\3\2\2\2DF\3\2\2\2E<\3\2\2\2EA\3\2\2\2F\22\3\2\2\2GL\5"+
		"\21\t\2HK\5\t\5\2IK\5\21\t\2JH\3\2\2\2JI\3\2\2\2KN\3\2\2\2LJ\3\2\2\2L"+
		"M\3\2\2\2M\24\3\2\2\2NL\3\2\2\2OR\5\r\7\2PR\5\17\b\2QO\3\2\2\2QP\3\2\2"+
		"\2R\26\3\2\2\2ST\t\3\2\2T\30\3\2\2\2UV\t\4\2\2V\32\3\2\2\2WY\t\5\2\2X"+
		"W\3\2\2\2YZ\3\2\2\2ZX\3\2\2\2Z[\3\2\2\2[\\\3\2\2\2\\]\b\16\2\2]\34\3\2"+
		"\2\2\16\2+\60\629>CEJLQZ\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}