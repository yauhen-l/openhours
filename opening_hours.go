//line opening_hours.y:4
package main

import __yyfmt__ "fmt"

//line opening_hours.y:4
import (
	"fmt"
)

var wholeWeek = [...]int{0, 1, 2, 3, 4, 5, 6}

type TimeRange struct {
	Start int
	End   int
}

func NewTimeRange(start, end int) TimeRange {
	return TimeRange{Start: start, End: end}
}

func makeWeekly() [][]TimeRange {
	weekly := make([][]TimeRange, 7)
	for i, _ := range weekly {
		weekly[i] = make([]TimeRange, 0)
	}
	return weekly
}

func appendWeeklyTimeRanges(weekly [][]TimeRange, weeks []int, trs []TimeRange) [][]TimeRange {
	for _, w := range weeks {
		weekly[w] = append(weekly[w], trs...)
	}
	return weekly
}

func mergeWeeklyTimeRanges(w1, w2 [][]TimeRange) [][]TimeRange {
	for i, _ := range w1 {
		w1[i] = append(w1[i], w2[i]...)
	}
	return w1
}

//line opening_hours.y:45
type yySymType struct {
	yys        int
	num        int
	nums       []int
	continuous bool
	tr         TimeRange
	trs        []TimeRange
	weekly     [][]TimeRange
}

const WEEKDAY = 57346
const ALWAYS = 57347

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"WEEKDAY",
	"ALWAYS",
	"'-'",
	"':'",
	"';'",
	"' '",
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
const yyMaxDepth = 200

//line opening_hours.y:192

/*
func CompileOpeningHours(s string) ([]TimeRange, err) {
		defer
}
*/

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 33
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 54

var yyAct = [...]int{

	13, 6, 11, 10, 7, 2, 14, 5, 31, 27,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24,
	29, 35, 26, 34, 33, 32, 40, 10, 1, 37,
	25, 3, 36, 39, 4, 43, 41, 38, 15, 16,
	17, 18, 19, 20, 21, 22, 23, 24, 30, 9,
	28, 8, 42, 12,
}
var yyPact = [...]int{

	-1, -1000, -1000, 14, -1000, 0, -1000, -1000, 10, -2,
	19, 18, 16, -1000, 27, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1, 27, -1000, 23,
	-1000, 27, 22, 27, 27, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 2, 53, 52, 6, 0, 51, 7, 50, 49,
	1, 48, 34, 31, 5, 30, 28,
}
var yyR1 = [...]int{

	0, 16, 16, 14, 15, 15, 13, 13, 13, 12,
	7, 6, 6, 8, 8, 10, 11, 11, 9, 1,
	2, 3, 5, 4, 4, 4, 4, 4, 4, 4,
	4, 4, 4,
}
var yyR2 = [...]int{

	0, 0, 1, 2, 0, 2, 1, 3, 1, 1,
	2, 1, 3, 0, 2, 2, 0, 2, 3, 3,
	1, 1, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1,
}
var yyChk = [...]int{

	-1000, -16, -14, -13, -12, -7, -10, 5, -6, -9,
	4, -1, -2, -5, -4, 11, 12, 13, 14, 15,
	16, 17, 18, 19, 20, -15, 8, 9, -8, 10,
	-11, 10, 6, 6, 7, -4, -14, -10, -7, -10,
	4, -1, -3, -5,
}
var yyDef = [...]int{

	1, -2, 2, 4, 6, 0, 8, 9, 13, 16,
	11, 0, 0, 20, 0, 23, 24, 25, 26, 27,
	28, 29, 30, 31, 32, 3, 0, 0, 10, 0,
	15, 0, 0, 0, 0, 22, 5, 7, 14, 17,
	12, 18, 19, 21,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 9, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 10, 6, 3, 3, 11, 12,
	13, 14, 15, 16, 17, 18, 19, 20, 7, 8,
}
var yyTok2 = [...]int{

	2, 3, 4, 5,
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
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
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
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
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
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
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
			yychar = -1
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
		//line opening_hours.y:72
		{
			res := yylex.(*Lexer).parseResult
			if res == nil {
				yylex.(*Lexer).parseResult = yyDollar[1].weekly
			}
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line opening_hours.y:80
		{
			yyVAL.weekly = mergeWeeklyTimeRanges(yyDollar[1].weekly, yyDollar[2].weekly)
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line opening_hours.y:84
		{
			yyVAL.weekly = makeWeekly()
		}
	case 5:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line opening_hours.y:86
		{
			{
				yyVAL.weekly = yyDollar[2].weekly
			}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:90
		{
			yyVAL.weekly = appendWeeklyTimeRanges(makeWeekly(), wholeWeek, yyDollar[1].trs)
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line opening_hours.y:93
		{
			yyVAL.weekly = appendWeeklyTimeRanges(makeWeekly(), yyDollar[1].nums, yyDollar[3].trs)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:96
		{
			yyVAL.weekly = appendWeeklyTimeRanges(makeWeekly(), wholeWeek, yyDollar[1].trs)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:99
		{
			yyVAL.trs = []TimeRange{NewTimeRange(0, 1440)}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line opening_hours.y:103
		{
			yyVAL.nums = append(yyDollar[1].nums, yyDollar[2].nums...)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:109
		{
			yyVAL.nums = []int{yyDollar[1].num}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line opening_hours.y:114
		{
			if yyDollar[1].num > yyDollar[3].num {
				yylex.Error(fmt.Sprintf("invalid weekdays range: %d - %d\n", yyDollar[1].num, yyDollar[3].num))
			}
			yyVAL.nums = make([]int, 0)
			for i := yyDollar[1].num; i <= yyDollar[3].num; i++ {
				yyVAL.nums = append(yyVAL.nums, i)
			}
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line opening_hours.y:126
		{
			yyVAL.nums = []int{}
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line opening_hours.y:131
		{
			yyVAL.nums = yyDollar[2].nums
		}
	case 15:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line opening_hours.y:136
		{
			yyVAL.trs = append(yyDollar[2].trs, yyDollar[1].tr)
		}
	case 16:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line opening_hours.y:139
		{
			yyVAL.trs = []TimeRange{}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line opening_hours.y:141
		{
			yyVAL.trs = yyDollar[2].trs
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line opening_hours.y:145
		{
			ts := NewTimeRange(yyDollar[1].num, yyDollar[3].num)

			if ts.Start >= ts.End {
				yylex.Error(fmt.Sprintf("invalid timerange: %v\n", ts))
			}
			yyVAL.tr = ts
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line opening_hours.y:156
		{
			t := yyDollar[1].num + yyDollar[3].num
			if t > 1440 { // > 24:00
				yylex.Error(fmt.Sprintf("invalid time: %d\n", t))
			}
			yyVAL.num = t
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:166
		{
			if yyDollar[1].num < 0 || yyDollar[1].num > 24 {
				yylex.Error(fmt.Sprintf("invalid hour: %d\n", yyDollar[1].num))
			}
			yyVAL.num = yyDollar[1].num * 60
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:175
		{
			if yyDollar[1].num < 0 || yyDollar[1].num > 59 {
				yylex.Error(fmt.Sprintf("invalid minutes: %d\n", yyDollar[1].num))
			}
			yyVAL.num = yyDollar[1].num
		}
	case 22:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line opening_hours.y:184
		{
			yyVAL.num = yyDollar[1].num*10 + yyDollar[2].num
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:187
		{
			yyVAL.num = 0
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:187
		{
			yyVAL.num = 1
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:187
		{
			yyVAL.num = 2
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:188
		{
			yyVAL.num = 3
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:188
		{
			yyVAL.num = 4
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:188
		{
			yyVAL.num = 5
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:189
		{
			yyVAL.num = 6
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:189
		{
			yyVAL.num = 7
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:189
		{
			yyVAL.num = 8
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line opening_hours.y:189
		{
			yyVAL.num = 9
		}
	}
	goto yystack /* stack new state and value */
}
