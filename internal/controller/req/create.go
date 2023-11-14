package req

type Create struct {
	Writer  int    `json:"writer_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
