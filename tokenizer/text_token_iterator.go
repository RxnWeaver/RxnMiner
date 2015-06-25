package tokenizer

import (
	"bytes"
	"io"
	"strings"
	"unicode"
)

// TextToken represents a piece of text extracted from a larger input.
// It holds information regarding its beginning and ending offsets in
// the input text.  A text token may span the entire input.
type TextToken struct {
	text  string
	begin int
	end   int
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

// TextTokenIterator helps in retrieving consecutive text tokens from
// an input text.
type TextTokenIterator struct {
	in  string
	ct  *TextToken
	idx int
	buf bytes.Buffer
}

func NewTextTokenIterator(input string) *TextTokenIterator {
	ti := &TextTokenIterator{}
	ti.in = input
	return ti
}

func NewTextTokenIteratorWithRunningIndex(input string, idx int) *TextTokenIterator {
	ti := &TextTokenIterator{}
	ti.in = input
	ti.idx = idx
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
	rd := strings.NewReader(ti.in[ti.idx:])
	begin, end := ti.idx, ti.idx

loop:
	for r, n, err := rd.ReadRune(); err == nil; r, n, err = rd.ReadRune() {
		switch {
		case r == '.', r == '!', r == '?', unicode.IsSpace(r):
			{
				if r == '.' && inNum {
					ti.buf.WriteRune(r)
					end += n
					inNum = false
					continue loop
				}

				if ti.buf.Len() > 0 {
					ti.ct = &TextToken{ti.in[begin:end], begin, end - 1}
					ti.buf.Reset()
					ti.idx = end
					return err
				}

				begin = end
				end = begin + n
				ti.ct = &TextToken{string(r), begin, end - 1}
				ti.idx = end
				return err
			}

		case unicode.IsPunct(r), unicode.IsSymbol(r):
			{
				if ti.buf.Len() > 0 {
					ti.ct = &TextToken{ti.in[begin:end], begin, end - 1}
					ti.buf.Reset()
					ti.idx = end
					return err
				}

				begin = end
				end = begin + n
				ti.ct = &TextToken{string(r), begin, end - 1}
				ti.idx = end
				return err
			}

		case unicode.IsNumber(r):
			{
				if inStr {
					ti.ct = &TextToken{ti.in[begin:end], begin, end - 1}
					ti.buf.Reset()
					ti.idx = end
					return err
				} else {
					inNum = true
					end += n
				}
				inStr = false
			}

		case unicode.IsLetter(r):
			{
				if inNum {
					ti.ct = &TextToken{ti.in[begin:end], begin, end - 1}
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
		ti.ct = &TextToken{ti.in[begin:end], begin, end - 1}
		ti.buf.Reset()
		ti.idx = end
		return nil
	}

	return io.EOF
}
