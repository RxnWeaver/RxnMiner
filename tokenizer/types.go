// Copyright (c) 2015 RxnWeaver
//
// Part of the RxnWeaver suite of projects.  See README.md and LICENSE
// for more details.

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
	TokMayBeTerm
	TokTerm
	TokPause
	TokParenOpen
	TokParenClose
	TokBracketOpen
	TokBracketClose
	TokBraceOpen
	TokBraceClose
	TokSquote
	TokDquote
	TokIniQuote
	TokFinQuote
	TokPunct
	TokSymbol

	// Others
	TokMayBeWord
	TokWord
	TokSentence
)

var TtDescriptions map[TokenType]string = map[TokenType]string{
	TokOther:        "TokOther",
	TokSpace:        "TokSpace",
	TokLetter:       "TokLetter",
	TokNumber:       "TokLetter",
	TokMayBeTerm:    "TokMayBeTerm",
	TokTerm:         "TokTerm",
	TokPause:        "TokPause",
	TokParenOpen:    "TokParenOpen",
	TokParenClose:   "TokParenClose",
	TokBracketOpen:  "TokBracketOpen",
	TokBracketClose: "TokBracketClose",
	TokBraceOpen:    "TokBraceOpen",
	TokBraceClose:   "TokBraceClose",
	TokSquote:       "TokSquote",
	TokDquote:       "TokDquote",
	TokIniQuote:     "TokIniQuote",
	TokFinQuote:     "TokFinQuote",
	TokPunct:        "TokPunct",
	TokSymbol:       "TokSymbol",
	TokMayBeWord:    "TokMayBeWord",
	TokWord:         "TokWord",
	TokSentence:     "TokSentence",
}

//

func RuneType(r rune) TokenType {
	switch {
	case uni.IsSpace(r):
		return TokSpace

	case uni.IsLetter(r):
		return TokLetter

	case uni.IsNumber(r):
		return TokNumber

	case r == '!', r == '?',
		r == '\u3002', // Ideographic full stop
		r == '\uff01', // Full width exclamation mark
		r == '\uff1f': // Full width question mark
		return TokTerm

	case r == '.',
		r == '\uff0e': // Full width full stop
		return TokMayBeTerm

	case r == ',', r == ':', r == ';':
		return TokPause

	case r == '(':
		return TokParenOpen

	case r == ')':
		return TokParenClose

	case r == '[':
		return TokBracketOpen

	case r == ']':
		return TokBracketClose

	case r == '{':
		return TokBraceOpen

	case r == '}':
		return TokBraceClose

	case r == '\'':
		return TokSquote

	case r == '"':
		return TokDquote

	case uni.Is(uni.Pi, r):
		return TokIniQuote

	case uni.Is(uni.Pf, r):
		return TokFinQuote

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
