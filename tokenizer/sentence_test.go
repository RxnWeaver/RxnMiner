package tokenizer

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"
)

//

type Offsets struct {
	Begin int
	End int
}

type DocOffsets struct {
	Id string
	TOffs []Offsets
	AOffs []Offsets
}

//

func TestPatent7k(t *testing.T) {
	fn1 := "testdata/patent_7k_text.txt.gz"
	fn2 := "testdata/patent_7k_ref.txt"
	fn3 := "testdata/patent_7k_output.txt"

	runPatent7k(fn1, fn2, fn3, t)

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

func runPatent7k(fn1, fn2, fn3 string, t *testing.T) {
	f1, err := os.Open(fn1)
	if err != nil {
		t.Fatalf("!! Unable to read file : %s\n", fn1)
	}
	gr, err := gzip.NewReader(f1)
	if err != nil {
		t.Fatalf("!! Unable to read file : %s\n", fn1)
	}
	defer gr.Close()

	f2, err := os.Open(fn2)
	if err != nil {
		t.Fatalf("!! Unable to read file : %s\n", fn2)
	}
	defer f2.Close()

	f3, err := os.Create(fn3)
	if err != nil {
		t.Fatalf("!! Unable to read file : %s\n", fn3)
	}
	defer f3.Close()

	bf1 := bufio.NewReader(gr)
	bf2 := bufio.NewReader(f2)
	bf3 := bufio.NewWriter(f3)

	offs := offsets(bf2)
	tokenize(bf1, bf3, offs)
	_ = bf3.Flush()
}

//

func offsets(bf *bufio.Reader) map[string]*DocOffsets {
	doffs := make(map[string]*DocOffsets, 7000)

	for s, err := bf.ReadString('\n'); err == nil; s, err = bf.ReadString('\n') {
		fs := strings.Split(s, "\t")
		d := &DocOffsets{Id: fs[0]}
		fs1 := strings.Split(fs[1], ",")
		for _, s2 := range fs1 {
			s3 := strings.Split(s2, ":")
			i, _ := strconv.Atoi(s3[0])
			j, _ := strconv.Atoi(s3[1])
			d.TOffs = append(d.TOffs, Offsets{i, j})
		}
		fs2 := strings.Split(fs[2], ",")
		for _, s2 := range fs2 {
			s3 := strings.Split(s2, ":")
			i, _ := strconv.Atoi(s3[0])
			j, _ := strconv.Atoi(s3[1])
			d.AOffs = append(d.AOffs, Offsets{i, j})
		}

		doffs[d.Id] = d
	}

	return doffs
}

//

func tokenize(bf2 *bufio.Reader, bf3 *bufio.Writer, offs map[string]*DocOffsets) {
	for s, err := bf2.ReadString('\n'); err == nil; s, err = bf2.ReadString('\n') {
		fs := strings.Split(s, "\t")

		doc, _ := NewDocument(fs[0])
		doc.SetInput("T", fs[1])
		doc.SetInput("A", fs[2])
		doc.Tokenize()

		var sents []*Sentence
		var soffs []string
		si := NewSentenceIterator(doc.TokensInSection("T"))
		for err = si.MoveNext(); err == nil; err = si.MoveNext() {
			sents = append(sents, si.Item())
			soffs = append(soffs, fmt.Sprintf("%d:%d", si.Item().Begin(), si.Item().End()))
		}
		bf3.WriteString(fmt.Sprintf("%s\t%s\t", fs[0], strings.Join(soffs, ",")))

		sents = sents[:0]
		soffs = soffs[:0]
		si = NewSentenceIterator(doc.TokensInSection("A"))
		for err = si.MoveNext(); err == nil; err = si.MoveNext() {
			sents = append(sents, si.Item())
			soffs = append(soffs, fmt.Sprintf("%d:%d", si.Item().Begin(), si.Item().End()))
		}
		bf3.WriteString(fmt.Sprintf("%s\n", strings.Join(soffs, ",")))
	}
}
