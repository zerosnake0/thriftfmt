package format

import (
	"reflect"

	"github.com/zerosnake0/gothrift/pkg/ast"
)

func (f *Formatter) startBrace(lbrace, next ast.Pos, singleLine bool) {
	f.encodeKeyword(lbrace, "{")
	f.newScope(next, false)
}

func (f *Formatter) endBrace(rbrace ast.Pos) {
	f.forward(false, rbrace)
	f.endScope()
	f.encodeKeyword(rbrace, "}")
}

func (f *Formatter) startChevron(lchevron, next ast.Pos, singleLine bool) {
	f.encodeKeyword(lchevron, "(")
	f.newScope(next, singleLine)
}

func (f *Formatter) endChevron(rchevron ast.Pos) {
	f.forwardAndEmptySep(false, rchevron)
	f.endScope()
	f.encodeKeyword(rchevron, ")")
}

func (f *Formatter) startBracket(lbracket, next ast.Pos, singleLine bool) {
	f.encodeKeyword(lbracket, "[")
	f.newScope(next, singleLine)
}

func (f *Formatter) endBracket(rbracket ast.Pos) {
	f.forwardAndEmptySep(false, rbracket)
	f.endScope()
	f.encodeKeyword(rbracket, "]")
}

type coupleOpt struct {
	startFunc   func(left, next ast.Pos, singleLine bool)
	endFunc     func(right ast.Pos)
	attachFirst bool
}

func (f *Formatter) encodeCouple(opt *coupleOpt, l, r ast.Pos, sep string, array interface{}, cb func(span ast.Span)) {
	v := reflect.ValueOf(array)
	singleLine := l.Line == r.Line
	if v.Len() == 0 {
		opt.startFunc(l, r, singleLine)
	} else {
		span := v.Index(0).Addr().Interface().(ast.Span)
		opt.startFunc(l, span.StartPos(), singleLine)
		for i := 0; i < v.Len(); i++ {
			item := v.Index(i).Addr().Interface().(ast.Span)
			start := item.StartPos()
			if i > 0 {
				if sep != "" {
					f.print(sep)
				}
				f.outputRemainingComments(start)
			}

			attach := false
			if i == 0 && opt.attachFirst {
				attach = true
				if f.cmtIdx > 0 {
					cmt := f.Doc.Comments[f.cmtIdx-1]
					if l.Before(cmt.StartPos()) {
						attach = false
					}
				}
			}
			if attach {
				f.forwardAndEmptySep(false, item.StartPos())
			} else {
				f.forward(false, item.StartPos())
			}
			if !singleLine {
				f.newLineIfNot()
			}
			cb(item)
		}
	}
	f.outputRemainingComments(r)
	if !singleLine {
		f.newLineIfNot()
	}
	opt.endFunc(r)
}

func (f *Formatter) encodeChevron(l, r ast.Pos, sep string, array interface{}, cb func(span ast.Span)) {
	f.encodeCouple(&coupleOpt{
		startFunc:   f.startChevron,
		endFunc:     f.endChevron,
		attachFirst: true,
	}, l, r, sep, array, cb)
}

func (f *Formatter) encodeBrace(l, r ast.Pos, sep string, array interface{}, cb func(span ast.Span)) {
	f.encodeCouple(&coupleOpt{
		startFunc: f.startBrace,
		endFunc:   f.endBrace,
	}, l, r, sep, array, cb)
}

func (f *Formatter) encodeBracket(l, r ast.Pos, sep string, array interface{}, cb func(span ast.Span)) {
	f.encodeCouple(&coupleOpt{
		startFunc:   f.startBracket,
		endFunc:     f.endBracket,
		attachFirst: true,
	}, l, r, sep, array, cb)
}
