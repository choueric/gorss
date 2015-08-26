package gorss

import (
	"fmt"
	//"github.com/gorilla/feeds"
	"io"
	"net/http"
)

func GenerateFeed(client *http.Client, w io.Writer, id string) {
	feed, err := NewFeed(id)
	if err != nil {
		fmt.Fprintf(w, "NewFeed failed: %v\n", err)
		return
	}

	err = FetchList(client, id, feed)
	if err != nil {
		fmt.Fprintf(w, "FetchList failed: %v\n", err)
		return
	}

	atom, err := feed.ToAtom()
	if err != nil {
		fmt.Fprintf(w, "error %v\n", err)
		return
	}
	fmt.Fprint(w, atom)
}
