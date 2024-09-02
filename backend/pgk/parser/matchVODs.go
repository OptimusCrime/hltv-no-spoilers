package parser

import (
	"github.com/optimuscrime/hltv-no-spoilers/pgk/ttokenizer"
	"golang.org/x/net/html"
)

type VOD struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

func ParseMatchVODs(body string) ([]VOD, error) {
	tokenizer := ttokenizer.CreateTokenizerFromString(body)

	for {
		if !isVODContainerStart(tokenizer) {
			tokenizer.Next()
			continue
		}

		return parseVodCollection(tokenizer)
	}
}

func parseVodCollection(t *ttokenizer.Ttokenizer) ([]VOD, error) {
	var vods []VOD

	for {
		t.Next()
		tokenType := *t.TokenType
		token := *t.Token

		switch {
		case tokenType == html.ErrorToken:
			return vods, nil
		default:
			attributes := token.Attr
			if len(attributes) == 0 {
				continue
			}

			for _, attribute := range attributes {
				if isVodContainerEnd(attribute) {
					return vods, nil
				}

				if isStreamLinkDiv(attribute) {
					vod, err := parseVOD(t)
					if err != nil {
						continue
					}

					vods = append(vods, *vod)
				}
			}
		}
	}
}

func parseVOD(t *ttokenizer.Ttokenizer) (*VOD, error) {
	streamUrl, err := parseVodStreamUrl(t)
	if err != nil {
		return nil, err
	}

	return &VOD{
		Title: "Foobar",
		Url:   streamUrl,
	}, nil
}

func parseVodStreamUrl(t *ttokenizer.Ttokenizer) (string, error) {
	attributes := t.Token.Attr
	for _, attribute := range attributes {
		if isStreamLinkUrl(attribute) {
			return attribute.Val, nil
		}
	}

	return "", nil
}
