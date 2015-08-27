package gorss

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/feeds"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
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

// for printf %v
func (p *DisplayXml) String() string {
	return fmt.Sprintf("  Title: %s\n  URL: %s\n  Content: %s\n  Date: %s\n",
		p.Title, p.Url, p.Content, p.Date)
}

func (q *QueryElement) buildURL() string {
	return fmt.Sprintf("%s%s?cb=%s&openid=%s&eqs=%s&ekv=%s&page=%s&t=%s",
		BaseURL, queryURL, q.cb, q.openid, q.eqs, q.ekv, q.page, q.t)
}

func NewFeed(index int) (*feeds.Feed, error) {
	q := IDQuerys[index]
	feed := &feeds.Feed{
		Title:       q.name + "-公众号RSS",
		Link:        &feeds.Link{Href: "http://gorss-1047.appspot.com/" + q.id + "_rss"},
		Description: "GoRSS feed for " + q.name,
		Author:      &feeds.Author{q.id, ""},
		Updated:     time.Now(),
	}

	return feed, nil
}

func FetchList(client *http.Client, index int) ([]*feeds.Item, error) {
	query := IDQuerys[index]
	query.t = strconv.FormatInt(time.Now().Unix(), 10)
	url := query.buildURL()

	data, err := getPage(client, url)
	if err != nil {
		fmt.Printf("getPage failed: %v\n", err)
		return nil, err
	}

	return parsePage(data)
}

func getPage(client *http.Client, url string) ([]byte, error) {
	var res *http.Response
	var err error
	if client == nil {
		// use http package's default client
		res, err = http.Get(url)
	} else {
		// use GAE's http client
		res, err = client.Get(url)
	}
	if err != nil {
		fmt.Printf("get page failed: %v\n", err)
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Printf("read page body failed: %v\n", err)
		return nil, err
	}

	return data, nil
}

func parsePage(data []byte) ([]*feeds.Item, error) {
	var page PageJson
	data = fetchJsonBody(data)
	err := json.Unmarshal(data, &page)
	if err != nil {
		fmt.Printf("json page unmarshal failed: %v\n", err)
		return nil, err
	}

	items := make([]*feeds.Item, len(page.Items))
	for i, item := range page.Items {
		items[i] = parseItemXml(item)
	}

	return items, nil
}

// skip some charactors
func fetchJsonBody(data []byte) []byte {
	i := bytes.IndexByte(data, '}')
	return data[5 : i+1]
}

func parseItemXml(str string) *feeds.Item {
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

	return &feeds.Item{
		Title:       entry.Item.Display.Title,
		Link:        &feeds.Link{Href: entry.Item.Display.Url},
		Description: entry.Item.Display.Content,
		Id:          entry.Item.Display.Docid,
		Author:      &feeds.Author{Name: entry.Item.Display.Source},
		//Created:     entry.Item.Display.Date,
		Updated: modifyTime(entry.Item.Display.Update),
	}
}

func modifyTime(ts string) time.Time {
	t, _ := strconv.ParseInt(ts, 10, 64)
	return time.Unix(t, 0)
}
