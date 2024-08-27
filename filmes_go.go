package main  

import "fmt"

var idAtual = 1
var idFilme = 1
var listaUsuarios []*Usuario


type Usuario struct{
	Nome string
	Email string
	ID int
	Filmes []Filme
}

type Filme struct{
	Titulo string
	Diretor string
	Ano int
	Genero string
	ID int

}

func novoUsuario(nome, email string)*Usuario{
	usuario := &Usuario{
		Nome: nome,
		Email: email,
		ID: idAtual,
	}
	idAtual++
	listaUsuarios = append(listaUsuarios,usuario )
	
	return usuario

}

func (u *Usuario) criarFilme(titulo string, diretor string, ano int, genero string){
	filme := Filme{
		Titulo: titulo,
		Diretor: diretor,
		Ano: ano,
		Genero: genero,
		ID: idFilme,
	}
	idFilme++
	u.Filmes = append(u.Filmes, filme)
		
	
}

func filmesPorUsuario(id int)[]Filme{
	for _, usuario := range listaUsuarios{
		if usuario.ID == id{
			for _, filme := range usuario.Filmes{
				fmt.Println(filme)
			}
			break
		}
	}

}




func main() {

	guilherme := novoUsuario("Guilherme Gaspar", "guilhermegaspar@gmail.com")
	fmt.Println(guilherme)


	
	joao := novoUsuario("Joao Saraceni", "joaosaraceni@gmail.com")
	fmt.Println(joao)
	
	for _, usuario := range listaUsuarios{
		fmt.Println(usuario.Nome)
	}
	

	guilherme.criarFilme("Titanic", "James Cameron", 1997, "Drama/Romance")
	fmt.Println(guilherme)

	

	joao.criarFilme("Interestellar", "Christopher Nolan", 2014, "Ficção Cientifica")
	fmt.Println(joao)

	joao.criarFilme("Bad Boy", "Jordy Nichols",2007, "Ação")



	for _, usuario := range listaUsuarios{
		fmt.Printf("Filmes de %s: \n", usuario.Nome)
		for _, filme := range usuario.Filmes{
			fmt.Println(filme.Titulo)
		}
	}

	
	filmesPorUsuario(2)




}