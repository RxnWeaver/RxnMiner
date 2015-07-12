### tokenizer -- A text tokenizer and sentence assembler
`tokenizer` is a small Go package for:

- splitting input text into quasi-atomic tokens,
- assembling sentences from those tokens, and
- annotating a set of one or more (consecutive) token(s) as words or phrases.

Rather than split the input text into sentences first, and tokenize the sentences next, `tokenizer` *assembles* them from tokens.  For the purposes of **RxnMiner** (the containing project of this package) - which processes technical text - the conventional approach (as followed by most leading NLP engines) produced too many incorrect sentence breaks, leading to mis-applied annotations downstream.  Hence this inverted design.

`tokenizer` is rule-based.

### Installation

Preferred:

```sh
go get -u 'github.com/RxnWeaver/RxnMiner/tokenizer'
cd $GOPATH/src/github.com/RxnWeaver/RxnMiner
git checkout <tag>
go test -v ./...
go install ./...
```

where `<tag>` represents the most-recently tagged release.

For the adventurous:

```sh
go get -u 'github.com/RxnWeaver/RxnMiner/tokenizer'
```

### Status

This package is being used already, and is - consequently - reasonably battle-tested.  The repository itself includes tests for over **7,000** real life test input texts.

Abbreviation handling is both English-centric and very limited.  This is likely to improve in future.

See open issues on GitHub for currently known issues and corner cases.

### Usage

In most cases, instantiating a document is a good place to start.  Here is a trivial example.

```go
doc, err := tokenizer.NewDocument("MyDoc-1")
if err != nil {
    return err
}
doc.SetInput("Section-1", someText)
doc.Tokenize()
doc.AssembleSentences()

toks := doc.SectionTokens("Section-1")
for _, tok := range toks {
    fmt.Printf("%v\n", tok)
}

sents := doc.SectionSentences("Section-1")
for _, sent := range sents {
    fmt.Printf("%v\n", sent)
}
```

The tokens obtained by splitting the given input text can, of course, be used for purposes other than sentence assembly as well.

Refer to tests for a few more interesting usages, and examples of applying annotations and extracting the resulting words and phrases.
