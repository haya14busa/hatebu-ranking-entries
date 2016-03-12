package parser

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"

	"golang.org/x/net/html"

	"github.com/PuerkitoBio/goquery"
)

const (
	BASE_HATEBU_URL = "http://b.hatena.ne.jp/"
)

func main() {
	fmt.Println("start")

	es, err := NewEntries("http://b.hatena.ne.jp/ranking/daily/")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(es))

	es, err = NewEntries("http://b.hatena.ne.jp/ranking/monthly/201602/it")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(len(es))
	for _, e := range es {
		fmt.Println(e.Title)
	}

}

type Entry struct {
	Eid            string
	Rank           int
	Bookmarkcount  int
	Title          string
	Url            string
	HatebuEntryUrl string
	Date           string
	Category       string
	Tags           []string
	Description    string
}

// NewEntries returns entries from hatena bookmark ranking pages.
// It supports daily, weekly and monthly ranking pages for each categories.
// Example:
// - http://b.hatena.ne.jp/ranking/daily/
// - http://b.hatena.ne.jp/ranking/daily/20160311/it
// - http://b.hatena.ne.jp/ranking/weekly/20160229/it
// - http://b.hatena.ne.jp/ranking/monthly/201602/it
func NewEntries(url string) ([]*Entry, error) {
	html, err := getHtml(url)
	if err != nil {
		return nil, err
	}
	return entries(goquery.NewDocumentFromNode(html))
}

func getHtml(url string) (*html.Node, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.75 Safari/537.36")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	node, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func entries(doc *goquery.Document) ([]*Entry, error) {
	selector := ".entry-list-l .entrylist-unit"
	var es []*Entry
	var errors []string
	doc.Find(selector).Each(func(_ int, s *goquery.Selection) {
		entry, err := newEntryFromUnit(s)
		if err != nil {
			errors = append(errors, err.Error())
		} else {
			es = append(es, entry)
		}
	})
	if len(errors) > 0 {
		return nil, fmt.Errorf("Unable to create entries: %v", strings.Join(errors, "\n"))
	}
	if len(es) == 0 {
		return nil, fmt.Errorf("Unable to find entries from document: url: %v", doc.Url)
	}
	return es, nil
}

// newEntryFromUnit creates Entry from <li class="entrylist-unit">
func newEntryFromUnit(s *goquery.Selection) (*Entry, error) {
	eid, err := unitDataAttr(s, "data-eid")
	if err != nil {
		return nil, err
	}

	entryrankStr, err := unitDataAttr(s, "data-entryrank")
	entryrank, err := strconv.Atoi(entryrankStr)
	if err != nil {
		return nil, err
	}

	bookmarkcountStr, err := unitDataAttr(s, "data-bookmark-count")
	bookmarkcount, err := strconv.Atoi(bookmarkcountStr)
	if err != nil {
		return nil, err
	}

	entrylink := s.Find("a.entry-link").First()
	title, _ := entrylink.Attr("title")
	url, _ := entrylink.Attr("href")

	entrydata := s.Find(".entry-data").First()

	hatebuEntryUrlRel, _ := entrydata.Find(".users a").First().Attr("href")
	hatebuEntryUrl := path.Join(BASE_HATEBU_URL, hatebuEntryUrlRel)

	date := entrydata.Find(".date").First().Text()

	category := entrydata.Find(".category a").First().Text()

	var tags []string
	entrydata.Find(".tag a").Each(func(_ int, s *goquery.Selection) {
		tags = append(tags, s.Text())
	})

	// optional
	description := entrydata.Find(".description blockquote").First().Text()

	return &Entry{
		Eid:            eid,
		Rank:           entryrank,
		Bookmarkcount:  bookmarkcount,
		Title:          title,
		Url:            url,
		HatebuEntryUrl: hatebuEntryUrl,
		Date:           date,
		Category:       category,
		Tags:           tags,
		Description:    description,
	}, nil
}

// unitDataAttr returns attr value or descriptive error
func unitDataAttr(s *goquery.Selection, attr string) (string, error) {
	data, ok := s.Attr(attr)
	if !ok {
		html, err := s.Html()
		if err != nil {
			return "", fmt.Errorf("Unable to find %s from %v", attr, s.Text())
		}
		return "", fmt.Errorf("Unable to find %s from %v", attr, html)
	}
	return data, nil
}
