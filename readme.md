# Go Oportunities

Este é um projeto em Go que implementa um CRUD (Create, Read, Update, Delete) de Oppenings. O objetivo deste projeto é fornecer uma aplicação básica para gerenciar oportunidades (Oppenings) de alguma entidade, como vagas de emprego, projetos, eventos, etc.

### Requisitos
Certifique-se de ter o seguinte software instalado em sua máquina antes de executar este projeto:

* Go 
* Make

### Configuração

Certifique-se de ter o Go instalado em sua máquina. Você pode fazer o download da versão mais recente do Go em https://golang.org/dl/ e seguir as instruções de instalação específicas para o seu sistema operacional.

##### Clonando o repositório
Clone este repositório em sua máquina local:
```
git clone https://github.com/wesleybruno/go-opprtunities
```

##### Instalando as dependências
Navegue até o diretório do projeto e execute o seguinte comando para instalar as dependências necessárias:
```
go mod download
```
##### Executando a aplicação
Para executar a aplicação, execute o seguinte comando no diretório do projeto:
```
go run main.go
```
A aplicação será executada e estará disponível no endereço http://localhost:3000.


#### Contribuindo
Se você quiser contribuir para este projeto, siga as etapas abaixo:

Faça um fork deste repositório.
Crie uma branch para sua nova funcionalidade ou correção de bug: git checkout -b minha-branch.
Faça suas alterações e faça commit delas: git commit -m "Minha alteração".
Envie para o repositório remoto: git push origin minha-branch.
Envie um pull request para este repositório.

#### Basic Commands

- Build and run
```
go build -o NOME_EXECUTAVEL LOCAL_DE_DESTINO ( . PASTA LOCAL)

ou 

go run NOME_ARQUIVO
```

- Lint imports
```
go mod tidy
```
- Generate Docs
```
make docs
```