// Code generated by goyacc -o parse.go parse.y. DO NOT EDIT.

//line parse.y:2
package parser

import __yyfmt__ "fmt"

//line parse.y:2

import (
//"fmt"
)

func addStatement(yylex yyLexer, stmt Statement) {
	yylex.(*lex).Query.Statements = append(yylex.(*lex).Query.Statements, stmt)
}

//line parse.y:14
type yySymType struct {
	yys        int
	stmt       Statement
	stmts      Statements
	selectStmt *SelectStatement

	str     string
	query   Query
	field   *Field
	fields  Fields
	source  *Source
	int     int
	int64   int64
	float64 float64
	bool    bool

	expr *ExprNode
}

const SELECT = 57346
const FROM = 57347
const WHERE = 57348
const AS = 57349
const STAR = 57350
const EQ = 57351
const NEQ = 57352
const LT = 57353
const LTE = 57354
const GT = 57355
const GTE = 57356
const TLIKE = 57357
const LEFTC = 57358
const RIGHTC = 57359
const NAME = 57360
const INTEGER = 57361
const FLOAT = 57362
const STRING = 57363
const COMMA = 57364
const SEMICOLON = 57365
const OR = 57366
const AND = 57367
const ADD = 57368
const DEC = 57369
const TIME = 57370
const DIV = 57371
const MOD = 57372

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"SELECT",
	"FROM",
	"WHERE",
	"AS",
	"STAR",
	"EQ",
	"NEQ",
	"LT",
	"LTE",
	"GT",
	"GTE",
	"TLIKE",
	"LEFTC",
	"RIGHTC",
	"NAME",
	"INTEGER",
	"FLOAT",
	"STRING",
	"COMMA",
	"SEMICOLON",
	"OR",
	"AND",
	"ADD",
	"DEC",
	"TIME",
	"DIV",
	"MOD",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parse.y:197

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 58

var yyAct = [...]int{
	5, 16, 14, 47, 30, 29, 29, 7, 22, 9,
	30, 29, 15, 23, 17, 19, 20, 18, 31, 10,
	13, 12, 11, 17, 19, 20, 18, 41, 44, 25,
	42, 43, 45, 46, 48, 21, 6, 24, 4, 32,
	6, 49, 33, 34, 35, 36, 37, 38, 39, 26,
	27, 28, 6, 40, 8, 3, 2, 1,
}

var yyPact = [...]int{
	34, -1000, -1000, -1000, 1, -1000, -4, 30, -9, -1000,
	22, 22, 22, 22, -20, -4, 33, -1000, -1000, -1000,
	-1000, 9, -1000, 1, -1000, 10, -1000, -1000, -1000, -4,
	-4, -14, 5, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	46, -1000, -1000, -1000, -1000, -1000, -19, -1000, -1000, -1000,
}

var yyPgo = [...]int{
	0, 57, 56, 55, 54, 7, 53, 2, 0, 1,
	37, 39,
}

var yyR1 = [...]int{
	0, 1, 2, 3, 3, 3, 5, 5, 4, 4,
	4, 4, 4, 10, 10, 6, 6, 8, 8, 7,
	7, 7, 7, 11, 11, 11, 11, 11, 11, 11,
	9, 9, 9, 9,
}

var yyR2 = [...]int{
	0, 1, 1, 5, 3, 1, 1, 3, 1, 2,
	2, 2, 2, 0, 2, 1, 1, 2, 0, 3,
	3, 3, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1,
}

var yyChk = [...]int{
	-1000, -1, -2, -3, 4, -8, 6, -5, -4, 8,
	18, 21, 20, 19, -7, 16, -9, 18, 21, 19,
	20, 5, -8, 22, -10, 7, -10, -10, -10, 25,
	24, -7, -11, 9, 10, 11, 12, 13, 14, 15,
	-6, 18, 21, -5, 18, -7, -7, 17, -9, -8,
}

var yyDef = [...]int{
	18, -2, 1, 2, 0, 5, 0, 18, 6, 8,
	13, 13, 13, 13, 17, 0, 0, 30, 31, 32,
	33, 0, 4, 0, 9, 0, 10, 11, 12, 0,
	0, 0, 0, 23, 24, 25, 26, 27, 28, 29,
	18, 15, 16, 7, 14, 21, 22, 19, 20, 3,
}

var yyTok1 = [...]int{
	1,
}

var yyTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30,
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
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:64
		{
			addStatement(yylex, yyDollar[1].stmt)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:68
		{
			yyVAL.stmt = yyDollar[1].selectStmt
		}
	case 3:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parse.y:73
		{
			sele := &SelectStatement{
				Fields:         yyDollar[2].fields,
				Source:         yyDollar[4].source,
				WhereCondition: yyDollar[5].expr,
			}
			yyVAL.selectStmt = sele
		}
	case 4:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:81
		{
			sele := &SelectStatement{
				Fields:         yyDollar[2].fields,
				WhereCondition: yyDollar[3].expr,
			}

			yyVAL.selectStmt = sele
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:89
		{
			sele := &SelectStatement{

				WhereCondition: yyDollar[1].expr,
			}

			yyVAL.selectStmt = sele
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:99
		{
			yyVAL.fields = []*Field{yyDollar[1].field}
		}
	case 7:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:102
		{
			yyVAL.fields = append(yyDollar[3].fields, yyDollar[1].field)
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:107
		{
			yyVAL.field = &Field{Type: "ALL"}
		}
	case 9:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:110
		{
			yyVAL.field = &Field{Type: "NAME", Name: yyDollar[1].str, Alias: yyDollar[2].str}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:113
		{
			yyVAL.field = &Field{Type: "STRING", ValueString: yyDollar[1].str, Alias: yyDollar[2].str}
		}
	case 11:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:116
		{
			yyVAL.field = &Field{Type: "FLOAT", ValueFloat: yyDollar[1].float64, Alias: yyDollar[2].str}
		}
	case 12:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:119
		{
			yyVAL.field = &Field{Type: "INT", ValueInt: yyDollar[1].int64, Alias: yyDollar[2].str}
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parse.y:123
		{
			yyVAL.str = ""
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:126
		{
			yyVAL.str = yyDollar[2].str
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:131
		{
			yyVAL.source = &Source{Name: yyDollar[1].str}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:134
		{
			yyVAL.source = &Source{Name: yyDollar[1].str}
		}
	case 17:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parse.y:139
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parse.y:142
		{
			yyVAL.expr = nil
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:147
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:150
		{
			yyVAL.expr = &ExprNode{Type: BinaryNode, Left: yyDollar[1].expr, Op: yyDollar[2].int, Right: yyDollar[3].expr}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:153
		{
			yyVAL.expr = &ExprNode{Type: BinaryNode, Left: yyDollar[1].expr, Op: yyDollar[2].int, Right: yyDollar[3].expr}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parse.y:156
		{
			yyVAL.expr = &ExprNode{Type: BinaryNode, Left: yyDollar[1].expr, Op: yyDollar[2].int, Right: yyDollar[3].expr}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:161
		{
			yyVAL.int = yyDollar[1].int
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:164
		{
			yyVAL.int = yyDollar[1].int
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:167
		{
			yyVAL.int = yyDollar[1].int
		}
	case 26:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:170
		{
			yyVAL.int = yyDollar[1].int
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:173
		{
			yyVAL.int = yyDollar[1].int
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:176
		{
			yyVAL.int = yyDollar[1].int
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:179
		{
			yyVAL.int = yyDollar[1].int
		}
	case 30:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:184
		{
			yyVAL.expr = &ExprNode{Type: FieldNode, Name: yyDollar[1].str}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:187
		{
			yyVAL.expr = &ExprNode{Type: StringNode, StrVal: yyDollar[1].str}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:190
		{
			yyVAL.expr = &ExprNode{Type: IntegerNode, IntVal: yyDollar[1].int64}
		}
	case 33:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parse.y:193
		{
			yyVAL.expr = &ExprNode{Type: FloatNode, FloVal: yyDollar[1].float64}
		}
	}
	goto yystack /* stack new state and value */
}
