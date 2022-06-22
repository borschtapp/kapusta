package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html/charset"

	"borscht.app/kapusta"
	"borscht.app/kapusta/testdata"
	"borscht.app/kapusta/utils"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [options] [url]:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nAdds new recipe webpage to testdata catalog. Use to automate some routine.\n")
	}
	flag.Parse()

	switch len(flag.Args()) {
	case 1:
		recipeUrl := flag.Args()[0]
		createWebsiteTestdata(recipeUrl)
		fmt.Println("Done, testdata page added!")
	default:
		flag.Usage()
		os.Exit(1)
	}
}

func createWebsiteTestdata(recipeUrl string) {
	alias := utils.ParserAlias(recipeUrl)
	websiteFileName := testdata.WebsitesDir + alias + testdata.WebsiteExt
	recipeFileName := testdata.WebsitesDir + alias + testdata.RecipeExt

	if _, err := os.Stat(websiteFileName); err == nil {
		log.Fatal("Testdata already exists for the alias: " + alias)
	}

	resp, err := http.Get(recipeUrl)
	if err != nil {
		log.Fatal("Unable to fetch content: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Bad response status: " + resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	reader, err := charset.NewReader(resp.Body, contentType)
	if err != nil {
		log.Fatal("Unable to read the page: " + err.Error())
	}

	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal("Unable to read the content: " + err.Error())
	}

	if err = ioutil.WriteFile(websiteFileName, content, 0644); err != nil {
		log.Fatal("Unable to create file: " + err.Error())
	}

	recipe, err := kapusta.ScrapeFile(websiteFileName)
	if err != nil {
		log.Fatal("Unable to scrape recipe: " + err.Error())
	}

	if err = ioutil.WriteFile(recipeFileName, []byte(recipe.String()), 0644); err != nil {
		log.Fatal("Unable to create recipe file: " + err.Error())
	}
}
