package html

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// NodeSearchAttr searches html node, find attr's value and return.
func NodeSearchAttr(node *html.Node, attrName string) string {
	if node == nil || attrName == "" {
		return ""
	}
	for _, attr := range node.Attr {
		if attr.Key == attrName {
			return attr.Val
		}
	}
	return ""
}

// NodeSearchAttrEq searches html node, find attr's value equal to attrValue in node.
// If found, return true, otherwise return false.
func NodeSearchAttrEq(node *html.Node, attrName string, attrValue string) bool {
	val := NodeSearchAttr(node, attrName)
	return val == attrValue
}

// HtmlNodeGrepAttr greps value of given attrName with reg.
// If attr value does not match reg, return an empty string.
// If reg is nil, use strings.Contains to match attr value.
func HtmlNodeGrepAttr(node *html.Node, attrName string, reg *regexp.Regexp) string {
	if node == nil {
		return ""
	}
	if reg == nil {
		for _, attr := range node.Attr {
			if attr.Key == attrName {
				return attr.Val
			}
		}
	} else {
		for _, attr := range node.Attr {
			if attr.Key == attrName && reg.MatchString(attr.Val) {
				return reg.FindString(attr.Val)
			}
		}
	}
	return ""
}

// NodeGrepHref greps a href link that matches regexp in node.
// Return link when regexp is nil.
func NodeGrepHref(node *html.Node, reg *regexp.Regexp) string {
	if node == nil {
		return ""
	}
	if reg == nil {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	} else {
		for _, attr := range node.Attr {
			if attr.Key == "href" && reg.MatchString(attr.Val) {
				return attr.Val
			}
		}
	}
	return ""
}

// NodeGrepAllHref greps all href inside node and its children.
func NodeGrepAllHref(node *html.Node) []string {
	if node == nil {
		return []string{}
	}
	var ret []string
	grepHref := func(node *html.Node) {
		link := NodeGrepHref(node, nil)
		if link != "" {
			ret = append(ret, link)
		}
	}
	traverseHtmlNode(node, grepHref)
	return ret
}

// NodeContainsText checks whether node contains text.
func NodeContainsText(node *html.Node, text string) bool {
	if node == nil {
		return false
	}
	containsText := false
	traverseHtmlNode(node, func(node *html.Node) {
		if node.Type == html.TextNode && strings.Contains(node.Data, text) {
			containsText = true
		}
	})
	return containsText
}

// traverseHtmlNode walks through node and its children.
// Runs function f on every html node.
func traverseHtmlNode(node *html.Node, f func(*html.Node)) {
	if node == nil {
		return
	}
	f(node)
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		traverseHtmlNode(n, f)
	}
}

// DocumentFromUrl builds a goquery.Document from urlRef.
func DocumentFromUrl(urlRef string) (*goquery.Document, error) {
	page, err := http.Get(urlRef)
	if err != nil {
		return nil, err
	}
	if page.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d:%s", page.StatusCode, page.Status)
	}
	defer page.Body.Close()
	if page.StatusCode != 200 {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(page.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}
