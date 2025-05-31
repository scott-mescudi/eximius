package parser

import (
	"bytes"
	"strings"

	"github.com/scott-mescudi/eximius/spider/pkg/models"
	"golang.org/x/net/html"
)

func extractText(n *html.Node) string {
	var sb strings.Builder

	var crawler func(*html.Node)
	crawler = func(n *html.Node) {
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
			return
		}

		if n.Type == html.TextNode {
			trimmed := strings.TrimSpace(n.Data)
			if trimmed != " " {
				sb.WriteString(trimmed)
				sb.WriteString(" ")
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}

	crawler(n)
	return sb.String()
}

func extractTitleAndDescription(n *html.Node) (title, description string) {
	var walker func(*html.Node)
	walker = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "title" && n.FirstChild != nil {
				title = n.FirstChild.Data
			}

			if n.Data == "meta" {
				var name, content string
				for _, attr := range n.Attr {
					switch strings.ToLower(attr.Key) {
					case "name":
						name = strings.ToLower(attr.Val)
					case "content":
						content = attr.Val
					}
				}

				if name == "description" {
					description = content
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walker(c)
		}
	}
	walker(n)
	return
}

func ParseDocument(document *models.Document) {
	dt, _ := html.Parse(bytes.NewReader([]byte(document.Text)))
	title, description := extractTitleAndDescription(dt)
	document.Title = title
	
	text := extractText(dt)
	if description == "" {
		document.Description = text[256:2560]
	} else {
		document.Description = description
	}

	document.Text = replacer.Replace(strings.ToLower(text))
	document.Text = replacer.Replace(document.Text)
}
