// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

// NonTermAbbrevs lists the common abbreviations that could end with a
// full stop, but without ending the sentence.  The abbrevs are in
// lowercase.
var NonTermAbbrevs map[string]struct{} = map[string]struct{}{
	"viz": {},
	"eg":  {},
	"fig": {},
	// Salutations
	"mr":   {},
	"ms":   {},
	"mrs":  {},
	"dr":   {},
	"prof": {},
}

// MayBeTermAbbrevs lists the common abbreviations that could end with
// a full stop, possibly without ending the sentence.  The abbrevs are
// in lowercase.
var MayBeTermAbbrevs map[string]struct{} = map[string]struct{}{
	"etc": {},
	"al":  {},
	"e":   {},
	"g":   {},
}
