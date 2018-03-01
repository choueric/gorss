package main

import (
	"strconv"
	"time"

	"github.com/gorilla/feeds"
)

// const webUrl = "http://ericnode.info"
const webUrl = "127.0.0.1"

func modifyTime(ts string) time.Time {
	t, _ := strconv.ParseInt(ts, 10, 64)
	return time.Unix(t, 0)
}

func feedNewItem() *feeds.Item {
	return &feeds.Item{
		Title:       "feed item title",
		Link:        &feeds.Link{Href: "127.0.0.1/feed_item_link"},
		Description: "feed item description",
		Id:          "feed item ID",
		Author:      &feeds.Author{Name: "feed item author"},
		Created:     time.Now(),
		Updated:     time.Now(),
	}
}

func getItems(feed *feeds.Feed, title string) error {
	itemNum := 10
	feed.Items = make([]*feeds.Item, itemNum)

	for i := 0; i < 10; i++ {
		feed.Items[i] = feedNewItem()
	}
	return nil
}

func feedNew(title string) (*feeds.Feed, error) {
	feed := &feeds.Feed{
		Title:       title + "_RSS",
		Id:          "TODO",
		Link:        &feeds.Link{Href: webUrl + "/" + title},
		Description: "GoRSS feed for " + title,
		Author:      &feeds.Author{"TODO", ""},
		Updated:     time.Now(),
	}

	err := getItems(feed, title)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
