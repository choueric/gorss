package gorss

import (
	"fmt"
	"github.com/gorilla/feeds"
	"io"
	"time"
)

func TestGenerateFeed(w io.Writer) {
	now := time.Now()
	feed := &feeds.Feed{
		Title:       "jmoiron.net blog",
		Link:        &feeds.Link{Href: "http://jmoiron.net/blog"},
		Description: "discussion about tech, footie, photos",
		Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
		Created:     now,
	}

	feed.Items = []*feeds.Item{
		&feeds.Item{
			Title:       "Limiting Concurrency in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/limiting-concurrency-in-go/"},
			Description: "A discussion on controlled parallelism in golang",
			Author:      &feeds.Author{"Jason Moiron", "jmoiron@jmoiron.net"},
			Created:     now,
		},
		&feeds.Item{
			Title:       "Logic-less Template Redux",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/logicless-template-redux/"},
			Description: "More thoughts on logicless templates",
			Created:     now,
		},
		&feeds.Item{
			Title:       "Idiomatic Code Reuse in Go",
			Link:        &feeds.Link{Href: "http://jmoiron.net/blog/idiomatic-code-reuse-in-go/"},
			Description: "How to use interfaces <em>effectively</em>",
			Created:     now,
		},
	}

	atom, err := feed.ToAtom()

	if err != nil {
		fmt.Fprintf(w, "error %v\n", err)
	} else {
		fmt.Fprint(w, atom)
	}
}

func GenerateFeed(w io.Writer, id string) {
	openid := FetchOpenID(id)
	fmt.Fprintf(w, "openid = %s\n", openid)

	feed, err := NewFeed(openid)
	if err != nil {
		fmt.Fprintf(w, "NewFeed failed: %v\n", err)
		return
	}

	err = FetchList(openid, feed)
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
