package main

import (
	"encoding/json"
	"fmt"
	"github.com/andrewstueart/goq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Year     int       `goquery:"div.intro h1 i"`
	Articles []Article `goquery:"div.article-card"`
}

type Article struct {
	Name   string   `goquery:"h3"`
	Author string   `goquery:"p.credits"`
	Link   string   `goquery:"h3 a,[href]"`
	Date   string   `goquery:"p.date"`
	Tags   []string `goquery:"span.tag-link"`
}

func (a Article) ToMap(baseUrl string) map[string]interface{} {
	return map[string]interface{}{
		"Description": a.Name,
		"Author":      a.Author,
		"Link":        fmt.Sprintf("%s%s", baseUrl, a.Link),
		"Date":        a.Date,
		"Tags":        a.Tags,
	}
}

func main() {
	articles := map[int][]interface{}{}
	baseUrl := "https://martinfowler.com"
	for year := 1997; year <= 2021; year++ {
		res, err := http.Get(fmt.Sprintf("%s/tags/%d.html", baseUrl, year))
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		var ex Page

		err = goq.NewDecoder(res.Body).Decode(&ex)
		if err != nil {
			log.Fatal(err)
		}
		var reversedArticles []interface{}
		for i := len(ex.Articles) - 1; i >= 0; i-- {
			reversedArticles = append(reversedArticles, ex.Articles[i].ToMap(baseUrl))
		}
		articles[year] = reversedArticles
	}
	data, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = ioutil.WriteFile("martin-fowler-articles.json", data, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
}
