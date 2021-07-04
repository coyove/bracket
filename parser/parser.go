// Code generated by goyacc -o parser.go parser.go.y. DO NOT EDIT.

//line parser.go.y:2
package parser

import __yyfmt__ "fmt"

//line parser.go.y:2

//line parser.go.y:25
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
const TReturn = 57357
const TReturnVoid = 57358
const TRepeat = 57359
const TUntil = 57360
const TNot = 57361
const TLabel = 57362
const TGoto = 57363
const TIn = 57364
const TLsh = 57365
const TRsh = 57366
const TURsh = 57367
const TOr = 57368
const TAnd = 57369
const TEqeq = 57370
const TNeq = 57371
const TLte = 57372
const TGte = 57373
const TIdent = 57374
const TNumber = 57375
const TString = 57376
const TIDiv = 57377
const ASSIGN = 57378
const FUNC = 57379
const UNARY = 57380

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
	"TReturn",
	"TReturnVoid",
	"TRepeat",
	"TUntil",
	"TNot",
	"TLabel",
	"TGoto",
	"TIn",
	"TLsh",
	"TRsh",
	"TURsh",
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
	"'|'",
	"'~'",
	"TIDiv",
	"'T'",
	"ASSIGN",
	"FUNC",
	"UNARY",
	"';'",
	"'}'",
	"','",
	"')'",
	"']'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.go.y:343

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 25,
	38, 48,
	58, 48,
	59, 48,
	-2, 82,
	-1, 104,
	38, 49,
	58, 49,
	59, 49,
	-2, 82,
}

const yyPrivate = 57344

const yyLast = 1260

var yyAct = [...]int{
	89, 31, 175, 16, 134, 87, 47, 88, 101, 202,
	45, 101, 171, 196, 30, 47, 166, 164, 146, 101,
	137, 51, 53, 105, 49, 56, 82, 83, 84, 16,
	147, 142, 100, 85, 187, 135, 159, 92, 93, 94,
	48, 95, 157, 149, 73, 74, 75, 76, 78, 106,
	98, 51, 99, 101, 103, 77, 16, 102, 206, 180,
	107, 49, 178, 167, 148, 27, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 131, 132, 138, 39, 54, 25,
	110, 172, 139, 136, 73, 74, 75, 76, 78, 211,
	75, 76, 78, 141, 143, 77, 144, 25, 37, 77,
	153, 154, 201, 186, 16, 25, 42, 150, 40, 44,
	170, 90, 33, 34, 35, 91, 32, 151, 112, 160,
	41, 38, 47, 161, 145, 46, 158, 104, 26, 162,
	36, 47, 25, 158, 96, 55, 52, 86, 158, 168,
	165, 29, 28, 64, 16, 190, 60, 16, 4, 59,
	18, 3, 179, 61, 6, 5, 15, 14, 182, 43,
	181, 16, 57, 2, 152, 189, 37, 155, 192, 193,
	50, 195, 1, 188, 0, 0, 16, 16, 0, 26,
	33, 34, 35, 16, 32, 0, 0, 0, 0, 38,
	25, 16, 16, 0, 0, 213, 0, 215, 36, 0,
	0, 16, 16, 0, 16, 169, 16, 221, 0, 0,
	16, 0, 0, 0, 0, 16, 0, 0, 0, 0,
	0, 0, 0, 37, 0, 0, 184, 185, 0, 0,
	25, 0, 191, 25, 0, 0, 90, 33, 34, 35,
	91, 32, 199, 200, 0, 0, 38, 25, 0, 0,
	0, 0, 0, 0, 0, 36, 0, 209, 210, 212,
	0, 214, 25, 25, 97, 0, 37, 218, 0, 25,
	0, 0, 0, 0, 0, 0, 223, 25, 25, 90,
	33, 34, 35, 91, 32, 0, 0, 25, 25, 38,
	25, 0, 25, 0, 0, 0, 25, 0, 36, 0,
	0, 25, 82, 83, 84, 65, 66, 71, 72, 70,
	69, 0, 0, 0, 0, 0, 0, 0, 67, 68,
	73, 74, 75, 76, 78, 81, 0, 0, 79, 80,
	0, 77, 0, 0, 0, 0, 0, 0, 0, 0,
	194, 82, 83, 84, 65, 66, 71, 72, 70, 69,
	0, 0, 0, 0, 0, 0, 0, 67, 68, 73,
	74, 75, 76, 78, 81, 0, 0, 79, 80, 0,
	77, 0, 0, 0, 0, 0, 0, 0, 0, 163,
	82, 83, 84, 65, 66, 71, 72, 70, 69, 0,
	0, 0, 0, 82, 83, 84, 67, 68, 73, 74,
	75, 76, 78, 81, 203, 0, 79, 80, 0, 77,
	0, 73, 74, 75, 76, 78, 81, 0, 140, 79,
	80, 0, 77, 82, 83, 84, 65, 66, 71, 72,
	70, 69, 0, 0, 0, 0, 0, 0, 0, 67,
	68, 73, 74, 75, 76, 78, 81, 0, 0, 79,
	80, 0, 77, 0, 0, 0, 0, 0, 0, 204,
	82, 83, 84, 65, 66, 71, 72, 70, 69, 0,
	0, 0, 0, 0, 0, 0, 67, 68, 73, 74,
	75, 76, 78, 81, 0, 0, 79, 80, 0, 77,
	0, 0, 0, 0, 0, 0, 0, 133, 82, 83,
	84, 65, 66, 71, 72, 70, 69, 0, 0, 0,
	0, 0, 0, 0, 67, 68, 73, 74, 75, 76,
	78, 81, 0, 0, 79, 80, 0, 77, 0, 220,
	62, 17, 177, 0, 174, 9, 176, 23, 21, 0,
	24, 13, 12, 22, 0, 0, 11, 10, 82, 83,
	84, 65, 66, 71, 72, 70, 69, 0, 26, 0,
	0, 19, 0, 0, 67, 68, 73, 74, 75, 76,
	78, 81, 205, 0, 79, 80, 0, 77, 0, 0,
	0, 0, 0, 63, 0, 0, 0, 0, 82, 83,
	84, 65, 66, 71, 72, 70, 69, 0, 0, 0,
	0, 0, 111, 0, 67, 68, 73, 74, 75, 76,
	78, 81, 0, 0, 79, 80, 0, 77, 82, 83,
	84, 65, 66, 71, 72, 70, 69, 0, 0, 108,
	0, 0, 0, 0, 67, 68, 73, 74, 75, 76,
	78, 81, 0, 0, 79, 80, 0, 77, 82, 83,
	84, 65, 66, 71, 72, 70, 69, 0, 0, 0,
	0, 0, 0, 0, 67, 68, 73, 74, 75, 76,
	78, 81, 0, 0, 79, 80, 0, 77, 62, 17,
	0, 0, 224, 9, 0, 23, 21, 0, 24, 13,
	12, 22, 0, 0, 11, 10, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 26, 62, 17, 19,
	0, 222, 9, 0, 23, 21, 0, 24, 13, 12,
	22, 0, 0, 11, 10, 0, 0, 0, 0, 0,
	0, 63, 0, 0, 0, 26, 0, 0, 19, 0,
	0, 0, 0, 0, 0, 82, 83, 84, 65, 66,
	71, 72, 70, 69, 0, 0, 0, 0, 0, 0,
	63, 67, 68, 73, 74, 75, 76, 78, 81, 0,
	0, 79, 80, 0, 77, 62, 17, 0, 0, 219,
	9, 0, 23, 21, 0, 24, 13, 12, 22, 0,
	0, 11, 10, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 26, 62, 17, 19, 0, 217, 9,
	0, 23, 21, 0, 24, 13, 12, 22, 0, 0,
	11, 10, 0, 0, 0, 0, 0, 0, 63, 0,
	0, 0, 26, 62, 17, 19, 0, 216, 9, 0,
	23, 21, 0, 24, 13, 12, 22, 0, 0, 11,
	10, 0, 0, 0, 0, 0, 0, 63, 0, 0,
	0, 26, 62, 17, 19, 0, 208, 9, 0, 23,
	21, 0, 24, 13, 12, 22, 0, 0, 11, 10,
	0, 0, 0, 0, 0, 0, 63, 0, 0, 0,
	26, 62, 17, 19, 0, 207, 9, 0, 23, 21,
	0, 24, 13, 12, 22, 0, 0, 11, 10, 0,
	0, 0, 0, 0, 0, 63, 0, 0, 0, 26,
	62, 17, 19, 0, 198, 9, 0, 23, 21, 0,
	24, 13, 12, 22, 0, 0, 11, 10, 0, 0,
	0, 0, 0, 0, 63, 0, 0, 0, 26, 62,
	17, 19, 0, 197, 9, 0, 23, 21, 0, 24,
	13, 12, 22, 0, 0, 11, 10, 0, 0, 0,
	0, 0, 0, 63, 0, 0, 0, 26, 62, 17,
	19, 0, 183, 9, 0, 23, 21, 0, 24, 13,
	12, 22, 0, 0, 11, 10, 0, 0, 0, 0,
	0, 0, 63, 0, 0, 0, 26, 62, 17, 19,
	0, 173, 9, 0, 23, 21, 0, 24, 13, 12,
	22, 0, 0, 11, 10, 0, 0, 0, 0, 0,
	0, 63, 0, 0, 0, 26, 62, 17, 19, 0,
	156, 9, 0, 23, 21, 0, 24, 13, 12, 22,
	0, 0, 11, 10, 0, 0, 0, 0, 0, 0,
	63, 0, 0, 0, 26, 62, 17, 19, 0, 0,
	9, 0, 23, 21, 0, 24, 13, 12, 22, 109,
	0, 11, 10, 0, 0, 0, 0, 0, 0, 63,
	0, 0, 0, 26, 62, 17, 19, 0, 58, 9,
	0, 23, 21, 0, 24, 13, 12, 22, 0, 0,
	11, 10, 0, 0, 0, 0, 0, 0, 63, 0,
	0, 0, 26, 7, 17, 19, 0, 0, 9, 0,
	23, 21, 20, 24, 13, 12, 22, 0, 0, 11,
	10, 0, 0, 0, 0, 0, 0, 63, 0, 0,
	0, 26, 62, 17, 19, 0, 0, 9, 0, 23,
	21, 0, 24, 13, 12, 22, 0, 0, 11, 10,
	0, 0, 0, 0, 0, 0, 8, 0, 0, 0,
	26, 0, 0, 19, 0, 0, 0, 0, 0, 0,
	82, 83, 84, 0, 66, 71, 72, 70, 69, 0,
	0, 0, 0, 0, 0, 63, 67, 68, 73, 74,
	75, 76, 78, 81, 0, 0, 79, 80, 0, 77,
	82, 83, 84, 0, 0, 71, 72, 70, 69, 0,
	0, 0, 0, 0, 0, 0, 67, 68, 73, 74,
	75, 76, 78, 81, 0, 0, 79, 80, 0, 77,
}

var yyPact = [...]int{
	-1000, 1129, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	120, 119, -1000, 157, -1000, -1000, 82, 100, 2, 106,
	114, 157, -1000, 113, 157, -1000, -1000, 1100, -1000, 133,
	732, 82, 157, -1000, -1000, 89, 157, 157, 157, -1000,
	157, 112, -1000, -1000, 214, -6, 109, -1000, 157, 106,
	-35, 82, 12, 635, 1071, 52, 605, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 157, 157, 157, 157, 157,
	157, 157, 157, 157, 157, 157, 157, 157, 157, 157,
	157, 157, 157, 157, 157, 447, -1000, -24, -39, 732,
	48, 157, -1000, -1000, -1000, 367, -1000, -1000, -28, -39,
	157, 102, -40, -29, -1000, 26, -17, 95, -1000, 157,
	157, -1000, 1042, 1177, 1207, 380, 380, 380, 380, 380,
	380, 57, 57, -1000, -1000, -1000, -1000, 3, 3, 3,
	53, 53, 53, -1000, -16, 157, -22, 97, 157, 328,
	-1000, -43, 257, -44, -29, -1000, 25, 157, 157, 86,
	-48, 54, 1013, 732, 485, 536, -1000, -1000, 732, -1000,
	24, 157, 732, 21, -1000, -39, -1000, 157, 732, 984,
	-1000, 79, -26, -1000, 157, 147, -1000, 157, 157, 289,
	157, -47, 732, -1000, 955, 926, -1000, 78, -51, 410,
	-1000, 1158, 575, 732, 20, 732, -1000, -1000, -1000, 897,
	868, -1000, 65, -1000, 157, -1000, 157, -1000, -1000, 839,
	810, -1000, 781, 535, 536, 732, -1000, -1000, 713, -1000,
	-1000, -1000, -1000, 684, -1000,
}

var yyPgo = [...]int{
	0, 182, 65, 173, 172, 87, 160, 10, 0, 5,
	7, 1, 169, 163, 167, 166, 2, 159, 156, 164,
	4,
}

var yyR1 = [...]int{
	0, 1, 1, 2, 2, 3, 3, 3, 3, 3,
	3, 4, 4, 4, 4, 4, 18, 18, 13, 13,
	13, 13, 13, 13, 14, 14, 14, 14, 15, 16,
	16, 16, 19, 19, 19, 19, 19, 19, 19, 19,
	17, 17, 17, 17, 17, 5, 5, 5, 6, 6,
	7, 7, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 8, 8, 8, 8, 8, 8, 8, 8,
	8, 8, 11, 11, 11, 12, 12, 12, 12, 9,
	9, 10, 10, 10, 10, 20, 20,
}

var yyR2 = [...]int{
	0, 0, 2, 0, 2, 1, 1, 1, 1, 3,
	1, 1, 1, 1, 3, 1, 1, 1, 1, 2,
	4, 6, 3, 5, 5, 4, 9, 11, 6, 0,
	2, 5, 6, 7, 7, 8, 8, 9, 9, 10,
	1, 2, 3, 1, 2, 1, 4, 3, 1, 3,
	1, 3, 1, 3, 1, 1, 2, 4, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 1, 2, 2, 2, 4, 4, 6, 1,
	3, 3, 5, 5, 7, 0, 1,
}

var yyChk = [...]int{
	-1000, -1, -3, -17, -18, -13, -19, 4, 57, 9,
	21, 20, 16, 15, -14, -15, -11, 5, -6, 35,
	13, 12, 17, 11, 14, -5, 32, -2, 32, 32,
	-8, -11, 37, 33, 34, 35, 51, 19, 42, -5,
	36, 48, 34, -12, 37, -7, 35, 32, 38, 59,
	-6, -11, 32, -8, -2, 32, -8, -4, 8, -17,
	-18, -13, 4, 57, 20, 26, 27, 39, 40, 31,
	30, 28, 29, 41, 42, 43, 44, 52, 45, 49,
	50, 46, 23, 24, 25, -8, 58, -9, -10, -8,
	32, 36, -8, -8, -8, -8, 32, 60, -9, -10,
	38, 59, -7, -9, -5, 58, 37, 48, 4, 18,
	38, 7, -2, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, 60, -20, 59, -20, 59, 38, -8,
	61, -20, 59, -20, -9, 32, 58, 59, 38, 60,
	-7, 32, -2, -8, -8, -2, 8, 58, -8, 58,
	32, 36, -8, 61, 60, -10, 60, 38, -8, -2,
	34, 60, 37, 8, 59, -16, 10, 6, 38, -8,
	38, -20, -8, 8, -2, -2, 34, 60, -7, -8,
	8, -2, -8, -8, 61, -8, 60, 8, 8, -2,
	-2, 34, 60, 4, 59, 7, 38, 8, 8, -2,
	-2, 34, -2, -8, -2, -8, 8, 8, -2, 8,
	4, -16, 8, -2, 8,
}

var yyDef = [...]int{
	1, -2, 2, 5, 6, 7, 8, 3, 10, 40,
	0, 0, 43, 0, 16, 17, 18, 0, 0, 0,
	0, 0, 3, 0, 0, -2, 45, 0, 41, 0,
	44, 52, 0, 54, 55, 0, 0, 0, 0, 82,
	0, 0, 83, 84, 0, 19, 0, 50, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 4, 9, 11,
	12, 13, 3, 15, 42, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 56, 95, 95, 89,
	45, 0, 79, 80, 81, 0, 47, 85, 95, 95,
	0, 0, 0, 22, -2, 0, 0, 0, 3, 0,
	0, 3, 0, 59, 60, 61, 62, 63, 64, 65,
	66, 67, 68, 69, 70, 71, 72, 73, 74, 75,
	76, 77, 78, 53, 0, 96, 0, 96, 0, 0,
	46, 0, 96, 0, 20, 51, 0, 0, 0, 3,
	0, 0, 0, 25, 0, 29, 14, 57, 90, 58,
	0, 0, 91, 0, 86, 95, 87, 0, 23, 0,
	3, 3, 0, 24, 0, 0, 3, 0, 0, 0,
	0, 0, 21, 32, 0, 0, 3, 3, 0, 0,
	28, 30, 0, 93, 0, 92, 88, 34, 33, 0,
	0, 3, 3, 3, 0, 3, 0, 35, 36, 0,
	0, 3, 0, 0, 29, 94, 38, 37, 0, 26,
	3, 31, 39, 0, 27,
}

var yyTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 47, 3, 45, 49, 3,
	37, 60, 43, 41, 59, 42, 48, 44, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 57,
	40, 38, 39, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 53, 3, 3, 3, 3, 3,
	3, 36, 3, 61, 46, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 35, 50, 58, 51,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 52, 54, 55, 56,
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
//line parser.go.y:54
		{
			yyVAL.expr = __chain()
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:60
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[2].expr)
			if l, ok := yylex.(*Lexer); ok {
				l.Stmts = yyVAL.expr
			}
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:68
		{
			yyVAL.expr = __chain()
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:71
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[2].expr)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:76
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:77
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:78
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:79
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 9:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:80
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:81
		{
			yyVAL.expr = emptyNode
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:84
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:85
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:86
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:87
		{
			yyVAL.expr = __do(yyDollar[2].expr)
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:88
		{
			yyVAL.expr = emptyNode
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:91
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:92
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:95
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:98
		{
			yyVAL.expr = __chain()
			for _, v := range yyDollar[2].expr.Nodes {
				yyVAL.expr = yyVAL.expr.append(__set(v, NewSymbol(ANil)).SetPos(yyDollar[1].token.Pos))
			}
		}
	case 20:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:104
		{
			yyVAL.expr = __local(yyDollar[2].expr.Nodes, yyDollar[4].expr.Nodes, yyDollar[1].token.Pos)
		}
	case 21:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:107
		{
			tmp := randomVarname()
			yyVAL.expr = __chain(__local([]Node{tmp}, []Node{yyDollar[6].expr}, yyDollar[1].token.Pos))
			for i, ident := range yyDollar[3].expr.Nodes {
				yyVAL.expr = yyVAL.expr.append(__local([]Node{ident}, []Node{__load(tmp, NewNumberFromInt(int64(i))).SetPos(yyDollar[1].token.Pos)}, yyDollar[1].token.Pos))
			}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:114
		{
			yyVAL.expr = __moveMulti(yyDollar[1].expr.Nodes, yyDollar[3].expr.Nodes, yyDollar[2].token.Pos)
		}
	case 23:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:117
		{
			tmp := randomVarname()
			yyVAL.expr = __chain(__local([]Node{tmp}, []Node{yyDollar[5].expr}, yyDollar[1].token.Pos))
			for i, decl := range yyDollar[2].expr.Nodes {
				x := decl.moveLoadStore(__move, __load(tmp, NewNumberFromInt(int64(i))).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos)
				yyVAL.expr = yyVAL.expr.append(__local([]Node{decl}, []Node{x}, yyDollar[1].token.Pos))
			}
		}
	case 24:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:127
		{
			yyVAL.expr = __loop(__if(yyDollar[2].expr, yyDollar[4].expr, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:130
		{
			yyVAL.expr = __loop(
				__chain(
					yyDollar[2].expr,
					__if(yyDollar[4].expr, breakNode, emptyNode).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			).SetPos(yyDollar[1].token.Pos)
		}
	case 26:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:138
		{
			forVar, forEnd := NewSymbolFromToken(yyDollar[2].token), randomVarname()
			yyVAL.expr = __do(
				__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
				__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
				__loop(
					__if(
						__less(forVar, forEnd),
						__chain(yyDollar[8].expr, __inc(forVar, oneNode).SetPos(yyDollar[1].token.Pos)),
						breakNode,
					).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos),
			)
		}
	case 27:
		yyDollar = yyS[yypt-11 : yypt+1]
//line parser.go.y:152
		{
			forVar, forEnd, forStep := NewSymbolFromToken(yyDollar[2].token), randomVarname(), randomVarname()
			body := __chain(yyDollar[10].expr, __inc(forVar, forStep))
			yyVAL.expr = __do(
				__set(forVar, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos),
				__set(forEnd, yyDollar[6].expr).SetPos(yyDollar[1].token.Pos),
				__set(forStep, yyDollar[8].expr).SetPos(yyDollar[1].token.Pos))

			if yyDollar[8].expr.IsNumber() { // step is a static number, easy case
				if yyDollar[8].expr.IsNegativeNumber() {
					yyVAL.expr = yyVAL.expr.append(__loop(__if(__less(forEnd, forVar), body, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos))
				} else {
					yyVAL.expr = yyVAL.expr.append(__loop(__if(__less(forVar, forEnd), body, breakNode).SetPos(yyDollar[1].token.Pos)).SetPos(yyDollar[1].token.Pos))
				}
			} else {
				yyVAL.expr = yyVAL.expr.append(__loop(
					__if(
						__less(zeroNode, forStep).SetPos(yyDollar[1].token.Pos),
						// +step
						__if(__lessEq(forEnd, forVar), breakNode, body).SetPos(yyDollar[1].token.Pos),
						// -step
						__if(__lessEq(forVar, forEnd), breakNode, body).SetPos(yyDollar[1].token.Pos),
					).SetPos(yyDollar[1].token.Pos),
				).SetPos(yyDollar[1].token.Pos))
			}
		}
	case 28:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:180
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:185
		{
			yyVAL.expr = NewComplex()
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:188
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:191
		{
			yyVAL.expr = __if(yyDollar[2].expr, yyDollar[4].expr, yyDollar[5].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 32:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:196
		{
			yyVAL.expr = __func(yyDollar[2].token, emptyNode, "", yyDollar[5].expr)
		}
	case 33:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:197
		{
			yyVAL.expr = __func(yyDollar[2].token, yyDollar[4].expr, "", yyDollar[6].expr)
		}
	case 34:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:198
		{
			yyVAL.expr = __func(yyDollar[2].token, emptyNode, yyDollar[5].token.Str, yyDollar[6].expr)
		}
	case 35:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:199
		{
			yyVAL.expr = __func(yyDollar[2].token, yyDollar[4].expr, yyDollar[6].token.Str, yyDollar[7].expr)
		}
	case 36:
		yyDollar = yyS[yypt-8 : yypt+1]
//line parser.go.y:200
		{
			yyVAL.expr = __store(NewSymbolFromToken(yyDollar[2].token), NewString(yyDollar[4].token.Str), __func(yyDollar[4].token, emptyNode, "", yyDollar[7].expr))
		}
	case 37:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:203
		{
			yyVAL.expr = __store(NewSymbolFromToken(yyDollar[2].token), NewString(yyDollar[4].token.Str), __func(yyDollar[4].token, yyDollar[6].expr, "", yyDollar[8].expr))
		}
	case 38:
		yyDollar = yyS[yypt-9 : yypt+1]
//line parser.go.y:206
		{
			yyVAL.expr = __store(NewSymbolFromToken(yyDollar[2].token), NewString(yyDollar[4].token.Str), __func(yyDollar[4].token, emptyNode, yyDollar[7].token.Str, yyDollar[8].expr))
		}
	case 39:
		yyDollar = yyS[yypt-10 : yypt+1]
//line parser.go.y:209
		{
			yyVAL.expr = __store(NewSymbolFromToken(yyDollar[2].token), NewString(yyDollar[4].token.Str), __func(yyDollar[4].token, yyDollar[6].expr, yyDollar[8].token.Str, yyDollar[9].expr))
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:214
		{
			yyVAL.expr = NewComplex(NewSymbol(ABreak)).SetPos(yyDollar[1].token.Pos)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:217
		{
			yyVAL.expr = NewComplex(NewSymbol(AGoto), NewSymbolFromToken(yyDollar[2].token)).SetPos(yyDollar[1].token.Pos)
		}
	case 42:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:220
		{
			yyVAL.expr = NewComplex(NewSymbol(ALabel), NewSymbolFromToken(yyDollar[2].token))
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:223
		{
			yyVAL.expr = NewComplex(NewSymbol(AReturn), NewSymbol(ANil)).SetPos(yyDollar[1].token.Pos)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:226
		{
			if len(yyDollar[2].expr.Nodes) == 3 && yyDollar[2].expr.Nodes[0].SymbolValue() == ACall {
				// return call(...) -> return tailcall(...)
				yyDollar[2].expr.Nodes[0].strSym = ATailCall
			}
			yyVAL.expr = NewComplex(NewSymbol(AReturn), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:235
		{
			yyVAL.expr = NewSymbolFromToken(yyDollar[1].token)
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:238
		{
			yyVAL.expr = __load(yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:241
		{
			yyVAL.expr = __load(yyDollar[1].expr, NewString(yyDollar[3].token.Str)).SetPos(yyDollar[2].token.Pos)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:246
		{
			yyVAL.expr = NewComplex(yyDollar[1].expr)
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:249
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[3].expr)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:254
		{
			yyVAL.expr = NewComplex(NewSymbolFromToken(yyDollar[1].token))
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:257
		{
			yyVAL.expr = yyDollar[1].expr.append(NewSymbolFromToken(yyDollar[3].token))
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:262
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:263
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:264
		{
			yyVAL.expr = NewNumberFromString(yyDollar[1].token.Str)
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:265
		{
			yyVAL.expr = NewString(yyDollar[1].token.Str)
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:266
		{
			yyVAL.expr = NewComplex(NewSymbol(AArrayMap), emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:267
		{
			yyVAL.expr = NewComplex(NewSymbol(AArray), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:268
		{
			yyVAL.expr = NewComplex(NewSymbol(AArrayMap), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:269
		{
			yyVAL.expr = NewComplex(NewSymbol(AOr), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:270
		{
			yyVAL.expr = NewComplex(NewSymbol(AAnd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:271
		{
			yyVAL.expr = NewComplex(NewSymbol(ALess), yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:272
		{
			yyVAL.expr = NewComplex(NewSymbol(ALess), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:273
		{
			yyVAL.expr = NewComplex(NewSymbol(ALessEq), yyDollar[3].expr, yyDollar[1].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:274
		{
			yyVAL.expr = NewComplex(NewSymbol(ALessEq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:275
		{
			yyVAL.expr = NewComplex(NewSymbol(AEq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:276
		{
			yyVAL.expr = NewComplex(NewSymbol(ANeq), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:277
		{
			yyVAL.expr = NewComplex(NewSymbol(AAdd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:278
		{
			yyVAL.expr = NewComplex(NewSymbol(ASub), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:279
		{
			yyVAL.expr = NewComplex(NewSymbol(AMul), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:280
		{
			yyVAL.expr = NewComplex(NewSymbol(ADiv), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:281
		{
			yyVAL.expr = NewComplex(NewSymbol(AIDiv), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:282
		{
			yyVAL.expr = NewComplex(NewSymbol(AMod), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:283
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitAnd), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:284
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitOr), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:285
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitXor), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:286
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitLsh), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:287
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitRsh), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:288
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitURsh), yyDollar[1].expr, yyDollar[3].expr).SetPos(yyDollar[2].token.Pos)
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:289
		{
			yyVAL.expr = NewComplex(NewSymbol(ABitNot), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:290
		{
			yyVAL.expr = NewComplex(NewSymbol(ANot), yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:291
		{
			yyVAL.expr = NewComplex(NewSymbol(ASub), zeroNode, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:294
		{
			yyVAL.expr = yyDollar[1].expr
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:297
		{
			yyVAL.expr = __call(yyDollar[1].expr, NewComplex(NewString(yyDollar[2].token.Str))).SetPos(yyDollar[1].expr.Pos())
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:300
		{
			yyDollar[2].expr.Nodes[1] = yyDollar[1].expr
			yyVAL.expr = yyDollar[2].expr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser.go.y:306
		{
			yyVAL.expr = __call(emptyNode, emptyNode).SetPos(yyDollar[1].token.Pos)
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:309
		{
			yyVAL.expr = __call(emptyNode, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 87:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parser.go.y:312
		{
			yyVAL.expr = __callMap(emptyNode, emptyNode, yyDollar[2].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 88:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parser.go.y:315
		{
			yyVAL.expr = __callMap(emptyNode, yyDollar[2].expr, yyDollar[4].expr).SetPos(yyDollar[1].token.Pos)
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:320
		{
			yyVAL.expr = NewComplex(yyDollar[1].expr)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:323
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[3].expr)
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser.go.y:328
		{
			yyVAL.expr = NewComplex(NewString(yyDollar[1].token.Str), yyDollar[3].expr)
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:331
		{
			yyVAL.expr = NewComplex(yyDollar[2].expr, yyDollar[5].expr)
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser.go.y:334
		{
			yyVAL.expr = yyDollar[1].expr.append(NewString(yyDollar[3].token.Str)).append(yyDollar[5].expr)
		}
	case 94:
		yyDollar = yyS[yypt-7 : yypt+1]
//line parser.go.y:337
		{
			yyVAL.expr = yyDollar[1].expr.append(yyDollar[4].expr).append(yyDollar[7].expr)
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser.go.y:341
		{
			yyVAL.expr = emptyNode
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser.go.y:341
		{
			yyVAL.expr = emptyNode
		}
	}
	goto yystack /* stack new state and value */
}
