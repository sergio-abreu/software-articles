package vladimir_khorikov

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
	articlesListHtml, err := getArticlesListPage()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get articles list page")
	}
	var page ArticlePage
	err = goq.Unmarshal(articlesListHtml, &page)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse article list")
	}
	articles, err := ConvertArticles(page.Articles)
	return articles, errors.Wrap(err, "failed to parse to default articles")
}

func getArticlesListPage() ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("%s%s", curation.VladimirKhorikovBlog, "/archives"))
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
