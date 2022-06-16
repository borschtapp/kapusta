package model

import (
	"time"
)

// refer to this doc for basic schema
// https://developers.google.com/search/docs/advanced/structured-data/recipe

type Author struct {
	Name        string `json:"name,omitempty"`
	JobTitle    string `json:"jobTitle,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}

type Step struct {
	Name  string `json:"name,omitempty"`
	Text  string `json:"text,omitempty"`
	Url   string `json:"url,omitempty"`
	Image string `json:"image,omitempty"`
}

type Instruction struct {
	Step
	Steps []*Step `json:"itemListElement,omitempty"`
}

// TODO: parse numbers and convert
type Nutrition struct {
	Calories              string `json:"calories,omitempty"`              // The number of calories.
	ServingSize           string `json:"servingSize,omitempty"`           // The serving size, in terms of the number of volume or mass.
	CarbohydrateContent   string `json:"carbohydrateContent,omitempty"`   // The number of grams of carbohydrates.
	CholesterolContent    string `json:"cholesterolContent,omitempty"`    // The number of milligrams of cholesterol.
	FatContent            string `json:"fatContent,omitempty"`            // The number of grams of fat.
	FiberContent          string `json:"fiberContent,omitempty"`          // The number of grams of fiber.
	ProteinContent        string `json:"proteinContent,omitempty"`        // The number of grams of protein.
	SaturatedFatContent   string `json:"saturatedFatContent,omitempty"`   // The number of grams of saturated fat.
	SodiumContent         string `json:"sodiumContent,omitempty"`         // The number of milligrams of sodium.
	SugarContent          string `json:"sugarContent,omitempty"`          // The number of grams of sugar.
	TransFatContent       string `json:"transFatContent,omitempty"`       // The number of grams of trans fat.
	UnsaturatedFatContent string `json:"unsaturatedFatContent,omitempty"` // The number of grams of unsaturated fat.
}

type Rating struct {
	RatingValue float32 `json:"ratingValue,omitempty"`
	RatingCount int     `json:"ratingCount,omitempty"`
	ReviewCount int     `json:"reviewCount,omitempty"`
}

type Video struct {
	Name         string     `json:"name,omitempty"`
	Description  string     `json:"description,omitempty"`
	ThumbnailURL []string   `json:"thumbnailUrl,omitempty"`
	ContentURL   string     `json:"contentUrl,omitempty"`
	EmbedURL     string     `json:"embedUrl,omitempty"`
	UploadDate   *time.Time `json:"uploadDate,omitempty"`
	Duration     string     `json:"duration,omitempty"`
	Expires      *time.Time `json:"expires,omitempty"`
}

// Recipe is the basic struct for the recipe
type Recipe struct {
	Name            string         `json:"name,omitempty"`
	Image           []string       `json:"image,omitempty"`
	Author          *Author        `json:"author,omitempty"`
	Description     string         `json:"description,omitempty"`
	PrepTime        int            `json:"prepTime,omitempty"`
	CookTime        int            `json:"cookTime,omitempty"`
	TotalTime       int            `json:"totalTime,omitempty"`
	Keywords        string         `json:"keywords,omitempty"`
	Yield           int            `json:"recipeYield,omitempty"`
	Category        string         `json:"recipeCategory,omitempty"`
	Cuisine         string         `json:"recipeCuisine,omitempty"`
	Nutrition       *Nutrition     `json:"nutrition,omitempty"`
	Ingredients     []string       `json:"recipeIngredient,omitempty"`
	Instructions    []*Instruction `json:"recipeInstructions,omitempty"`
	DateModified    *time.Time     `json:"dateModified,omitempty"`
	DatePublished   *time.Time     `json:"datePublished,omitempty"`
	AggregateRating *Rating        `json:"aggregateRating,omitempty"`
	Video           *Video         `json:"video,omitempty"`

	Url       string   `json:"url,omitempty"`
	SiteName  string   `json:"siteName,omitempty"`
	Language  string   `json:"language,omitempty"`
	Footnotes []string `json:"footnotes,omitempty"`
	Links     []string `json:"links,omitempty"`
}
