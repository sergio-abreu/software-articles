package vladimir_khorikov

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/curation"
	"time"
)

type ArticlePage struct {
	Articles []Article `goquery:"div.container div.col-md-12 div.row div.postIndexItem"`
}

type Article struct {
	Description string `goquery:"div.title a"`
	Author      string
	Link        string `goquery:"div.title a,[href]"`
	Date        string `goquery:"div.date"`
	Tags        []string
}

func (a Article) MarshalJSON() ([]byte, error) {
	date, err := time.Parse("02 Jan 2006", a.Date)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse date")
	}
	a.Date = date.Format("2006-01-02")
	a.Author = "Vladimir Khorikov"
	a.Link = fmt.Sprintf("%s%s", curation.VladimirKhorikov, a.Link)
	type marshaledArticle Article
	article := marshaledArticle(a)
	return json.Marshal(article)
}
