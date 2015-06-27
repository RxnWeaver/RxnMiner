package tokenizer

// WordProps qualifies a token whose type is one of `TokMayBeWord` or
// `TokWord`.
//
// It holds information regarding the so-called IOB (Inside, Outside,
// Beginning) status of the token, its lemma form (in case of a word),
// its part of speech (in case of a word), etc.
type WordProps struct {
	Word   Token  // The token being qualified.
	IsWord bool   // TokMayBeWord, etc. (false) vs. TokWord (true)
	IOB    byte   // 'B', 'I' or 'O'
	Lemma  string // Lemma form
	PoS    string // Part of Speech
}
