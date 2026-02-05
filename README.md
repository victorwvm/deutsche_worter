## Deutsche Wörter

Projeto simples em Go para armazenar e listar palavras em alemão usando MySQL.

### Tecnologias
- Go
- MySQL
- Docker

### Como rodar
1. Suba o MySQL com Docker
2. Crie o banco e tabela
3. Execute `go run main.go`

## Configuração

Este projeto utiliza variáveis de ambiente para configurar a conexão com o banco de dados.

Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

DB_USER  
DB_PASSWORD  
DB_HOST  
DB_PORT  
DB_NAME  
