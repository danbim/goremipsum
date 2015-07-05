package goremipsum

import "fmt"
import "unicode"
import "strings"
import "math/rand"

type Set map[string]bool

func (set *Set) Add(elem string) {
	(map[string]bool(*set)[elem]) = true
}

func (set *Set) Remove(elem string) {
	delete(map[string]bool(*set), elem)
}

func (set *Set) Contains(elem string) bool {
	return (map[string]bool(*set))[elem]
}

var words Set = Set(make(map[string]bool))
var dict []string

func GenWords(wordCnt int) string {
	initWords()
	return strings.Join(genWordSlice(wordCnt), " ")
}

func genWordSlice(wordCnt int) []string {
	wordSlice := make([]string, wordCnt)
	for word := 0; word < wordCnt; word++ {
		fmt.Printf("%v", len(dict))
		r := rand.Int31n(int32(len(dict)))
		wordSlice[word] = dict[r]
	}
	return wordSlice
}

func GenSentence(wordCnt int) string {
	initWords()
	words := genWordSlice(wordCnt)
	firstWord := []rune(words[0])
	words[0] = string(unicode.ToUpper(firstWord[0])) + string(firstWord[1:])
	return strings.Join(words, " ") + "."
}

func GenSentences(sentenceCnt int) string {
	initWords()
	sentences := make([]string, sentenceCnt)
	for i := 0; i < sentenceCnt; i++ {
		wordCnt := int(rand.Int31n(10) + 3)
		sentences = append(sentences, GenSentence(wordCnt))
	}
	return strings.Join(sentences, " ")
}

func initWords() {
	for _, word := range strings.Split(text, " ") {
		if len(word) > 1 {
			word = strings.TrimRight(word, " .,;:")
			if len(word) > 1 {
				words.Add(word)
			}
		}
	}
	for word, _ := range words {
		dict = append(dict, word)
	}
}
