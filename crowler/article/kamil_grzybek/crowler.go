package kamil_grzybek

import (
	"astuart.co/goq"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/article"
	"github.com/sergio-vaz-abreu/software-articles/curation"
	"io/ioutil"
	"net/http"
)

func ExtractArticles() ([]article.Article, error) {
	articlesListHtml, err := getArticlesListPage(string(curation.KamilGrzybekBlog))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get articles list page")
	}
	var articleList ArticleList
	err = goq.Unmarshal(articlesListHtml, &articleList)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse article list")
	}
	var articles []article.Article
	for _, link := range articleList.Link {
		articleHtml, err := getArticlesListPage(link)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get article html")
		}
		var page ArticlePage
		err = goq.Unmarshal(articleHtml, &page)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse articles")
		}
		convertArticles, err := ConvertArticles(page.Articles)
		if err != nil {
			return nil, errors.Wrap(err, "failed to parse to default article")
		}
		articles = append(articles, convertArticles...)
	}
	return articles, nil
}

func getArticlesListPage(link string) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("%s", link))
	if err != nil {
		return nil, errors.Wrap(err, "failed to get html")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status code is not 200")
	}
	return body, nil
}
