package main

// Heap to get k most polular items from results
import (
	"container/heap"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

// HEAP
type MaxHeap []WordCount

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Count > h[j].Count }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(WordCount))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//

type Website struct {
	URL   string
	Words map[string]int
}

type WordCount struct {
	Word  string
	Count int
}

func specialSignRemoval(r rune) rune {
	specialSymbols := "!@#$%^&*()_-+={[}]\\|;:\"<>?/.,"
	for _, symbol := range specialSymbols {
		if r == rune(symbol) {
			return -1
		}
	}
	return r
}

func getKLargestWordCount(maxHeap *MaxHeap, k int) []WordCount {
	KLargest := make(MaxHeap, 0)
	for i := 0; i < min(maxHeap.Len(), k); i++ {
		wc := heap.Pop(maxHeap).(WordCount)
		KLargest = append(KLargest, wc)
	}
	return KLargest
}

func scrapePage(url string, result chan Website, wg *sync.WaitGroup) {
	defer wg.Done()

	c := colly.NewCollector()

	website := Website{
		URL:   url,
		Words: make(map[string]int),
	}

	c.OnHTML("body", func(e *colly.HTMLElement) {
		words := strings.Fields(e.Text)
		for _, word := range words {
			word = strings.ToLower(word)
			word = strings.Map(specialSignRemoval, word)
			website.Words[word]++
		}
	})
	err := c.Visit(url)
	if err != nil {
		log.Printf("cannot scrape page [url]: %s [error]: %v", url, err)
		return
	}
	result <- website
}

func main() {

	urls := []string{
		"https://scrapeme.live/shop",
		"https://scrapeme.live/shop/page/2/",
		"https://scrapeme.live/shop/page/3/",
		"https://scrapeme.live/shop/page/4/",
		"https://scrapeme.live/shop/page/5/",
	}

	results := make(chan Website)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go scrapePage(url, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	websiteMap := make(map[string]Website)

	// Count total

	for website := range results {
		websiteMap[website.URL] = website
	}

	words := make(map[string]WordCount)

	for _, v := range websiteMap {
		for word, count := range v.Words {
			if entity, ok := words[word]; ok {
				entity.Count += count
				words[word] = entity
			} else {
				words[word] = WordCount{
					Word:  word,
					Count: count,
				}
			}
		}
	}

	globalMaxHeap := &MaxHeap{}
	for _, v := range words {
		heap.Push(globalMaxHeap, v)
	}
	k := 5
	res := getKLargestWordCount(globalMaxHeap, 5)
	fmt.Print("most popular words: ", k, res)
}
