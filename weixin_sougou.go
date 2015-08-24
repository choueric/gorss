package gorss

import (
	"fmt"
	"github.com/gorilla/feeds"
	"time"
)

func FetchOpenID(id string) string {
	return fmt.Sprintf("%s_openid", id)
}

func NewFeed(openid string) (*feeds.Feed, error) {
	feed := &feeds.Feed{
		Title:       openid + " 公众号RSS",
		Link:        &feeds.Link{Href: "http://gorss-1047.appspot.com/"}, // TODO
		Description: "TODO",
		Author:      &feeds.Author{"TODO", "TODO@TODO.com"},
		Created:     time.Now(),
	}

	return feed, nil
}

func FetchList(openid string, feed *feeds.Feed) error {
	now := time.Now()
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

	return nil
}
