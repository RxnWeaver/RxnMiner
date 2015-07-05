package tokenizer

import (
	"io"
	"strings"
	"unicode"
)

// Sentence represents a logical sentence.
//
// It holds information about its text, its offsets and its
// constituent text tokens.
type Sentence struct {
	token   TextToken // Actual text and its properties
	bTokIdx int       // Index of beginning token of this sentence
	eTokIdx int       // Index of ending token of this sentence
}

// newSentence creates and initialises a sentence with its text and
// offsets.
func newSentence(text string, b, e, btok, etok int) *Sentence {
	s := &Sentence{}
	s.token.text = text
	s.token.begin = b
	s.token.end = e
	s.token.ttype = TokSentence
	s.bTokIdx = btok
	s.eTokIdx = etok

	return s
}

func (s *Sentence) Text() string {
	return s.token.text
}

func (s *Sentence) Begin() int {
	return s.token.begin
}

func (s *Sentence) End() int {
	return s.token.end
}

func (s *Sentence) Type() TokenType {
	return s.token.ttype
}

func (s *Sentence) BeginToken() int {
	return s.bTokIdx
}

func (s *Sentence) EndToken() int {
	return s.eTokIdx
}

// groupIndex helps in keeping track of opening and closing grouping
// tokens: parentheses, brackets and braces.
type groupIndex struct {
	tokIndex int
	tokType  TokenType
}

// SentenceIterator helps in assembling consecutive sentences from the
// underlying text tokens.
type SentenceIterator struct {
	toks        []*TextToken
	isTech      bool
	cs          *Sentence
	idx         int
	idxTerm     int
	buf         string
	inTerm      bool
	inMayBeTerm bool
	inTermSpc   bool
	grpStack    []groupIndex
}

// NewSentenceIterator creates and initialises a sentence iterator
// over the given text tokens.
func NewSentenceIterator(toks []*TextToken) *SentenceIterator {
	si := &SentenceIterator{}
	si.toks = toks
	return si
}

// NewTechnicalSentenceIterator creates and initialises a sentence
// iterator in technical mode, over the given text tokens.
func NewTechnicalSentenceIterator(toks []*TextToken) *SentenceIterator {
	si := &SentenceIterator{}
	si.toks = toks
	si.isTech = true
	return si
}

// Item answers the current sentence.
func (si *SentenceIterator) Item() *Sentence {
	return si.cs
}

// MoveNext detects the next sentence formed from the given input
// tokens, should one be available.  Otherwise, it answers an error
// describing the problem.
func (si *SentenceIterator) MoveNext() error {
	begin, end := si.idx, si.idx
	size := len(si.toks)

	commonProc := func() {
		si.cs = newSentence(si.buf,
			si.toks[begin].Begin(), si.toks[si.idxTerm].End(),
			begin, si.idxTerm)
		si.buf = ""
		si.inTerm = false
		si.inTermSpc = false
		si.inMayBeTerm = false
	}

	for end < size {
		t := si.toks[end]

		switch t.ttype {
		case TokSpace:
			{
				switch {
				case si.inTerm, si.inMayBeTerm:
					si.inTerm = false
					si.inMayBeTerm = false
					si.inTermSpc = true
					end++

				case si.inTermSpc:
					end++
					begin = end

				default:
					si.buf += t.text
					end++
				}
			}

		case TokTerm, TokMayBeTerm:
			{
				var prev string
				if end-1 >= 0 {
					prevt := si.toks[end-1]
					prev = strings.ToLower(prevt.text)
					if prevt.ttype == TokSymbol || prevt.ttype == TokPunct {
						si.inTerm = false
						si.inMayBeTerm = true
					} else {
						if _, ok := NonTermAbbrevs[prev]; ok {
							si.inTerm = false
							si.inMayBeTerm = true
						} else {
							si.inTerm = true
						}
					}
				} else {
					si.inTerm = false
				}
				si.inTermSpc = false
				si.buf += t.text
				si.idxTerm = end
				end++
			}

		case TokParenOpen, TokBracketOpen, TokBraceOpen:
			si.grpStack = append(si.grpStack, groupIndex{end, t.ttype})
			if si.inTerm || si.inTermSpc {
				commonProc()
				si.idx = end
				return nil
			}
			end++

		case TokParenClose, TokBracketClose, TokBraceClose:
			{
				if len(si.grpStack) > 0 {
					lsta := si.grpStack[len(si.grpStack)-1]
					switch t.ttype {
					case TokParenClose:
						if lsta.tokType == TokParenOpen {
							si.grpStack = si.grpStack[:len(si.grpStack)-1]
						}
					case TokBracketClose:
						if lsta.tokType == TokBracketOpen {
							si.grpStack = si.grpStack[:len(si.grpStack)-1]
						}
					case TokBraceClose:
						if lsta.tokType == TokBraceOpen {
							si.grpStack = si.grpStack[:len(si.grpStack)-1]
						}
					}
				}

				if si.inTerm || si.inTermSpc {
					nt := si.nextNonSpaceToken(end)
					if nt != -1 {
						ntt := si.toks[nt].ttype
						if ntt != TokPause &&
							ntt != TokPunct &&
							ntt != TokSymbol {
							var r rune
							for _, r = range t.text {
								break
							}
							if unicode.IsUpper(r) {
								commonProc()
								si.idx = end + 1
								return nil
							}
						}
					}
				}
				end++
			}

		default:
			{
				if si.inTermSpc {
					var r rune
					for _, r = range t.text {
						break
					}
					if unicode.IsUpper(r) {
						commonProc()
						si.idx = end
						return nil
					} else {
						if end > si.idxTerm {
							for i := si.idxTerm + 1; i < end; i++ {
								si.buf += si.toks[i].Text()
							}
						}
					}
				}

				si.inTerm = false
				si.inTermSpc = false
				si.inMayBeTerm = false
				si.buf += t.text
				end++
			}
		}
	}

	if len(si.buf) > 0 {
		skip := false
		chars := strings.Split(si.buf, "")
		if len(chars) == 1 {
			rdr := strings.NewReader(si.buf)
			r, _, _ := rdr.ReadRune()
			if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
				skip = true
			}
		}
		if !skip {
			commonProc()
			si.idx = end
			return nil
		}
	}

	return io.EOF
}

// nextNonSpaceToken answers the index of the first token after that
// at the given index that represents a non-space token.  If none such
// exists, it answers -1.
func (si *SentenceIterator) nextNonSpaceToken(idx int) int {
	l := len(si.toks)
	for i := idx + 1; i < l; i++ {
		if si.toks[i].ttype != TokSpace {
			return i
		}
	}

	return -1
}
