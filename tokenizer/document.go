package tokenizer

import (
	"fmt"
)

// Document represents the entirety of input text of one logical
// document -- usually a file.
//
// It holds information about its sections, tokens in them, and the
// words that were recognised by other processors.  In case the
// document has associated training annotations, it holds them as
// well.
type Document struct {
	id     string // Must be unique within a run
	input  map[string]string
	tokens map[string][]*TextToken
	words  map[string][]*Word
	annos  map[string][]*Annotation
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
	d.input = make(map[string]string, 2)
	d.tokens = make(map[string][]*TextToken, 2)
	d.words = make(map[string][]*Word, 2)
	d.annos = make(map[string][]*Annotation, 2)

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

// Tokenize breaks the text in the title, abstract and body into
// quasi-atomic tokens.
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

// Annotate records the given annotation against the applicable
// sequence of tokens in the appropriate section of the document.
// Currently known sections are "T" for title, "A" for abstract and
// "B" for body.  It creates and stores a `Word` corresponding to the
// text in the annotation.
func (d *Document) Annotate(a *Annotation) error {
	toks, ok := d.tokens[a.Section]
	if !ok {
		return fmt.Errorf("Annotation for unrecognised section : %s", a.Section)
	}

	w, err := d.wordForAnnotation(a, toks)
	if err != nil {
		return err
	}

	words := d.words[a.Section]
	words = append(words, w)
	d.words[a.Section] = words

	annos := d.annos[a.Section]
	annos = append(annos, a)
	d.annos[a.Section] = annos

	return nil
}

//

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

// TokensInSection answers any recognised tokens in the given section.
func (d *Document) TokensInSection(sec string) []*TextToken {
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

// WordsInSection answers any recognised words in the given section.
func (d *Document) WordsInSection(sec string) []*Word {
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

// AnnotationsForSection answers the registered annotations for the
// given section.
func (d *Document) AnnotationsForSection(sec string) []*Annotation {
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
