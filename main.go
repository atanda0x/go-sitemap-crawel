package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type SeoData struct {
	URL             string
	Tile            string
	H1              string
	MetaDescription string
	StatusCode      int
}
type parser interface{}

var userAgents = []string{}

type DefaultParser struct {
}

func randomUserAgents() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func isSitemap(urls []string) ([]string, []string) {
	sitemapFiles := []string{}
	pages := []string{}
	for _, page := range urls {
		if foundSitmap == true {
			fmt.Println("Found sitemap", page)
			sitemapFiles = append(sitemapFiles, page)
		} else {
			pages = append(pages, page)
		}
	}
	return sitemapFiles, pages
}

func extraSiteMapURLs(startURL string) []string {
	worklist := make(chan []string)
	toCrawel := []string{}
	var n int
	n++
	go func() {
		worklist <- []string{startURL}
	}()

	for ; n > 0; n-- {

		list := <-worklist
		for _, link := range list {
			n++
			go func(link string) {
				response, err := makeRequest(link)
				if err != nil {
					log.Printf("Error retrieving URL:%S", link)
				}
				urls, _ := extractUrls(response)
				if err != nil {
					log.Printf("Error extracting doc from response, URL:%s", link)
				}

				siteMapFiles, Pages := isSitemap(urls)
				if siteMapFiles != nil {
					worklist <- siteMapFiles
				}

				for _, page := range Pages {
					toCrawel = append(toCrawel, page)
				}
			}(link)
		}
	}
	return toCrawel
}

func makeRequest() {

}

func scrapeURLs() {

}

func crawelPage() {

}

func getSEOData() {

}

func scrapeSiteMap(url string) []SeoData {
	results := extraSiteMapURLs(url)
	res := scrapeURLs(results)
	return res
}

func main() {
	p := DefaultParser{}
	results := scrapeSiteMap("")
	for _, res := range results {
		fmt.Println(res)
	}
}
