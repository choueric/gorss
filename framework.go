package main

import (
	"errors"

	"github.com/gorilla/feeds"
)

type rssSiter interface {
	title() string
	newFeed() (*feeds.Feed, error)
	fetchItems() ([]*feeds.Item, error)
}

type siteMapT map[string]rssSiter

var siteMap siteMapT

func siteMapAdd(title string, site rssSiter) {
	siteMap[title] = site
}

func siteMapFind(title string) (rssSiter, error) {
	site, ok := siteMap[title]
	if !ok {
		return nil, errors.New("Site does not exist in map")
	}
	return site, nil
}
