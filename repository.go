package main  
import "fmt"

var idAtual = 1

type Usuario struct{
	Nome string
	Email string
	ID int
}

func novoUsuario(nome, email string)Usuario{
	usuario := Usuario{
		Nome: nome,
		Email: email,
		ID: idAtual,
	}
	idAtual++
	return usuario
}


func main() {

	guilherme := novoUsuario("Guilherme", "guilhermegaspar@gmail.com")
	fmt.Println(guilherme)

	joao := novoUsuario("Joao", "joaosaraceni@gmail,com")
	fmt.Println(joao)


}