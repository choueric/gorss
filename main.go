package main

import (
	"fmt"
	"net/http"

	"github.com/choueric/clog"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is gorss!\n"+
		"Maybe this path dose not exist!\n"+
		"Please go to the right path for the target site.\n")
}

func siteHandler(w http.ResponseWriter, r *http.Request) {
	siteName := r.URL.Path[1:]
	clog.Printf("Site name: %s\n", siteName)

	site, err := siteMapFind(siteName)
	if err != nil {
		fmt.Fprintf(w, "error %v\n", err)
		return
	}

	feed, err := site.newFeed()
	if err != nil {
		fmt.Fprintf(w, "error %v\n", err)
		return
	}

	feed.Items, err = site.fetchItems()
	if err != nil {
		fmt.Fprintf(w, "error %v\n", err)
		return
	}

	page, err := feed.ToRss()
	if err != nil {
		fmt.Fprintf(w, "error %v\n", err)
		return
	}
	fmt.Fprint(w, page)
}

func main() {
	const (
		testUrlPath = "test.xml"
		aacUrlPath  = "aacfree.xml"
	)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/"+testUrlPath, siteHandler)
	http.HandleFunc("/"+aacUrlPath, siteHandler)

	siteMap = make(map[string]rssSiter)
	siteMapAdd(testUrlPath, &test)
	siteMapAdd(aacUrlPath, &aacfree)

	port := ":2888"
	clog.Printf("start listen at %v ...\n", port)
	clog.Fatal(http.ListenAndServe(port, nil))
}
