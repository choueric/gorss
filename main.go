package main

import (
	"fmt"
	"log"
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

	atom, err := feed.ToAtom()
	if err != nil {
		fmt.Fprintf(w, "error %v\n", err)
		return
	}
	fmt.Fprint(w, atom)
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/test", siteHandler)

	siteMap = make(map[string]rssSiter)
	siteMapAdd("test", &test)

	port := ":2888"
	fmt.Printf("start listen at %v ...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
