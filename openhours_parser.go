//line openhours.y:2
package openhours

import __yyfmt__ "fmt"

//line openhours.y:2
import "fmt"

//line openhours.y:7
type yySymType struct {
	yys     int
	num     int
	nums    []int
	tr      TimeRange
	trs     []TimeRange
	weekly  Weekly
	monthly Monthly
}

const WEEKDAY = 57346
const MONTH = 57347
const ALWAYS = 57348
const SEP = 57349

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"WEEKDAY",
	"MONTH",
	"ALWAYS",
	"SEP",
	"'-'",
	"':'",
	"';'",
	"','",
	"'0'",
	"'1'",
	"'2'",
	"'3'",
	"'4'",
	"'5'",
	"'6'",
	"'7'",
	"'8'",
	"'9'",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line openhours.y:225

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 14,
	9, 33,
	-2, 35,
}

const yyNprod = 47
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 120

var yyAct = [...]int{

	7, 48, 14, 2, 10, 16, 11, 5, 40, 38,
	34, 9, 15, 6, 4, 61, 17, 30, 44, 42,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	41, 36, 47, 49, 43, 62, 51, 35, 51, 32,
	52, 31, 50, 53, 56, 55, 59, 14, 57, 60,
	54, 15, 45, 1, 29, 3, 15, 6, 4, 39,
	13, 33, 8, 63, 19, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 15, 46, 37, 12, 58, 18,
	0, 0, 19, 20, 21, 22, 23, 24, 25, 26,
	27, 28, 15, 0, 0, 0, 0, 0, 0, 0,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
}
var yyPact = [...]int{

	52, -1000, -1000, 7, -1000, 34, 32, -1000, -1, 30,
	-1000, 23, -2, -3, -1000, 22, 11, 98, 9, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 70, 88, -1000, 98, 98, 98, -1000, 47, -1000,
	98, 40, 98, -1000, 98, 8, 28, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 88, -1000,
}
var yyPgo = [...]int{

	0, 5, 79, 78, 6, 16, 1, 77, 11, 76,
	7, 62, 61, 60, 4, 59, 55, 3, 54, 0,
	53, 52,
}
var yyR1 = [...]int{

	0, 20, 20, 17, 18, 18, 16, 16, 16, 16,
	16, 16, 16, 10, 11, 11, 12, 12, 19, 19,
	19, 8, 7, 7, 9, 9, 21, 21, 14, 15,
	15, 13, 1, 2, 3, 4, 6, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5,
}
var yyR2 = [...]int{

	0, 0, 1, 2, 0, 3, 1, 1, 3, 5,
	3, 3, 1, 2, 1, 3, 0, 2, 1, 3,
	1, 2, 1, 3, 0, 2, 0, 2, 2, 0,
	2, 3, 3, 1, 1, 1, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -20, -17, -16, 6, -10, 5, -19, -11, -8,
	-14, -4, -7, -13, -6, 4, -1, -5, -2, 12,
	13, 14, 15, 16, 17, 18, 19, 20, 21, -18,
	10, 7, 7, -12, 11, 7, 8, -9, 11, -15,
	11, 8, 8, -5, 9, -21, 5, -19, -6, -19,
	-10, -6, -14, -4, -8, -14, 4, -1, -3, -6,
	-17, 7, 7, -19,
}
var yyDef = [...]int{

	1, -2, 2, 4, 6, 7, 0, 12, 16, 18,
	20, 14, 24, 29, -2, 22, 0, 0, 0, 37,
	38, 39, 40, 41, 42, 43, 44, 45, 46, 3,
	26, 0, 0, 13, 0, 0, 0, 21, 0, 28,
	0, 0, 0, 36, 0, 0, 8, 10, 33, 11,
	17, 35, 19, 15, 25, 30, 23, 31, 32, 34,
	5, 27, 0, 9,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 11, 8, 3, 3, 12, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 9, 10,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7,
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

	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:36
		{
			res := yylex.(*Lexer).parseResult
			if res == nil {
				yylex.(*Lexer).parseResult = yyDollar[1].monthly
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line openhours.y:44
		{
			yyVAL.monthly = mergeMonthly(yyDollar[1].monthly, yyDollar[2].monthly)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line openhours.y:48
		{
			yyVAL.monthly = make(Monthly)
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:50
		{
			yyVAL.monthly = yyDollar[3].monthly
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:55
		{
			yyVAL.monthly = setWeekly(makeMonthly(any, []int{any}), anyTime)
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:60
		{
			yyVAL.monthly = setWeekly(makeMonthly(any, yyDollar[1].nums), anyTime)
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:65
		{
			yyVAL.monthly = setWeekly(makeMonthly(yyDollar[3].num, yyDollar[1].nums), anyTime)
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line openhours.y:70
		{
			yyVAL.monthly = setWeekly(makeMonthly(yyDollar[3].num, yyDollar[1].nums), yyDollar[5].weekly)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:75
		{
			yyVAL.monthly = setWeekly(makeMonthly(any, yyDollar[1].nums), yyDollar[3].weekly)
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:80
		{
			yyVAL.monthly = setWeekly(makeMonthly(yyDollar[1].num, []int{any}), yyDollar[3].weekly)
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:85
		{
			yyVAL.monthly = setWeekly(makeMonthly(any, []int{any}), yyDollar[1].weekly)
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line openhours.y:92
		{
			yyVAL.nums = append(yyDollar[1].nums, yyDollar[2].nums...)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:98
		{
			yyVAL.nums = []int{yyDollar[1].num}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:101
		{
			if yyDollar[1].num > yyDollar[3].num {
				yylex.Error(fmt.Sprintf("invalid days range: %d - %d\n", yyDollar[1].num, yyDollar[3].num))
			}
			yyVAL.nums = make([]int, 0)
			for i := yyDollar[1].num; i <= yyDollar[3].num; i++ {
				yyVAL.nums = append(yyVAL.nums, i)
			}
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line openhours.y:113
		{
			yyVAL.nums = []int{}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line openhours.y:116
		{
			yyVAL.nums = yyDollar[2].nums
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:121
		{
			yyVAL.weekly = appendWeeklyTimeRanges(make(Weekly), yyDollar[1].nums, wholeDay)
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:123
		{
			yyVAL.weekly = appendWeeklyTimeRanges(make(Weekly), yyDollar[1].nums, yyDollar[3].trs)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:125
		{
			yyVAL.weekly = appendWeeklyTimeRanges(make(Weekly), wholeWeek, yyDollar[1].trs)
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line openhours.y:129
		{
			yyVAL.nums = append(yyDollar[1].nums, yyDollar[2].nums...)
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:135
		{
			yyVAL.nums = []int{yyDollar[1].num}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:140
		{
			if yyDollar[1].num > yyDollar[3].num {
				yylex.Error(fmt.Sprintf("invalid weekdays range: %d - %d\n", yyDollar[1].num, yyDollar[3].num))
			}
			yyVAL.nums = make([]int, 0)
			for i := yyDollar[1].num; i <= yyDollar[3].num; i++ {
				yyVAL.nums = append(yyVAL.nums, i)
			}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line openhours.y:152
		{
			yyVAL.nums = []int{}
		}
	case 25:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line openhours.y:157
		{
			yyVAL.nums = yyDollar[2].nums
		}
	case 28:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line openhours.y:164
		{
			yyVAL.trs = append(yyDollar[2].trs, yyDollar[1].tr)
		}
	case 29:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line openhours.y:167
		{
			yyVAL.trs = []TimeRange{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line openhours.y:169
		{
			yyVAL.trs = yyDollar[2].trs
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:173
		{
			ts := NewTimeRange(yyDollar[1].num, yyDollar[3].num)

			if ts.Start >= ts.End {
				yylex.Error(fmt.Sprintf("invalid timerange: %v\n", ts))
			}
			yyVAL.tr = ts
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line openhours.y:184
		{
			t := yyDollar[1].num + yyDollar[3].num
			if t > 1440 { // > 24:00
				yylex.Error(fmt.Sprintf("invalid time: %d\n", t))
			}
			yyVAL.num = t
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:194
		{
			if yyDollar[1].num < 0 || yyDollar[1].num > 24 {
				yylex.Error(fmt.Sprintf("invalid hour: %d\n", yyDollar[1].num))
			}
			yyVAL.num = yyDollar[1].num * 60
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:203
		{
			if yyDollar[1].num < 0 || yyDollar[1].num > 59 {
				yylex.Error(fmt.Sprintf("invalid minutes: %d\n", yyDollar[1].num))
			}
			yyVAL.num = yyDollar[1].num
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:212
		{
			if yyDollar[1].num < 1 || yyDollar[1].num > 31 {
				yylex.Error(fmt.Sprintf("invalid day: %d\n", yyDollar[1].num))
			}
			yyVAL.num = yyDollar[1].num
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line openhours.y:219
		{
			yyVAL.num = yyDollar[1].num*10 + yyDollar[2].num
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 0
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 1
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 2
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 3
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 4
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 5
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 6
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 7
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 8
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line openhours.y:222
		{
			yyVAL.num = 9
		}
	}
	goto yystack /* stack new state and value */
}
