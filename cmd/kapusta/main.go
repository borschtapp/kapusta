package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"borscht.app/kapusta"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [options] [url]:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nScrapes a Recipe data from a given webpage. Provide an URL to a valid HTML5 document.\n")
	}
	flag.Parse()

	switch len(flag.Args()) {
	case 1:
		recipeUrl := flag.Args()[0]

		recipe, err := kapusta.ScrapeUrl(recipeUrl)
		if err != nil {
			log.Fatal("Unable to scrape target: " + err.Error())
		}

		fmt.Println(recipe)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
