package tokenizer

import (
	"io"
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

// SentenceIterator helps in assembling consecutive sentences from the
// underlying text tokens.
type SentenceIterator struct {
	toks     []*TextToken
	cs       *Sentence
	idx      int
	idxTerm  int
	buf      string
	inTerm   bool
	inTermSp bool
}

// NewSentenceIterator creates and initialises a sentence iterator
// over the given text tokens.
func NewSentenceIterator(toks []*TextToken) *SentenceIterator {
	si := &SentenceIterator{}
	si.toks = toks
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

	for end < size {
		t := si.toks[end]
		switch t.ttype {
		case TokSpace:
			{
				switch {
				case si.inTerm:
					si.inTerm = false
					si.inTermSp = true
					end++

				case si.inTermSp:
					end++
					begin = end

				default:
					si.buf += t.text
					end++
				}
			}

		case TokTerm, TokMayBeTerm: // TODO(js): Handle 'Mr.'/'etc.'/...
			si.inTerm = true
			si.inTermSp = false
			si.buf += t.text
			si.idxTerm = end
			end++

		default:
			{
				if si.inTermSp {
					var r rune
					for _, r = range t.text {
						break
					}
					if unicode.IsUpper(r) {
						si.cs = newSentence(si.buf,
							si.toks[begin].Begin(), si.toks[si.idxTerm].End(),
							begin, si.idxTerm)
						si.buf = ""
						si.idx = end
						si.inTerm = false
						si.inTermSp = false
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
				si.inTermSp = false
				si.buf += t.text
				end++
			}
		}
	}

	if len(si.buf) > 0 {
		si.cs = newSentence(si.buf,
			si.toks[begin].Begin(), si.toks[end-1].End(),
			begin, end)
		si.buf = ""
		si.idx = end
		si.inTerm = false
		si.inTermSp = false
		return nil
	}

	return io.EOF
}
