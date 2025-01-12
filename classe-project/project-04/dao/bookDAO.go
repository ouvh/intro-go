package dao

type BookDAO struct {
	ID       int64    `json:"id"`
	Title    string   `json:"title"`
	AuthorId int64    `json:"author_id"`
	Genres   []string `json:"genres"`
	Price    float64  `json:"price"`
	Stock    int      `json:"stock"`
}
