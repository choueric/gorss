package main

import (
	"fmt"
	"io"
	"net/http"
)

func GenerateFeed(client *http.Client, w io.Writer, index int) {
	feed, err := NewFeed(index)
	if err != nil {
		fmt.Fprintf(w, "NewFeed failed: %v\n", err)
		return
	}

	feed.Items, err = FetchList(client, index)
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
