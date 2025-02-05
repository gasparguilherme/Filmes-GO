package entities

type Film struct {
	Title    string `json: "title"`
	Director string `json: "director"`
	Year     int    `json: "year"`
	Gender   string `json: "gender"`
	ID       int    `json: "id"`
	UserID   int    `json: "user_id"`
}
