package tokenizer

import (
// "fmt"
)

// Sentence represents a logical sentence.
//
// It holds information about its text, its offsets and its
// constituent text tokens.
type Sentence struct {
	token  TextToken    // Actual text and its properties
	tokens []*TextToken // Constituent text tokens
}

// newSentence creates and initialises a sentence with its text and
// offsets.
func newSentence(text string, b int, e int, toks []*TextToken) *Sentence {
	s := &Sentence{}
	s.token.text = text
	s.token.begin = b
	s.token.end = e
	s.token.ttype = TokSentence
	s.tokens = append(s.tokens, toks...)

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

func (s *Sentence) Tokens() []*TextToken {
	return s.tokens
}
