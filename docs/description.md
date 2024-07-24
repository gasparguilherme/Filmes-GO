# O projeto

Este repositório contém um projeto de estudo para entender o funcionamento - e criar do zero - uma API de "filmes favoritos".

Esta documentação tem como objetivo dar um panorama dos próximos passos, e indicações de alguns materiais de estudo. Vamos encorpando esta documentação conforme formos avançando.

Durante o desenvolvimento do projeto, espera-se aprofundar conhecimentos em:

* Git
* Go
* CRUD e API RESTful
* Code Review
* Logs
* Boas práticas de desenvolvimento de software
* Kanban
* Arquitetura de software (clean arch)
* Bancos da dados relacionais (postgres)
* Cache (redis)
* Documentação de API (OpenAPI)
* Conteinerização (Docker)
* Teste unitários em Go
* CI/CD

## Workflow

Simularemos um fluxo de trabalho aplicando [metodologias ágeis](https://www.totvs.com/blog/negocios/metodologia-agil/). Para isso, faremos um [Kanban](https://www.totvs.com/blog/negocios/kanban/) simplificado, utilizando o Github Projects. Teremos cinco colunas.

### Colunas Kanban

1. **Backlog**: Aqui nascerão as tarefas, talvez sem muito detalhamento no primeiro momento.
2. **Ready**: As tarefas que estão prontas para serem atacadas.
3. **In Progress**: O que estamos trabalhando no momento.
4. **In Review**: Depois que o desenvolvimento é concluído, tarefas que estão em *code review*, ou seja, com *pull request (PR)* aberto.
5. **Done**: Tarefa concluída.

Colunas novas poderão ser adicionadas ao longo do projeto (como Blocked, Waiting Validation, Icebox, etc...)

### Transições entre colunas Kanban

1. **Refinamento** (Backlog -> Ready): Refinamento é o ato de detalhar uma tarefa, elencando seus *requisitos*, trazendo o *contexto* que gera a necessidade de que ela seja feita, e o *resultado* esperado ao final do desenvolvimento. Exemplo: "Precisamos implementar um id sequencial nos registros de usuários. Esse id irá servir para identificar a quem um filme favorito pertence, sendo um identificador único e imutável. Espera-se que ao final dessa tarefa o sistema crie um id único e autoincrementável a cada novo usuário gerado."
2. **"Puxar" uma tarefa** (Ready -> In Progress): Sem muito mistério, aqui é pegar para desenvolver a funcionalidade descrita. **Antes de começar a desenvolver a tarefa, abra uma nova branch no git**

```shell
git checkout -b id-autoincrementavel
```

Este comando irá separar o seu desenvolvimento do *código estável*, que está na main.

3. **Abrir um Pull Request** (In Progress -> In Review): Agora que o seu desenvolvimento foi concluído (e testado 😉), abra um *pull request*, para que ele possa ser revisado pelos seus colegas. Sugestão: faça pelo github (site) mesmo. Quando você der push na sua nova branch, deve aparecer um botão de "criar pull request" na página inicial do seu repositório.
4. **Finalizar**: Depois que você receber o approve, faça o *merge* do seu código, ou seja, junte o código da sua branch de desenvolvimento com o da branch estável. Sugestão: também fazer pelo github. Ao final do PR, você verá um botão de "fazer merge". Este passo está bem simplificado. Em uma aplicação real, após o code review teria todo o processo de "colocar em produção" e validar, mas isso fica para um próximo momento.

#### Atualizar a main local

Antes de começar a desenvolver a próxima tarefa, não se esqueça de *puxar a main*, ou seja, atualizar o git do seu computador com o novo estado da main do repositório remoto depois de mergear o seu desenvolvimento.

Exemplo: minha main tem a função `fazer bola`. Abri uma branch `minha-branch-para-fazer-quadrado`, onde implementei `fazer quadrado`, e fiz o push. Nesse momento no meu git local e remoto tenho `fazer bola` na main, e `fazer bola` e `fazer quadrado` na branch de desenvolvimento. Depois que eu abrir PR e mergear a branch de desenvolvimento na main *pelo github*, a main remota vai ter `fazer bola` e `fazer quadrado`, mas a main local vai ter só `fazer bola`. Para atualizar, basta alterar a minha branch local para a main e fazer um pull.

```shell
git checkout main
git pull
```
