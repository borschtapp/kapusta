package parser

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
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
	Headers  *http.Header
	RootNode *html.Node
	Document *goquery.Document
	Schema   *microdata.Item
}

func FileInput(fileName string, options Options) (*InputData, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("unable to read the file: " + err.Error())
	}
	defer file.Close()

	contentType, err := getContentType(file)
	if err != nil {
		return nil, errors.New("unable to detect content type: " + err.Error())
	}

	if strings.HasPrefix(contentType, "text/html") {
		root, err := html.Parse(file)
		if err != nil {
			return nil, errors.New("unable to parse html tree: " + err.Error())
		}

		url := "file://" + fileName
		return NodeInput(root, url, options)
	} else {
		return &InputData{}, nil
	}
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
			}
		}
	}

	var schema *microdata.Item
	if !options.SkipSchema {
		data, err := microdata.ParseNode(root, url)
		if err != nil {
			return nil, errors.New("unable to parse microdata on the page: " + err.Error())
		}

		schema = data.GetFirstOfType("http://schema.org/Recipe", "Recipe")
		if schema == nil {
			return nil, errors.New("no embedded recipe Schema found")
		}
	}

	return &InputData{
		Url:      url,
		RootNode: root,
		Document: doc,
		Schema:   schema,
	}, nil
}

func UrlInput(url string, options Options) (*InputData, error) {
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

	options.SkipUrl = true
	input, err := NodeInput(root, url, options)
	if err != nil {
		return nil, err
	}

	input.Headers = &resp.Header
	return input, nil
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
