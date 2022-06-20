package model

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"

	"borscht.app/kapusta/microdata"
)

type Options struct {
	SkipUrl      bool
	SkipText     bool
	SkipDocument bool
	SkipSchema   bool
}

type InputData struct {
	Url      string
	Text     string
	Headers  *http.Header
	RootNode *html.Node
	Document *goquery.Document
	Schema   *microdata.Item
}
