// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

import (
	// "strings"
	"testing"
)

func TestSingleWord001(t *testing.T) {
	doc, _ := NewDocument("SingleWord001")
	doc.SetInput("Para1", "Raamu is a good boy.")
	doc.Tokenize()

	a, _ := NewAnnotation("SingleWord001\tPara1\t0\t4\tRaamu\tNOUN")
	err := doc.Annotate(a, "POS")
	if err != nil {
		t.Fatalf("Failed to annotate : %v", a)
	}
	c, _ := doc.SectionWordCount("Para1")
	if c != 1 {
		t.Fatalf("Expect word count : 1, observed : %d", c)
	}

	a, _ = NewAnnotation("SingleWord001\tPara1\t16\t18\tboy\tTRIVIAL")
	err = doc.Annotate(a, "CLS")
	if err != nil {
		t.Fatalf("Failed to annotate : %v", a)
	}
	c, _ = doc.SectionWordCount("Para1")
	if c != 2 {
		t.Fatalf("Expect word count : 1, observed : %d", c)
	}

	a, _ = NewAnnotation("SingleWord001\tPara1\t0\t4\tRaamu\tIDENTIFIER")
	err = doc.Annotate(a, "CLS")
	if err != nil {
		t.Fatalf("Failed to annotate : %v", a)
	}
	c, _ = doc.SectionWordCount("Para1")
	if c != 2 {
		t.Fatalf("Expect word count : 1, observed : %d", c)
	}
}

//

func TestDoubleWord001(t *testing.T) {
	doc, _ := NewDocument("DoubleWord001")
	doc.SetInput("Para1", "Raamu is a good boy.")
	doc.Tokenize()

	a, _ := NewAnnotation("DoubleWord001\tPara1\t11\t18\tgood boy\tQUALIFIER")
	err := doc.Annotate(a, "CLS")
	if err != nil {
		t.Fatalf("Failed to annotate : %v", a)
	}
	c, _ := doc.SectionWordCount("Para1")
	if c != 1 {
		t.Fatalf("Expect word count : 1, observed : %d", c)
	}
}
