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
	"os"
	"time"
)

type QueryElement struct {
	name   string
	cb     string
	openid string
	eqs    string
	ekv    string
	page   string
	t      string
}

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
}

func (p *DisplayXml) String() string {
	return fmt.Sprintf("  Title: %s\n  URL: %s\n  Content: %s\n  Date: %s\n",
		p.Title, p.Url, p.Content, p.Date)
}

var BaseURL = "http://weixin.sogou.com"
var queryURL = "/gzhjs?"

var zhiJapanQuery = QueryElement{
	name:   "知日",
	cb:     "sogou.weixin.gzhcb",
	openid: "oIWsFt3YfRKPuRZmMDZAdlPJgIPU",
	eqs:    "vVszo3Bguw%2BpoUyfUb7gSu7N7CSPLLzqm1DpF5tvTnfaP1JKRtX%2BIxaW3PH%2BFZuKmHrTW",
	ekv:    "3",
	page:   "1",
	t:      "1440596043703",
}

var idQueryMap = map[string]*QueryElement{
	"zhi_japan": &zhiJapanQuery,
}

func (q *QueryElement) buildURL() string {
	return fmt.Sprintf("%s%s?cb=%s&openid=%s&eqs=%s&ekv=%s&page=%s&t=%s",
		BaseURL, queryURL, q.cb, q.openid, q.eqs, q.ekv, q.page, q.t)
}

// TODO
func FetchOpenID(id string) string {
	return fmt.Sprintf("%s_openid", id)
}

func NewFeed(id string) (*feeds.Feed, error) {
	feed := &feeds.Feed{
		Title:       idQueryMap[id].name + "-公众号RSS",
		Link:        &feeds.Link{Href: "http://gorss-1047.appspot.com/"}, // TODO
		Description: "Description(TODO)",
		Author:      &feeds.Author{"Author(TODO)", "TODO@TODO.com"},
		Created:     time.Now(),
	}

	return feed, nil
}

func getSavePage(url, filename string) {
	data, err := getPage(nil, url)
	if err != nil {
		return
	}
	saveFile(data, filename)
}

func saveFile(data []byte, filename string) {
	f, err1 := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err1 != nil {
		return
	}
	defer f.Close()
	f.WriteString(string(data[:len(data)]))
}

func getPage(client *http.Client, url string) ([]byte, error) {
	var res *http.Response
	var err error
	if client == nil {
		res, err = http.Get(url)
	} else {
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

func fetchJsonBody(data []byte) []byte {
	i := bytes.IndexByte(data, '}')
	return data[5 : i+1]
}

func parseItemXml(str string) *feeds.Item {
	var entry EntryXml

	/*
		data := []byte(str)
		err := xml.Unmarshal(data, &entry)
	*/

	d := xml.NewDecoder(bytes.NewReader([]byte(str)))
	d.CharsetReader = func(s string, r io.Reader) (io.Reader, error) {
		return charset.NewReader(r, s)
	}
	err := d.Decode(&entry)

	if err != nil {
		fmt.Printf("xml entryXml unmarshal failed: %v\n", err)
		return nil
	}

	now := time.Now()
	return &feeds.Item{
		Title:       entry.Item.Display.Title,
		Link:        &feeds.Link{Href: entry.Item.Display.Url},
		Description: entry.Item.Display.Content,
		Created:     now,
	}
}

func ParsePage(data []byte) ([]*feeds.Item, error) {
	var listInfo PageJson
	data = fetchJsonBody(data)
	//err = json.NewDecoder(data).Decode(&listInfo)
	err := json.Unmarshal(data, &listInfo)
	if err != nil {
		fmt.Printf("json page unmarshal failed: %v\n", err)
		return nil, err
	}

	feedItems := make([]*feeds.Item, len(listInfo.Items))
	for i, item := range listInfo.Items {
		feedItems[i] = parseItemXml(item)
	}

	return feedItems, nil
}

func FetchList(client *http.Client, id string, feed *feeds.Feed) error {
	query := idQueryMap[id]
	url := query.buildURL()

	data, err := getPage(client, url)
	if err != nil {
		fmt.Printf("getPage failed: %v\n", err)
		return err
	}

	// For test
	saveFile(data, "zhiJapan.html")

	feed.Items, err = ParsePage(data)
	if err != nil {
		fmt.Printf("ParsePage failed: %v\n", err)
		return err
	}
	return nil
}
