package opengraph

import (
	"strings"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser"
)

func Parse(p *parser.InputData, r *model.Recipe) error {
	if p.Document == nil {
		return nil
	}

	if len(r.Url) == 0 {
		if val, ok := p.Document.Find("meta[property='og:url']").Attr("content"); ok {
			r.Url = val
		} else if val, ok := p.Document.Find("link[rel='canonical']").Attr("href"); ok {
			r.Url = val
		}
	}

	if len(r.Name) == 0 {
		if val, ok := p.Document.Find("meta[property='og:name']").Attr("content"); ok {
			r.Name = val
		} else if val, ok := p.Document.Find("meta[property='og:title']").Attr("content"); ok {
			r.Name = val
		}
	}

	if len(r.Description) == 0 {
		if val, ok := p.Document.Find("meta[property='og:description']").Attr("content"); ok {
			r.Description = val
		}
	}

	if len(r.Image) == 0 {
		// TODO: parse array of images
		if val, ok := p.Document.Find("meta[property='og:image']").Attr("content"); ok {
			r.Image = append(r.Image, val)
		}
	}

	if len(r.Language) == 0 {
		if val, ok := p.Document.Find("meta[property='og:locale']").Attr("content"); ok {
			r.Language = strings.Split(val, ",")[0]
			// try content-language header
		} else if p.Headers != nil && len(p.Headers.Get("Content-Language")) != 0 {
			r.Language = strings.Split(p.Headers.Get("Content-Language"), ",")[0]
			// try content-language meta attribute (deprecated)
		} else if attr, ok := p.Document.Find("meta[http-equiv='content-language']").Attr("content"); ok {
			r.Language = strings.Split(attr, ",")[0]
			// retrieve html lang attribute
		} else if val, ok = p.Document.Find("html").Attr("lang"); ok {
			r.Language = val
		}
	}

	if len(r.SiteName) == 0 {
		if val, ok := p.Document.Find("meta[property='og:site_name']").Attr("content"); ok {
			r.SiteName = val
		}
	}

	return nil
}
