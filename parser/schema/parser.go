package schema

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"borscht.app/kapusta/model"
	"borscht.app/kapusta/utils"
)

func Parse(p *model.InputData, r *model.Recipe) error {
	if p.Schema == nil {
		return nil
	}

	if val, ok := getPropertyString(p.Schema, "url"); ok {
		r.Url = val
	} else if val, ok := getPropertyStringOrChild(p.Schema, "mainEntityOfPage", "@id"); ok {
		r.Url = val
	}

	if val, ok := getPropertyString(p.Schema, "name"); ok {
		r.Name = utils.CleanupInline(val)
	}

	if val, ok := getPropertyString(p.Schema, "recipeCategory"); ok {
		r.Category = utils.CleanupInline(val)
	}

	if val, ok := getPropertyDuration(p.Schema, "totalTime"); ok {
		r.TotalTime = int(val.Minutes())
	}

	if val, ok := getPropertyDuration(p.Schema, "cookTime"); ok {
		r.CookTime = int(val.Minutes())
	}

	if val, ok := getPropertyDuration(p.Schema, "prepTime"); ok {
		r.CookTime = int(val.Minutes())
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
			return errors.New("unable to parse recipeYield: " + fmt.Sprint(val))
		}
	}

	if nested, ok := p.Schema.GetNested("image"); ok {
		for _, item := range nested.Items {
			if val, ok := getPropertyString(item, "url"); ok {
				r.Image = append(r.Image, val)
			}
		}
	} else if val, ok := getPropertyString(p.Schema, "image"); ok {
		r.Image = append(r.Image, val)
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

	if val, ok := getPropertyString(p.Schema, "inLanguage", "language"); ok {
		r.Language = val
	}

	if ingredients, ok := p.Schema.GetProperties("recipeIngredient", "ingredients"); ok {
		for _, val := range ingredients {
			if text, ok := getStringOrChild(val, "name"); ok {
				r.Ingredients = append(r.Ingredients, utils.CleanupInline(text))
			}
		}
	}

	if nested, ok := p.Schema.GetNested("recipeInstructions"); ok {
		for _, item := range nested.Items {
			if item.IsOfType("http://schema.org/HowToStep", "HowToStep") {
				var step = parseInstructionSteps(item)
				r.Instructions = append(r.Instructions, &model.Instruction{Step: step})
			} else if item.IsOfType("http://schema.org/HowToSection", "HowToSection") {
				var section model.Instruction
				if name, ok := getPropertyString(item, "name"); ok {
					section.Name = utils.CleanupInline(name)
				}

				if nested, ok := item.GetNested("itemListElement"); ok {
					for _, item := range nested.Items {
						var step = parseInstructionSteps(item)
						section.Steps = append(section.Steps, &step)
					}
				}
				r.Instructions = append(r.Instructions, &section)
			} else if item.IsOfType("http://schema.org/ItemList", "ItemList") {
				if nested, ok := item.GetNested("itemListElement"); ok {
					for _, item := range nested.Items {
						var step = parseInstructionSteps(item)
						r.Instructions = append(r.Instructions, &model.Instruction{Step: step})
					}
				}
			} else {
				return errors.New("unknown instruction type: " + fmt.Sprint(item.Types))
			}
		}
	} else if val, ok := getPropertyString(p.Schema, "recipeInstructions"); ok {
		for _, step := range utils.SplitParagraphs(val) {
			r.Instructions = append(r.Instructions, &model.Instruction{Step: model.Step{Text: step}})
		}
	}

	if item, ok := p.Schema.GetNestedItem("aggregateRating"); ok {
		var rating model.Rating
		if val, ok := getPropertyFloat(item, "ratingValue"); ok {
			rating.RatingValue = val
		}
		if val, ok := getPropertyInt(item, "ratingCount"); ok {
			rating.RatingCount = val
		}
		if val, ok := getPropertyInt(item, "reviewCount"); ok {
			rating.ReviewCount = val
		}
		r.AggregateRating = &rating
	}

	if item, ok := p.Schema.GetNestedItem("author", "creator"); ok {
		var author model.Author
		if val, ok := getPropertyString(item, "name"); ok {
			author.Name = utils.CleanupInline(val)
		}
		if val, ok := getPropertyString(item, "jobTitle"); ok {
			author.JobTitle = utils.CleanupInline(val)
		}
		if val, ok := getPropertyString(item, "description"); ok {
			author.Description = utils.CleanupInline(val)
		}
		if val, ok := getPropertyString(item, "url"); ok {
			author.Url = val
		}
		r.Author = &author
	} else if val, ok := getPropertyString(p.Schema, "author"); ok {
		r.Author = &model.Author{Name: utils.CleanupInline(val)}
	}

	if val, ok := getPropertyString(p.Schema, "recipeCuisine"); ok {
		r.Cuisine = utils.CleanupInline(val)
	}

	if val, ok := getPropertyString(p.Schema, "description"); ok {
		r.Description = utils.CleanupInline(val)
	}

	if keywords, ok := p.Schema.GetProperties("keywords"); ok {
		var arr []string
		for _, v := range keywords {
			arr = append(arr, utils.CleanupInline(v.(string)))
		}
		r.Keywords = strings.Join(arr, ", ")
	}

	if val, ok := getPropertyString(p.Schema, "datePublished"); ok {
		if val, err := time.Parse(time.RFC3339, val); err == nil {
			r.DatePublished = &val
		}
	}

	if val, ok := getPropertyString(p.Schema, "dateModified"); ok {
		if val, err := time.Parse(time.RFC3339, val); err == nil {
			r.DateModified = &val
		}
	}

	return nil
}
