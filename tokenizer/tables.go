// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

package tokenizer

// NonTermAbbrevs lists the common abbreviations that could end with a
// full stop, but without ending the sentence.  The abbrevs are in
// lowercase.
var NonTermAbbrevs = map[string]struct{}{
	"viz": {},
	"eg":  {},
	"ex":  {},
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
var MayBeTermAbbrevs = map[string]struct{}{
	"etc": {},
}

// MayBeTermGroupAbbrevs lists the common abbreviations that are
// compound, i.e. they involve more than one token.  The table omits
// any intervening period.  The abbrevs are in lowercase.
var MayBeTermGroupAbbrevs = map[string][]string{
	"e": {"i"},
	"g": {"e"},
}
