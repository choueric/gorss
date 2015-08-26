package gorss

import (
	"fmt"
	"net/http"
)

/*
var ids = map[string]func(http.ResponseWriter, *http.Request){
	"zhi_japan": weixinIDHandler,
}
*/

var ids = []string{
	"zhi_japan",
}

func init() {
	http.HandleFunc("/", rootHandler)

	/*
		for k, v := range ids {
			http.HandleFunc("/"+k+"_rss", v)
		}
	*/

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
	GenerateFeed(w, id)
}

func idHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1 : len(r.URL.Path)-4]
	GenerateFeed(w, id)
}
