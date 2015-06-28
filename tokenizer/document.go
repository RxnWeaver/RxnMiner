package tokenizer

import (
	"fmt"
)

// Document represents the entirety of input text of one logical
// document -- usually a file.
//
// It holds information about its sections.  In case the document has
// associated training annotations, it holds them as well.
type Document struct {
	id       string // Must be unique within a run
	title    string
	abstract string
	body     string
	tTokens  []*TextToken
	aTokens  []*TextToken
	bTokens  []*TextToken
	tWords   []*Word
	aWords   []*Word
	bWords   []*Word
	tAnnos   []*Annotation
	aAnnos   []*Annotation
	bAnnos   []*Annotation
}

// NewDocument creates and initialises a document with the given
// identifier.
//
// It holds information about its sections: currently title, abstract
// and body.  It also holds information of their constituent words and
// annotations.
func NewDocument(id string) (*Document, error) {
	if id == "" {
		return nil, fmt.Errorf("Empty identifier given")
	}

	d := &Document{}
	d.id = id

	return d, nil
}

// SetTitle provides the text of the title of the document.
func (d *Document) SetTitle(title string) error {
	if title == "" {
		return fmt.Errorf("Empty title given")
	}

	d.title = title
	return nil
}

// SetAbstract provides the text of the abstract of the document.
func (d *Document) SetAbstract(abstract string) error {
	if abstract == "" {
		return fmt.Errorf("Empty abstract given")
	}

	d.abstract = abstract
	return nil
}

// SetBody provides the text of the body of the document.
func (d *Document) SetBody(body string) error {
	if body == "" {
		return fmt.Errorf("Empty body given")
	}

	d.body = body
	return nil
}

// Tokenize breaks the text in the title, abstract and body into
// quasi-atomic tokens.
//
// These tokens can be matched against any available annotations.
// They can also be combined into logical words for named entity
// recognition and part of speech recognition purposes.
func (d *Document) Tokenize() {
	var ti *TextTokenIterator
	var err error

	if len(d.title) > 0 {
		ti = NewTextTokenIterator(d.title)
		for err = ti.MoveNext(); err == nil; err = ti.MoveNext() {
			d.tTokens = append(d.tTokens, ti.Item())
		}
	}

	if len(d.abstract) > 0 {
		ti = NewTextTokenIterator(d.abstract)
		for err = ti.MoveNext(); err == nil; err = ti.MoveNext() {
			d.aTokens = append(d.aTokens, ti.Item())
		}
	}

	if len(d.body) > 0 {
		ti = NewTextTokenIterator(d.body)
		for err = ti.MoveNext(); err == nil; err = ti.MoveNext() {
			d.bTokens = append(d.bTokens, ti.Item())
		}
	}
}

// Annotate records the given annotation against the applicable
// sequence of tokens in the appropriate section of the document.
// Currently known sections are "T" for title, "A" for abstract and
// "B" for body.  It creates and stores a `Word` corresponding to the
// text in the annotation.
func (d *Document) Annotate(a *Annotation) error {
	switch a.Section {
	case "T":
		if len(d.tTokens) == 0 {
			return fmt.Errorf("Title is not tokenized, but annotation provided")
		}
		w, err := d.wordForAnnotation(a, d.tTokens)
		if err != nil {
			return err
		}
		d.tWords = append(d.tWords, w)
		d.tAnnos = append(d.tAnnos, a)

	case "A":
		if len(d.aTokens) == 0 {
			return fmt.Errorf("Abstract is not tokenized, but annotation provided")
		}
		w, err := d.wordForAnnotation(a, d.aTokens)
		if err != nil {
			return err
		}
		d.aWords = append(d.aWords, w)
		d.aAnnos = append(d.aAnnos, a)

	case "B":
		if len(d.bTokens) == 0 {
			return fmt.Errorf("Body is not tokenized, but annotation provided")
		}
		w, err := d.wordForAnnotation(a, d.bTokens)
		if err != nil {
			return err
		}
		d.bWords = append(d.bWords, w)
		d.bAnnos = append(d.bAnnos, a)

	default:
		return fmt.Errorf("Annotation for unrecognised section : %s", a.Section)
	}

	return nil
}

func (d *Document) wordForAnnotation(a *Annotation, toks []*TextToken) (*Word, error) {
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

	w := NewWord(a.Entity, a.Begin, a.End)
	w.etype = a.Type
	return w, nil
}
