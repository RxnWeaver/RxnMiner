package tokenizer

import (
	"bytes"
	"io"
	"strings"
)

// TextToken represents a piece of text extracted from a larger input.
// It holds information regarding its beginning and ending offsets in
// the input text.  A text token may span the entire input.
type TextToken struct {
	text  string
	begin int
	end   int
	ttype TokenType
}

func (tt *TextToken) Text() string {
	return tt.text
}

func (tt *TextToken) Begin() int {
	return tt.begin
}

func (tt *TextToken) End() int {
	return tt.end
}

func (tt *TextToken) Type() TokenType {
	return tt.ttype
}

// TextTokenIterator helps in retrieving consecutive text tokens from
// an input text.
type TextTokenIterator struct {
	in  string
	ct  *TextToken
	idx int
	buf bytes.Buffer
}

// NewTextTokenIterator creates and initialises a token iterator over
// the given input text.
func NewTextTokenIterator(input string) *TextTokenIterator {
	ti := &TextTokenIterator{}
	ti.in = input
	return ti
}

// NewTextTokenIteratorWithOffset creates and initialises a token
// iterator over the given input text.
//
// It treats the given offset - rather than 0 - as the starting index
// from which to track all subsequent indices.
func NewTextTokenIteratorWithOffset(input string, n int) *TextTokenIterator {
	ti := &TextTokenIterator{}
	ti.in = input
	ti.idx = n
	return ti
}

// Item answers the current token.
func (ti *TextTokenIterator) Item() *TextToken {
	return ti.ct
}

// MoveNext detects the next token in the input, should one be
// available.  Otherwise, it answers an error describing the problem.
func (ti *TextTokenIterator) MoveNext() error {
	inStr := false
	inNum := false
	rd := strings.NewReader(ti.in[ti.idx:])
	begin, end := ti.idx, ti.idx

	for r, n, err := rd.ReadRune(); err == nil; r, n, err = rd.ReadRune() {
		tt := TypeOfRune(r)
		switch tt {
		case TokTerm, TokMayBeTerm, TokPause,
			TokParenOpen, TokParenClose,
			TokBracketOpen, TokBracketClose,
			TokBraceOpen, TokBraceClose,
			TokPunct,
			TokSymbol,
			TokSpace:
			{
				if ti.buf.Len() > 0 {
					ti.ct = &TextToken{string(ti.in[begin:end]), begin, end - 1, TokMayBeWord}
					ti.buf.Reset()
					ti.idx = end
					return err
				}

				begin = end
				end = begin + n
				ti.ct = &TextToken{string(r), begin, end - 1, tt}
				ti.idx = end
				return err
			}

		case TokNumber:
			{
				if inStr {
					ti.ct = &TextToken{string(ti.in[begin:end]), begin, end - 1, TokMayBeWord}
					ti.buf.Reset()
					ti.idx = end
					return err
				} else {
					inNum = true
					end += n
				}
				inStr = false
			}

		case TokLetter:
			{
				if inNum {
					ti.ct = &TextToken{string(ti.in[begin:end]), begin, end - 1, TokMayBeWord}
					ti.buf.Reset()
					ti.idx = end
					return err
				} else {
					end += n
				}
				inStr = true
				inNum = false
			}

		default:
			end += n
		}

		ti.buf.WriteRune(r)
	}

	if ti.buf.Len() > 0 {
		ti.ct = &TextToken{string(ti.in[begin:end]), begin, end - 1, TokMayBeWord}
		ti.buf.Reset()
		ti.idx = end
		return nil
	}

	return io.EOF
}
