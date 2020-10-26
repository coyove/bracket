// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2

//line parser.go.y:27
type yySymType struct {
	yys   int
	token Token
	expr  Node
}

const TDo = 57346
const TLocal = 57347
const TElseIf = 57348
const TThen = 57349
const TEnd = 57350
const TBreak = 57351
const TElse = 57352
const TFor = 57353
const TWhile = 57354
const TFunc = 57355
const TIf = 57356
const TLen = 57357
const TReturn = 57358
const TReturnVoid = 57359
const TYield = 57360
const TYieldVoid = 57361
const TRepeat = 57362
const TUntil = 57363
const TNot = 57364
const TLabel = 57365
const TGoto = 57366
const TOr = 57367
const TAnd = 57368
const TEqeq = 57369
const TNeq = 57370
const TLte = 57371
const TGte = 57372
const TIdent = 57373
const TNumber = 57374
const TString = 57375
const TAddEq = 57376
const TSubEq = 57377
const TMulEq = 57378
const TDivEq = 57379
const TModEq = 57380
const TSquare = 57381
const TDotDot = 57382
const ASSIGN = 57383
const FUNC = 57384
const UNARY = 57385

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TDo",
	"TLocal",
	"TElseIf",
	"TThen",
	"TEnd",
	"TBreak",
	"TElse",
	"TFor",
	"TWhile",
	"TFunc",
	"TIf",
	"TLen",
	"TReturn",
	"TReturnVoid",
	"TYield",
	"TYieldVoid",
	"TRepeat",
	"TUntil",
	"TNot",
	"TLabel",
	"TGoto",
	"TOr",
	"TAnd",
	"TEqeq",
	"TNeq",
	"TLte",
	"TGte",
	"TIdent",
	"TNumber",
	"TString",
	"'{'",
	"'['",
	"'('",
	"'='",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'#'",
	"'.'",
	"'&'",
	"TAddEq",
	"TSubEq",
	"TMulEq",
	"TDivEq",
	"TModEq",
	"TSquare",
	"TDotDot",
	"'T'",
	"ASSIGN",
	"FUNC",
	"UNARY",
	"';'",
	"']'",
	"','",
	"')'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:374

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 27,
	37, 55,
	62, 55,
	-2, 80,
	-1, 97,
	37, 56,
	62, 56,
	-2, 80,
}

const yyPrivate = 57344

const yyLast = 850

var yyAct = [...]int{

	33, 36, 50, 18, 155, 57, 61, 32, 95, 151,
	40, 37, 27, 52, 68, 135, 94, 68, 53, 100,
	41, 34, 35, 44, 56, 28, 41, 59, 110, 60,
	38, 28, 105, 18, 4, 39, 52, 137, 86, 87,
	88, 95, 27, 54, 131, 139, 90, 62, 63, 64,
	65, 66, 92, 80, 81, 82, 98, 93, 91, 18,
	58, 96, 113, 55, 4, 97, 43, 42, 27, 114,
	115, 116, 117, 118, 119, 120, 121, 122, 123, 124,
	125, 126, 127, 128, 129, 130, 77, 79, 80, 81,
	82, 83, 77, 79, 80, 81, 82, 83, 134, 89,
	161, 78, 136, 140, 142, 138, 104, 141, 3, 37,
	99, 144, 47, 143, 84, 49, 6, 145, 41, 34,
	35, 22, 47, 28, 45, 49, 85, 106, 38, 5,
	17, 16, 51, 39, 148, 149, 46, 19, 3, 147,
	18, 48, 21, 18, 102, 18, 2, 18, 158, 27,
	52, 0, 27, 1, 27, 160, 27, 0, 163, 5,
	0, 31, 162, 0, 18, 164, 168, 0, 0, 18,
	167, 18, 169, 27, 172, 18, 0, 173, 27, 0,
	27, 0, 0, 0, 27, 0, 69, 70, 75, 76,
	74, 73, 0, 0, 0, 0, 0, 0, 0, 71,
	72, 77, 79, 80, 81, 82, 83, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 78, 69, 70, 75,
	76, 74, 73, 165, 0, 0, 0, 0, 0, 0,
	71, 72, 77, 79, 80, 81, 82, 83, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 78, 0, 0,
	0, 0, 0, 146, 133, 69, 70, 75, 76, 74,
	73, 0, 0, 0, 0, 0, 0, 0, 71, 72,
	77, 79, 80, 81, 82, 83, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 78, 0, 0, 0, 0,
	0, 132, 133, 69, 70, 75, 76, 74, 73, 0,
	0, 0, 0, 0, 0, 0, 71, 72, 77, 79,
	80, 81, 82, 83, 107, 109, 157, 0, 0, 11,
	156, 25, 23, 78, 26, 0, 15, 14, 9, 10,
	24, 112, 0, 13, 12, 0, 0, 0, 0, 0,
	0, 29, 0, 0, 0, 0, 28, 69, 70, 75,
	76, 74, 73, 0, 0, 0, 0, 0, 0, 0,
	71, 72, 77, 79, 80, 81, 82, 83, 0, 0,
	108, 0, 0, 0, 0, 0, 0, 78, 7, 20,
	0, 0, 67, 11, 154, 25, 23, 30, 26, 0,
	15, 14, 9, 10, 24, 0, 0, 13, 12, 107,
	109, 0, 0, 174, 11, 29, 25, 23, 0, 26,
	28, 15, 14, 9, 10, 24, 0, 0, 13, 12,
	0, 0, 0, 0, 0, 0, 29, 0, 0, 0,
	0, 28, 107, 109, 8, 0, 170, 11, 0, 25,
	23, 0, 26, 0, 15, 14, 9, 10, 24, 0,
	0, 13, 12, 0, 0, 108, 0, 0, 0, 29,
	0, 0, 0, 0, 28, 69, 70, 75, 76, 74,
	73, 0, 0, 0, 0, 0, 0, 0, 71, 72,
	77, 79, 80, 81, 82, 83, 0, 0, 108, 0,
	0, 0, 107, 109, 0, 78, 153, 11, 0, 25,
	23, 159, 26, 0, 15, 14, 9, 10, 24, 0,
	0, 13, 12, 107, 109, 0, 0, 152, 11, 29,
	25, 23, 0, 26, 28, 15, 14, 9, 10, 24,
	0, 0, 13, 12, 0, 0, 0, 0, 0, 0,
	29, 0, 0, 0, 0, 28, 107, 109, 108, 0,
	150, 11, 0, 25, 23, 0, 26, 0, 15, 14,
	9, 10, 24, 0, 0, 13, 12, 107, 109, 108,
	0, 0, 11, 29, 25, 23, 0, 26, 28, 15,
	14, 9, 10, 24, 103, 0, 13, 12, 0, 0,
	0, 0, 0, 0, 29, 0, 0, 0, 0, 28,
	7, 20, 108, 0, 0, 11, 0, 25, 23, 30,
	26, 0, 15, 14, 9, 10, 24, 0, 0, 13,
	12, 107, 109, 108, 0, 0, 11, 29, 25, 23,
	0, 26, 28, 15, 14, 9, 10, 24, 0, 0,
	13, 12, 0, 0, 171, 0, 0, 0, 29, 0,
	0, 0, 0, 28, 0, 0, 8, 0, 0, 0,
	0, 0, 0, 0, 0, 69, 70, 75, 76, 74,
	73, 166, 0, 0, 0, 0, 0, 108, 71, 72,
	77, 79, 80, 81, 82, 83, 0, 0, 0, 69,
	70, 75, 76, 74, 73, 78, 0, 0, 0, 0,
	0, 0, 71, 72, 77, 79, 80, 81, 82, 83,
	111, 0, 0, 0, 0, 0, 0, 0, 0, 78,
	0, 0, 0, 0, 0, 0, 0, 0, 69, 70,
	75, 76, 74, 73, 0, 0, 0, 0, 101, 0,
	0, 71, 72, 77, 79, 80, 81, 82, 83, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 78, 69,
	70, 75, 76, 74, 73, 0, 0, 0, 0, 0,
	0, 0, 71, 72, 77, 79, 80, 81, 82, 83,
	69, 70, 75, 76, 74, 73, 0, 0, 0, 78,
	0, 0, 0, 71, 72, 77, 79, 80, 81, 82,
	83, 70, 75, 76, 74, 73, 0, 0, 0, 0,
	78, 0, 0, 71, 72, 77, 79, 80, 81, 82,
	83, 75, 76, 74, 73, 0, 0, 0, 0, 0,
	78, 0, 71, 72, 77, 79, 80, 81, 82, 83,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 78,
}
var yyPact = [...]int{

	-1000, 596, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 87,
	-1000, -1000, 36, 35, -1000, 87, -1000, -1000, 89, -1000,
	119, -19, 32, 87, -1000, 29, 87, -1000, 87, -2,
	-1000, 374, -45, 755, -1000, -1000, 79, 87, 87, 87,
	-1000, -1000, -1000, 76, -45, 87, 27, -1000, -1000, -11,
	-21, -1000, -1000, 87, -5, -17, 734, 563, -9, 703,
	268, 87, -1000, -1000, -1000, -1000, -1000, -1000, 87, 87,
	87, 87, 87, 87, 87, 87, 87, 87, 87, 87,
	87, 87, 87, 87, 87, 13, -1000, -1000, -1000, -1000,
	230, -2, -1000, -48, 87, 6, -45, -1000, 79, -1000,
	-18, -1000, -1000, 87, -1000, -1000, -1000, -1000, -1000, 5,
	87, -1000, -1000, 755, 755, 775, 794, 46, 46, 46,
	46, 46, 46, 11, 52, 11, -1000, -1000, -1000, 11,
	192, -1000, -2, 87, 87, -1000, -45, -1000, 542, -1000,
	-54, 509, 755, 488, 322, 310, -1000, 87, 440, 755,
	-1000, -1000, -1000, -1000, 87, 92, -1000, 87, 755, -1000,
	161, -1000, 617, 664, -1000, 87, -1000, 428, 640, 310,
	-1000, -1000, -1000, 395, -1000,
}
var yyPgo = [...]int{

	0, 153, 5, 146, 144, 10, 142, 2, 7, 141,
	0, 137, 6, 1, 127, 131, 130, 4, 106, 32,
	121, 116, 110,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 4, 4, 4, 4, 4, 19, 19, 12, 12,
	12, 12, 12, 14, 14, 14, 14, 14, 11, 11,
	11, 15, 15, 15, 15, 16, 17, 17, 17, 20,
	20, 21, 22, 22, 18, 18, 18, 18, 18, 18,
	18, 5, 5, 5, 5, 6, 6, 7, 7, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	13, 13, 13, 13, 8, 8, 9, 9,
}
var yyR2 = [...]int{

	0, 0, 2, 0, 2, 1, 1, 1, 1, 3,
	1, 1, 1, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 4, 3, 3, 6,
	5, 5, 4, 9, 11, 6, 0, 2, 5, 1,
	2, 5, 2, 3, 2, 1, 1, 2, 3, 1,
	2, 1, 4, 6, 3, 1, 3, 1, 3, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	1, 2, 2, 3, 1, 3, 2, 3,
}
var yyChk = [...]int{

	-1000, -1, -3, -18, -19, -14, -21, 4, 60, 18,
	19, 9, 24, 23, 17, 16, -15, -16, -13, -11,
	5, -6, -20, 12, 20, 11, 14, -5, 36, 31,
	13, -1, -8, -10, 32, 33, -13, 22, 41, 46,
	-5, 31, 31, 31, -8, 35, 47, 33, -9, 36,
	-7, 13, 31, 37, 62, 31, -10, -2, 31, -10,
	-10, -12, 49, 50, 51, 52, 53, 8, 62, 25,
	26, 38, 39, 30, 29, 27, 28, 40, 55, 41,
	42, 43, 44, 45, 35, 47, -10, -10, -10, 23,
	-10, 31, 63, -8, 37, 62, -8, -5, -13, -22,
	36, 4, -4, 21, -18, -19, -14, 4, 60, 5,
	37, 7, 63, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, 31, 61, 62, -12, 63, -8, 31, -2, 63,
	-7, -2, -10, -2, -10, -2, 61, -12, -10, -10,
	8, 63, 8, 8, 62, -17, 10, 6, -10, 61,
	-10, 8, -2, -10, 4, 62, 7, -2, -10, -2,
	8, 4, -17, -2, 8,
}
var yyDef = [...]int{

	1, -2, 2, 5, 6, 7, 8, 1, 10, 0,
	45, 46, 0, 0, 49, 0, 16, 17, 23, 24,
	0, 0, 0, 0, 3, 0, 0, -2, 0, 51,
	39, 0, 44, 84, 59, 60, 61, 0, 0, 0,
	80, 51, 47, 0, 50, 0, 0, 81, 82, 0,
	25, 40, 57, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 18, 19, 20, 21, 22, 9, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 77, 78, 79, 48,
	0, 54, 86, 0, 0, 0, 27, -2, 0, 3,
	0, 3, 4, 0, 11, 12, 13, 3, 15, 0,
	0, 3, 83, 28, 85, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 72, 73, 74, 75, 76,
	0, 54, 52, 0, 0, 87, 26, 58, 0, 42,
	0, 0, 32, 0, 0, 36, 52, 0, 0, 30,
	41, 43, 31, 14, 0, 0, 3, 0, 29, 53,
	0, 35, 37, 0, 3, 0, 3, 0, 0, 36,
	33, 3, 38, 0, 34,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 46, 3, 44, 48, 3,
	36, 63, 42, 40, 62, 41, 47, 43, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 60,
	39, 37, 38, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 56, 3, 3, 3, 3, 3,
	3, 35, 3, 61, 45, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 34,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 49, 50, 51, 52, 53, 54, 55, 57,
	58, 59,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:57
		{
			yyVAL.expr = __chain()
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:63
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[2].expr)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:71
		{
			yyVAL.expr = __chain()
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:74
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[2].expr)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:79
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:80
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:81
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:82
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:83
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:84
		{
			yyVAL.expr = emptyNode
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:87
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:88
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:89
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:90
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:91
		{
			yyVAL.expr = emptyNode
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:94
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:95
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:98
		{
			yyVAL.expr = NewSymbol(AAdd).SetPos(yyDollar[1].token.Pos)
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:99
		{
			yyVAL.expr = NewSymbol(ASub).SetPos(yyDollar[1].token.Pos)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:100
		{
			yyVAL.expr = NewSymbol(AMul).SetPos(yyDollar[1].token.Pos)
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:101
		{
			yyVAL.expr = NewSymbol(ADiv).SetPos(yyDollar[1].token.Pos)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:102
		{
			yyVAL.expr = NewSymbol(AMod).SetPos(yyDollar[1].token.Pos)
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:105
		{
			if yyDollar[1].expr.isCallStat() {
				// Single call statement, clear env.V to avoid side effects
				yyVAL.expr = __chain(yyDollar[1].expr, popvClearNode)
			} else {
				yyVAL.expr = yyDollar[1].expr
			}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:113
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:116
		{
			yyVAL.expr = __chain()
			for _, v := range yyDollar[2].expr.Nodes {
				yyVAL.expr = yyVAL.expr.append(__set(v, NewSymbol(ANil)).SetPos(yyDollar[1].token.Pos))
			}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:122
		{
			m, n := len(yyDollar[2].expr.Nodes), len(yyDollar[4].expr.Nodes)
			for i, count := 0, m-n; i < count; i++ {
				if i == count-1 {
					yyDollar[4].expr = yyDollar[4].expr.append(__chain(popvNode, popvClearNode))
				} else {
					yyDollar[4].expr = yyDollar[4].expr.append(popvNode)
				}
			}

			yyVAL.expr = __chain()
			for i, v := range yyDollar[2].expr.Nodes {
				if v.IsSymbolDotDotDot() {
					yyVAL.expr = yyVAL.expr.append(__set(v, __popvAll(i, yyDollar[4].expr.Nodes[i])).SetPos(yyDollar[1].token.Pos))
				} else {
					yyVAL.expr = yyVAL.expr.append(__set(v, yyDollar[4].expr.Nodes[i]).SetPos(yyDollar[1].token.Pos))
				}
			}

			if m == 1 && n == 1 && yyDollar[4].expr.Nodes[0].isCallStat() {
				// Single call statement with single assignment, clear env.V to avoid side effects
				yyVAL.expr = yyVAL.expr.append(popvClearNode)
			}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:146
		{
			nodes := yyDollar[1].expr.Nodes
			m, n := len(nodes), len(yyDollar[3].expr.Nodes)
			for i, count := 0, m-n; i < count; i++ {
				if i == count-1 {
					yyDollar[3].expr = yyDollar[3].expr.append(__chain(popvNode, popvClearNode))
				} else {
					yyDollar[3].expr = yyDollar[3].expr.append(popvNode)
				}
			}

			if head := nodes[0]; len(nodes) == 1 && !nodes[0].IsSymbolDotDotDot() {
				yyVAL.expr = head.moveLoadStore(__move, yyDollar[3].expr.Nodes[0]).SetPos(yyDollar[2].token.Pos)
			} else {
				// a0, ..., an = b0, ..., bn
				yyVAL.expr = __chain()
				names, retaddr := []Node{}, NewComplex(NewSymbol(ARetAddr))
				for i := range nodes {
					names = append(names, randomVarname())
					retaddr = retaddr.append(names[i])
					if nodes[i].IsSymbolDotDotDot() {
						yyVAL.expr = yyVAL.expr.append(__set(names[i], __popvAll(i, yyDollar[3].expr.Nodes[i])).SetPos(yyDollar[2].token.Pos))
					} else {
						yyVAL.expr = yyVAL.expr.append(__set(names[i], yyDollar[3].expr.Nodes[i]).SetPos(yyDollar[2].token.Pos))
					}
				}
				for i, v := range nodes {
					yyVAL.expr = yyVAL.expr.append(v.moveLoadStore(__move, names[i]).SetPos(yyDollar[2].token.Pos))
				}
				yyVAL.expr = yyVAL.expr.append(retaddr)
			}

			if m == 1 && n == 1 && yyDollar[3].expr.Nodes[0].isCallStat() {
				// Single call statement with single assignment, clear env.V to avoid side effects
				yyVAL.expr = __chain(yyVAL.expr, popvClearNode)
			}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:185
		{
			yyVAL.expr = __move(NewSymbolFromToken(yyDollar[1].token), NewComplex(yyDollar[2].expr, NewSymbolFromToken(yyDollar[1].token), yyDollar[3].expr)).SetPos(yyDollar[2].expr.Pos())
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:188
		{
			yyVAL.expr = __store(yyDollar[1].expr, yyDollar[3].expr, NewComplex(yyDollar[5].expr, __load(yyDollar[1].expr, yyDollar[3].expr), yyDollar[6].expr).SetPos(yyDollar[5].expr.Pos()))
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:191
		{
			i := NewString(yyDollar[3].token.Str)
			yyVAL.expr = __store(yyDollar[1].expr, i, NewComplex(yyDollar[4].expr, __load(yyDollar[1].expr, i), yyDollar[5].expr).SetPos(yyDollar[4].expr.Pos()))
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:197
		{
			yyVAL.expr = __loop(__if(yyDollar[2].expr, yyDollar[4].expr, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos)
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:200
		{
			yyVAL.expr = __loop(
				__chain(
					yyDollar[2].expr,
					__if(yyDollar[4].expr, breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			).SetPos(yyDollar[1].token.Pos)
		}
	case 33:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:208
		{
			forVar, forEnd := NewSymbolFromToken(yyDollar[2].token), randomVarname()
			yyVAL.expr = __do(
				__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
				__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
				__loop(
					__if(
						__lessEq(forVar, forEnd),
						__chain(yyDollar[8].expr, __inc(forVar, oneNode).SetPos(yyDollar[1].token.Pos)),
						breakNode,
					).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			)
		}
	case 34:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:222
		{
			forVar, forEnd := NewSymbolFromToken(yyDollar[2].token), randomVarname()
			if yyDollar[8].expr.IsNumber() { // step is a static number, easy case
				var cond Node
				if yyDollar[8].expr.IsNegativeNumber() {
					cond = __lessEq(forEnd, forVar)
				} else {
					cond = __lessEq(forVar, forEnd)
				}
				yyVAL.expr = __do(
					__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
					__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
					__loop(
						__chain(
							__if(
								cond,
								__chain(yyDollar[10].expr, __inc(forVar, yyDollar[8].expr)),
								breakNode,
							).SetPos(yyDollar[1].token.Pos),
						),
					).SetPos(yyDollar[1].token.Pos),
				)
			} else {
				forStep := randomVarname()
				yyVAL.expr = __do(
					__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
					__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
					__set(forStep, yyDollar[8].expr).SetPos(yyDollar[1].token.Pos),
					__loop(
						__chain(
							__if(
								__less(zeroNode, forStep).SetPos(yyDollar[1].token.Pos),
								// +step
								__if(__less(forEnd, forVar), breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
								// -step
								__if(__less(forVar, forEnd), breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
							).SetPos(yyDollar[1].token.Pos),
							yyDollar[10].expr,
							__inc(forVar, forStep),
						),
					).SetPos(yyDollar[1].token.Pos),
				)
			}

		}
	case 35:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:269
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:274
		{
			yyVAL.expr = NewComplex()
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:277
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:280
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:285
		{
			yyVAL.expr = NewSymbol(AMove).SetPos(yyDollar[1].token.Pos)
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:286
		{
			yyVAL.expr = NewSymbol(ASet).SetPos(yyDollar[1].token.Pos)
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:289
		{
			__findTailCall(yyDollar[4].expr.Nodes)
			funcname := NewSymbolFromToken(yyDollar[2].token)
			x := __move
			if yyDollar[1].expr.SymbolValue() == ASet {
				x = __set
			}
			yyVAL.expr = __chain(
				x(funcname, NewSymbol(ANil)).SetPos(yyDollar[1].expr.Pos()),
				__move(funcname, __func(funcname, yyDollar[3].expr, yyDollar[4].expr).SetPos(yyDollar[1].expr.Pos())).SetPos(yyDollar[1].expr.Pos()),
			)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:303
		{
			yyVAL.expr = NewComplex()
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:304
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:307
		{
			yyVAL.expr = NewComplex(NewSymbol(AYield), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:308
		{
			yyVAL.expr = NewComplex(NewSymbol(AYield), emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:309
		{
			yyVAL.expr = NewComplex(NewSymbol(ABreak)).SetPos(yyDollar[1].token.Pos)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:310
		{
			yyVAL.expr = NewComplex(NewSymbol(AGoto), NewSymbolFromToken(yyDollar[2].token)).SetPos(yyDollar[1].token.Pos)
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:311
		{
			yyVAL.expr = NewComplex(NewSymbol(ALabel), NewSymbolFromToken(yyDollar[2].token))
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:312
		{
			yyVAL.expr = NewComplex(NewSymbol(AReturn), emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:313
		{
			if len(yyDollar[2].expr.Nodes) == 1 {
				x := yyDollar[2].expr.Nodes[0]
				if len(x.Nodes) == 3 && x.Nodes[0].SymbolValue() == ACall {
					x.Nodes[0].strSym = ATailCall
				}
			}
			yyVAL.expr = NewComplex(NewSymbol(AReturn), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:324
		{
			yyVAL.expr = NewSymbolFromToken(yyDollar[1].token)
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:325
		{
			yyVAL.expr = __load(yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:326
		{
			yyVAL.expr = NewComplex(NewSymbol(ASlice), yyDollar[1].expr, yyDollar[3].expr, yyDollar[5].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:327
		{
			yyVAL.expr = __load(yyDollar[1].expr, NewString(yyDollar[3].token.Str)).SetPos(yyDollar[2].token.Pos)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:330
		{
			yyVAL.expr = NewComplex(yyDollar[1].expr)
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:331
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[3].expr)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:334
		{
			yyVAL.expr = NewComplex(NewSymbolFromToken(yyDollar[1].token))
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:335
		{
			yyVAL.expr = yyDollar[1].expr.append(NewSymbolFromToken(yyDollar[3].token))
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:338
		{
			yyVAL.expr = NewNumberFromString(yyDollar[1].token.Str)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:339
		{
			yyVAL.expr = NewString(yyDollar[1].token.Str)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:340
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:341
		{
			yyVAL.expr = NewComplex(NewSymbol(AOr), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:342
		{
			yyVAL.expr = NewComplex(NewSymbol(AAnd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:343
		{
			yyVAL.expr = NewComplex(NewSymbol(ALess), yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:344
		{
			yyVAL.expr = NewComplex(NewSymbol(ALess), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:345
		{
			yyVAL.expr = NewComplex(NewSymbol(ALessEq), yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:346
		{
			yyVAL.expr = NewComplex(NewSymbol(ALessEq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:347
		{
			yyVAL.expr = NewComplex(NewSymbol(AEq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:348
		{
			yyVAL.expr = NewComplex(NewSymbol(ANeq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:349
		{
			yyVAL.expr = NewComplex(NewSymbol(AAdd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:350
		{
			yyVAL.expr = NewComplex(NewSymbol(AConcat), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:351
		{
			yyVAL.expr = NewComplex(NewSymbol(ASub), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:352
		{
			yyVAL.expr = NewComplex(NewSymbol(AMul), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:353
		{
			yyVAL.expr = NewComplex(NewSymbol(ADiv), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:354
		{
			yyVAL.expr = NewComplex(NewSymbol(AMod), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:355
		{
			yyVAL.expr = NewComplex(NewSymbol(APow), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:356
		{
			yyVAL.expr = NewComplex(NewSymbol(ANot), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:357
		{
			yyVAL.expr = NewComplex(NewSymbol(ASub), zeroNode, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:358
		{
			yyVAL.expr = NewComplex(NewSymbol(ALen), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:361
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:362
		{
			yyVAL.expr = __call(yyDollar[1].expr, NewComplex(NewString(yyDollar[2].token.Str))).SetPos(yyDollar[1].expr.Pos())
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:363
		{
			yyVAL.expr = __call(yyDollar[1].expr, yyDollar[2].expr).SetPos(yyDollar[1].expr.Pos())
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:364
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:367
		{
			yyVAL.expr = NewComplex(yyDollar[1].expr)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:368
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[3].expr)
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:371
		{
			yyVAL.expr = NewComplex()
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:372
		{
			yyVAL.expr = yyDollar[2].expr
		}
	}
	goto yystack /* stack new state and value */
}
