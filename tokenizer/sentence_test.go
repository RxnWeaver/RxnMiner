// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

//

type Offsets struct {
	Begin int
	End   int
}

type DocOffsets struct {
	Id    string
	TOffs []Offsets
	AOffs []Offsets
}

//

func TestPatent7k(t *testing.T) {
	fn1 := "testdata/patent_7k_text.txt.gz"
	fn2 := "testdata/patent_7k_ref.txt"
	fn3 := "testdata/patent_7k_output.txt"

	runPatent7k(fn1, fn3, t)

	//

	output, err := exec.Command("diff", "-Nau", fn2, fn3).Output()
	if err != nil {
		t.Fatalf("!! Unable to diff the reference and output files : %s\n", err.Error())
	}
	soutput := string(output)
	if soutput != "" {
		t.Fatalf("!! Differences between reference and output observed.\n%s\n", soutput)
	}

	err = exec.Command("rm", "-f", fn3).Run()
}

//

func runPatent7k(fn1, fn3 string, t *testing.T) {
	f1, err := os.Open(fn1)
	if err != nil {
		t.Fatalf("!! Unable to read file : %s\n", fn1)
	}
	gr, err := gzip.NewReader(f1)
	if err != nil {
		t.Fatalf("!! Unable to read file : %s\n", fn1)
	}
	defer gr.Close()

	f3, err := os.Create(fn3)
	if err != nil {
		t.Fatalf("!! Unable to read file : %s\n", fn3)
	}
	defer f3.Close()

	bf1 := bufio.NewReader(gr)
	bf3 := bufio.NewWriter(f3)

	tokenize(bf1, bf3)
	_ = bf3.Flush()
}

//

func tokenize(bf2 *bufio.Reader, bf3 *bufio.Writer) {
	for s, err := bf2.ReadString('\n'); err == nil; s, err = bf2.ReadString('\n') {
		fs := strings.Split(s, "\t")

		doc, _ := NewDocument(fs[0])
		doc.SetInput("T", fs[1])
		doc.SetInput("A", fs[2])
		doc.Tokenize()
		doc.AssembleSentences()

		var soffs []string
		sents := doc.SectionSentences("T")
		for _, sent := range sents {
			soffs = append(soffs, fmt.Sprintf("%d:%d", sent.Begin(), sent.End()))
		}
		bf3.WriteString(fmt.Sprintf("%s\t%s\t", fs[0], strings.Join(soffs, ",")))

		soffs = soffs[:0]
		sents = doc.SectionSentences("A")
		for _, sent := range sents {
			soffs = append(soffs, fmt.Sprintf("%d:%d", sent.Begin(), sent.End()))
		}
		bf3.WriteString(fmt.Sprintf("%s\n", strings.Join(soffs, ",")))
	}
}
