package tokenizer

import (
	"testing"
)

//

func TestSentences001(t *testing.T) {
	t.Logf("%-8s : %s", "TEST", "CN1854125A")
	d, _ := NewDocument("CN1854125A")
	d.SetInput("T", "Aryl-group-substituted acrylonitrile compound, its production and use")
	d.SetInput("A", "An aryl-substituted acrylonitrile compound, its production and use are disclosed. It can be used for antineoplastic and to treat leukemia, hepatocarcinoma, gastric carcinoma and mastopathy.")

	d.Tokenize()

	si := NewSentenceIterator(d.TokensInSection("T"))
	var sents []*Sentence
	for err := si.MoveNext(); err == nil; err = si.MoveNext() {
		sents = append(sents, si.Item())
	}
	if len(sents) != 1 {
		t.Errorf("Incorrect number of sentences.  Expected : %d, observed : %d", 1, len(sents))
	}

	si = NewSentenceIterator(d.TokensInSection("A"))
	sents = sents[:0]
	for err := si.MoveNext(); err == nil; err = si.MoveNext() {
		sents = append(sents, si.Item())
	}
	if len(sents) != 2 {
		t.Errorf("Incorrect number of sentences.  Expected : %d, observed : %d", 2, len(sents))
	}
	t.Logf("%-8s : %s", "SUCCESS", "CN1854125A")
}

//

func TestSentences002(t *testing.T) {
	t.Logf("%-8s : %s", "TEST", "CN1857207A")
	d, _ := NewDocument("CN1857207A")
	d.SetInput("T", "Slow released compound anticancer injection containing blood vessel inhibitor")
	d.SetInput("A", "The slow released compound anticancer injection containing blood vessel inhibitor and its synergist consists of slow released microsphere and solvent. The slow released microsphere includes effective anticancer component and slow releasing supplementary material, and the solvent is common solvent or special solvent containing suspending agent. The effective anticancer component is blood vessel inhibitor and/or blood vessel inhibitor synergist selected from antitumor antibiotic and antimetabolite. The slow releasing supplementary material is selected from difatty acid-sebacic acid copolymer, poly(erucic acid dipolymer-sebacic acid), poly(fumaric acid-sebacic acid), etc or their composition. The suspending agent is carboxymethyl cellulose, etc. and has viscosity of 80-3000 cp at 20-30 deg.c. The slow released microsphere may be also prepared into slow released implanting agent for use alone or together with chemotherapeutic and/or radiotherapeutic medicine..")

	d.Tokenize()

	si := NewSentenceIterator(d.TokensInSection("T"))
	var sents []*Sentence
	for err := si.MoveNext(); err == nil; err = si.MoveNext() {
		sents = append(sents, si.Item())
	}
	if len(sents) != 1 {
		t.Errorf("Incorrect number of sentences.  Expected : %d, observed : %d", 1, len(sents))
	}

	si = NewSentenceIterator(d.TokensInSection("A"))
	sents = sents[:0]
	for err := si.MoveNext(); err == nil; err = si.MoveNext() {
		sents = append(sents, si.Item())
	}
	if len(sents) != 7 {
		t.Errorf("Incorrect number of sentences.  Expected : %d, observed : %d", 2, len(sents))
	}
	t.Logf("%-8s : %s", "SUCCESS", "CN1857207A")
}
