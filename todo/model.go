package todo

type TODO struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
