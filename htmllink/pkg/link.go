package pkg

import (
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func GetLinks(htmltxt []byte) ([]Link, error) {
	txt := string(htmltxt)
	reader := strings.NewReader(txt)
	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}
	var link []Link
	var f1 func(*html.Node) string
	f1 = func(n *html.Node) string {
		text := ""
		if n.Type == html.TextNode {
			text = strings.TrimSpace(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if len(f1(c)) != 0 {
				text = text + " " + f1(c)
			}
		}

		return strings.TrimSpace(text)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					key := attr.Val
					text := f1(n)
					l := Link{Href: key, Text: text}
					link = append(link, l)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return link, nil
}
