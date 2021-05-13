package items

type Item struct {
	ID          string    `json:"id"`
	Tittle      string    `json:"title"`
	Pictures    []Picture `json:"pictures"`
	Text        string    `json:"text"`
	Author      string    `json:"author"`
	Status      string    `json:"status"`
	DateCreated string    `json:"status"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}
