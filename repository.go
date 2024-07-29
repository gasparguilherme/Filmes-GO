package main  
import "fmt"

var idAtual = 1
var listaUsuarios []Usuario


type Usuario struct{
	Nome string
	Email string
	ID int
}

type Filme struct{
	Titulo string
	Diretor string
	Ano int
	Genero string

}

func novoUsuario(nome, email string)Usuario{
	usuario := Usuario{
		Nome: nome,
		Email: email,
		ID: idAtual,
	}
	idAtual++
	listaUsuarios = append(listaUsuarios,usuario )
	return usuario

}


func main() {

	guilherme := novoUsuario("Guilherme", "guilhermegaspar@gmail.com")
	fmt.Println(guilherme)
	
	joao := novoUsuario("Joao", "joaosaraceni@gmail.com")
	fmt.Println(joao)

	fmt.Println(listaUsuarios)


}