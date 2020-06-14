package jack_compiler

import (
	"regexp"
	"strings"
)

func ContainsInt(slice []int, n int) bool {
	for _, v := range slice {
		if v == n {
			return true
		}
	}
	return false
}

func ContainsString(slice []string, n string) bool {
	for _, v := range slice {
		if v == n {
			return true
		}
	}
	return false
}

func EscapeXml(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), ">", "&gt;"), "<", "&lt;")
}

var stringReg = regexp.MustCompile(`^".*"$`)
var identifierReg = regexp.MustCompile(`^[a-zA-Z]\w*$`)

func isWord(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || r == '_'
}

func isNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

var commentReg = regexp.MustCompile(`(/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/)|(//.*)|(^\*[^\/]).*|(^(\/\*\*).*)|^(\*\/)`)

func removeComments(line string) string {
	return strings.TrimSpace(commentReg.ReplaceAllString(strings.TrimSpace(line), ""))
}

type TokenIterator struct {
	tokens []Token
	i      int
}

func NewTokenIterator(ts []Token) *TokenIterator {
	return &TokenIterator{
		tokens: ts,
	}
}

func (it *TokenIterator) Size() int {
	return len(it.tokens)
}

func (it *TokenIterator) Next() Token {
	defer func() {
		it.i++
	}()
	return it.Peek()
}

func (it *TokenIterator) HasNext() bool {
	return len(it.tokens) > it.i
}

func (it *TokenIterator) Peek() Token {
	if !it.HasNext() {
		return nil
	}
	return it.tokens[it.i]
}
