package ttokenizer

import (
	"errors"
	"golang.org/x/net/html"
	"strings"
)

// This package thing is only used because we are not allowed to call .Token() multiple times without advancing the
// cursor with .Next() between each .Token() call. Temporarily storing the token and type in this struct to work around
// this without having to pass other structs back and forth all the time. Annoying.

type Ttokenizer struct {
	Tokenizer *html.Tokenizer
	Token     *html.Token
	TokenType *html.TokenType
}

func CreateTokenizerFromString(body string) *Ttokenizer {
	return &Ttokenizer{
		Tokenizer: html.NewTokenizer(strings.NewReader(body)),
		Token:     nil,
		TokenType: nil,
	}
}

func (t *Ttokenizer) Next() {
	tokenType := t.Tokenizer.Next()
	token := t.Tokenizer.Token()

	t.TokenType = &tokenType
	t.Token = &token
}

func (t *Ttokenizer) GetTokenString() string {
	return strings.Trim(html.EscapeString(t.Token.String()), "")
}

func (t *Ttokenizer) GetNextTokenString() (string, error) {
	t.Next()

	newTokenType := *t.TokenType

	if newTokenType == html.ErrorToken {
		return "", errors.New("failed to parse document")
	}

	return t.GetTokenString(), nil
}
