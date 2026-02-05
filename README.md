## Deutsche Wörter

Conexão simples com banco de dados em Go para armazenar e listar palavras em alemão.
Código feito apenas para estudo e prática de Go com MySQL.

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
```
DB_USER  
DB_PASSWORD  
DB_HOST  
DB_PORT  
DB_NAME  
```

<img width="698" height="146" alt="{162031F2-AE97-4B21-B3FE-FF05EDC5862D}" src="https://github.com/user-attachments/assets/23351bd0-876c-4660-94cb-4de6604b678b" />
