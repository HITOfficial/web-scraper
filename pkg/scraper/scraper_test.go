package scraper_test

import (
	"sync"
	"time"

	"github.com/hitofficial/web-scraper/pkg/scraper"
	website "github.com/hitofficial/web-scraper/pkg/website"
	ginkgo "github.com/onsi/ginkgo/v2"
	gomega "github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Scraper", ginkgo.Ordered, func() {
	var (
		wg *sync.WaitGroup
	)

	ginkgo.BeforeEach(func() {
		wg = &sync.WaitGroup{}
	})

	ginkgo.AfterEach(func() {
	})

	ginkgo.It("should scrape the website and count the words", func() {
		url := "https://example.com/"
		result := make(chan website.Website)
		wg.Add(1)
		go scraper.ScrapePage(url, result, wg)
		var receivedWebsite website.Website
		gomega.Eventually(result, time.Second*5, time.Second).Should(gomega.Receive(&receivedWebsite))

		ginkgo.By("checking the result")
		gomega.Expect(receivedWebsite.Error).To(gomega.BeNil())
		gomega.Expect(receivedWebsite.URL).To(gomega.Equal(url))
		ginkgo.By("total number of words should be 21")
		gomega.Expect(len(receivedWebsite.Words)).To(gomega.Equal(21))
		ginkgo.By("checking the count of some words")
		gomega.Expect(receivedWebsite.Words["in"]).To(gomega.Equal(3))
		gomega.Expect(receivedWebsite.Words["domain"]).To(gomega.Equal(3))

		gomega.Expect(receivedWebsite.Words["this"]).To(gomega.Equal(2))
		gomega.Expect(receivedWebsite.Words["for"]).To(gomega.Equal(2))
		gomega.Expect(receivedWebsite.Words["use"]).To(gomega.Equal(2))

		gomega.Expect(receivedWebsite.Words["more"]).To(gomega.Equal(1))
		gomega.Expect(receivedWebsite.Words["information"]).To(gomega.Equal(1))

		ginkgo.By("checking the count of some words which doesnt occur")
		gomega.Expect(receivedWebsite.Words["wrongWord"]).To(gomega.Equal(0))
		gomega.Expect(receivedWebsite.Words["my"]).To(gomega.Equal(0))
		gomega.Expect(receivedWebsite.Words["ididntoccur"]).To(gomega.Equal(0))

	})

	ginkgo.It("should handle error when url is not reachable", func() {
		url := "https://invalidurl"
		result := make(chan website.Website)
		wg.Add(1)
		go scraper.ScrapePage(url, result, wg)
		var receivedWebsite website.Website
		gomega.Eventually(result, time.Second*5, time.Second).Should(gomega.Receive(&receivedWebsite))

		ginkgo.By("checking the result")
		gomega.Expect(receivedWebsite.Error).To(gomega.HaveOccurred())
		gomega.Expect(receivedWebsite.URL).To(gomega.Equal(url))
		gomega.Expect(receivedWebsite.Words).To(gomega.BeEmpty())
	})

})
