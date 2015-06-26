package tokenizer

import (
	"bytes"
	"io"
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
	in  []byte
	ct  *TextToken
	idx int
	buf bytes.Buffer
}

// NewTextTokenIterator creates and initialises a token iterator over
// the given input text.
func NewTextTokenIterator(input []byte) *TextTokenIterator {
	ti := &TextTokenIterator{}
	ti.in = input
	return ti
}

// NewTextTokenIteratorWithOffset creates and initialises a token
// iterator over the given input text.
//
// It treats the given offset - rather than 0 - as the starting index
// from which to track all subsequent indices.
func NewTextTokenIteratorWithOffset(input []byte, n int) *TextTokenIterator {
	ti := &TextTokenIterator{}
	ti.in = input
	ti.idx = n
	return ti
}

func (ti *TextTokenIterator) Item() Token {
	return ti.ct
}

// MoveNext is an iterator method towards implementation of the
// `TokenIterator` interface.
func (ti *TextTokenIterator) MoveNext() error {
	inStr := false
	inNum := false
	rd := bytes.NewReader(ti.in[ti.idx:])
	begin, end := ti.idx, ti.idx

	for r, n, err := rd.ReadRune(); err == nil; r, n, err = rd.ReadRune() {
		tt := TypeOfRune(r)
		switch tt {
		case TokTerm, TokPause, TokGroupOpen, TokGroupClose, TokPunct, TokSymbol, TokSpace:
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
