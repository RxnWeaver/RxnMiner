package tokenizer

import (
	uni "unicode"
)

// TokenType represents types that a token can have.  The granularity
// of a token is variable: character, smallest logical unit, word,
// sentence, etc.  Accordingly, the corresponding tokens use
// appropriate token types.
type TokenType byte

// List of defined token types.
const (
	TokOther TokenType = iota

	// Rune types
	TokSpace
	TokLetter
	TokNumber
	TokTerm
	TokPause
	TokGroupOpen
	TokGroupClose
	TokPunct
	TokSymbol

	// Others
	TokMayBeWord
	TokWord
	TokSentence
)

var TtDescriptions map[TokenType]string = map[TokenType]string{
	TokOther:      "TokOther",
	TokSpace:      "TokSpace",
	TokLetter:     "TokLetter",
	TokNumber:     "TokLetter",
	TokTerm:       "TokTerm",
	TokPause:      "TokPause",
	TokGroupOpen:  "TokGroupOpen",
	TokGroupClose: "TokGroupClose",
	TokPunct:      "TokPunct",
	TokSymbol:     "TokSymbol",
	TokMayBeWord:  "TokMayBeWord",
	TokWord:       "TokWord",
	TokSentence:   "TokSentence",
}

//

func TypeOfRune(r rune) TokenType {
	switch {
	case uni.IsSpace(r):
		return TokSpace

	case uni.IsLetter(r):
		return TokLetter

	case uni.IsNumber(r):
		return TokNumber

	case r == '.', r == '!', r == '?':
		return TokTerm

	case r == ',', r == ':', r == ';':
		return TokPause

	case r == '(', r == '[', r == '{':
		return TokGroupOpen

	case r == ')', r == ']', r == '}':
		return TokGroupClose

	case uni.IsPunct(r):
		return TokPunct

	case uni.IsSymbol(r):
		return TokSymbol
	}

	return TokOther
}

// EntityType represents types that a logical word can have.  These
// types are crucial to subsequent mining.
type EntityType byte

// List of defined entity types.
const (
	EntOther EntityType = iota
	EntAbbreviation
	EntFamily
	EntFormula
	EntIdentifier
	EntMultiple
	EntSystematic
	EntTrivial
)

var EtDescriptions map[EntityType]string = map[EntityType]string{
	EntOther:        "OTHER",
	EntAbbreviation: "ABBREVIATION",
	EntFamily:       "FAMILY",
	EntFormula:      "FORMULA",
	EntIdentifier:   "IDENTIFIER",
	EntMultiple:     "MULTIPLE",
	EntSystematic:   "SYSTEMATIC",
	EntTrivial:      "TRIVIAL",
}

var EtDescriptionsInv map[string]EntityType = map[string]EntityType{
	"OTHER":        EntOther,
	"ABBREVIATION": EntAbbreviation,
	"FAMILY":       EntFamily,
	"FORMULA":      EntFormula,
	"IDENTIFIER":   EntIdentifier,
	"MULTIPLE":     EntMultiple,
	"SYSTEMATIC":   EntSystematic,
	"TRIVIAL":      EntTrivial,
}
