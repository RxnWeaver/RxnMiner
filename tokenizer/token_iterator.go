package tokenizer

// Token represents a piece of text extracted from a larger input.  It
// holds information regarding its beginning and ending offsets in the
// input text.  A token may span the entire input.
type Token interface {
	Text() string
	Begin() int
	End() int
	Type() TokenType
}

// TokenIterator helps in retrieving consecutive tokens.  Side effects
// of the iteration happen in `MoveNext`.  Therefore, `Item` can be
// called multiple times, should there be such a need.
type TokenIterator interface {
	MoveNext() error
	Item() Token
}
