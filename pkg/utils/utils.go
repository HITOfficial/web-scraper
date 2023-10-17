package utils

import (
	"fmt"

	heap "github.com/hitofficial/web-scraper/pkg/heap"
)

// SpecialSignRemoval removes special signs from string
func SpecialSignRemoval(r rune) rune {
	specialSymbols := "!@#$%^&*()_-+={[}]\\|;:\"<>?/., "
	for _, symbol := range specialSymbols {
		if r == rune(symbol) {
			return -1
		}
	}
	return r
}

// SummaryOfURL prints summary of url
func SummaryOfURL(url string, urlError error, k int, wc *[]heap.WordCount) {
	fmt.Println("\n")
	if urlError != nil {
		fmt.Printf("couldn't scrap url '%v' error '%v'  \n", url, urlError)
		return
	}
	fmt.Printf("url '%v', %v most popular words \n", url, k)
	for _, v := range *wc {
		fmt.Printf("word '%v'  occurs '%v' \n", v.Word, v.Count)
	}

}
