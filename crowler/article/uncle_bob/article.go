package uncle_bob

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/article"
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

func ConvertArticles(customArticles []Article) ([]article.Article, error) {
	var articles []article.Article
	for _, customArticle := range customArticles {
		a, err := ToArticle(customArticle)
		if err != nil {
			return nil, err
		}
		articles = append(articles, a)
	}
	return articles, nil
}

func ToArticle(customArticle Article) (article.Article, error) {
	date, err := time.Parse("01-02-2006", customArticle.Date)
	if err != nil {
		return article.Article{}, errors.Wrap(err, "failed to parse date")
	}
	return article.Article{
		Description: customArticle.Description,
		Author:      curation.GetCuratorName(curation.UncleBobBlog),
		Link:        fmt.Sprintf("%s%s", curation.UncleBobBlog, customArticle.Link),
		Date:        date,
		Tags:        article.SanitizeTags(customArticle.Tags),
		Site:        curation.UncleBobBlog,
	}, nil
}
