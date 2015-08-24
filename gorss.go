package gorss

import (
	"fmt"
	"net/http"
)

var ids = map[string]func(http.ResponseWriter, *http.Request){
	"test":      testHandler,
	"zhi_japan": weixinIDHandler,
}

func init() {
	http.HandleFunc("/", rootHandler)
	for k, v := range ids {
		http.HandleFunc("/"+k+"_rss", v)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, this is gorss!\n")
	fmt.Fprint(w, "Please go to the specificed path for rss.\n")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	TestGenerateFeed(w)
}

func weixinIDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v\n", r.URL.Path)
	id := r.URL.Path[1 : len(r.URL.Path)-4]
	fmt.Fprintf(w, "ID = %s\n", id)
	GenerateFeed(w, id)
}
