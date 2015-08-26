package gorss

import (
	"fmt"
	//"github.com/gorilla/feeds"
	"io"
)

func GenerateFeed(w io.Writer, id string) {
	feed, err := NewFeed(id)
	if err != nil {
		fmt.Fprintf(w, "NewFeed failed: %v\n", err)
		return
	}

	err = FetchList(id, feed)
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
