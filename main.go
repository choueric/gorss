package gorss

import (
	"appengine"
	"appengine/urlfetch"
	"errors"
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", rootHandler)
	for _, k := range IDQuerys {
		http.HandleFunc("/"+k.id+"_rss", idHandler)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is gorss!\n")
	fmt.Fprint(w, "Maybe this path dose not exist!\n")
	fmt.Fprint(w, "Please go to the specificed path for rss.\n")
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	id := r.URL.Path[1 : len(r.URL.Path)-4]
	index, err := getIndexFromQuerys(id)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n")
		return
	}
	GenerateFeed(client, w, index)
}

func getIndexFromQuerys(id string) (int, error) {
	for i, v := range IDQuerys {
		if id == v.id {
			return i, nil
		}
	}
	return -1, errors.New("not find match id in QueryArray")
}
