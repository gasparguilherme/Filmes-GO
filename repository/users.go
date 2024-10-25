package repository

var id_users int

func GetNextID_users() int {
	id_users++
	return id_users
}
