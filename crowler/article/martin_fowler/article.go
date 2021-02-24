package martin_fowler

import (
	"encoding/json"
	"fmt"
	"github.com/sergio-vaz-abreu/software-articles/curation"
	"strings"
	"time"
)

type Page struct {
	Articles []Article `goquery:"div.article-card"`
}

type Article struct {
	Description string   `goquery:"h3"`
	Author      string   `goquery:"p.credits"`
	Link        string   `goquery:"h3 a,[href]"`
	Date        string   `goquery:"p.date"`
	Tags        []string `goquery:"span.tag-link"`
}

func (a Article) MarshalJSON() ([]byte, error) {
	date, err := time.Parse("_2 Jan 2006", a.Date)
	if err != nil {
		date, err = time.Parse("Jan 2006", a.Date)
		if err != nil {
			return nil, err
		}
	}
	a.Date = date.Format("2006-01-02")
	a.Author = strings.Replace(a.Author, "by ", "", 1)
	a.Author = strings.Replace(a.Author, "with ", "", 1)
	a.Link = fmt.Sprintf("%s%s", curation.MartinFowler, a.Link)
	type marshaledArticle Article
	article := marshaledArticle(a)
	return json.Marshal(article)
}
