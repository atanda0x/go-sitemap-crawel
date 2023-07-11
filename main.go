package main

type SeoData struct {
	URL  string 
	Title	string
	Hl	string
	MetaDescription	string
	StatusCode	int
}

type parser interface{

} 

type DefaultParser struct {
	
}

userAgents

func randomAgent()  {
	
}

func extractSiteMapURLs()  {
	makeRequest
}

func makeRequest()  {
	
}

func scrapeURLs()  {
	
}

func scrapePage()  {
	
}


func crawlPag()  {
	
}

func scrapeSiteMap()  {
	results := extractSiteMapURLs(url)
	res := scrapeURLs(results)
}


func getSEOData()  {
	
}


func main()  {
	p := DefaultParser{}
	result := scrapesitemap("")

	for _, res := range result {
		fmt.Println(res)
	}
}