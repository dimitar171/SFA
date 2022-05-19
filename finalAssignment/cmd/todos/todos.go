package todos

// List is a struct containing list data
type List struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	UserId int    `json:"userId"`
}

// Task is a struct containing Task data
type Task struct {
	Id        int    `json:"id"`
	Text      string `json:"text"`
	ListId    int    `json:"listId"`
	Completed bool   `json:"completed"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Pass string `json:"pass"`
}
