package tokenizer

// NonTermAbbrevs lists the common abbreviations that could end with a
// full stop, possibly without ending the sentence.  The abbrevs are
// in lowercase.
var NonTermAbbrevs map[string]struct{} = map[string]struct{}{
	"etc": {},
	"al":  {},
	"viz": {},
	"eg":  {},
	"e":   {},
	"g":   {},
	"fig": {},
	// Salutations
	"mr":   {},
	"ms":   {},
	"mrs":  {},
	"dr":   {},
	"prof": {},
}
