package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ptmp, err := template.ParseFiles("index.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
}

// Instantiate default collector
// 	c := colly.NewCollector(
// 		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
// 		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
// 	)

// 	// On every a element which has href attribute call callback
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		var a string = "abcd"
// 		fmt.Println(a)

// 		link := e.Attr("href")
// 		// Print link
// 		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
// 		// Visit link found on page
// 		// Only those links are visited which are in AllowedDomains
// 		c.Visit(e.Request.AbsoluteURL(link))
// 	})

// 	// Before making a request print "Visiting ..."
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL.String())
// 	})

// 	// Start scraping on https://hackerspaces.org
// 	c.Visit("https://hackerspaces.org/")
// }
