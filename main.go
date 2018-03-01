package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func init() {
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is gorss!\n"+
		"Maybe this path dose not exist!\n"+
		"Please go to the specificed path for rss.\n")
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	/*
		c := appengine.NewContext(r)
		client := urlfetch.Client(c)
		id := r.URL.Path[1 : len(r.URL.Path)-4]
		index, err := getIndexFromQuerys(id)
		if err != nil {
			fmt.Fprintf(w, "error: %v\n")
			return
		}
		GenerateFeed(client, w, index)
	*/
	fmt.Fprintf(w, "feed name: %s\n", r.URL.Path)
}

func getIndexFromQuerys(id string) (int, error) {
	for i, v := range IDQuerys {
		if id == v.id {
			return i, nil
		}
	}
	return -1, errors.New("not find match id in QueryArray")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/aacfree", idHandler)

	port := ":2888"
	fmt.Printf("start listen at %v ...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
