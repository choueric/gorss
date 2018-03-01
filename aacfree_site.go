package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gorilla/feeds"
	"github.com/headzoo/surf"
	"github.com/headzoo/surf/agent"
	"github.com/headzoo/surf/browser"
)

type aacfreeSite struct {
	feed *feeds.Feed
}

func saveToFile(bow *browser.Browser, filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = bow.Download(w)
	if err != nil {
		return err
	}

	return nil
}

var aacfree aacfreeSite

func (s *aacfreeSite) title() string {
	return "iTunes分享店"
}

func (s *aacfreeSite) newFeed() (*feeds.Feed, error) {
	if s.feed == nil {
		s.feed = &feeds.Feed{
			Title:       s.title(),
			Link:        &feeds.Link{Href: "https://www.aacfree.com"},
			Description: s.title(),
			Author:      &feeds.Author{"aacfree", "no email"},
			Updated:     time.Now(),
			Created:     time.Now(),
		}
	} else {
		s.feed.Updated = time.Now()
	}

	return s.feed, nil
}

func (s *aacfreeSite) fetchItems() ([]*feeds.Item, error) {
	bow := surf.NewBrowser()
	bow.SetUserAgent(agent.Firefox())
	err := bow.Open("https://www.aacfree.com")
	if err != nil {
		return nil, err
	}

	items := make([]*feeds.Item, 0)
	sel := bow.Find("article")
	sel.Each(func(_ int, s *goquery.Selection) {
		item := &feeds.Item{
			Author:      &feeds.Author{Name: "aacfree"},
			Description: "aacfree",
		}

		if id, ok := s.Attr("id"); ok {
			item.Id = id
		}

		if a := s.Find("h2").Find("a"); a != nil {
			if href, ok := a.Attr("href"); ok {
				item.Link = &feeds.Link{Href: href}
			}
		}

		if img := s.Find("img"); img != nil {
			if alt, ok := img.Attr("alt"); ok {
				item.Title = alt
			}
		}

		t := s.Find("footer").Find("time")
		t.Each(func(_ int, s *goquery.Selection) {
			if class, ok := s.Attr("class"); ok {
				if strings.Contains(class, "updated") {
					if r, ok := s.Attr("datetime"); ok {
						v, err := time.Parse(time.RFC3339, r)
						if err == nil {
							item.Created = v
							item.Updated = v
						}
					}
				}
			}
		})
		items = append(items, item)
	})

	return items, nil
}
