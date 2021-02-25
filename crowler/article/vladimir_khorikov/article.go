package vladimir_khorikov

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/article"
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
	date, err := time.Parse("02 Jan 2006", customArticle.Date)
	if err != nil {
		return article.Article{}, errors.Wrap(err, "failed to parse date")
	}
	return article.Article{
		Description: customArticle.Description,
		Author:      curation.GetCuratorName(curation.VladimirKhorikovBlog),
		Link:        fmt.Sprintf("%s%s", curation.VladimirKhorikovBlog, customArticle.Link),
		Date:        date,
		Tags:        customArticle.Tags,
		Site:        curation.VladimirKhorikovBlog,
	}, nil
}
