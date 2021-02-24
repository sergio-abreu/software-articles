package main

import (
	"encoding/json"
	"fmt"
	"github.com/sergio-vaz-abreu/software-articles/article/kamil_grzybek"
	"github.com/sergio-vaz-abreu/software-articles/article/martin_fowler"
	"github.com/sergio-vaz-abreu/software-articles/article/uncle_bob"
	"github.com/sergio-vaz-abreu/software-articles/article/vladimir_khorikov"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	martinFowlerArticles, err := martin_fowler.ExtractArticles()
	if err != nil {
		log.Fatalf("failed to extract martin fowler articles: %s", err.Error())
		return
	}
	uncleBobArticles, err := uncle_bob.ExtractArticles()
	if err != nil {
		log.Fatalf("failed to extract martin fowler articles: %s", err.Error())
		return
	}
	kamilGrzybekArticles, err := kamil_grzybek.ExtractArticles()
	if err != nil {
		log.Fatalf("failed to extract kamil grzybek articles: %s", err.Error())
		return
	}
	vladimirKhorikovArticles, err := vladimir_khorikov.ExtractArticles()
	if err != nil {
		log.Fatalf("failed to extract vladimir khorikov articles: %s", err.Error())
		return
	}
	articles := mergeArticles(martinFowlerArticles, uncleBobArticles, kamilGrzybekArticles, vladimirKhorikovArticles)
	data, err := json.MarshalIndent(articles, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(data))
	fmt.Println("uncle bob articles", len(uncleBobArticles))
	fmt.Println("martin fowler articles", len(martinFowlerArticles))
	fmt.Println("kamil grzybek articles", len(kamilGrzybekArticles))
	fmt.Println("vladimir khorikov articles", len(vladimirKhorikovArticles))
	fmt.Println("total articles", len(articles))
	err = ioutil.WriteFile("./../articles.json", data, os.ModePerm)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func mergeArticles(martinFowlerArticles []martin_fowler.Article, uncleBobArticles []uncle_bob.Article, kamilGrzybekArticles []kamil_grzybek.Article, vladimirKhorikovArticles []vladimir_khorikov.Article) []interface{} {
	var data []interface{}
	for _, article := range martinFowlerArticles {
		data = append(data, article)
	}
	for _, article := range uncleBobArticles {
		data = append(data, article)
	}
	for _, article := range kamilGrzybekArticles {
		data = append(data, article)
	}
	for _, article := range vladimirKhorikovArticles {
		data = append(data, article)
	}
	return data
}
