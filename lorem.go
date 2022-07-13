// Copyright 2012 Derek A. Rhodes.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package englishgen

import (
	"math/rand"
	"strconv"
	"strings"
)

// Generate a natural word len.
func genWordLen() int {
	return rand.Intn(wordLenMax) + 1
}

func intRange(min, max int) int {
	if min == max {
		return intRange(min, min+1)
	}
	if min > max {
		return intRange(max, min)
	}
	n := rand.Int() % (max - min)
	return n + min
}

func word(wordLen int) string {
	if wordLen < 1 {
		wordLen = 1
	}
	if wordLen > wordLenMax {
		wordLen = wordLenMax
	}

	words, _ := wordMapByLen.Load(strconv.Itoa(wordLen))

	return words[rand.Intn(len(words))]
}

// Generate a word in a specfied range of letters.
func Word(min, max int) string {
	n := intRange(min, max)
	return word(n)
}

// Generate a sentence with a specified range of words.
func Sentence(min, max int) string {
	n := intRange(min, max)

	// grab some words
	ws := []string{}
	maxcommas := 2
	numcomma := 0
	for i := 0; i < n; i++ {
		ws = append(ws, word(genWordLen()))

		// maybe insert a comma, if there are currently < 2 commas, and
		// the current word is not the last or first
		if (rand.Int()%n == 0) && numcomma < maxcommas && i < n-1 && i > 2 {
			ws[i-1] += ","
			numcomma += 1
		}

	}

	ws[0] = strings.Title(ws[0])
	sentence := strings.Join(ws, " ") + "."
	return sentence
}

// Generate a paragraph with a specified range of sentenences.
const (
	minwords = 5
	maxwords = 30
)

func Paragraph(min, max int) string {
	n := intRange(min, max)

	p := []string{}
	for i := 0; i < n; i++ {
		p = append(p, Sentence(minwords, maxwords))
	}
	return strings.Join(p, " ")
}

// Generate a random URL
func Url() string {
	n := intRange(0, 3)

	base := `http://www.` + Host()

	switch n {
	case 0:
		break
	case 1:
		base += "/" + Word(2, 8)
	case 2:
		base += "/" + Word(2, 8) + "/" + Word(2, 8) + ".html"
	}
	return base
}

// Host
func Host() string {
	n := intRange(0, 3)
	tld := ""
	switch n {
	case 0:
		tld = ".com"
	case 1:
		tld = ".net"
	case 2:
		tld = ".org"
	}

	parts := []string{Word(2, 8), Word(2, 8), tld}
	return strings.Join(parts, ``)
}

// Email
func Email() string {
	return Word(4, 10) + `@` + Host()
}
