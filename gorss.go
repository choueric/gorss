package gorss

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"net/http"
)

var ids = []string{
	"zhi_japan",
}

func init() {
	http.HandleFunc("/", rootHandler)

	for _, k := range ids {
		http.HandleFunc("/"+k+"_rss", idHandler)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is gorss!\n")
	fmt.Fprint(w, "Please go to the specificed path for rss.\n")
}

func weixinIDHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1 : len(r.URL.Path)-4]
	GenerateFeed(nil, w, id)
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	client := urlfetch.Client(c)
	id := r.URL.Path[1 : len(r.URL.Path)-4]
	GenerateFeed(client, w, id)
}
