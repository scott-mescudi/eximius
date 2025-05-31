package internal

import (
	"bytes"
	"strings"

	"github.com/scott-mescudi/eximius/parser/pkg/models"
	"github.com/scott-mescudi/eximius/parser/pkg/regex"
	"golang.org/x/net/html"
)

func extractText(n *html.Node) string {
	var sb strings.Builder

	var crawler func(*html.Node)
	crawler = func(n *html.Node) {
		if n.Type == html.TextNode {
			trimmed := strings.TrimSpace(n.Data)
			if trimmed != " " {
				sb.WriteString(trimmed)
				sb.WriteString(" ")
			}
		}
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
			return // Skip script and style content
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}

	crawler(n)
	return sb.String()
}

var replacer = strings.NewReplacer(
	" a ", " ",
	" an ", " ",
	" and ", " ",
	" are ", " ",
	" as ", " ",
	" at ", " ",
	" be ", " ",
	" by ", " ",
	" for ", " ",
	" from ", " ",
	" has ", " ",
	" he ", " ",
	" in ", " ",
	" is ", " ",
	" it ", " ",
	"its", " ",
	" of ", " ",
	" on ", " ",
	" that ", " ",
	" the ", " ",
	" to ", " ",
	" was ", " ",
	" were ", " ",
	" will ", " ",
	" with ", " ",
)

func parseDocument(document string) (parsedDocument models.ParsedDocument) {
	title := regex.TitleRegex.FindStringSubmatch(document)
	if len(title) < 1 {
		parsedDocument.Title = ""
	} else {
		parsedDocument.Title = title[1]
	}

	description := regex.DescriptionRegex.FindStringSubmatch(document)
	if len(description) < 1 {
		parsedDocument.Description = ""
	} else {
		parsedDocument.Description = description[1]
	}

	node, err := html.Parse(bytes.NewReader([]byte(document)))
	text := extractText(node)
	if err != nil {
		parsedDocument.Text = ""
	} else {
		parsedDocument.Text = replacer.Replace(text)
	}

	return parsedDocument
}
