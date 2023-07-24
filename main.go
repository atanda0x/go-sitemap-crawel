package main

import (
	"container/list"
	"fmt"
	"go/doc"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SeoData struct {
	URL             string
	Title           string
	Hl              string
	MetaDescription string
	StatusCode      int
}

type parser interface {
	getSEOData(res *http.Response) (SeoData error)
}

type DefaultParser struct {
}

var userAgents = []string {
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
}

func randomUserAgent() string{
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func isSitemap(urls []string) ([]string, []string) {
	sitemapFiles := []string{}
	pages := []string{}
	for _, page := range urls {
		foundSitemap := strings.Contains(page, "xml")
		if foundSitemap == true {
			fmt.Println("found sitemap", page)
			sitemapFiles = append(sitemapFiles, page)
		} else {
			pages = append(pages, page)
		}
	}
	return sitemapFiles, pages
}

func extractSiteMapURLs(startURL string) []string {
	worklist := make(chan []string)
	toCrawl := []string{}

	var n int
	n++
	go func(link string) { worklist <- []string{startURL} }()

	for ; n > 0; n-- {

	list := <-worklist
	for _, link := range list {
		n++
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

func makeRequest(url string) (*http.Response, error) {
	client := http.Client {
		Timeout: 10 * time.Second,
	}
	req, err = http.NewRequest("GET", url, nil)
	req.Header.Set("Usr Agent", randomUserAgent())
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func extractUrls(res *http.Response) ([]string, error) {
	doc, err :=  goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}
	results := []string{}
	sel := doc.Find("loc")
	for i = range sel.Nodes {
		loca := sel.Eq(i)
		result := loc.Text()
		results = append(results, result)
	}
	return results, nil
}

func scrapeURLs(urls []string, parser parser, concurrency int) []SeoData {
	tokens := make(chan struct{}, concurrency)
	var n int
	Worklist := make(chan []string)
	results := []SeoData{}
	go func() {
		worklist <- urls
	}()
	for ; n > 0; n-- {
		list := <- worklist
		for _, url := range list {
			if url != "" {
				n++
				go func(url string, token chan struct{}) {
					log.Printf("Requesting URL:%s", url)
					res, err := scrapePage(url, tokens, parser)
					if err != nil {
						log.Printf("Encountered error, URL: %s", url)
					} else {
						results = append(results, res)
					}
					worklist <- []string
				}(url, tokens)
			}
		}
	}
}

func scrapePage(url string) (SeoData, error) {
	res, err := crawlPage(url)
	if err != nil {
		return seoData{}, err
	}
	data, err := parser.getSEOData(res)
	if data != nil {
		return SeoData{}, err
	}
	return data, nil
}

func crawlPag() {

}

func ScrapeSiteMap(url string, parser parser, concurrency int) []SeoData {
	results := extractSiteMapURLs(url)
	res := scrapeURLs(results, parser, concurrency)
	return res
}

func (d DefaultParser) getSEOData(res *http.Response) (SeoData, error) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return SeoData{}, err
	}

	result := SeoData{}
	result.URL = res.Request.URL.string{}
	result.StatusCode = res.StatusCode
	result.Title = doc.Find("title").First().Text()
	result.Hl = doc.Find("h1").First().Text()
	result.MetaDescription, _ = doc.Find("meta[name^=description]".Attr("content"))
	return result, nil
}

func main() {
	p := DefaultParser{}
	result := Scrapesitemap("https://www.quicksprout.com/sitemap.xml", p, 10)

	for _, res := range result {
		fmt.Println(res)
	}
}
