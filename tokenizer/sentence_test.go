package tokenizer

import (
	// "fmt"
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
		// fmt.Printf("#  %s\n", si.Item().token.Text())
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
		// fmt.Printf("#  %s\n", si.Item().token.Text())
	}
	if len(sents) != 6 {
		t.Errorf("Incorrect number of sentences.  Expected : %d, observed : %d", 6, len(sents))
	}
	t.Logf("%-8s : %s", "SUCCESS", "CN1857207A")
}

//

func TestSentences003(t *testing.T) {
	t.Logf("%-8s : %s", "TEST", "CA2156289C")
	d, _ := NewDocument("CA2156289C")
	d.SetInput("T", "Drug composition containing nucleic acid copolymer")
	d.SetInput("A", "This invention has for its object to insure an effective utilization of single-stranded nucleic acid copolymers, particularly poly(adenylic acid-uridylic acid), and to provide a pharmaceutical composition having antitumor activity.      The invention typically relates to a pharmaceutical composition comprising a lipid device such as Lipofectin (trademark), 3-O-(4-dimethylaminobutanoyl)-1,2-O-dioleylgycerol, 3-O-(2-dimethylamino-ethyl)carbamoyl-1,2-O-dioleylglycerol, 3-O-(2-diethylaminoethyl) carbamoyl-1,2-O-dioleylgycerol, or 2-O-(2-diethylaminoethyl)carbamoyl-1,3-O-dioleoylglycerol and poly(adenylic acid-uridylic acid).")

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
		// fmt.Printf("#  %s\n", si.Item().token.Text())
	}
	if len(sents) != 2 {
		t.Errorf("Incorrect number of sentences.  Expected : %d, observed : %d", 2, len(sents))
	}
	t.Logf("%-8s : %s", "SUCCESS", "CA2156289C")
}

//

func TestSentences004(t *testing.T) {
	t.Logf("%-8s : %s", "TEST", "CA2157792C")
	d, _ := NewDocument("CA2157792C")
	d.SetInput("T", "N-(3-piperidinyl-carbonyl)-beta-alanine derivatives as paf antagonists")
	d.SetInput("A", "This invention relates to .beta.-alanine derivatives represented by formula (I), wherein each symbol is as defined in the specification and pharmaceutically acceptable salt thereof which is glycoprotein IIb/IIIa antagonist, inhibitor of blood platelets aggregation and inhibitor of the binding of fibrinogen to blood platelets, to processes for the preparation thereof, to a pharmaceutical composition comprising the same and to a method for the prevention and/or treatment of diseases indicated in the specification to a human being or an animal.")

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
		// fmt.Printf("#  %s\n", si.Item().token.Text())
	}
	if len(sents) != 1 {
		t.Errorf("Incorrect number of sentences.  Expected : %d, observed : %d", 1, len(sents))
	}
	t.Logf("%-8s : %s", "SUCCESS", "CA2157792C")
}

//

func TestSentences005(t *testing.T) {
	t.Logf("%-8s : %s", "TEST", "CA2252050C")
	d, _ := NewDocument("CA2252050C")
	d.SetInput("T", "Buccal, non-polar spray or capsule")
	d.SetInput("A", "A buccal aerosol spray or capsule using a non-polar solvent has now been developed which provides biologically active compounds for rapid absorption through the oral mucosa, resulting in fast onset of effect. The buccal aerosol spray of the invention comprises formulation (I): propellant 50-95 %, non-polar solvent 5-50 %, active compound 0.0025- 40 %, flavoring agent 0.05-5 %. The soft bite gelatin capsule of the invention comprises formulation (II): non-polar solvent 30-99.8 %, emulsifier 0-20 %, active compound 0.0003-32 %, and flavoring agent 0.05-60.")

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
		// fmt.Printf("#  %s\n", si.Item().token.Text())
	}
	if len(sents) != 3 {
		t.Errorf("Incorrect number of sentences.  Expected : %d, observed : %d", 3, len(sents))
	}
	t.Logf("%-8s : %s", "SUCCESS", "CA2252050C")
}
