package tokenizer

import (
	"io/ioutil"
	"testing"
)

func TestEnglish001(t *testing.T) {
	files := []string{
		"testdata/input-org-syn.txt",
		"testdata/input-te-wiki.txt",
	}

	for _, fn := range files {
		t.Logf("%-8s : %s", "TEST", fn)
		bs, err := ioutil.ReadFile(fn)
		if err != nil {
			t.Fatalf("%-8s : Input data file '%s' could not be read : %s", "FATAL", fn, err.Error())
		}

		size := len(bs)
		ti := NewTextTokenIterator(string(bs))
		var toks []*TextToken
		for err = ti.MoveNext(); err == nil; err = ti.MoveNext() {
			toks = append(toks, ti.Item())
		}

		lt := toks[len(toks)-1]
		if lt.End() != size-1 {
			t.Errorf("%-8s : Token offset drift by EOF.  Expected : %d, observed : %d", "FAIL", size, lt.End())
		}
		t.Logf("%-8s : %s", "SUCCESS", fn)
	}
}
