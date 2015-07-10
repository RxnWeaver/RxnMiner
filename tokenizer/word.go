// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

// Word represents a token whose type is one of `TokMayBeWord` or
// `TokWord`, and qualifies it.
//
// It holds information regarding the so-called IOB (Inside, Outside,
// Beginning) status of the token, its lemma form (in case of a word),
// its part of speech (in case of a word), etc.
type Word struct {
	token TextToken // Actual text and its properties
	iob   byte      // 'B', 'I' or 'O'
	pos   string    // Part of Speech
	lemma string    // Lemma form
	class string    // Assigned after learning
}

// newWord creates and initialises a word with its properties set to
// reasonable defaults.
func newWord(text string, b int, e int) *Word {
	w := &Word{}
	w.token.text = text
	w.token.begin = b
	w.token.end = e
	w.token.ttype = TokMayBeWord
	w.iob = 'O'

	return w
}

func (w *Word) Text() string {
	return w.token.text
}

func (w *Word) Begin() int {
	return w.token.begin
}

func (w *Word) End() int {
	return w.token.end
}

func (w *Word) Type() TokenType {
	return w.token.ttype
}

func (w *Word) IOB() byte {
	return w.iob
}

func (w *Word) POS() string {
	return w.pos
}

func (w *Word) Lemma() string {
	return w.lemma
}

func (w *Word) Class() string {
	return w.class
}
