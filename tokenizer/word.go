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
	lemma string    // Lemma form
	pos   string    // Part of Speech
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

func (w *Word) Lemma() string {
	return w.lemma
}

func (w *Word) POS() string {
	return w.pos
}
