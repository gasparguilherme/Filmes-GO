package login

import (
	"fmt"

	"github.com/gasparguilherme/my-repository/domain/entities"
	"golang.org/x/crypto/bcrypt"
)

func (u Usecase) Login(data Login) (*entities.User, error) {
	user, err := u.repository.FindUserByEmail(data.Email)
	if err != nil {
		return nil, fmt.Errorf("usuário não encontrado")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)) != nil {
		return nil, fmt.Errorf("senha incorreta")
	}

	return user, nil
}
