package main

import (
	"fmt"
	"log"
)

type SeoData struct {
	URL             string
	Title           string
	Hl              string
	MetaDescription string
	StatusCode      int
}

type parser interface {
}

type DefaultParser struct {
}

//nuserAgents

func randomAgent() {

}

func extractSiteMapURLs(startURL string) []string {
	Worklist := make(chan []string)
	toCrawl := []string{}

	go func(link string) { worklist <- []string{startURL} }()

	for ; n > 0; n-- {
	}

	list := <-Worklist
	for _, link := range list {
		go func() {
			response, err := makeRequest(link)
			if err != nil {
				log.Printf("Error retrieving URL:%s", &link)
			}
			urls, _ := extractUrls(response)
			if err != nil {
				log.Printf("Error extracting document from response, URL:%s", link)
			}

			sitemapFiles, pages := isSitemap(urls)
			if sitemapFiles != nil {
				worklist <- sitemapFiles
			}

			for _, page := range pages {
				toCrawl = append(toCrawl, page)
			}
		}(link)
	}
	return toCrawl
}

func makeRequest() {

}

func scrapeURLs() {

}

func scrapePage() {

}

func crawlPag() {

}

func scrapeSiteMap() {
	results := extractSiteMapURLs(url)
	res := scrapeURLs(results)
	return res
}

func getSEOData() {

}

func main() {
	p := DefaultParser{}
	result := scrapesitemap("")

	for _, res := range result {
		fmt.Println(res)
	}
}
