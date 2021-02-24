package martin_fowler

import (
	"astuart.co/goq"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/curation"
	"io/ioutil"
	"net/http"
	"sync"
)

func ExtractArticles() ([]Article, error) {
	articlesCh := make(chan []Article, 100)
	errCh := make(chan error)
	var wg sync.WaitGroup
	for year := 1996; year <= 2021; year++ {
		year := year
		wg.Add(1)
		go func() {
			defer wg.Done()
			htmlPage, err := getArticlesListPage(year)
			if err != nil {
				errCh <- errors.Wrap(err, "failed to get html page")
				return
			}
			var page Page
			err = goq.Unmarshal(htmlPage, &page)
			if err != nil {
				errCh <- errors.Wrap(err, "failed to extract articles from html")
				return
			}
			articlesCh <- page.Articles
		}()
	}
	wg.Wait()
	close(articlesCh)
	close(errCh)
	var articles []Article
	for article := range articlesCh {
		articles = append(articles, article...)
	}
	return articles, nil
}

func ExtractArticles2() ([]Article, error) {
	var articles []Article
	for year := 1996; year <= 2021; year++ {
		htmlPage, err := getArticlesListPage(year)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get html page")
		}
		var page Page
		err = goq.Unmarshal(htmlPage, &page)
		if err != nil {
			return nil, errors.Wrap(err, "failed to extract articles from html")
		}
		articles = append(articles, page.Articles...)
	}
	return articles, nil
}

func getArticlesListPage(year int) ([]byte, error) {
	res, err := http.Get(fmt.Sprintf("%s/tags/%d.html", curation.MartinFowler, year))
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
