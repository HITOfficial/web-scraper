package scraper_test

import (
	"testing"

	ginkgo "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func TestScrapePage(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ScrapePage Suite")
}
