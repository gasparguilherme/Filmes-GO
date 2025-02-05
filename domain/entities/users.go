package entities

type User struct {
	Name  string `json: "name"`
	Email string `json: "email"`
	ID    int    `json: "id"`
}
