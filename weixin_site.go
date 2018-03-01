package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/feeds"
	"golang.org/x/net/html/charset"
)

/*
 * the query page of weixin_sougou is a json data, which includes
 * a xml data containg these entries for each essays.
 * So, PageJson's Items are xml data for each entry, which is presented by
 * EntryXml.
 * In EntryXml, we just fetch these elements we need. These elements presented
 * by DisplayXml are in the <Display> tag.
 */

const BaseURL = "http://weixin.sogou.com"
const queryURL = "/gzhjs?"

type PageJson struct {
	TotalItems int      `json:"totalItems"`
	TotalPages int      `json:"totalPages"`
	Page       int      `json:"page"`
	Items      []string `json:"items"`
}

type EntryXml struct {
	Item ItemXml `xml:"item"`
}

type ItemXml struct {
	Display DisplayXml `xml:"display"`
}

type DisplayXml struct {
	Title   string `xml:"title"`
	Url     string `xml:"url"`
	Content string `xml:"content"`
	Date    string `xml:"date"`
	Docid   string `xml:"docid"`
	Source  string `xml:"sourcename"`
	Update  string `xml:"lastModified"`
}

// RSS query info for a gzh. gzh is the "public account"
type QueryElement struct {
	id     string
	name   string
	openid string
	eqs    string
	cb     string // fixed
	ekv    string // fixed
	page   string // fixed
	t      string
}

// for printf %v
func (p *DisplayXml) String() string {
	return fmt.Sprintf("  Title: %s\n  URL: %s\n  Content: %s\n  Date: %s\n",
		p.Title, p.Url, p.Content, p.Date)
}

func (q *QueryElement) buildURL() string {
	return fmt.Sprintf("%s%s?cb=%s&openid=%s&eqs=%s&ekv=%s&page=%s&t=%s",
		BaseURL, queryURL, q.cb, q.openid, q.eqs, q.ekv, q.page, q.t)
}

/* parse page data from URL */
func parsePage(client *http.Client, data []byte) ([]*feeds.Item, error) {
	var page PageJson

	data = fetchJsonBody(data)
	err := json.Unmarshal(data, &page)
	if err != nil {
		fmt.Printf("json page unmarshal failed: %v\n", err)
		return nil, err
	}

	items := make([]*feeds.Item, len(page.Items))
	for i, item := range page.Items {
		items[i] = parseItemXml(client, item)
	}

	return items, nil
}

// skip some charactors
func fetchJsonBody(data []byte) []byte {
	i := bytes.IndexByte(data, '}')
	return data[19 : i+1]
}

func fetchFeedUrl(client *http.Client, requestUrl string) (string, error) {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("fetchFeedUrl failed: %v\n", err)
		return "error url", err
	}

	return res.Request.URL.String(), nil
}

func parseItemXml(client *http.Client, str string) *feeds.Item {
	var entry EntryXml

	// change from gbk to utf8
	d := xml.NewDecoder(bytes.NewReader([]byte(str)))
	d.CharsetReader = func(s string, r io.Reader) (io.Reader, error) {
		return charset.NewReader(r, s)
	}
	err := d.Decode(&entry)
	if err != nil {
		fmt.Printf("xml entryXml unmarshal failed: %v\n", err)
		return nil
	}

	url, err := fetchFeedUrl(client, BaseURL+entry.Item.Display.Url)
	if err != nil {
		return nil
	}

	return &feeds.Item{
		Title:       entry.Item.Display.Title,
		Link:        &feeds.Link{Href: url},
		Description: entry.Item.Display.Content,
		Id:          entry.Item.Display.Docid,
		Author:      &feeds.Author{Name: entry.Item.Display.Source},
		Updated:     time.Now(),
		//Created:     entry.Item.Display.Date,
	}
}
