package main

// Heap to get k most polular items from results
import (
	"container/heap"
	"sync"

	maxheap "github.com/hitofficial/web-scraper/pkg/heap"
	scraper "github.com/hitofficial/web-scraper/pkg/scraper"
	utils "github.com/hitofficial/web-scraper/pkg/utils"
	website "github.com/hitofficial/web-scraper/pkg/website"
)

func main() {

	urls := []string{
		"https://webscraper.io/test-sites/e-commerce/allinone",
		"https://scrapeme.live/shop",
		"https://scrapeme.live/shop/page/2/",
		"https://scrapeme.live/shop/page/3/",
		"https://scrapeme.live/shop/page/4/",
		"https://scrapeme.live/shop/page/5/",
	}
	k := 5

	results := make(chan website.Website)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go scraper.ScrapePage(url, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()

	websiteArrayWC := make([]website.Website, 0)

	for website := range results {
		websiteArrayWC = append(websiteArrayWC, website)
	}

	words := make(map[string]maxheap.WordCount)
	sliceOfWordCountHeap := make([]maxheap.MaxHeap, 0)
	urlSlice := make([]string, 0)
	errorSlice := make([]error, 0)
	for idx, v := range websiteArrayWC {

		sliceOfWordCountHeap = append(sliceOfWordCountHeap, maxheap.MaxHeap{})
		urlSlice = append(urlSlice, v.URL)
		errorSlice = append(errorSlice, v.Error)

		if v.Error != nil {
			continue
		}

		// init empty heap
		for word, count := range v.Words {
			entity, ok := words[word]
			// add to global counter
			if ok {
				entity.Count += count
				words[word] = entity
			} else {
				words[word] = maxheap.WordCount{
					Word:  word,
					Count: count,
				}
			}
			heap.Push(
				&sliceOfWordCountHeap[idx], maxheap.WordCount{
					Word:  word,
					Count: count,
				})
		}
	}

	globalMaxHeap := &maxheap.MaxHeap{}
	for _, v := range words {
		heap.Push(globalMaxHeap, v)
	}

	// Most common K words per URL
	for i := 0; i < len(urlSlice); i++ {
		wc := maxheap.PopKLargestWordCounts(&sliceOfWordCountHeap[i], k)
		utils.SummaryOfURL(urlSlice[i], errorSlice[i], k, &wc)
	}
	res := maxheap.PopKLargestWordCounts(globalMaxHeap, 5)
	utils.SummaryOfURL("SUMMARY", nil, k, &res)
}
