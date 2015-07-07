// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

import (
	"io/ioutil"
	"testing"
)

func TestEnglish001(t *testing.T) {
	files := []string{
		"testdata/input-article.txt",
		"testdata/input-te-wiki.txt",
	}

	for _, fn := range files {
		bs, err := ioutil.ReadFile(fn)
		if err != nil {
			t.Fatalf("Input data file '%s' could not be read : %s", fn, err.Error())
		}

		size := len(bs)
		ti := NewTextTokenIterator(string(bs))
		var toks []*TextToken
		for err = ti.MoveNext(); err == nil; err = ti.MoveNext() {
			toks = append(toks, ti.Item())
		}

		lt := toks[len(toks)-1]
		if lt.End() != size-1 {
			t.Errorf("Token offset drift by EOF.  Expected : %d, observed : %d", size, lt.End())
		}
	}
}
