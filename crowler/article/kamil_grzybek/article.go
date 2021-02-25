package kamil_grzybek

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/article"
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
	date, err := time.Parse(time.RFC3339, customArticle.Date)
	if err != nil {
		return article.Article{}, errors.Wrap(err, "failed to parse date")
	}
	return article.Article{
		Description: customArticle.Description,
		Author:      curation.GetCuratorName(curation.KamilGrzybekBlog),
		Link:        fmt.Sprintf("%s%s", curation.KamilGrzybekBlog, customArticle.Link),
		Date:        date,
		Tags:        article.SanitizeTags(customArticle.Tags),
		Site:        curation.KamilGrzybekBlog,
	}, nil
}
