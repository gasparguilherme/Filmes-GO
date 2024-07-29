package main  
import "fmt"

var idAtual = 1
var listaUsuarios []Usuario
var listaFilmes []Filme


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

func criarFilme(titulo string, diretor string, ano int, genero string)Filme{
	filmes := Filme{
		Titulo: titulo,
		Diretor: diretor,
		Ano: ano,
		Genero: genero,
	}
	listaFilmes = append(listaFilmes, filmes)
	return filmes
	
}


func main() {

	guilherme := novoUsuario("Guilherme Gaspar", "guilhermegaspar@gmail.com")
	fmt.Println(guilherme)
	
	joao := novoUsuario("Joao Saraceni", "joaosaraceni@gmail.com")
	fmt.Println(joao)
	
	for _, usuario := range listaUsuarios{
		fmt.Println(usuario.Nome)
	}
	


	titanic := criarFilme("Titanic", "James Cameron", 1997, "Drama/Romance")
	fmt.Println(titanic)

	interestellar := criarFilme("Interestellar", "Christopher Nolan", 2014, "Ficção Cientifica")
	fmt.Println(interestellar)

	for _, filmes := range listaFilmes{
		fmt.Println(filmes.Titulo)
	}


}