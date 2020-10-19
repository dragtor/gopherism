package pkg

import (
	"fmt"
	//"io"
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	Href string
	Text string
}

func GetLinks(htmltxt []byte) ([]Link, error) {
	fmt.Println(string(htmltxt))
	txt := string(htmltxt)
	reader := strings.NewReader(txt)
	/*
	   z := html.NewTokenizer(reader)
	   for {
	       tt := z.Next()
	       if tt == html.ErrorToken {
	           fmt.Println("Error : html.ErrorToken")
	           break
	       }
	       fmt.Printf("%v\n",tt)
	   }
	*/
	doc, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("doc %#v", doc)
	var link []Link
	var f func(*html.Node)
	f1 := func(n *html.Node) {
		nList, err := n.ParseFragment(reader, n)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%#v", nList)
	}
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			//fmt.Printf("\n%#v\n", n)
			f1(n)
			//l := Link{Href: "hello", Text: "text"}
			//link = append(link, l)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return link, nil
}
