package parser

import (
	"github.com/optimuscrime/hltv-no-spoilers/pgk/ttokenizer"
	"golang.org/x/net/html"
	"strings"
)

func isVODContainerStart(t *ttokenizer.Ttokenizer) bool {
	if t.Token == nil {
		return false
	}

	attributes := t.Token.Attr

	if len(attributes) != 1 {
		return false
	}

	return strings.Contains(attributes[0].Val, "streams") && !strings.Contains(attributes[0].Val, "streams-")
}

func isVodContainerEnd(attr html.Attribute) bool {
	return attr.Key == "class" && strings.Contains(attr.Val, "no-spoiler")
}

func isStreamLinkDiv(attr html.Attribute) bool {
	return attr.Key == "class" && strings.Contains(attr.Val, "stream-box")
}

func isStreamDemoLinkButton(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "data-demo-link-button" {
			return true
		}
	}

	return false
}

func isStreamLinkUrl(attr html.Attribute) bool {
	return attr.Key == "data-stream-embed"
}

func isSpoilerTag(attr html.Attribute) bool {
	return attr.Key == "class" && strings.Contains(attr.Val, "spoiler")
}
