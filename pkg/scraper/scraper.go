package scraper

import (
	"strings"
	"sync"

	"github.com/gocolly/colly"
	utils "github.com/hitofficial/web-scraper/pkg/utils"
	website "github.com/hitofficial/web-scraper/pkg/website"
)

// ScrapePage scrapes the page and counts the words using colly
func ScrapePage(url string, result chan website.Website, wg *sync.WaitGroup) {
	defer wg.Done()

	c := colly.NewCollector()

	website := website.Website{
		URL:   url,
		Words: make(map[string]int),
		Error: nil,
	}

	c.OnHTML("body", func(e *colly.HTMLElement) {
		words := strings.Fields(e.Text)
		for _, word := range words {
			word = strings.ToLower(word)
			word = strings.Map(utils.SpecialSignRemoval, word)
			if word == "" {
				continue
			}
			website.Words[word]++
		}
	})
	err := c.Visit(url)
	if err != nil {
		website.Error = err
	}
	result <- website
}
