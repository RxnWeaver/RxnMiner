package tokenizer

//

type Token interface {
	Text() string
	Begin() int
	End() int
}

//

type TokenIterator interface {
	MoveNext() error
	Item() Token
}
