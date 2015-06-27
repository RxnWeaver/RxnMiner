package tokenizer

import (
	"fmt"
)

// Sentence represents a logical sentence, with possibly multiple
// words in it.
//
// It holds information about its constituent words as well.
type Sentence struct {
	token TextToken // Actual text and its properties
	words []*Word   // Constituent words, possibly just one
}

// NewSentence creates and initialises a sentence with its text and
// offsets.
func NewSentence(text string, b int, e int) *Sentence {
	s := &Sentence{}
	s.token.text = text
	s.token.begin = b
	s.token.end = e
	s.token.ttype = TokSentence
	s.words = make([]*Word, 4) // TODO(js): May need to tune this later.

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

func (s *Sentence) Words() []*Word {
	return s.words
}

func (s *Sentence) WordCount() int {
	return len(s.words)
}

//

func (s *Sentence) WordAt(idx int) (*Word, error) {
	if idx > len(s.words) {
		return nil, fmt.Errorf("Index our of bounds : %d", idx)
	}

	return s.words[idx], nil
}

func (s *Sentence) AddWord(w *Word) error {
	if w == nil {
		return fmt.Errorf("Given word is nil")
	}

	s.words = append(s.words, w)
	return nil
}

func (s *Sentence) AddWords(ws []*Word) error {
	s.words = append(s.words, ws...)
	return nil
}

func (s *Sentence) AddWordAt(idx int, w *Word) error {
	if idx > len(s.words) {
		return fmt.Errorf("Index our of bounds : %d", idx)
	}

	t := s.words[idx:]
	s.words = append(s.words[:idx], w)
	s.words = append(s.words, t...)
	return nil
}

func (s *Sentence) RemoveWordAt(idx int) error {
	if idx > len(s.words) {
		return fmt.Errorf("Index our of bounds : %d", idx)
	}

	s.words = append(s.words[:idx], s.words[idx+1:]...)
	return nil
}
