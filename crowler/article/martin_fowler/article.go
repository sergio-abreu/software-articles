package martin_fowler

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/article"
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
	date, err := time.Parse("_2 Jan 2006", customArticle.Date)
	if err != nil {
		date, err = time.Parse("Jan 2006", customArticle.Date)
		if err != nil {
			return article.Article{}, errors.Wrap(err, "failed to parse date")
		}
	}
	author := strings.Replace(customArticle.Author, "by ", "", 1)
	author = strings.Replace(author, "with ", "", 1)
	if len(author) == 0 {
		author = curation.GetCuratorName(curation.MartinFowlerBlog)
	}
	return article.Article{
		Description: customArticle.Description,
		Author:      author,
		Link:        fmt.Sprintf("%s%s", curation.MartinFowlerBlog, customArticle.Link),
		Date:        date,
		Tags:        article.SanitizeTags(customArticle.Tags),
		Site:        curation.MartinFowlerBlog,
	}, nil
}
