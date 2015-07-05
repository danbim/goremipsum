package goremipsum

import "strings"
import "math/rand"
import "testing"
import . "gopkg.in/check.v1"

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func (s *TestSuite) TestGenSentence(c *C) {
	cntExpected := int(rand.Int31n(500))
	text := GenSentence(cntExpected)
	cntActual := cntWords(text)
	c.Assert(cntActual, Equals, cntExpected)
}

func (s *TestSuite) TestGenSentences(c *C) {
	cntExpected := int(rand.Int31n(100))
	text := GenSentences(cntExpected)
	cntActual := cntSentences(text)
	c.Assert(cntActual, Equals, cntExpected)
}

func cntWords(text string) int {
	cnt := 0
	for _, word := range strings.Split(text, " ") {
		if len(word) > 1 {
			cnt += 1
		}
	}
	return cnt
}

func cntSentences(text string) int {
	cnt := 0
	for _, elem := range []byte(text) {
		if elem == '.' {
			cnt += 1
		}
	}
	return cnt
}
