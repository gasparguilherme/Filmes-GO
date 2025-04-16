package user

var idUsers int

func (Repository) GetNextUserID() int {
	idUsers++
	return idUsers
}
