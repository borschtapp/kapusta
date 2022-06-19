package parser

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/jaytaylor/html2text"
	"golang.org/x/net/html"
	"golang.org/x/net/html/charset"

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

var html2textOptions = html2text.Options{PrettyTables: false, OmitLinks: true}

func FileInput(fileName string, options Options) (*InputData, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("unable to read the file: " + err.Error())
	}
	defer file.Close()

	contentType, err := getContentType(file)
	if err == nil && strings.HasPrefix(contentType, "text/html") {
		root, err := html.Parse(file)
		if err != nil {
			return nil, errors.New("unable to parse html tree: " + err.Error())
		}

		url := "file://" + strings.ReplaceAll(fileName, "\\", "/")
		return NodeInput(root, url, options)
	} else {
		content, err := html2text.FromReader(file, html2textOptions)
		if err != nil {
			return nil, errors.New("failed to convert html to text: " + err.Error())
		}

		return &InputData{
			Text: content,
		}, nil
	}
}

func UrlInput(url string, options Options) (*InputData, error) {
	options.SkipUrl = true

	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("request to the url failed: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	reader, err := charset.NewReader(resp.Body, contentType)
	if err != nil {
		return nil, errors.New("unable read the page: " + err.Error())
	}

	root, err := html.Parse(reader)
	if err != nil {
		return nil, errors.New("unable to parse html tree: " + err.Error())
	}

	input, err := NodeInput(root, url, options)
	if err != nil {
		return nil, err
	}

	input.Headers = &resp.Header
	return input, nil
}

func NodeInput(root *html.Node, url string, options Options) (i *InputData, err error) {
	var doc *goquery.Document
	if !options.SkipDocument {
		doc = goquery.NewDocumentFromNode(root)

		if !options.SkipUrl {
			if val, ok := doc.Find("link[rel='canonical']").Attr("href"); ok {
				url = val
			} else if val, ok := doc.Find("meta[property='og:url']").Attr("content"); ok {
				url = val
			} else if val, ok := doc.Find("link[rel='alternate']").Attr("href"); ok {
				url = val
			}
		}
	}

	var content string
	if !options.SkipText {
		content, err = html2text.FromHTMLNode(root, html2textOptions)
		if err != nil {
			log.Println("failed to convert html to text: " + err.Error())
		}
	}

	var schema *microdata.Item
	if !options.SkipSchema {
		data, err := microdata.ParseNode(root, url)
		if err != nil {
			log.Println("unable to parse microdata on the page: " + err.Error())
		} else {
			schema = data.GetFirstOfType("http://schema.org/Recipe", "Recipe")
			if schema == nil {
				log.Println("no embedded recipe schema found on: " + url)
			}
		}
	}

	return &InputData{
		Url:      url,
		Text:     content,
		RootNode: root,
		Document: doc,
		Schema:   schema,
	}, nil
}

func getContentType(file *os.File) (string, error) {
	// to sniff the content type only the first 512 bytes are used.
	buf := make([]byte, 512)

	_, err := file.Read(buf)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buf)
	return contentType, nil
}
