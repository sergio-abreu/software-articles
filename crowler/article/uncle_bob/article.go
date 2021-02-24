package uncle_bob

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/curation"
	"time"
)

type Page struct {
	Articles []Article `goquery:"aside li"`
}

type Article struct {
	Description string `goquery:"a"`
	Author      string
	Link        string `goquery:"a,[href]"`
	Date        string `goquery:"div.tiny-date"`
	Tags        []string
}

func (a Article) MarshalJSON() ([]byte, error) {
	date, err := time.Parse("01-02-2006", a.Date)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse date")
	}
	a.Date = date.Format("2006-01-02")
	a.Author = "Robert C. Martin (Uncle Bob)"
	a.Link = fmt.Sprintf("%s%s", curation.UncleBob, a.Link)
	type marshaledArticle Article
	article := marshaledArticle(a)
	return json.Marshal(article)
}
