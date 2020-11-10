// Code generated by goyacc -o expr.go -p expr expr.y. DO NOT EDIT.

//line expr.y:2

package parser

import __yyfmt__ "fmt"

//line expr.y:3

import (
	//"log"

	"github.com/zerosnake0/gothrift/pkg/ast"
)

//line expr.y:14
type exprSymType struct {
	yys      int
	document ast.Document

	headers []ast.Header
	header  ast.Header

	definitions []ast.Definition
	definition  ast.Definition

	fieldType ast.FieldType

	simpleContainerType ast.SimpleContainerType
	cppType             *ast.CppType

	constValue ast.ConstValue

	annotations     *ast.Annotations
	annotationList  []ast.Annotation
	annotation      ast.Annotation
	annotationValue ast.AnnotationValue

	constListContent []ast.ConstListItem
	constMapContent  []ast.ConstMapItem

	enumDefList []ast.EnumDef
	enumDef     ast.EnumDef
	enumValue   ast.EnumValue

	senumDefList []ast.SenumDef
	senumDef     ast.SenumDef

	fieldList       []ast.Field
	field           ast.Field
	fieldIdentifier *ast.FieldIdentifier
	fieldValue      *ast.FieldValue

	xsdAttrs *ast.XsdAttrs

	extends *ast.Extends

	functionList []ast.Function
	function     ast.Function
	functionType ast.FunctionType

	throws *ast.Throws

	literal    ast.Literal
	identifier *ast.Identifier
	number     ast.Number
}

const INCLUDE = 57346
const NAMESPACE = 57347
const CONST = 57348
const BASETYPE = 57349
const MAP = 57350
const SET = 57351
const LIST = 57352
const CPPTYPE = 57353
const TYPEDEF = 57354
const ENUM = 57355
const SENUM = 57356
const STRUCTHEAD = 57357
const REQUIRED = 57358
const OPTIONAL = 57359
const XSDALL = 57360
const XSDOPTIONAL = 57361
const XSDNILLABLE = 57362
const XSDNAMESPACE = 57363
const XSDATTRS = 57364
const EXCEPTION = 57365
const SERVICE = 57366
const EXTENDS = 57367
const ONEWAY = 57368
const VOID = 57369
const THROWS = 57370
const Identifier = 57371
const Literal = 57372
const Number = 57373

var exprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'('",
	"')'",
	"'='",
	"','",
	"';'",
	"'*'",
	"'<'",
	"'>'",
	"'['",
	"']'",
	"'{'",
	"'}'",
	"':'",
	"'&'",
	"INCLUDE",
	"NAMESPACE",
	"CONST",
	"BASETYPE",
	"MAP",
	"SET",
	"LIST",
	"CPPTYPE",
	"TYPEDEF",
	"ENUM",
	"SENUM",
	"STRUCTHEAD",
	"REQUIRED",
	"OPTIONAL",
	"XSDALL",
	"XSDOPTIONAL",
	"XSDNILLABLE",
	"XSDNAMESPACE",
	"XSDATTRS",
	"EXCEPTION",
	"SERVICE",
	"EXTENDS",
	"ONEWAY",
	"VOID",
	"THROWS",
	"Identifier",
	"Literal",
	"Number",
}

var exprStatenames = [...]string{}

const exprEofCode = 1
const exprErrCode = 2
const exprInitialStackSize = 16

//line expr.y:838

//line yacctab:1
var exprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const exprPrivate = 57344

const exprLast = 198

var exprAct = [...]int{
	94, 50, 29, 78, 79, 53, 85, 144, 86, 134,
	85, 132, 86, 176, 179, 85, 160, 86, 33, 38,
	39, 40, 101, 137, 42, 126, 87, 106, 69, 28,
	33, 38, 39, 40, 97, 26, 52, 82, 81, 80,
	30, 82, 81, 80, 108, 55, 82, 81, 80, 65,
	142, 103, 30, 108, 171, 108, 108, 108, 155, 71,
	14, 74, 100, 27, 89, 150, 21, 22, 23, 24,
	73, 90, 59, 91, 49, 48, 47, 25, 20, 46,
	109, 104, 45, 58, 44, 43, 41, 165, 167, 117,
	162, 129, 130, 54, 63, 8, 9, 152, 116, 121,
	148, 122, 131, 124, 125, 175, 77, 72, 127, 64,
	61, 60, 149, 136, 120, 133, 135, 138, 115, 139,
	92, 70, 141, 143, 68, 56, 95, 96, 145, 114,
	159, 146, 123, 113, 147, 66, 51, 169, 154, 1,
	119, 166, 161, 62, 151, 128, 164, 140, 118, 93,
	57, 170, 112, 153, 156, 88, 67, 158, 157, 107,
	105, 102, 76, 99, 163, 98, 168, 75, 111, 172,
	110, 84, 174, 173, 83, 177, 37, 36, 35, 178,
	34, 32, 31, 13, 19, 18, 17, 16, 15, 12,
	11, 10, 4, 7, 6, 5, 3, 2,
}

var exprPact = [...]int{
	-1000, -1000, -1000, 77, 40, -1000, -1000, -1000, -9, 20,
	-1000, -1000, -1000, -1000, -3, -1000, -1000, -1000, -1000, -1000,
	43, -3, 42, 41, 39, 36, -1000, 33, 32, 31,
	-1000, -1000, -1000, 132, 132, -1000, -1000, -1000, 68, 68,
	115, 44, 29, 97, 96, 62, 95, 132, -1000, 129,
	-1000, -1000, -1000, 114, -16, 111, -3, 93, 27, 132,
	-1000, -1000, 92, -1000, -1000, -1000, 3, 21, -3, -1000,
	-3, 109, -1000, -1000, 119, 19, 7, -1000, 12, 119,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 127,
	122, 107, 68, 74, -1000, -1000, -1000, 132, -1000, 132,
	126, 132, -1000, 119, 10, -1000, 132, 61, 86, -1000,
	-2, -6, 119, -21, -3, -1000, -1000, 132, -1000, 9,
	-1000, -1000, 119, -38, -1000, -1000, 132, -1000, -3, -1000,
	-1000, -1000, -1000, 119, -1000, 84, -1000, -1000, 101, -1000,
	22, -1000, -1000, -1000, -1000, -1000, 80, -1000, 3, -1000,
	134, 15, -1000, 119, -1000, 124, -1000, 11, 57, 3,
	45, 54, -1000, -1000, 132, 133, 18, -1000, 119, -1000,
	132, 91, -1000, 8, 119, -1000, -1000, -1000, -1, -1000,
}

var exprPgo = [...]int{
	0, 197, 196, 195, 194, 193, 192, 191, 190, 189,
	188, 187, 186, 185, 184, 183, 2, 182, 181, 180,
	178, 177, 176, 5, 4, 174, 171, 170, 168, 167,
	165, 163, 162, 161, 3, 160, 159, 157, 1, 156,
	155, 152, 151, 150, 149, 148, 147, 146, 0, 145,
	144, 143, 142, 141, 140, 139,
}

var exprR1 = [...]int{
	0, 55, 1, 2, 2, 3, 3, 4, 5, 5,
	6, 6, 7, 7, 7, 8, 16, 16, 16, 17,
	18, 19, 19, 19, 20, 21, 22, 23, 23, 24,
	24, 24, 24, 24, 25, 27, 27, 26, 28, 28,
	38, 38, 39, 39, 40, 41, 48, 48, 48, 9,
	9, 9, 9, 9, 10, 11, 29, 29, 30, 31,
	31, 12, 32, 32, 33, 13, 34, 34, 35, 36,
	36, 49, 49, 49, 50, 50, 37, 37, 51, 51,
	52, 52, 53, 53, 42, 42, 14, 15, 43, 43,
	44, 44, 45, 54, 54, 46, 46, 47, 47,
}

var exprR2 = [...]int{
	0, 1, 2, 2, 0, 1, 1, 2, 4, 3,
	2, 0, 1, 1, 1, 6, 1, 1, 1, 2,
	2, 1, 1, 1, 7, 5, 5, 2, 0, 1,
	1, 1, 1, 1, 3, 3, 0, 3, 5, 0,
	3, 0, 2, 0, 3, 2, 1, 1, 0, 1,
	1, 1, 1, 1, 5, 6, 2, 0, 3, 3,
	1, 6, 2, 0, 2, 7, 2, 0, 11, 2,
	0, 1, 1, 0, 1, 0, 2, 0, 1, 0,
	1, 0, 1, 0, 4, 0, 6, 7, 2, 0,
	2, 0, 9, 1, 0, 1, 1, 4, 0,
}

var exprChk = [...]int{
	-1000, -55, -1, -2, -6, -3, -4, -5, 18, 19,
	-7, -8, -9, -15, 20, -10, -11, -12, -13, -14,
	38, 26, 27, 28, 29, 37, 44, 43, 9, -16,
	43, -17, -18, 21, -19, -20, -21, -22, 22, 23,
	24, 43, -16, 43, 43, 43, 43, 43, 43, 43,
	-38, 4, -38, -23, 25, -23, 10, -43, 39, 43,
	14, 14, -51, 32, 14, -38, 6, -39, 10, 44,
	10, -16, 14, 43, -38, -29, -32, 14, -34, -24,
	45, 44, 43, -25, -26, 12, 14, 5, -40, 43,
	-16, -16, 11, -44, -48, 7, 8, 15, -30, -31,
	43, 15, -33, 44, -34, -35, 15, -36, 45, -48,
	-27, -28, -41, 6, 7, 11, -23, 15, -45, -54,
	40, -38, -38, 6, -38, -48, 15, -38, -49, 30,
	31, 16, 13, -24, 15, -24, -48, 44, -16, -38,
	-46, -16, 41, -48, 45, -38, -16, -48, 16, 11,
	43, -50, 17, -24, 4, 43, -48, -34, -37, 6,
	5, -52, 33, -24, -47, 42, -53, 34, -38, 4,
	-42, 36, -48, -34, -38, 14, 5, -48, -34, 15,
}

var exprDef = [...]int{
	4, -2, 1, 11, 2, 3, 5, 6, 0, 0,
	10, 12, 13, 14, 0, 49, 50, 51, 52, 53,
	0, 0, 0, 0, 0, 0, 7, 0, 0, 0,
	16, 17, 18, 41, 41, 21, 22, 23, 28, 28,
	0, 89, 0, 0, 0, 79, 0, 41, 9, 0,
	19, 43, 20, 0, 0, 0, 0, 0, 0, 41,
	57, 63, 0, 78, 67, 8, 0, 0, 0, 27,
	0, 0, 91, 88, 48, 0, 0, 67, 70, 48,
	29, 30, 31, 32, 33, 36, 39, 40, 42, 0,
	0, 0, 28, 94, 54, 46, 47, 41, 56, 41,
	60, 41, 62, 48, 70, 66, 41, 73, 0, 15,
	0, 0, 48, 0, 0, 25, 26, 41, 90, 0,
	93, 55, 48, 0, 61, 64, 41, 86, 0, 71,
	72, 69, 34, 48, 37, 0, 44, 45, 0, 87,
	0, 95, 96, 58, 59, 65, 75, 35, 0, 24,
	0, 0, 74, 48, 67, 77, 38, 70, 81, 0,
	98, 83, 80, 76, 41, 0, 85, 82, 48, 67,
	41, 0, 92, 70, 48, 67, 97, 68, 70, 84,
}

var exprTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 17, 3,
	4, 5, 9, 3, 7, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 16, 8,
	10, 6, 11, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 12, 3, 13, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 14, 3, 15,
}

var exprTok2 = [...]int{
	2, 3, 18, 19, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, 38, 39, 40, 41, 42, 43, 44, 45,
}

var exprTok3 = [...]int{
	0,
}

var exprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	exprDebug        = 0
	exprErrorVerbose = false
)

type exprLexer interface {
	Lex(lval *exprSymType) int
	Error(s string)
}

type exprParser interface {
	Parse(exprLexer) int
	Lookahead() int
}

type exprParserImpl struct {
	lval  exprSymType
	stack [exprInitialStackSize]exprSymType
	char  int
}

func (p *exprParserImpl) Lookahead() int {
	return p.char
}

func exprNewParser() exprParser {
	return &exprParserImpl{}
}

const exprFlag = -1000

func exprTokname(c int) string {
	if c >= 1 && c-1 < len(exprToknames) {
		if exprToknames[c-1] != "" {
			return exprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func exprStatname(s int) string {
	if s >= 0 && s < len(exprStatenames) {
		if exprStatenames[s] != "" {
			return exprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func exprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !exprErrorVerbose {
		return "syntax error"
	}

	for _, e := range exprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + exprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := exprPact[state]
	for tok := TOKSTART; tok-1 < len(exprToknames); tok++ {
		if n := base + tok; n >= 0 && n < exprLast && exprChk[exprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if exprDef[state] == -2 {
		i := 0
		for exprExca[i] != -1 || exprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; exprExca[i] >= 0; i += 2 {
			tok := exprExca[i]
			if tok < TOKSTART || exprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if exprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += exprTokname(tok)
	}
	return res
}

func exprlex1(lex exprLexer, lval *exprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = exprTok1[0]
		goto out
	}
	if char < len(exprTok1) {
		token = exprTok1[char]
		goto out
	}
	if char >= exprPrivate {
		if char < exprPrivate+len(exprTok2) {
			token = exprTok2[char-exprPrivate]
			goto out
		}
	}
	for i := 0; i < len(exprTok3); i += 2 {
		token = exprTok3[i+0]
		if token == char {
			token = exprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = exprTok2[1] /* unknown char */
	}
	if exprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", exprTokname(token), uint(char))
	}
	return char, token
}

func exprParse(exprlex exprLexer) int {
	return exprNewParser().Parse(exprlex)
}

func (exprrcvr *exprParserImpl) Parse(exprlex exprLexer) int {
	var exprn int
	var exprVAL exprSymType
	var exprDollar []exprSymType
	_ = exprDollar // silence set and not used
	exprS := exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	exprstate := 0
	exprrcvr.char = -1
	exprtoken := -1 // exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		exprstate = -1
		exprrcvr.char = -1
		exprtoken = -1
	}()
	exprp := -1
	goto exprstack

ret0:
	return 0

ret1:
	return 1

exprstack:
	/* put a state and value onto the stack */
	if exprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", exprTokname(exprtoken), exprStatname(exprstate))
	}

	exprp++
	if exprp >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprS[exprp] = exprVAL
	exprS[exprp].yys = exprstate

exprnewstate:
	exprn = exprPact[exprstate]
	if exprn <= exprFlag {
		goto exprdefault /* simple state */
	}
	if exprrcvr.char < 0 {
		exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
	}
	exprn += exprtoken
	if exprn < 0 || exprn >= exprLast {
		goto exprdefault
	}
	exprn = exprAct[exprn]
	if exprChk[exprn] == exprtoken { /* valid shift */
		exprrcvr.char = -1
		exprtoken = -1
		exprVAL = exprrcvr.lval
		exprstate = exprn
		if Errflag > 0 {
			Errflag--
		}
		goto exprstack
	}

exprdefault:
	/* default state action */
	exprn = exprDef[exprstate]
	if exprn == -2 {
		if exprrcvr.char < 0 {
			exprrcvr.char, exprtoken = exprlex1(exprlex, &exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if exprExca[xi+0] == -1 && exprExca[xi+1] == exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			exprn = exprExca[xi+0]
			if exprn < 0 || exprn == exprtoken {
				break
			}
		}
		exprn = exprExca[xi+1]
		if exprn < 0 {
			goto ret0
		}
	}
	if exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			exprlex.Error(exprErrorMessage(exprstate, exprtoken))
			Nerrs++
			if exprDebug >= 1 {
				__yyfmt__.Printf("%s", exprStatname(exprstate))
				__yyfmt__.Printf(" saw %s\n", exprTokname(exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for exprp >= 0 {
				exprn = exprPact[exprS[exprp].yys] + exprErrCode
				if exprn >= 0 && exprn < exprLast {
					exprstate = exprAct[exprn] /* simulate a shift of "error" */
					if exprChk[exprstate] == exprErrCode {
						goto exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if exprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", exprS[exprp].yys)
				}
				exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if exprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", exprTokname(exprtoken))
			}
			if exprtoken == exprEofCode {
				goto ret1
			}
			exprrcvr.char = -1
			exprtoken = -1
			goto exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production exprn */
	if exprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", exprn, exprStatname(exprstate))
	}

	exprnt := exprn
	exprpt := exprp
	_ = exprpt // guard against "declared and not used"

	exprp -= exprR2[exprn]
	// exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if exprp+1 >= len(exprS) {
		nyys := make([]exprSymType, len(exprS)*2)
		copy(nyys, exprS)
		exprS = nyys
	}
	exprVAL = exprS[exprp+1]

	/* consult goto table to find next state */
	exprn = exprR1[exprn]
	exprg := exprPgo[exprn]
	exprj := exprg + exprS[exprp].yys + 1

	if exprj >= exprLast {
		exprstate = exprAct[exprg]
	} else {
		exprstate = exprAct[exprj]
		if exprChk[exprstate] != -exprn {
			exprstate = exprAct[exprg]
		}
	}
	// dummy call; replaced with literal code
	switch exprnt {

	case 1:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:167
		{
			setExprValDoc(exprlex, exprDollar[1].document)
		}
	case 2:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:173
		{
			exprVAL.document.Headers = exprDollar[1].headers
			exprVAL.document.Definitions = exprDollar[2].definitions
		}
	case 3:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:180
		{
			exprVAL.headers = append(exprDollar[1].headers, exprDollar[2].header)
		}
	case 4:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:184
		{
			exprVAL.headers = nil
		}
	case 5:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:190
		{
			exprVAL.header = exprDollar[1].header
		}
	case 6:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:194
		{
			exprVAL.header = exprDollar[1].header
		}
	case 7:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:200
		{
			exprVAL.header = ast.Include{
				Keyword: *exprDollar[1].identifier,
				Literal: exprDollar[2].literal,
			}
		}
	case 8:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line expr.y:209
		{
			exprVAL.header = ast.Namespace{
				Start:       exprDollar[1].identifier.Start,
				Language:    *exprDollar[2].identifier,
				Namespace:   *exprDollar[3].identifier,
				Annotations: exprDollar[4].annotations,
			}
		}
	case 9:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:218
		{
			exprVAL.header = ast.Namespace{
				Start:     exprDollar[1].identifier.Start,
				Language:  *exprDollar[2].identifier,
				Namespace: *exprDollar[3].identifier,
			}
		}
	case 10:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:228
		{
			exprVAL.definitions = append(exprDollar[1].definitions, exprDollar[2].definition)
		}
	case 11:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:232
		{
			exprVAL.definitions = nil
		}
	case 12:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:238
		{
			exprVAL.definition = exprDollar[1].definition
		}
	case 13:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:242
		{
			exprVAL.definition = exprDollar[1].definition
		}
	case 14:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:246
		{
			exprVAL.definition = exprDollar[1].definition
		}
	case 15:
		exprDollar = exprS[exprpt-6 : exprpt+1]
//line expr.y:252
		{
			exprVAL.definition = ast.Const{
				Start:     exprDollar[1].identifier.Start,
				FieldType: exprDollar[2].fieldType,
				Key:       *exprDollar[3].identifier,
				Eq:        exprDollar[4].identifier.Start,
				Value:     exprDollar[5].constValue,
				End:       exprDollar[6].identifier,
			}
		}
	case 16:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:265
		{
			exprVAL.fieldType = *exprDollar[1].identifier
		}
	case 17:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:269
		{
			exprVAL.fieldType = exprDollar[1].fieldType
		}
	case 18:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:273
		{
			exprVAL.fieldType = exprDollar[1].fieldType
		}
	case 19:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:279
		{
			exprVAL.fieldType = ast.BaseType{
				Identifier:  *exprDollar[1].identifier,
				Annotations: exprDollar[2].annotations,
			}
		}
	case 20:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:289
		{
			exprVAL.fieldType = ast.ContainerType{
				SimpleContainerType: exprDollar[1].simpleContainerType,
				Annotations:         exprDollar[2].annotations,
			}
		}
	case 21:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:298
		{
			exprVAL.simpleContainerType = exprDollar[1].simpleContainerType
		}
	case 22:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:302
		{
			exprVAL.simpleContainerType = exprDollar[1].simpleContainerType
		}
	case 23:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:306
		{
			exprVAL.simpleContainerType = exprDollar[1].simpleContainerType
		}
	case 24:
		exprDollar = exprS[exprpt-7 : exprpt+1]
//line expr.y:312
		{
			exprVAL.simpleContainerType = ast.MapType{
				Start:    exprDollar[1].identifier.Start,
				CppType:  exprDollar[2].cppType,
				LChevron: exprDollar[3].identifier.Start,
				Key:      exprDollar[4].fieldType,
				Comma:    exprDollar[5].identifier.Start,
				Value:    exprDollar[6].fieldType,
				RChevron: exprDollar[7].identifier.Start,
			}
		}
	case 25:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line expr.y:326
		{
			exprVAL.simpleContainerType = ast.SetType{
				Start:    exprDollar[1].identifier.Start,
				CppType:  exprDollar[2].cppType,
				LChevron: exprDollar[3].identifier.Start,
				Elem:     exprDollar[4].fieldType,
				RChevron: exprDollar[5].identifier.Start,
			}
		}
	case 26:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line expr.y:338
		{
			exprVAL.simpleContainerType = ast.ListType{
				Start:    exprDollar[1].identifier.Start,
				LChevron: exprDollar[2].identifier.Start,
				Elem:     exprDollar[3].fieldType,
				RChevron: exprDollar[4].identifier.Start,
				CppType:  exprDollar[5].cppType,
			}
		}
	case 27:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:350
		{
			exprVAL.cppType = &ast.CppType{
				Start:   exprDollar[1].identifier.Start,
				Literal: exprDollar[2].literal,
			}
		}
	case 28:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:357
		{
			exprVAL.cppType = nil
		}
	case 29:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:363
		{
			exprVAL.constValue = exprDollar[1].number
		}
	case 30:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:367
		{
			exprVAL.constValue = exprDollar[1].literal
		}
	case 31:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:371
		{
			exprVAL.constValue = *exprDollar[1].identifier
		}
	case 32:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:375
		{
			exprVAL.constValue = exprDollar[1].constValue
		}
	case 33:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:379
		{
			exprVAL.constValue = exprDollar[1].constValue
		}
	case 34:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:385
		{
			exprVAL.constValue = ast.ConstList{
				Node: ast.Node{
					Start: exprDollar[1].identifier.Start,
					End:   exprDollar[3].identifier.End,
				},
				Content: exprDollar[2].constListContent,
			}
		}
	case 35:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:397
		{
			exprVAL.constListContent = append(exprDollar[1].constListContent, ast.ConstListItem{
				ConstValue: exprDollar[2].constValue,
				End:        exprDollar[3].identifier,
			})
		}
	case 36:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:404
		{
			exprVAL.constListContent = nil
		}
	case 37:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:410
		{
			exprVAL.constValue = ast.ConstMap{
				Node: ast.Node{
					Start: exprDollar[1].identifier.Start,
					End:   exprDollar[3].identifier.End,
				},
				Content: exprDollar[2].constMapContent,
			}
		}
	case 38:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line expr.y:422
		{
			exprVAL.constMapContent = append(exprDollar[1].constMapContent, ast.ConstMapItem{
				Key:   exprDollar[2].constValue,
				Colon: exprDollar[3].identifier.Start,
				Value: exprDollar[4].constValue,
				End:   exprDollar[5].identifier,
			})
		}
	case 39:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:431
		{
			exprVAL.constMapContent = nil
		}
	case 40:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:437
		{
			exprVAL.annotations = &ast.Annotations{}
			exprVAL.annotations.Start = exprDollar[1].identifier.Start
			exprVAL.annotations.Annotations = exprDollar[2].annotationList
			exprVAL.annotations.End = exprDollar[3].identifier.End
		}
	case 41:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:444
		{
			exprVAL.annotations = nil
		}
	case 42:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:450
		{
			exprVAL.annotationList = append(exprDollar[1].annotationList, exprDollar[2].annotation)
		}
	case 43:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:454
		{
			exprVAL.annotationList = nil
		}
	case 44:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:460
		{
			exprVAL.annotation.Key = *exprDollar[1].identifier
			exprVAL.annotation.Value = exprDollar[2].annotationValue
			exprVAL.annotation.End = exprDollar[3].identifier
		}
	case 45:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:468
		{
			exprVAL.annotationValue.Start = exprDollar[1].identifier.Start
			exprVAL.annotationValue.Value = exprDollar[2].literal
		}
	case 46:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:475
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 47:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:479
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 48:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:483
		{
			exprVAL.identifier = nil
		}
	case 49:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:489
		{
			exprVAL.definition = exprDollar[1].definition
		}
	case 50:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:493
		{
			exprVAL.definition = exprDollar[1].definition
		}
	case 51:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:497
		{
			exprVAL.definition = exprDollar[1].definition
		}
	case 52:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:501
		{
			exprVAL.definition = exprDollar[1].definition
		}
	case 53:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:505
		{
			exprVAL.definition = exprDollar[1].definition
		}
	case 54:
		exprDollar = exprS[exprpt-5 : exprpt+1]
//line expr.y:511
		{
			exprVAL.definition = ast.TypeDef{
				Start:       exprDollar[1].identifier.Start,
				FieldType:   exprDollar[2].fieldType,
				Identifier:  *exprDollar[3].identifier,
				Annotations: exprDollar[4].annotations,
				End:         exprDollar[5].identifier,
			}
		}
	case 55:
		exprDollar = exprS[exprpt-6 : exprpt+1]
//line expr.y:523
		{
			exprVAL.definition = ast.Enum{
				Start:       exprDollar[1].identifier.Start,
				Identifier:  *exprDollar[2].identifier,
				LBrace:      exprDollar[3].identifier.Start,
				List:        exprDollar[4].enumDefList,
				RBrace:      exprDollar[5].identifier.Start,
				Annotations: exprDollar[6].annotations,
			}
		}
	case 56:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:536
		{
			exprVAL.enumDefList = append(exprDollar[1].enumDefList, exprDollar[2].enumDef)
		}
	case 57:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:540
		{
			exprVAL.enumDefList = nil
		}
	case 58:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:546
		{
			exprVAL.enumDef = ast.EnumDef{
				Value:       exprDollar[1].enumValue,
				Annotations: exprDollar[2].annotations,
				End:         exprDollar[3].identifier,
			}
		}
	case 59:
		exprDollar = exprS[exprpt-3 : exprpt+1]
//line expr.y:556
		{
			exprVAL.enumValue = ast.EnumValueWithNumber{
				Identifier: *exprDollar[1].identifier,
				Eq:         exprDollar[2].identifier.Start,
				Number:     exprDollar[3].number,
			}
		}
	case 60:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:564
		{
			exprVAL.enumValue = *exprDollar[1].identifier
		}
	case 61:
		exprDollar = exprS[exprpt-6 : exprpt+1]
//line expr.y:570
		{
			exprVAL.definition = ast.Senum{
				Start:       exprDollar[1].identifier.Start,
				Identifier:  *exprDollar[2].identifier,
				LBrace:      exprDollar[3].identifier.Start,
				List:        exprDollar[4].senumDefList,
				RBrace:      exprDollar[5].identifier.Start,
				Annotations: exprDollar[6].annotations,
			}
		}
	case 62:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:583
		{
			exprVAL.senumDefList = append(exprDollar[1].senumDefList, exprDollar[2].senumDef)
		}
	case 63:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:587
		{
			exprVAL.senumDefList = nil
		}
	case 64:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:593
		{
			exprVAL.senumDef = ast.SenumDef{
				Literal: exprDollar[1].literal,
				End:     exprDollar[2].identifier,
			}
		}
	case 65:
		exprDollar = exprS[exprpt-7 : exprpt+1]
//line expr.y:602
		{
			exprVAL.definition = ast.Struct{
				Head:        *exprDollar[1].identifier,
				Identifier:  *exprDollar[2].identifier,
				XsdAll:      exprDollar[3].identifier,
				LBrace:      exprDollar[4].identifier.Start,
				List:        exprDollar[5].fieldList,
				RBrace:      exprDollar[6].identifier.Start,
				Annotations: exprDollar[7].annotations,
			}
		}
	case 66:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:616
		{
			exprVAL.fieldList = append(exprDollar[1].fieldList, exprDollar[2].field)
		}
	case 67:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:620
		{
			exprVAL.fieldList = nil
		}
	case 68:
		exprDollar = exprS[exprpt-11 : exprpt+1]
//line expr.y:626
		{
			exprVAL.field = ast.Field{
				FieldIdentifier: exprDollar[1].fieldIdentifier,
				Requiredness:    exprDollar[2].identifier,
				FieldType:       exprDollar[3].fieldType,
				Reference:       exprDollar[4].identifier,
				Identifier:      *exprDollar[5].identifier,
				FieldValue:      exprDollar[6].fieldValue,
				XsdOptional:     exprDollar[7].identifier,
				XsdNillable:     exprDollar[8].identifier,
				XsdAttrs:        exprDollar[9].xsdAttrs,
				Annotations:     exprDollar[10].annotations,
				End:             exprDollar[11].identifier,
			}
		}
	case 69:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:644
		{
			exprVAL.fieldIdentifier = &ast.FieldIdentifier{
				Number: exprDollar[1].number,
				Colon:  exprDollar[2].identifier.Start,
			}
		}
	case 70:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:651
		{
			exprVAL.fieldIdentifier = nil
		}
	case 71:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:657
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 72:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:661
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 73:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:665
		{
			exprVAL.identifier = nil
		}
	case 74:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:671
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 75:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:675
		{
			exprVAL.identifier = nil
		}
	case 76:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:681
		{
			exprVAL.fieldValue = &ast.FieldValue{
				Start: exprDollar[1].identifier.Start,
				Value: exprDollar[2].constValue,
			}
		}
	case 77:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:688
		{
			exprVAL.fieldValue = nil
		}
	case 78:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:694
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 79:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:698
		{
			exprVAL.identifier = nil
		}
	case 80:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:704
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 81:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:708
		{
			exprVAL.identifier = nil
		}
	case 82:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:714
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 83:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:718
		{
			exprVAL.identifier = nil
		}
	case 84:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line expr.y:724
		{
			exprVAL.xsdAttrs = &ast.XsdAttrs{
				Start:     exprDollar[1].identifier.Start,
				LBrace:    exprDollar[2].identifier.Start,
				FieldList: exprDollar[3].fieldList,
				RBrace:    exprDollar[4].identifier.Start,
			}
		}
	case 85:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:733
		{
			exprVAL.xsdAttrs = nil
		}
	case 86:
		exprDollar = exprS[exprpt-6 : exprpt+1]
//line expr.y:739
		{
			exprVAL.definition = ast.Exception{
				Start:       exprDollar[1].identifier.Start,
				Identifier:  *exprDollar[2].identifier,
				LBrace:      exprDollar[3].identifier.Start,
				FieldList:   exprDollar[4].fieldList,
				RBrace:      exprDollar[5].identifier.Start,
				Annotations: exprDollar[6].annotations,
			}
		}
	case 87:
		exprDollar = exprS[exprpt-7 : exprpt+1]
//line expr.y:752
		{
			exprVAL.definition = ast.Service{
				Start:        exprDollar[1].identifier.Start,
				Identifier:   *exprDollar[2].identifier,
				Extends:      exprDollar[3].extends,
				LBrace:       exprDollar[4].identifier.Start,
				FunctionList: exprDollar[5].functionList,
				RBrace:       exprDollar[6].identifier.Start,
				Annotations:  exprDollar[7].annotations,
			}
		}
	case 88:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:766
		{
			exprVAL.extends = &ast.Extends{
				Start:      exprDollar[1].identifier.Start,
				Identifier: *exprDollar[2].identifier,
			}
		}
	case 89:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:773
		{
			exprVAL.extends = nil
		}
	case 90:
		exprDollar = exprS[exprpt-2 : exprpt+1]
//line expr.y:779
		{
			exprVAL.functionList = append(exprDollar[1].functionList, exprDollar[2].function)
		}
	case 91:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:783
		{
			exprVAL.functionList = nil
		}
	case 92:
		exprDollar = exprS[exprpt-9 : exprpt+1]
//line expr.y:789
		{
			exprVAL.function = ast.Function{
				Oneway:       exprDollar[1].identifier,
				FunctionType: exprDollar[2].functionType,
				Identifier:   *exprDollar[3].identifier,
				LChevron:     exprDollar[4].identifier.Start,
				FieldList:    exprDollar[5].fieldList,
				RChevron:     exprDollar[6].identifier.Start,
				Throws:       exprDollar[7].throws,
				Annotations:  exprDollar[8].annotations,
				End:          exprDollar[9].identifier,
			}
		}
	case 93:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:805
		{
			exprVAL.identifier = exprDollar[1].identifier
		}
	case 94:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:809
		{
			exprVAL.identifier = nil
		}
	case 95:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:815
		{
			exprVAL.functionType = exprDollar[1].fieldType
		}
	case 96:
		exprDollar = exprS[exprpt-1 : exprpt+1]
//line expr.y:819
		{
			exprVAL.functionType = *exprDollar[1].identifier
		}
	case 97:
		exprDollar = exprS[exprpt-4 : exprpt+1]
//line expr.y:825
		{
			exprVAL.throws = &ast.Throws{
				Start:     exprDollar[1].identifier.Start,
				LChevron:  exprDollar[2].identifier.Start,
				FieldList: exprDollar[3].fieldList,
				RChevron:  exprDollar[4].identifier.Start,
			}
		}
	case 98:
		exprDollar = exprS[exprpt-0 : exprpt+1]
//line expr.y:834
		{
			exprVAL.throws = nil
		}
	}
	goto exprstack /* stack new state and value */
}
