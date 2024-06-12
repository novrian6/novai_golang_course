package models


type article struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}


var articleList = []article{
	article{ID: 1, Title:"Article 1",Content:"Article 1 body"},
}

func GetAllArticles() []article {
	return articleList
}