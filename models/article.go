package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var Articles = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func GetAllArtciles() []Article {
	return Articles
}

func GetArticleById(id int) (*Article, error) {
	for _, a := range Articles {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
