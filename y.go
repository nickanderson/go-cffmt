// Code generated by goyacc -v  parse.y. DO NOT EDIT.

//line parse.y:2
package cf

import __yyfmt__ "fmt"

//line parse.y:2

import (
	"github.com/miekg/cf/ast"
)

//line parse.y:17
type yySymType struct {
	yys   int
	token ast.Token
}

const IDENTIFIER = 57346
const QSTRING = 57347
const CLASSGUARD = 57348
const PROMISEGUARD = 57349
const BUNDLE = 57350
const BODY = 57351
const PROMISE = 57352
const FATARROW = 57353
const THINARROW = 57354
const NAKEDVAR = 57355

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"QSTRING",
	"CLASSGUARD",
	"PROMISEGUARD",
	"BUNDLE",
	"BODY",
	"PROMISE",
	"FATARROW",
	"THINARROW",
	"NAKEDVAR",
	"'('",
	"')'",
	"','",
	"'}'",
	"'{'",
	"';'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parse.y:636

//line yacctab:1
var yyExca = [...]int16{
	-1, 0,
	1, 1,
	-2, 0,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 2,
	-2, 0,
	-1, 52,
	15, 45,
	16, 45,
	-2, 38,
	-1, 59,
	17, 50,
	-2, 0,
	-1, 61,
	15, 42,
	-2, 0,
	-1, 62,
	17, 87,
	-2, 0,
	-1, 72,
	17, 88,
	-2, 0,
	-1, 90,
	7, 59,
	17, 59,
	-2, 61,
	-1, 95,
	4, 71,
	19, 71,
	-2, 0,
	-1, 106,
	19, 74,
	-2, 0,
	-1, 111,
	14, 121,
	-2, 104,
	-1, 113,
	14, 122,
	-2, 106,
	-1, 120,
	19, 75,
	-2, 0,
	-1, 127,
	19, 74,
	-2, 0,
	-1, 130,
	16, 113,
	17, 113,
	-2, 0,
	-1, 131,
	14, 121,
	-2, 116,
	-1, 133,
	14, 122,
	-2, 118,
	-1, 150,
	15, 128,
	16, 128,
	-2, 0,
	-1, 153,
	15, 129,
	16, 129,
	-2, 0,
	-1, 154,
	14, 121,
	-2, 132,
	-1, 156,
	14, 122,
	-2, 134,
}

const yyPrivate = 57344

const yyLast = 179

var yyAct = [...]uint8{
	115, 153, 110, 130, 99, 119, 122, 74, 73, 91,
	65, 50, 51, 56, 57, 41, 30, 116, 87, 111,
	112, 135, 103, 131, 132, 142, 141, 135, 113, 131,
	132, 49, 133, 117, 84, 86, 148, 35, 133, 81,
	161, 39, 128, 138, 52, 163, 55, 48, 61, 44,
	17, 54, 145, 45, 43, 68, 47, 54, 158, 58,
	154, 155, 104, 60, 22, 26, 98, 76, 162, 156,
	7, 101, 107, 69, 66, 82, 8, 9, 10, 68,
	100, 85, 80, 125, 79, 124, 76, 70, 40, 55,
	32, 92, 36, 31, 32, 32, 27, 23, 19, 19,
	18, 3, 19, 143, 11, 137, 160, 152, 150, 144,
	92, 126, 109, 136, 118, 129, 114, 108, 134, 88,
	78, 77, 75, 72, 71, 62, 151, 139, 123, 121,
	120, 106, 127, 140, 105, 97, 96, 95, 94, 93,
	102, 90, 89, 134, 147, 146, 149, 83, 67, 64,
	63, 157, 59, 53, 159, 42, 38, 25, 34, 21,
	29, 16, 157, 164, 37, 24, 14, 33, 20, 13,
	46, 28, 15, 12, 6, 5, 4, 2, 1,
}

var yyPact = [...]int16{
	68, -1000, 68, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 98, 95, 94, 91, -1000, -1000, -1000, -1000,
	90, -1000, -1000, -1000, 86, -1000, -1000, -1000, 40, -1000,
	-1000, -1000, -1000, 40, -1000, -1000, -1000, 40, -1000, -1000,
	-1000, 29, 42, -1000, 29, 29, -1000, -1000, -1000, -1000,
	36, -1000, -1000, 32, -1000, -1000, -1000, -1000, -1000, 72,
	-1000, 85, 80, 22, 48, -1000, -1000, -1000, -1000, -1000,
	-1000, 17, 80, -1000, -1000, -1000, -1000, 16, -1000, -1000,
	-1000, -1000, -1000, 61, -1000, -1000, -1000, -1000, 69, -1000,
	-1000, -1000, -1000, -1000, 3, 60, -1000, -1000, -1000, -1000,
	-1000, -1000, 61, -1000, -1000, 15, 81, -1000, 15, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 25, -1000, -1000,
	103, 27, -1000, -1000, -1000, -1000, -1000, 81, -1000, 9,
	101, -1000, -1000, -1000, -1000, -1000, 38, -1000, 81, 69,
	-1000, -1000, 19, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	56, 15, 24, 66, -1000, -1000, -1000, -1000, -1000, -1000,
	30, 56, -1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 178, 177, 101, 176, 175, 174, 173, 172, 171,
	15, 170, 169, 168, 167, 13, 166, 165, 164, 161,
	50, 160, 16, 159, 158, 157, 156, 155, 11, 12,
	153, 14, 152, 150, 149, 10, 148, 147, 142, 141,
	9, 140, 7, 139, 138, 137, 136, 135, 134, 2,
	132, 5, 131, 130, 129, 6, 128, 127, 4, 126,
	125, 124, 123, 8, 122, 121, 120, 119, 117, 116,
	0, 115, 3, 114, 113, 109, 108, 107, 106, 1,
}

var yyR1 = [...]int8{
	0, 1, 1, 2, 2, 3, 3, 3, 3, 7,
	4, 12, 5, 16, 6, 8, 19, 19, 9, 21,
	21, 13, 23, 23, 14, 24, 24, 17, 25, 25,
	18, 26, 26, 20, 22, 10, 10, 10, 10, 27,
	29, 28, 28, 28, 30, 30, 32, 11, 31, 31,
	33, 33, 34, 34, 34, 37, 35, 36, 38, 38,
	39, 41, 39, 40, 40, 43, 43, 44, 44, 50,
	46, 52, 47, 45, 51, 51, 51, 53, 54, 54,
	57, 59, 55, 56, 56, 60, 15, 61, 61, 62,
	62, 63, 63, 64, 64, 67, 68, 65, 66, 66,
	58, 58, 48, 42, 49, 49, 49, 49, 49, 49,
	69, 69, 69, 71, 71, 71, 72, 72, 72, 72,
	72, 73, 73, 74, 70, 76, 78, 75, 77, 77,
	77, 77, 79, 79, 79, 79, 79,
}

var yyR2 = [...]int8{
	0, 0, 1, 1, 2, 1, 1, 1, 1, 0,
	6, 0, 6, 0, 6, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 0, 3, 2, 2, 1,
	1, 1, 2, 3, 1, 1, 0, 4, 1, 1,
	0, 1, 1, 2, 1, 0, 3, 1, 0, 1,
	1, 0, 3, 1, 1, 2, 2, 1, 1, 0,
	5, 0, 3, 1, 0, 1, 2, 1, 1, 3,
	0, 0, 5, 1, 1, 0, 4, 0, 1, 1,
	2, 1, 1, 2, 2, 0, 0, 5, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 3, 4, 1, 3, 2, 1, 1, 1, 1,
	1, 1, 1, 0, 3, 0, 0, 5, 0, 1,
	3, 2, 1, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -1, -2, -3, -4, -5, -6, 2, 8, 9,
	10, -3, -7, -12, -16, -8, -19, -20, 2, 4,
	-13, -23, -20, 2, -17, -25, -20, 2, -9, -21,
	-22, 2, 4, -14, -24, -22, 2, -18, -26, -22,
	2, -10, -27, 14, -10, -10, -11, -31, 18, 2,
	-28, -29, 2, -30, 15, 4, -15, -31, -15, -32,
	-29, 16, -60, -33, -34, -35, 2, -36, 7, -28,
	2, -61, -62, -63, -42, -64, 6, -65, -66, 4,
	2, 17, -35, -37, 17, -63, 19, 2, -67, -38,
	-39, -40, -42, -43, -44, -45, -46, -47, 5, -58,
	11, 2, -41, 19, 2, -48, -52, 12, -68, -40,
	-49, 4, 5, 13, -69, -70, 2, 18, -73, -51,
	-53, -54, -55, -56, 4, 2, -49, -50, 17, -71,
	-72, 4, 5, 13, -70, 2, -74, 2, 16, -57,
	-51, 17, 16, 2, -75, 14, -55, -58, 17, -72,
	-76, -59, -77, -79, 4, 5, 13, -70, 2, -49,
	-78, 16, 2, 15, -79,
}

var yyDef = [...]int16{
	-2, -2, -2, 3, 5, 6, 7, 8, 9, 11,
	13, 4, 0, 0, 0, 0, 15, 16, 17, 33,
	0, 21, 22, 23, 0, 27, 28, 29, 35, 18,
	19, 20, 34, 35, 24, 25, 26, 35, 30, 31,
	32, 0, 0, 39, 0, 0, 10, 46, 48, 49,
	0, 37, -2, 41, 40, 44, 12, 85, 14, -2,
	36, -2, -2, 0, 51, 52, 54, 55, 57, 43,
	45, 0, -2, 89, 91, 92, 103, 0, 95, 98,
	99, 47, 53, 58, 86, 90, 93, 94, 0, 56,
	-2, 60, 63, 64, 0, -2, 67, 68, 73, 96,
	100, 101, 0, 65, 66, 0, -2, 102, 0, 62,
	69, -2, 105, -2, 107, 108, 109, 0, 123, 72,
	-2, 77, 78, 80, 83, 84, 97, -2, 110, 0,
	-2, -2, 117, -2, 119, 120, 0, 76, 0, 0,
	70, 111, 0, 115, 124, 125, 79, 81, 112, 114,
	-2, 0, 126, -2, -2, 133, -2, 135, 136, 82,
	0, 0, 131, 127, 130,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	14, 15, 3, 3, 16, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 19,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 18, 3, 17,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13,
}

var yyTok3 = [...]int8{
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
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
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
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
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
	yyn = int(yyPact[yystate])
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
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
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
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
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
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
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

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:25
		{
			yylex.(*Lexer).Spec = &ast.Specification{}
			yylex.(*Lexer).Spec.SetChildren([]ast.Node{ast.Up(yylex.(*Lexer).parent)})
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:34
		{
			yylex.(*Lexer).yydebug("block:bundle", yyVAL.token)
			// Here we find the actual token, but we get here last. Find original bundle and put
			// token contents in it. Mostly to get the comments out.
			bundle := ast.UpTo(yylex.(*Lexer).parent, &ast.Bundle{})
			if bundle != nil {
				bundle.SetToken(yyVAL.token)
			}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:44
		{
			yylex.(*Lexer).yydebug("block:body", yyVAL.token)
			body := ast.UpTo(yylex.(*Lexer).parent, &ast.Body{})
			if body != nil {
				body.SetToken(yyVAL.token)
			}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:53
		{
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:57
		{
			yylex.(*Lexer).yydebug("bundle:BUNDLE", yyVAL.token)
			spec := ast.UpTo(yylex.(*Lexer).parent, &ast.Specification{})
			yylex.(*Lexer).parent = spec
			b := ast.New(&ast.Bundle{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, b)
			yylex.(*Lexer).parent = b
		}
	case 10:
		yyDollar = yyS[yypt-6 : yypt+1]
//line parse.y:66
		{
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:70
		{
			yylex.(*Lexer).yydebug("body:BODY", yyVAL.token)
			spec := ast.UpTo(yylex.(*Lexer).parent, &ast.Specification{})
			yylex.(*Lexer).parent = spec
			b := ast.New(&ast.Body{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, b)
			yylex.(*Lexer).parent = b
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:81
		{
			yylex.(*Lexer).yydebug("promise:PROMISE")
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:89
		{
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:92
		{
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:96
		{
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:101
		{
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:105
		{
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:109
		{
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:112
		{
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:116
		{
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:121
		{
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:125
		{
			yylex.(*Lexer).yydebug("promisecomponent")
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:130
		{
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:133
		{
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:137
		{
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:142
		{
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:148
		{
			ast.Append(yylex.(*Lexer).parent, ast.New(&ast.Identifier{}, yyVAL.token))
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:155
		{
			ast.Append(yylex.(*Lexer).parent, ast.New(&ast.Identifier{}, yyVAL.token))
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:163
		{
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:167
		{
			yylex.(*Lexer).yydebug("arglist_begin:(", yyVAL.token)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:172
		{
			yylex.(*Lexer).yydebug("arglist_end:)", yyVAL.token)
			bundle := ast.UpTo(yylex.(*Lexer).parent, &ast.Bundle{})
			if bundle != nil {
				yylex.(*Lexer).parent = bundle
			} else { // maybe body?
				if body := ast.UpTo(yylex.(*Lexer).parent, &ast.Body{}); body != nil {
					yylex.(*Lexer).parent = body
				}
			}
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:185
		{
			if _, ok := yylex.(*Lexer).parent.(*ast.ArgList); !ok {
				a := ast.New(&ast.ArgList{})
				ast.Append(yylex.(*Lexer).parent, a)
				yylex.(*Lexer).parent = a
			}
			al := ast.New(&ast.ArgListItem{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, al)
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:196
		{
			if _, ok := yylex.(*Lexer).parent.(*ast.ArgList); !ok {
				a := ast.New(&ast.ArgList{})
				ast.Append(yylex.(*Lexer).parent, a)
				yylex.(*Lexer).parent = a
			}
			al := ast.New(&ast.ArgListItem{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, al)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:207
		{
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:210
		{
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:216
		{
		}
	case 47:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parse.y:222
		{
			// only here for comments.
			if bundle := ast.UpTo(yylex.(*Lexer).parent, &ast.Bundle{}); bundle != nil {
				bundle.SetToken(yyVAL.token)
			}
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:232
		{
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:235
		{
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:249
		{
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:254
		{
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:257
		{
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:261
		{
			yylex.(*Lexer).yydebug("promise_guard", yyVAL.token)
			pg := ast.New(&ast.PromiseGuard{}, yyVAL.token)
			// If there is previous promiseguard, this one closes it, and we can reyylex.(*Lexer).parent this new one, to that _parent_.
			prev := ast.UpTo(yylex.(*Lexer).parent, &ast.PromiseGuard{})
			if prev != nil {
				yylex.(*Lexer).parent = prev.Parent()
			}
			ast.Append(yylex.(*Lexer).parent, pg)
			yylex.(*Lexer).parent = pg
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parse.y:274
		{
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:277
		{
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:281
		{
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:284
		{
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:287
		{
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:291
		{
			yylex.(*Lexer).yydebug("classpromise", yyVAL.token)
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:295
		{
		}
	case 65:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:299
		{
			yylex.(*Lexer).yydebug("promise_decl", yyVAL.token)
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:303
		{
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:307
		{
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:310
		{
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:319
		{
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:325
		{
			yylex.(*Lexer).yydebug("promise_without_promisee: promiser", yyVAL.token)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:330
		{
			yylex.(*Lexer).yydebug("promise_without_promisee: promise_decl_constraints", yyVAL.token)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:335
		{
			yylex.(*Lexer).yydebug("promiser:QSTRING", yyVAL.token)
			// same level as previous Promiser, or PromiseGuard

			prev := ast.UpTo(yylex.(*Lexer).parent, &ast.Promiser{})
			if prev == nil {
				if prev = ast.UpTo(yylex.(*Lexer).parent, &ast.PromiseGuard{}); prev != nil {
					yylex.(*Lexer).parent = prev
				}
			} else {
				yylex.(*Lexer).parent = prev.Parent()
			}

			p := ast.New(&ast.Promiser{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, p)
			yylex.(*Lexer).parent = p
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:356
		{
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:360
		{
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:367
		{
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:370
		{
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parse.y:373
		{
			yylex.(*Lexer).yydebug("contraint:rval")
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:378
		{
			yylex.(*Lexer).yydebug("contraint_id:IDENTIFIER", yyVAL.token)

			prev := ast.UpTo(yylex.(*Lexer).parent, &ast.Promiser{})
			if prev != nil {
				yylex.(*Lexer).parent = prev
			}

			c := ast.New(&ast.Constraint{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, c)
			yylex.(*Lexer).parent = c
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:391
		{
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:395
		{
		}
	case 86:
		yyDollar = yyS[yypt-4 : yypt+1]
//line parse.y:401
		{
			// only here for comments.
			if body := ast.UpTo(yylex.(*Lexer).parent, &ast.Body{}); body != nil {
				body.SetToken(yyVAL.token)
			}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:415
		{
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:421
		{
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:425
		{
			yylex.(*Lexer).yydebug("selection:selection_id", yyVAL.token)
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:429
		{
			yylex.(*Lexer).yydebug("selection:assign_arrow")
		}
	case 97:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parse.y:433
		{
			yylex.(*Lexer).yydebug("selection:rval", yyVAL.token)
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:438
		{
			yylex.(*Lexer).yydebug("selection_id:IDENTIFIER", yyVAL.token)
			// need to be parent of body or classguard
			prev := ast.UpTo(yylex.(*Lexer).parent, &ast.ClassGuard{})
			if prev == nil {
				prev = ast.UpTo(yylex.(*Lexer).parent, &ast.Body{})
			}
			yylex.(*Lexer).parent = prev // should never be nil

			s := ast.New(&ast.Selection{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, s)
			yylex.(*Lexer).parent = s
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:452
		{
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:456
		{
			ast.Append(yylex.(*Lexer).parent, ast.New(&ast.FatArrow{}, yyVAL.token))
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:460
		{
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:464
		{
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:468
		{
			yylex.(*Lexer).yydebug("class")
			gc := ast.New(&ast.ClassGuard{}, yyVAL.token)
			// If there is previous classguard, this one closes it, and we can yylex.(*Lexer).parent this new one, to that _parent_.
			prev := ast.UpTo(yylex.(*Lexer).parent, &ast.ClassGuard{})
			// If there is no previous one, look for promise guard, and yylex.(*Lexer).parent to that.
			if prev == nil {
				prev = ast.UpTo(yylex.(*Lexer).parent, &ast.PromiseGuard{})
			}
			// still not found, then body, bundle
			if prev == nil {
				prev = ast.UpTo(yylex.(*Lexer).parent, &ast.Bundle{})
			}
			if prev == nil {
				prev = ast.UpTo(yylex.(*Lexer).parent, &ast.Body{})
			}
			// re-yylex.(*Lexer).parent if found
			if prev != nil {
				switch prev.(type) {
				case *ast.Body, *ast.Bundle: // no .Parent() for these.
					yylex.(*Lexer).parent = prev
				default:
					yylex.(*Lexer).parent = prev.Parent()
				}
			}

			ast.Append(yylex.(*Lexer).parent, gc)
			yylex.(*Lexer).parent = gc
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:499
		{
			// awkward that these are the only ones here..?
			yylex.(*Lexer).yydebug("rval:IDENTIFIER", yyVAL.token)
			i := ast.New(&ast.Identifier{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, i)
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:506
		{
			yylex.(*Lexer).yydebug("rval:QSTRING", yyVAL.token)
			q := ast.New(&ast.Qstring{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, q)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:512
		{
			yylex.(*Lexer).yydebug("rval:NAKEDVAR", yyVAL.token)
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:516
		{
			yylex.(*Lexer).yydebug("rval:LIST", yyVAL.token)
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:520
		{
			yylex.(*Lexer).yydebug("rval:usefunction", yyVAL.token)
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:524
		{
		}
	case 110:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:528
		{
			yylex.(*Lexer).yydebug("list:{}", yyVAL.token)
			// empty list, add, but do not make parent
			l := ast.New(&ast.List{})
			ast.Append(yylex.(*Lexer).parent, l)
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:539
		{
			yylex.(*Lexer).yydebug("Litems:Litem", yyVAL.token)
			if _, ok := yylex.(*Lexer).parent.(*ast.List); !ok {
				l := ast.New(&ast.List{})
				ast.Append(yylex.(*Lexer).parent, l)
				yylex.(*Lexer).parent = l
			}
			l := ast.New(&ast.ListItem{}, yyVAL.token)
			ast.Append(yylex.(*Lexer).parent, l)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:550
		{
			yylex.(*Lexer).yydebug("Litems:Litems,Litem", yyVAL.token)
			if _, ok := yylex.(*Lexer).parent.(*ast.List); !ok {
				l := ast.New(&ast.List{})
				ast.Append(yylex.(*Lexer).parent, l)
				yylex.(*Lexer).parent = l
			}
			l := ast.New(&ast.ListItem{}, yyDollar[3].token)
			ast.Append(yylex.(*Lexer).parent, l)
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:561
		{
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:565
		{
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:568
		{
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:571
		{
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:574
		{
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:577
		{
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:581
		{
			f := ast.New(&ast.Function{}, ast.Token{})
			ast.Append(yylex.(*Lexer).parent, f)
			yylex.(*Lexer).parent = f

			ast.Append(yylex.(*Lexer).parent, ast.New(&ast.Identifier{}, yyVAL.token))
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:589
		{
			f := ast.New(&ast.Function{}, ast.Token{})
			ast.Append(yylex.(*Lexer).parent, f)
			yylex.(*Lexer).parent = f

			ast.Append(yylex.(*Lexer).parent, ast.New(&ast.NakedVar{}, yyVAL.token))
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:598
		{
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:603
		{
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:607
		{
		}
	case 127:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parse.y:611
		{
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:616
		{
			yylex.(*Lexer).yydebug("gaitems:gaitem", yyVAL.token)
			l := ast.New(&ast.GiveArgItem{}, yyVAL.token) // single arg
			ast.Append(yylex.(*Lexer).parent, l)
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:622
		{
			yylex.(*Lexer).yydebug("gaitems:gaitems,gaitem", yyDollar[3].token)
			l := ast.New(&ast.GiveArgItem{}, yyDollar[3].token) // multiple args
			ast.Append(yylex.(*Lexer).parent, l)
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:628
		{
		}
	}
	goto yystack /* stack new state and value */
}
