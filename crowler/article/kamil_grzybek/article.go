package kamil_grzybek

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/curation"
	"time"
)

type ArticleList struct {
	Link []string `goquery:"aside#archives-2 li a,[href]"`
}

type ArticlePage struct {
	Articles []Article `goquery:"article"`
}

type Article struct {
	Description string `goquery:"h2 a"`
	Author      string
	Link        string   `goquery:"h2 a,[href]"`
	Date        string   `goquery:"time.entry-date,[datetime]"`
	Tags        []string `goquery:"span.tags-links a"`
	Folder      string   `goquery:"span.cat-links a" json:"-"`
}

func (a Article) MarshalJSON() ([]byte, error) {
	date, err := time.Parse(time.RFC3339, a.Date)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse date")
	}
	a.Date = date.Format("2006-01-02")
	a.Author = "Kamil Grzybek"
	a.Link = fmt.Sprintf("%s%s", curation.KamilGrzybek, a.Link)
	a.Tags = append(a.Tags, a.Folder)
	type marshaledArticle Article
	article := marshaledArticle(a)
	return json.Marshal(article)
}
