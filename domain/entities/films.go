package entities

type Film struct {
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
	Genre    string `json:"genre"`
	ID       int    `json:"id"`
	UserID   int    `json:"userID"`
}
