// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

import (
	"fmt"
)

// Document represents the entirety of input text of one logical
// document -- usually a file.
//
// It holds information about its sections, tokens in them, and the
// words and sentences that were recognised by other processors.  In
// case the document has associated training annotations, it holds
// them as well.
type Document struct {
	id     string // Must be unique within a run
	isTech bool   // Is this a technical document?
	input  map[string]string
	tokens map[string][]*TextToken
	words  map[string][]*Word
	annos  map[string][]*Annotation
	sents  map[string][]*Sentence
}

// NewDocument creates and initialises a document with the given
// identifier.
//
// It holds information about its sections.  It also holds information
// of their constituent tokens, words, sentences and annotations.
func NewDocument(id string) (*Document, error) {
	if id == "" {
		return nil, fmt.Errorf("Empty identifier given")
	}

	d := &Document{}
	d.id = id
	d.input = make(map[string]string, 2)
	d.tokens = make(map[string][]*TextToken, 2)
	d.words = make(map[string][]*Word, 2)
	d.annos = make(map[string][]*Annotation, 2)
	d.sents = make(map[string][]*Sentence, 2)

	return d, nil
}

// NewTechnicalDocument creates and initialises a document of
// technical nature, with the given identifier.
//
// It holds information about its sections.  It also holds information
// of their constituent tokens, words, sentences and annotations.
func NewTechnicalDocument(id string) (*Document, error) {
	d, err := NewDocument(id)
	if err != nil {
		return nil, err
	}

	d.isTech = true
	return d, nil
}

// SetInput registers the input text of the given section of the
// document.
func (d *Document) SetInput(sec, input string) error {
	if sec == "" || input == "" {
		return fmt.Errorf("Empty section name or body given.")
	}

	d.input[sec] = input
	return nil
}

// Input answers the registered input text of the given section, if
// one exists.
func (d *Document) Input(sec string) (string, error) {
	if s, ok := d.input[sec]; ok {
		return s, nil
	}

	return "", fmt.Errorf("No input text for section : %s", sec)
}

// Tokenize breaks the text in the various sections of the document
// into quasi-atomic tokens.
//
// These tokens can be matched against any available annotations.
// They can also be combined into logical words for named entity
// recognition and part of speech recognition purposes.
func (d *Document) Tokenize() {
	var ti *TextTokenIterator
	var err error

	for sec, inp := range d.input {
		ti = NewTextTokenIterator(inp)
		var toks []*TextToken
		for err = ti.MoveNext(); err == nil; err = ti.MoveNext() {
			toks = append(toks, ti.Item())
		}
		d.tokens[sec] = toks
	}
}

// AssembleSentences builds sentences the text tokens obtained as a
// result of tokenization of the sections in the document.
func (d *Document) AssembleSentences() {
	var si *SentenceIterator
	var err error

	for sec, toks := range d.tokens {
		si = NewSentenceIterator(toks)
		var sents []*Sentence
		for err = si.MoveNext(); err == nil; err = si.MoveNext() {
			sents = append(sents, si.Item())
		}
		d.sents[sec] = sents
	}
}

// Annotate records the given annotation against the applicable
// sequence of tokens in the appropriate section of the document.
//
// It creates or updates a `Word` corresponding to the text in the
// annotation.  The annotation can be for one of: (a) part of speech
// ("POS"), (b) lemma ("LEM") or (c) class/category ("CLS").
func (d *Document) Annotate(a *Annotation, what string) error {
	toks, ok := d.tokens[a.Section]
	if !ok {
		return fmt.Errorf("Annotation for unrecognised section : %s", a.Section)
	}

	_, err := d.annotateWord(a, what, toks)
	if err != nil {
		return err
	}

	annos := d.annos[a.Section]
	annos = append(annos, a)
	d.annos[a.Section] = annos

	return nil
}

//

func (d *Document) annotateWord(a *Annotation, what string, toks []*TextToken) (*Word, error) {
	bidx := -1
	eidx := -1
	for i, t := range toks {
		if t.Begin() == a.Begin {
			bidx = i
			break
		}
	}
	if bidx == -1 {
		return nil, fmt.Errorf("Annotation could not be matched : %v", *a)
	}
	l := len(toks)
	for i := bidx; i < l; i++ {
		if toks[i].End() == a.End {
			eidx = i
			break
		}
	}
	if eidx == -1 {
		return nil, fmt.Errorf("Annotation could not be matched : %v", *a)
	}

	var w *Word
	found := false
	words := d.words[a.Section]
	for _, tw := range words {
		if tw.Begin() == a.Begin && tw.End() == a.End {
			w = tw
			found = true
			break
		}
	}

	if !found {
		w = newWord(a.Entity, a.Begin, a.End)
	}
	switch what {
	case "POS":
		w.pos = a.Property
	case "LEM":
		w.lemma = a.Property
	case "CLS":
		w.class = a.Property
	default:
		return nil, fmt.Errorf("Unknown annotation type : %s", what)
	}

	if !found {
		words = append(words, w)
		d.words[a.Section] = words
	}

	return w, nil
}

// SectionTokens answers recognised tokens in the given section.
func (d *Document) SectionTokens(sec string) []*TextToken {
	if v, ok := d.tokens[sec]; ok {
		return v
	}

	return nil
}

// SectionTokenCount answers the number of recognised tokens in the
// given section.
func (d *Document) SectionTokenCount(sec string) (int, error) {
	if v, ok := d.tokens[sec]; ok {
		return len(v), nil
	}

	return -1, fmt.Errorf("Unknown section : %s", sec)
}

// SectionSentences answers assembled sentences in the given section.
func (d *Document) SectionSentences(sec string) []*Sentence {
	if v, ok := d.sents[sec]; ok {
		return v
	}

	return nil
}

// SectionSentenceCount answers the number of assembled sentences in
// the given section.
func (d *Document) SectionSentenceCount(sec string) (int, error) {
	if v, ok := d.sents[sec]; ok {
		return len(v), nil
	}

	return -1, fmt.Errorf("Unknown section : %s", sec)
}

// SectionWords answers recognised words in the given section.
func (d *Document) SectionWords(sec string) []*Word {
	if v, ok := d.words[sec]; ok {
		return v
	}

	return nil
}

// SectionWordCount answers the number of recognised words in the
// given section.
func (d *Document) SectionWordCount(sec string) (int, error) {
	if v, ok := d.words[sec]; ok {
		return len(v), nil
	}

	return -1, fmt.Errorf("Unknown section : %s", sec)
}

// SectionAnnotations answers registered annotations for the given
// section.
func (d *Document) SectionAnnotations(sec string) []*Annotation {
	if v, ok := d.annos[sec]; ok {
		return v
	}

	return nil
}

// SectionAnnotationCount answers the number of registered annotations
// in the given section.
func (d *Document) SectionAnnotationCount(sec string) (int, error) {
	if v, ok := d.annos[sec]; ok {
		return len(v), nil
	}

	return -1, fmt.Errorf("Unknown section : %s", sec)
}
