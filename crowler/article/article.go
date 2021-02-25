package article

import (
	"strings"
	"time"
)

type Article struct {
	Description string    `json:"description"`
	Author      string    `json:"author"`
	Link        string    `json:"link"`
	Date        time.Time `json:"date"`
	Tags        []string  `json:"tags"`
	Site        string    `json:"site"`
}

type Articles []Article

func (a Articles) Len() int {
	return len(a)
}

func (a Articles) Less(i, j int) bool {
	return a[i].Date.Before(a[j].Date)
}

func (a Articles) Swap(i, j int) {
	aux := a[i]
	a[i] = a[j]
	a[j] = aux
}

func SanitizeTags(rawTags []string) []string {
	var tags []string
	for _, tag := range rawTags {
		tags = append(tags, strings.ToLower(tag))
	}
	return tags
}
