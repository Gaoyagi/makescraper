package main

import (
	"fmt"
	"encoding/json"
	"github.com/gocolly/colly"
	"io/ioutil"
)


type Scraped struct {
	Content string	`json:"content"`
} 

//INVOKE SCRAPER IN THE ENDPOINt function isntead of the ain


//order is: c= colly.newcollector, c.Visit, c.onRequest, c.onResponse, c.onHtml or c.onError, c.onXML, c.onDcrapped

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	//
	holder := make([]Scraped, 0)
	// Instantiate default collector
	c := colly.NewCollector()

	// Before making a request print "Visiting ..."
	//whenever  there is  a request proc this
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// // On every a element which has href attribute call callback
	// // finds all tags/eselctors with an "a href"
	// c.OnHTML("a[href]", func(e *colly.HTMLElement){
	// 	// determines what attributes to scrape from the site
	// 	// everythign right before an equals sign in html
    //     link := e.Attr("href")
	// 	// Print link
	// 	fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	// 	//selector = := e.Attr("href")
	// })
	

	c.OnHTML("p strong", func(e *colly.HTMLElement) {
	//c.OnHTML("a[hrf] img", func(e *colly.HTMLElement) {
	//c.OnHTML("#post-37968", func(e *colly.HTMLElement) {
	//c.OnHTML("#menu-footer", func(e *colly.HTMLElement) {
	//c.OnHTML(".entry-title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		temp := Scraped { Content: e.Text}
		holder = append(holder, temp)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.bkacontent.com/40-popular-idioms-and-their-meanings/")


	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })
	
	// c.OnError(func(_ *colly.Response, err error) {
	// 	fmt.Println("Something went wrong:", err)
	// })
	
	// c.OnResponse(func(r *colly.Response) {
	// 	fmt.Println("Visited", r.Request.URL)
	// })
	
	// c.OnHTML("a[href]", func(e *colly.HTMLElement) {
	// 	e.Request.Visit(e.Attr("href"))
	// })
	
	// c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	// 	fmt.Println("First column of a table row:", e.Text)
	// })
	
	// c.OnXML("//h1", func(e *colly.XMLElement) {
	// 	fmt.Println(e.Text)
	// })
	
	


	file, _ := json.MarshalIndent(holder, "", " ")
	_ = ioutil.WriteFile("output.json", file, 0644)
}
