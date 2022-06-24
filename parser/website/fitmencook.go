package website

import (
	"github.com/PuerkitoBio/goquery"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/utils"
)

func ParseFitMenCook(p *model.InputData, r *model.Recipe) error {
	if p.Document != nil {
		if s := p.Document.Find(".recipe-ingredients h4 strong").First(); len(s.Nodes) != 0 {
			if val := utils.FindNumber(s.Text()); val > 0 {
				r.Yield = val
			}
		}

		if s := p.Document.Find("div.recipe-ingredients li"); len(s.Nodes) != 0 {
			s.Each(func(i int, s *goquery.Selection) {
				text := utils.CleanupInline(s.Text())
				if text != "" {
					r.Ingredients = append(r.Ingredients, text)
				}
			})
		}

		if s := p.Document.Find("div.recipe-steps > ol:first-of-type li"); len(s.Nodes) != 0 {
			s.Each(func(i int, s *goquery.Selection) {
				text := utils.CleanupInline(s.Text())
				if text != "" {
					r.Instructions = append(r.Instructions, &model.Instruction{Step: model.Step{Text: text}})
				}
			})
		}
	}

	return nil
}
