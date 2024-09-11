# API CRUD com Go, PostgreSQL e Gorilla Mux

Este projeto é uma API simples para realizar operações de CRUD (Create, Read, Update, Delete) de usuários, construída em Go utilizando PostgreSQL como banco de dados e o pacote Gorilla Mux para roteamento.

## Estrutura do Projeto

A organização do projeto está separada em diferentes camadas:

- **cmd/**: Contém o arquivo principal (`main.go`), responsável por iniciar a aplicação.
- **controllers/**: Contém os controladores responsáveis por lidar com as requisições HTTP.
- **services/**: Contém a lógica de negócios, separada da camada de controle.
- **models/**: Define as estruturas (structs) dos dados usados no projeto.
- **infra/**: Contém a configuração da conexão com o banco de dados.
- **middlewares/**: Contém middlewares para lidar com as requisições HTTP.

### Estrutura de Diretórios

```bash
first-go-crud/
│
├── cmd/
│   └── main.go
├── middlewares/
│   └── json_content_type.go
├── controller/
│   └── user_controller.go
├── infra/
│   └── db.go
├── models/
│   └── user.go
└── services/
    └── user_service.go
```
