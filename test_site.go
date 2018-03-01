package main

import (
	"strconv"
	"time"

	"github.com/gorilla/feeds"
)

type testSite struct {
	feed *feeds.Feed
}

var test testSite

func (s *testSite) title() string {
	return "Test Site"
}

func (s *testSite) newFeed() (*feeds.Feed, error) {
	if s.feed == nil {
		s.feed = &feeds.Feed{
			Title:       s.title(),
			Id:          "TODO",
			Link:        &feeds.Link{Href: "http://ericnode.info/test"},
			Description: "GoRSS test feed",
			Author:      &feeds.Author{"choueric", "zhssmail@gmail.com"},
			Updated:     time.Now(),
		}
	} else {
		s.feed.Updated = time.Now()
	}

	return s.feed, nil
}

func (s *testSite) fetchItems() ([]*feeds.Item, error) {
	itemNum := 10
	items := make([]*feeds.Item, itemNum)

	for i := 0; i < 10; i++ {
		id := strconv.Itoa(i)
		items[i] = &feeds.Item{
			Title:       "feed test title " + id,
			Link:        &feeds.Link{Href: s.feed.Link.Href + "/" + id},
			Description: "feed test item description",
			Id:          "feed item ID",
			Author:      &feeds.Author{Name: "feed item author"},
			Created:     time.Now(),
			Updated:     time.Now(),
		}
	}
	return items, nil
}
