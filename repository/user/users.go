package repository

var idUsers int

func GetNextUserID() int {
	idUsers++
	return idUsers
}