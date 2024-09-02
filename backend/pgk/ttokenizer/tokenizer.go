package ttokenizer

import (
	"golang.org/x/net/html"
	"strings"
)

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

// Next Stupid wrapper stuff because we can't call .Token() multiple times without calling .Next() in between, so we
// do this to store the token for multiple lookups
func (t *Ttokenizer) Next() {
	tokenType := t.Tokenizer.Next()
	token := t.Tokenizer.Token()

	t.TokenType = &tokenType
	t.Token = &token
}

func (t *Ttokenizer) GetNextTokenString() (string, error) {
	t.Next()
	newTokenType := *t.TokenType
	newToken := *t.Token
	if newTokenType == html.ErrorToken {
		return "", ErrFailedToParse
	}

	str := strings.Trim(html.EscapeString(newToken.String()), "")
	return str, nil
}
