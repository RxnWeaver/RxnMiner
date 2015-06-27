package tokenizer

// Sentence represents a logical sentence, with possibly multiple
// words in it.
//
// It holds information about its constituent words as well.
type Sentence struct {
	token TextToken // Actual text and its properties
	words []Word    // Constituent words, possibly just one
}

// NewSentence creates and initialises a sentence with its text and
// offsets.
func NewSentence(text string, b int, e int) *Sentence {
	s := &Sentence{}
	s.token.text = text
	s.token.begin = b
	s.token.end = e
	s.token.ttype = TokSentence
	s.words = make([]Word, 4) // May need to tune this later.

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

func (s *Sentence) Words() []Word {
	return s.words
}
