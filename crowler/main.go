package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sergio-vaz-abreu/software-articles/article"
	"github.com/sergio-vaz-abreu/software-articles/article/kamil_grzybek"
	"github.com/sergio-vaz-abreu/software-articles/article/martin_fowler"
	"github.com/sergio-vaz-abreu/software-articles/article/uncle_bob"
	"github.com/sergio-vaz-abreu/software-articles/article/vladimir_khorikov"
	"github.com/sergio-vaz-abreu/software-articles/curation"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	action := os.Args[1]
	var err error
	switch action {
	case "extract":
		err = extractArticles()
	case "markdown":
		err = createMarkdown()
	default:
		err = errors.New("not implemented action")
	}
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func extractArticles() error {
	martinFowlerArticles, err := martin_fowler.ExtractArticles()
	if err != nil {
		return errors.Wrap(err, "failed to extract martin fowler articles")
	}
	uncleBobArticles, err := uncle_bob.ExtractArticles()
	if err != nil {
		return errors.Wrap(err, "failed to extract uncle bob articles")
	}
	kamilGrzybekArticles, err := kamil_grzybek.ExtractArticles()
	if err != nil {
		return errors.Wrap(err, "failed to extract kamil grzybek articles")
	}
	vladimirKhorikovArticles, err := vladimir_khorikov.ExtractArticles()
	if err != nil {
		return errors.Wrap(err, "failed to extract vladimir khorikov articles")
	}
	articles := mergeArticles(martinFowlerArticles, uncleBobArticles, kamilGrzybekArticles, vladimirKhorikovArticles)
	data, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to encode articles")
	}
	fmt.Println(string(data))
	fmt.Println("uncle bob articles", len(uncleBobArticles))
	fmt.Println("martin fowler articles", len(martinFowlerArticles))
	fmt.Println("kamil grzybek articles", len(kamilGrzybekArticles))
	fmt.Println("vladimir khorikov articles", len(vladimirKhorikovArticles))
	fmt.Println("total articles", len(articles))
	err = ioutil.WriteFile("./../articles.json", data, os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to persist articles")
	}
	return nil
}

func createMarkdown() error {
	rawData, err := ioutil.ReadFile("./../articles.json")
	if err != nil {
		return errors.Wrap(err, "failed to retrieve articles")
	}
	var articles article.Articles
	err = json.Unmarshal(rawData, &articles)
	if err != nil {
		return errors.Wrap(err, "failed to decode articles")
	}
	data := markdownBySite(articles)
	err = ioutil.WriteFile("./../groupedByBlog.md", []byte(data), os.ModePerm)
	if err != nil {
		return errors.Wrap(err, "failed to create markdown file")
	}
	return nil
}

func markdownBySite(articles article.Articles) string {
	groupedArticles := groupArticlesByBlog(articles)
	markdown := "### Content By Blog\n\n"
	var authors []string
	for blog := range groupedArticles {
		authors = append(authors, curation.GetCuratorName(blog))
	}
	sort.Strings(authors)
	var postMarkdown string
	for _, author := range authors {
		siteArticles := groupedArticles[curation.GetBlog(author)]
		markdown += fmt.Sprintf("- [%s](#%s)\n", author, strings.ToLower(strings.ReplaceAll(author, " ", "-")))
		postMarkdown += fmt.Sprintf("## [%s](%s)\n", author, curation.GetBlog(author))
		for _, a := range siteArticles {
			postMarkdown += fmt.Sprintf("* %s - [%s](%s) [%s]\n", a.Date.Format("02 Jan 06"), a.Description, a.Link, strings.Join(a.Tags, ", "))
		}
	}
	markdown += "\n" + postMarkdown
	return markdown
}

func groupArticlesByBlog(articles article.Articles) map[string]article.Articles {
	group := map[string]article.Articles{}
	for _, a := range articles {
		group[a.Site] = append(group[a.Site], a)
	}
	return group
}

func mergeArticles(articles ...[]article.Article) article.Articles {
	var mergedArticles article.Articles
	for _, a := range articles {
		mergedArticles = append(mergedArticles, a...)
	}
	sort.Sort(mergedArticles)
	return mergedArticles
}
