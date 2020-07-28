package idf

import (
	"github.com/go-ego/gse"
)

// StopWordMap default contains some stop words.
var StopWordMap = map[string]bool{
	"the":   true,
	"of":    true,
	"is":    true,
	"and":   true,
	"to":    true,
	"in":    true,
	"that":  true,
	"we":    true,
	"for":   true,
	"an":    true,
	"are":   true,
	"by":    true,
	"be":    true,
	"as":    true,
	"on":    true,
	"with":  true,
	"can":   true,
	"if":    true,
	"from":  true,
	"which": true,
	"you":   true,
	"it":    true,
	"this":  true,
	"then":  true,
	"at":    true,
	"have":  true,
	"all":   true,
	"not":   true,
	"one":   true,
	"has":   true,
	"or":    true,
}

// StopWord is a dictionary for all stop words.
type StopWord struct {
	stopWordMap map[string]bool

	seg gse.Segmenter
}

// AddToken adds a token into StopWord dictionary.
func (s *StopWord) AddToken(text string) {
	s.stopWordMap[text] = true
}

// NewStopWord create a new StopWord with default stop words.
func NewStopWord() *StopWord {
	s := new(StopWord)
	s.stopWordMap = StopWordMap
	return s
}

// IsStopWord checks if a given word is stop word.
func (s *StopWord) IsStopWord(word string) bool {
	_, ok := s.stopWordMap[word]
	return ok
}

func (s *StopWord) loadDict(files ...string) error {
	err := s.seg.LoadStop(files...)
	if err != nil {
		return err
	}

	for k, v := range gse.StopWordMap {
		StopWordMap[k] = v
	}

	return nil
}
