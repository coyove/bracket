//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2
import (
	"bytes"
	"io/ioutil"
	"path/filepath"
)

//line parser.go.y:36
type yySymType struct {
	yys   int
	token Token
	expr  *Node
}

const TAssert = 57346
const TBreak = 57347
const TContinue = 57348
const TElse = 57349
const TFor = 57350
const TFunc = 57351
const TIf = 57352
const TNil = 57353
const TReturn = 57354
const TRequire = 57355
const TVar = 57356
const TYield = 57357
const TEqeq = 57358
const TNeq = 57359
const TLsh = 57360
const TRsh = 57361
const TLte = 57362
const TGte = 57363
const TIdent = 57364
const TNumber = 57365
const TString = 57366
const TOr = 57367
const TAnd = 57368
const UNARY = 57369

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TAssert",
	"TBreak",
	"TContinue",
	"TElse",
	"TFor",
	"TFunc",
	"TIf",
	"TNil",
	"TReturn",
	"TRequire",
	"TVar",
	"TYield",
	"TEqeq",
	"TNeq",
	"TLsh",
	"TRsh",
	"TLte",
	"TGte",
	"TIdent",
	"TNumber",
	"TString",
	"'{'",
	"'('",
	"TOr",
	"TAnd",
	"'|'",
	"'&'",
	"'^'",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"'~'",
	"'}'",
	"';'",
	"'='",
	"')'",
	"'['",
	"']'",
	"'.'",
	"','",
	"':'",
	"'!'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:538

func TokenName(c int) string {
	if c >= TAnd && c-TAnd < len(yyToknames) {
		if yyToknames[c-TAnd] != "" {
			return yyToknames[c-TAnd]
		}
	}
	return string([]byte{byte(c)})
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 100,
	42, 14,
	-2, 43,
}

const yyPrivate = 57344

const yyLast = 664

var yyAct = [...]int{

	64, 17, 126, 1, 27, 63, 39, 46, 9, 8,
	143, 34, 49, 33, 58, 50, 107, 106, 132, 57,
	106, 53, 48, 44, 58, 35, 131, 164, 86, 4,
	59, 60, 30, 159, 31, 28, 19, 36, 47, 45,
	20, 94, 83, 84, 85, 29, 92, 138, 26, 41,
	25, 91, 130, 87, 42, 96, 101, 128, 99, 93,
	55, 100, 8, 103, 43, 54, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	122, 123, 124, 125, 98, 102, 73, 74, 75, 76,
	77, 23, 129, 127, 134, 75, 76, 77, 137, 51,
	104, 61, 56, 157, 28, 140, 127, 141, 71, 72,
	79, 80, 70, 69, 21, 89, 38, 37, 32, 65,
	66, 81, 82, 78, 67, 68, 73, 74, 75, 76,
	77, 142, 144, 141, 145, 52, 6, 148, 147, 150,
	18, 153, 151, 5, 16, 40, 90, 2, 0, 9,
	8, 0, 79, 80, 158, 0, 9, 8, 0, 162,
	161, 0, 163, 9, 8, 0, 0, 165, 73, 74,
	75, 76, 77, 0, 0, 0, 0, 0, 0, 155,
	71, 72, 79, 80, 70, 69, 160, 0, 0, 0,
	0, 65, 66, 81, 82, 78, 67, 68, 73, 74,
	75, 76, 77, 0, 0, 0, 0, 0, 146, 0,
	0, 0, 136, 71, 72, 79, 80, 70, 69, 0,
	0, 0, 0, 0, 65, 66, 81, 82, 78, 67,
	68, 73, 74, 75, 76, 77, 0, 0, 0, 0,
	0, 135, 0, 0, 0, 136, 71, 72, 79, 80,
	70, 69, 0, 0, 0, 0, 0, 65, 66, 81,
	82, 78, 67, 68, 73, 74, 75, 76, 77, 71,
	72, 79, 80, 70, 69, 0, 0, 0, 0, 133,
	65, 66, 81, 82, 78, 67, 68, 73, 74, 75,
	76, 77, 48, 44, 0, 35, 0, 0, 7, 105,
	48, 44, 0, 35, 0, 0, 19, 36, 47, 45,
	20, 0, 0, 0, 19, 36, 47, 45, 20, 41,
	0, 0, 0, 0, 42, 0, 97, 41, 99, 0,
	0, 0, 42, 88, 43, 71, 72, 79, 80, 70,
	69, 0, 43, 0, 0, 0, 65, 66, 81, 82,
	78, 67, 68, 73, 74, 75, 76, 77, 0, 0,
	0, 0, 0, 154, 71, 72, 79, 80, 70, 69,
	0, 0, 0, 0, 0, 65, 66, 81, 82, 78,
	67, 68, 73, 74, 75, 76, 77, 0, 0, 0,
	0, 0, 139, 71, 72, 79, 80, 70, 69, 0,
	0, 0, 0, 0, 65, 66, 81, 82, 78, 67,
	68, 73, 74, 75, 76, 77, 48, 44, 0, 35,
	0, 95, 48, 44, 0, 35, 0, 0, 0, 0,
	19, 36, 47, 45, 20, 0, 19, 36, 47, 45,
	20, 48, 0, 41, 0, 0, 0, 7, 42, 41,
	0, 0, 62, 0, 42, 19, 0, 0, 43, 20,
	0, 0, 0, 0, 43, 71, 72, 79, 80, 70,
	69, 0, 0, 0, 0, 0, 65, 66, 81, 82,
	78, 67, 68, 73, 74, 75, 76, 77, 0, 0,
	0, 156, 71, 72, 79, 80, 70, 69, 0, 0,
	0, 0, 0, 65, 66, 81, 82, 78, 67, 68,
	73, 74, 75, 76, 77, 0, 0, 0, 149, 71,
	72, 79, 80, 70, 69, 0, 0, 0, 0, 0,
	65, 66, 81, 82, 78, 67, 68, 73, 74, 75,
	76, 77, 14, 12, 13, 0, 22, 24, 23, 0,
	10, 15, 7, 11, 0, 0, 0, 0, 0, 0,
	19, 0, 0, 0, 20, 0, 0, 0, 0, 71,
	72, 79, 80, 70, 69, 0, 0, 0, 0, 152,
	3, 66, 81, 82, 78, 67, 68, 73, 74, 75,
	76, 77, 14, 12, 13, 0, 22, 24, 23, 0,
	10, 15, 7, 11, 0, 0, 0, 0, 0, 0,
	19, 0, 0, 0, 20, 0, 0, 0, 71, 72,
	79, 80, 70, 69, 0, 0, 0, 0, 0, 0,
	3, 81, 82, 78, 67, 68, 73, 74, 75, 76,
	77, 71, 72, 79, 80, 70, 69, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 73,
	74, 75, 76, 77,
}
var yyPact = [...]int{

	-1000, 588, -1000, -1000, 8, 6, -1000, 82, 2, -13,
	414, 414, -1000, -1000, 414, 75, -1000, -1000, -1000, -1000,
	414, -1000, 39, 34, 80, -1000, -1000, -24, -1000, 414,
	414, 79, -1000, 408, 503, -1000, -1000, -1000, -1000, -13,
	-1000, 414, 414, 414, 27, 292, -1000, -1000, 33, 503,
	503, -1000, -3, 377, 284, 414, 27, 414, 78, 503,
	253, -1000, -1000, -28, 503, 414, 414, 414, 414, 414,
	414, 414, 414, 414, 414, 414, 414, 414, 414, 414,
	414, 414, 414, -1000, -1000, -1000, 68, 13, -1000, 11,
	-22, -30, 230, 14, -1000, -1000, 197, 414, 5, -1000,
	-13, 348, 68, -31, -1000, -1000, 414, -1000, 553, 602,
	134, 134, 134, 134, 134, 134, 59, 59, -1000, -1000,
	-1000, 625, 52, 52, 625, 625, -1000, -1000, -1000, -34,
	-1000, 414, 414, 414, 164, 68, 414, 476, 414, 68,
	-1000, 503, 538, -1000, 92, 503, -1000, -1000, 319, 433,
	449, 96, -1000, 414, -1000, -11, 433, 81, 503, 68,
	-17, -1000, -1000, -1000, 68, -1000,
}
var yyPgo = [...]int{

	0, 3, 2, 147, 7, 4, 5, 146, 0, 145,
	6, 29, 144, 1, 143, 140, 136, 114, 118, 117,
	28, 116, 115,
}
var yyR1 = [...]int{

	0, 1, 1, 2, 3, 3, 3, 3, 16, 16,
	16, 11, 11, 11, 11, 12, 12, 12, 13, 13,
	13, 15, 14, 14, 14, 14, 14, 14, 14, 14,
	4, 4, 4, 5, 5, 6, 6, 7, 7, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 9, 10, 10, 10,
	10, 17, 17, 17, 17, 18, 18, 19, 20, 20,
	21, 21, 22, 22, 22, 22,
}
var yyR2 = [...]int{

	0, 0, 2, 3, 1, 2, 2, 1, 1, 1,
	1, 4, 2, 3, 1, 5, 8, 9, 5, 7,
	7, 4, 1, 2, 1, 2, 1, 1, 2, 2,
	1, 4, 3, 1, 3, 1, 3, 3, 5, 1,
	1, 1, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 2, 2, 2, 1, 1, 3, 1,
	3, 3, 4, 6, 2, 2, 3, 3, 2, 3,
	2, 3, 1, 2, 1, 2,
}
var yyChk = [...]int{

	-1000, -1, -3, 42, -11, -14, -16, 14, -4, -10,
	12, 15, 5, 6, 4, 13, -12, -13, -15, 22,
	26, -17, 8, 10, 9, 42, 42, -5, 22, 43,
	45, 47, -18, 26, -8, 11, 23, -19, -21, -10,
	-9, 35, 40, 50, 9, 25, -4, 24, 8, -8,
	-8, 24, -17, -8, 26, 26, 22, 43, 48, -8,
	-8, 22, 44, -6, -8, 27, 28, 32, 33, 21,
	20, 16, 17, 34, 35, 36, 37, 38, 31, 18,
	19, 29, 30, -8, -8, -8, -20, 26, 41, -22,
	-7, -6, -8, 26, 44, 44, -8, 42, -11, 44,
	-10, -8, -20, -6, 22, 46, 48, 44, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -2, 25, 44, -5,
	41, 48, 48, 49, -8, 44, 48, -8, 42, 44,
	-2, -8, -1, 44, -8, -8, 44, -2, -8, 42,
	-8, -2, 41, 49, 44, -11, 42, 7, -8, 44,
	-11, -2, -13, -2, 44, -2,
}
var yyDef = [...]int{

	1, -2, 2, 4, 0, 0, 7, 0, 67, 14,
	22, 24, 26, 27, 0, 0, 8, 9, 10, 30,
	0, 69, 0, 0, 0, 5, 6, 12, 33, 0,
	0, 0, 74, 0, 23, 39, 40, 41, 42, 43,
	44, 0, 0, 0, 0, 0, 67, 66, 0, 25,
	28, 29, 69, 0, 0, 0, 0, 0, 0, 13,
	0, 32, 75, 0, 35, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 63, 64, 65, 0, 0, 80, 0,
	82, 84, 35, 0, 68, 70, 0, 0, 0, 71,
	-2, 0, 0, 11, 34, 31, 0, 76, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 56,
	57, 58, 59, 60, 61, 62, 77, 1, 78, 0,
	81, 83, 85, 0, 0, 72, 0, 0, 0, 0,
	21, 36, 0, 79, 0, 37, 72, 15, 0, 0,
	0, 18, 3, 0, 73, 0, 0, 0, 38, 0,
	0, 19, 20, 16, 0, 17,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 50, 3, 3, 3, 38, 30, 3,
	26, 44, 36, 34, 48, 35, 47, 37, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 49, 42,
	33, 43, 32, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 45, 3, 46, 31, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 25, 29, 41, 40,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 27, 28, 39,
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
		//line parser.go.y:61
		{
			yyVAL.expr = NewCompoundNode("chain")
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:67
		{
			yyDollar[1].expr.Compound = append(yyDollar[1].expr.Compound, yyDollar[2].expr)
			yyVAL.expr = yyDollar[1].expr
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:76
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:81
		{
			yyVAL.expr = NewCompoundNode()
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:84
		{
			if yyDollar[1].expr.isIsolatedDupCall() {
				yyDollar[1].expr.Compound[2].Compound[0] = NewNumberNode("0")
			}
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:90
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:93
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:98
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:101
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:104
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 11:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:109
		{
			yyVAL.expr = NewCompoundNode("chain")
			for i, name := range yyDollar[2].expr.Compound {
				var e *Node
				if i < len(yyDollar[4].expr.Compound) {
					e = yyDollar[4].expr.Compound[i]
				} else {
					e = yyDollar[4].expr.Compound[len(yyDollar[4].expr.Compound)-1]
				}
				c := NewCompoundNode("set", name, e)
				name.Pos, e.Pos = yyDollar[1].token.Pos, yyDollar[1].token.Pos
				c.Compound[0].Pos = yyDollar[1].token.Pos
				yyVAL.expr.Compound = append(yyVAL.expr.Compound, c)
			}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:124
		{
			yyVAL.expr = NewCompoundNode("chain")
			for _, name := range yyDollar[2].expr.Compound {
				c := NewCompoundNode("set", name, NewNilNode())
				c.Compound[0].Pos = yyDollar[1].token.Pos
				yyVAL.expr.Compound = append(yyVAL.expr.Compound, c)
			}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:132
		{
			yyVAL.expr = NewCompoundNode("move", yyDollar[1].expr, yyDollar[3].expr)
			if len(yyDollar[1].expr.Compound) > 0 {
				if c, _ := yyDollar[1].expr.Compound[0].Value.(string); c == "load" {
					yyVAL.expr = NewCompoundNode("store", yyDollar[1].expr.Compound[1], yyDollar[1].expr.Compound[2], yyDollar[3].expr)
				}
			}
			if c, _ := yyDollar[1].expr.Value.(string); c != "" && yyDollar[1].expr.Type == NTAtom {
				if a, b, s := yyDollar[3].expr.isSimpleAddSub(); a == c {
					yyDollar[3].expr.Compound[2].Value = yyDollar[3].expr.Compound[2].Value.(float64) * s
					yyVAL.expr = NewCompoundNode("inc", yyDollar[1].expr, yyDollar[3].expr.Compound[2])
					yyVAL.expr.Compound[1].Pos = yyDollar[1].expr.Pos
				} else if b == c {
					yyDollar[3].expr.Compound[1].Value = yyDollar[3].expr.Compound[1].Value.(float64) * s
					yyVAL.expr = NewCompoundNode("inc", yyDollar[1].expr, yyDollar[3].expr.Compound[1])
					yyVAL.expr.Compound[1].Pos = yyDollar[1].expr.Pos
				}
			}
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:152
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 15:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:157
		{
			yyVAL.expr = NewCompoundNode("for", yyDollar[3].expr, NewCompoundNode(), yyDollar[5].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.go.y:161
		{
			yyVAL.expr = NewCompoundNode("for", yyDollar[4].expr, NewCompoundNode("chain", yyDollar[6].expr), yyDollar[8].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 17:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line parser.go.y:165
		{
			yyVAL.expr = NewCompoundNode("chain",
				yyDollar[3].expr,
				NewCompoundNode("for", yyDollar[5].expr, NewCompoundNode("chain", yyDollar[7].expr), yyDollar[9].expr))
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 18:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:173
		{
			yyVAL.expr = NewCompoundNode("if", yyDollar[3].expr, yyDollar[5].expr, NewCompoundNode())
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:176
		{
			yyVAL.expr = NewCompoundNode("if", yyDollar[3].expr, yyDollar[5].expr, yyDollar[7].expr)
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.go.y:179
		{
			yyVAL.expr = NewCompoundNode("if", yyDollar[3].expr, yyDollar[5].expr, NewCompoundNode("chain", yyDollar[7].expr))
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:184
		{
			funcname := NewAtomNode(yyDollar[2].token)
			yyVAL.expr = NewCompoundNode(
				"chain",
				NewCompoundNode("set", funcname, NewNilNode()),
				NewCompoundNode("move", funcname, NewCompoundNode("lambda", yyDollar[3].expr, yyDollar[4].expr)))
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:193
		{
			yyVAL.expr = NewCompoundNode("ret")
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:197
		{
			if yyDollar[2].expr.isIsolatedDupCall() {
				if h, _ := yyDollar[2].expr.Compound[2].Compound[2].Value.(float64); h == 1 {
					yyDollar[2].expr.Compound[2].Compound[2] = NewNumberNode("2")
				}
			}
			yyVAL.expr = NewCompoundNode("ret", yyDollar[2].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:206
		{
			yyVAL.expr = NewCompoundNode("yield")
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:210
		{
			yyVAL.expr = NewCompoundNode("yield", yyDollar[2].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:214
		{
			yyVAL.expr = NewCompoundNode("break")
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:218
		{
			yyVAL.expr = NewCompoundNode("continue")
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:222
		{
			yyVAL.expr = NewCompoundNode("assert", yyDollar[2].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[2].expr.Pos
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:226
		{
			path := filepath.Dir(yyDollar[1].token.Pos.Source)
			path = filepath.Join(path, yyDollar[2].token.Str)
			filename := filepath.Base(yyDollar[2].token.Str)
			filename = filename[:len(filename)-len(filepath.Ext(filename))]

			code, err := ioutil.ReadFile(path)
			if err != nil {
				yylex.(*Lexer).Error(err.Error())
			}
			n, err := Parse(bytes.NewReader(code), path)
			if err != nil {
				yylex.(*Lexer).Error(err.Error())
			}

			// now the required code is loaded, for naming scope we will wrap them into a closure
			cls := NewCompoundNode("lambda", NewCompoundNode(), n)
			call := NewCompoundNode("call", cls, NewCompoundNode())
			yyVAL.expr = NewCompoundNode("set", filename, call)
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:248
		{
			yyVAL.expr = NewAtomNode(yyDollar[1].token)
		}
	case 31:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:251
		{
			yyVAL.expr = NewCompoundNode("load", yyDollar[1].expr, yyDollar[3].expr)
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:254
		{
			yyVAL.expr = NewCompoundNode("load", yyDollar[1].expr, NewStringNode(yyDollar[3].token.Str))
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:259
		{
			yyVAL.expr = NewCompoundNode(yyDollar[1].token.Str)
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:262
		{
			yyDollar[1].expr.Compound = append(yyDollar[1].expr.Compound, NewAtomNode(yyDollar[3].token))
			yyVAL.expr = yyDollar[1].expr
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:268
		{
			yyVAL.expr = NewCompoundNode(yyDollar[1].expr)
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:271
		{
			yyDollar[1].expr.Compound = append(yyDollar[1].expr.Compound, yyDollar[3].expr)
			yyVAL.expr = yyDollar[1].expr
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:277
		{
			yyVAL.expr = NewCompoundNode(yyDollar[1].expr, yyDollar[3].expr)
		}
	case 38:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.go.y:280
		{
			yyDollar[1].expr.Compound = append(yyDollar[1].expr.Compound, yyDollar[3].expr, yyDollar[5].expr)
			yyVAL.expr = yyDollar[1].expr
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:286
		{
			yyVAL.expr = NewNilNode()
			yyVAL.expr.Pos = yyDollar[1].token.Pos
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:290
		{
			yyVAL.expr = NewNumberNode(yyDollar[1].token.Str)
			yyVAL.expr.Pos = yyDollar[1].token.Pos
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:294
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:297
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:300
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:303
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:306
		{
			yyVAL.expr = NewCompoundNode("or", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:310
		{
			yyVAL.expr = NewCompoundNode("and", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:314
		{
			yyVAL.expr = NewCompoundNode("<", yyDollar[3].expr, yyDollar[1].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:318
		{
			yyVAL.expr = NewCompoundNode("<", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:322
		{
			yyVAL.expr = NewCompoundNode("<=", yyDollar[3].expr, yyDollar[1].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:326
		{
			yyVAL.expr = NewCompoundNode("<=", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:330
		{
			yyVAL.expr = NewCompoundNode("==", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:334
		{
			yyVAL.expr = NewCompoundNode("!=", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:338
		{
			yyVAL.expr = NewCompoundNode("+", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:342
		{
			yyVAL.expr = NewCompoundNode("-", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:346
		{
			yyVAL.expr = NewCompoundNode("*", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:350
		{
			yyVAL.expr = NewCompoundNode("/", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:354
		{
			yyVAL.expr = NewCompoundNode("%", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:358
		{
			yyVAL.expr = NewCompoundNode("^", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:362
		{
			yyVAL.expr = NewCompoundNode("<<", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:366
		{
			yyVAL.expr = NewCompoundNode(">>", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:370
		{
			yyVAL.expr = NewCompoundNode("|", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:374
		{
			yyVAL.expr = NewCompoundNode("&", yyDollar[1].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:378
		{
			yyVAL.expr = NewCompoundNode("-", NewNumberNode("0"), yyDollar[2].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[2].expr.Pos
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:382
		{
			yyVAL.expr = NewCompoundNode("bitnot", yyDollar[2].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[2].expr.Pos
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:386
		{
			yyVAL.expr = NewCompoundNode("!", yyDollar[2].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[2].expr.Pos
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:392
		{
			yyVAL.expr = NewStringNode(yyDollar[1].token.Str)
			yyVAL.expr.Pos = yyDollar[1].token.Pos
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:398
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:401
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:404
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:407
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:412
		{
			yyVAL.expr = NewCompoundNode("call", "foreach", NewCompoundNode(NewNumberNode("1"), NewNumberNode("1"), NewNumberNode("1")))
		}
	case 72:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.go.y:415
		{
			yyVAL.expr = NewCompoundNode("call", "foreach", NewCompoundNode(NewNumberNode("1"), yyDollar[3].expr, NewNumberNode("0")))
		}
	case 73:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.go.y:418
		{
			if yyDollar[5].expr.Type != NTCompound && yyDollar[5].expr.Type != NTAtom {
				yylex.(*Lexer).Error("the second argument of 'for' must be a closure")
			}
			yyVAL.expr = NewCompoundNode("call", "foreach", NewCompoundNode(NewNumberNode("1"), yyDollar[3].expr, yyDollar[5].expr))
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:424
		{
			switch c, _ := yyDollar[1].expr.Value.(string); c {
			case "typeof":
				switch len(yyDollar[2].expr.Compound) {
				case 0:
					yylex.(*Lexer).Error("typeof takes at least 1 argument")
				case 1:
					yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, NewCompoundNode(yyDollar[2].expr.Compound[0], NewNumberNode("255")))
				default:
					switch x, _ := yyDollar[2].expr.Compound[1].Value.(string); x {
					case "nil":
						yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, NewCompoundNode(yyDollar[2].expr.Compound[0], NewNumberNode("0")))
					case "number":
						yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, NewCompoundNode(yyDollar[2].expr.Compound[0], NewNumberNode("1")))
					case "string":
						yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, NewCompoundNode(yyDollar[2].expr.Compound[0], NewNumberNode("2")))
					case "map":
						yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, NewCompoundNode(yyDollar[2].expr.Compound[0], NewNumberNode("3")))
					case "closure":
						yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, NewCompoundNode(yyDollar[2].expr.Compound[0], NewNumberNode("4")))
					case "generic":
						yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, NewCompoundNode(yyDollar[2].expr.Compound[0], NewNumberNode("5")))
					default:
						yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, NewCompoundNode(yyDollar[2].expr.Compound[0], yyDollar[2].expr.Compound[1]))
					}
				}
			case "addressof":
				if len(yyDollar[2].expr.Compound) != 1 {
					yylex.(*Lexer).Error("addressof takes 1 argument")
				}
				if yyDollar[2].expr.Compound[0].Type != NTAtom {
					yylex.(*Lexer).Error("addressof can only get the address of a variable")
				}
				yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, yyDollar[2].expr)
			case "len":
				switch len(yyDollar[2].expr.Compound) {
				case 0:
					yylex.(*Lexer).Error("len takes 1 argument")
				default:
					yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, yyDollar[2].expr)
				}
			default:
				yyVAL.expr = NewCompoundNode("call", yyDollar[1].expr, yyDollar[2].expr)
			}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:471
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.expr = NewCompoundNode()
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:477
		{
			if yylex.(*Lexer).PNewLine {
				yylex.(*Lexer).TokenError(yyDollar[1].token, "ambiguous syntax (function call x new statement)")
			}
			yyVAL.expr = yyDollar[2].expr
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:485
		{
			yyVAL.expr = NewCompoundNode("lambda", yyDollar[2].expr, yyDollar[3].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:491
		{
			yyVAL.expr = NewCompoundNode()
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:494
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:499
		{
			yyVAL.expr = NewCompoundNode("map", NewCompoundNode())
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.go.y:503
		{
			yyVAL.expr = yyDollar[2].expr
			yyVAL.expr.Compound[0].Pos = yyDollar[1].token.Pos
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:509
		{
			yyVAL.expr = NewCompoundNode("map", yyDollar[1].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:513
		{
			yyVAL.expr = NewCompoundNode("map", yyDollar[1].expr)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.go.y:517
		{
			table := NewCompoundNode()
			for i, v := range yyDollar[1].expr.Compound {
				table.Compound = append(table.Compound,
					&Node{Type: NTNumber, Value: float64(i)},
					v)
			}
			yyVAL.expr = NewCompoundNode("map", table)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.go.y:527
		{
			table := NewCompoundNode()
			for i, v := range yyDollar[1].expr.Compound {
				table.Compound = append(table.Compound,
					&Node{Type: NTNumber, Value: float64(i)},
					v)
			}
			yyVAL.expr = NewCompoundNode("map", table)
			yyVAL.expr.Compound[0].Pos = yyDollar[1].expr.Pos
		}
	}
	goto yystack /* stack new state and value */
}
