package schema

import (
	"fmt"
	"log"
	"strings"
	"time"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/parser"
	"borscht.app/kapusta/utils"
)

func Parse(p *parser.InputData, r *model.Recipe) error {
	if p.Schema == nil {
		return nil
	}

	if val, ok := p.Schema.GetProperty("url"); ok {
		r.Url = val.(string)
	} else if val, ok := p.Schema.GetProperty("mainEntityOfPage"); ok {
		r.Url = val.(string)
	}

	if val, ok := p.Schema.GetProperty("name"); ok {
		r.Name = val.(string)
	}

	if val, ok := p.Schema.GetProperty("recipeCategory"); ok {
		r.Category = val.(string)
	}

	if val, ok := p.Schema.GetProperty("totalTime"); ok {
		if val, ok := utils.GetDurationMinutes(val.(string)); ok {
			r.TotalTime = val
		}
	}

	if val, ok := p.Schema.GetProperty("cookTime"); ok {
		if val, ok := utils.GetDurationMinutes(val.(string)); ok {
			r.CookTime = val
		}
	}

	if val, ok := p.Schema.GetProperty("prepTime"); ok {
		if val, ok := utils.GetDurationMinutes(val.(string)); ok {
			r.PrepTime = val
		}
	}

	if val, ok := p.Schema.GetProperty("recipeYield"); ok {
		switch val.(type) {
		case string:
			if val, err := utils.ParseInt(val.(string)); err == nil {
				r.Yield = val
			}
		case float64:
			r.Yield = int(val.(float64))
		default:
			log.Printf("Unable to parse recipeYield: %v", val)
		}
	}

	if nested, ok := p.Schema.GetNested("image"); ok {
		for _, item := range nested.Items {
			if val, ok := item.GetProperty("url"); ok {
				r.Image = append(r.Image, val.(string))
			}
		}
	} else if val, ok := p.Schema.GetProperty("image"); ok {
		r.Image = append(r.Image, val.(string))
	}

	if item, ok := p.Schema.GetNestedItem("nutrition"); ok {
		var nutrition model.Nutrition
		for key, val := range item.Properties {
			strVal := fmt.Sprint(val[0])

			switch key {
			case "calories":
				nutrition.Calories = strVal
			case "servingSize":
				nutrition.ServingSize = strVal
			case "carbohydrateContent":
				nutrition.CarbohydrateContent = strVal
			case "cholesterolContent":
				nutrition.CholesterolContent = strVal
			case "fatContent":
				nutrition.FatContent = strVal
			case "fiberContent":
				nutrition.FiberContent = strVal
			case "proteinContent":
				nutrition.ProteinContent = strVal
			case "saturatedFatContent":
				nutrition.SaturatedFatContent = strVal
			case "sodiumContent":
				nutrition.SodiumContent = strVal
			case "sugarContent":
				nutrition.SugarContent = strVal
			case "transFatContent":
				nutrition.TransFatContent = strVal
			case "unsaturatedFatContent":
				nutrition.UnsaturatedFatContent = strVal
			}
		}
		r.Nutrition = &nutrition
	}

	if val, ok := p.Schema.GetProperty("inLanguage"); ok {
		r.Language = val.(string)
	} else if val, ok := p.Schema.GetProperty("language"); ok {
		r.Language = val.(string)
	}

	if ingredients, ok := p.Schema.GetProperties("recipeIngredient"); ok {
		for _, v := range ingredients {
			r.Ingredients = append(r.Ingredients, v.(string))
		}
	} else if ingredients, ok := p.Schema.GetProperties("ingredients"); ok {
		for _, v := range ingredients {
			r.Ingredients = append(r.Ingredients, v.(string))
		}
	}

	if nested, ok := p.Schema.GetNested("recipeInstructions"); ok {
		for _, item := range nested.Items {
			var instr model.Instruction
			if item.IsOfType("HowToStep") {
				if val, ok := item.GetProperty("text"); ok {
					instr.Text = val.(string)
				}
				if val, ok := item.GetProperty("name"); ok && val != instr.Text {
					instr.Name = val.(string)
				}
				if val, ok := item.GetProperty("image"); ok {
					instr.Image = val.(string)
				}
			} else {
				log.Println("Unknown instruction type:", item.Types)
			}
			r.Instructions = append(r.Instructions, &instr)
		}
	} else if val, ok := p.Schema.GetProperty("recipeInstructions"); ok {
		r.Instructions = append(r.Instructions, &model.Instruction{Text: val.(string)})
	}

	if item, ok := p.Schema.GetNestedItem("aggregateRating"); ok {
		var rating model.Rating
		if val, ok := item.GetProperty("ratingValue"); ok {
			rating.RatingValue = val.(float64)
		}
		if val, ok := item.GetProperty("ratingCount"); ok {
			rating.RatingCount = val.(int)
		}
		if val, ok := item.GetProperty("reviewCount"); ok {
			rating.ReviewCount = val.(int)
		}
		r.AggregateRating = &rating
	}

	if item, ok := p.Schema.GetNestedItem("author"); ok {
		var author model.Author
		if val, ok := item.GetProperty("name"); ok {
			author.Name = val.(string)
		}
		if val, ok := item.GetProperty("jobTitle"); ok {
			author.JobTitle = val.(string)
		}
		if val, ok := item.GetProperty("description"); ok {
			author.Description = val.(string)
		}
		if val, ok := item.GetProperty("url"); ok {
			author.Url = val.(string)
		}
		r.Author = &author
	} else if val, ok := p.Schema.GetProperty("author"); ok {
		r.Author = &model.Author{Name: val.(string)}
	} else if item, ok := p.Schema.GetNestedItem("creator"); ok {
		var author model.Author
		if val, ok := item.GetProperty("name"); ok {
			author.Name = val.(string)
		}
		if val, ok := item.GetProperty("jobTitle"); ok {
			author.JobTitle = val.(string)
		}
		if val, ok := item.GetProperty("description"); ok {
			author.Description = val.(string)
		}
		if val, ok := item.GetProperty("url"); ok {
			author.Url = val.(string)
		}
		r.Author = &author
	}

	if val, ok := p.Schema.GetProperty("recipeCuisine"); ok {
		r.Cuisine = val.(string)
	}

	if val, ok := p.Schema.GetProperty("description"); ok {
		r.Description = strings.Trim(val.(string), "\n")
	}

	if keywords, ok := p.Schema.GetProperties("keywords"); ok {
		var arr []string
		for _, v := range keywords {
			arr = append(arr, v.(string))
		}
		r.Keywords = strings.Join(arr, ", ")
	}

	if val, ok := p.Schema.GetProperty("datePublished"); ok {
		if val, err := time.Parse(time.RFC3339, val.(string)); err == nil {
			r.DatePublished = &val
		}
	}

	if val, ok := p.Schema.GetProperty("dateModified"); ok {
		if val, err := time.Parse(time.RFC3339, val.(string)); err == nil {
			r.DateModified = &val
		}
	}

	return nil
}
