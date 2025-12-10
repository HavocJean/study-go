# Study Golang (Basic MVC)

## User API


### Create User
Esta é a primeira rota do projeto em Go, usando Gin. Ela recebe os dados de um usuário, valida, converte para o domínio e envia ao service responsável pela regra de negócio.
- `POST /users` 
- Body:
    ```json
    {
        "email": "user@example.com",
        "password": "P@ssw0rd!",
        "name": "John Doe",
        "age": 25
    }
    ```

# Proccess
1. A requisição POST /users chega em:
    `internal/routes/routes.go`
2. O controller é chamado:
    `internal/controller/create_user.go`
3. O JSON recebido é validado conforme regras definidas em:
    `internal/controller/model/request/user_request.go`
4. Os dados são convertidos em domínio (UserDomain):
    `internal/model/user.go`
5. O service executa a regra de negócio:
    `internal/model/service/user_interface.go`

## Structure
internal/
    ├─ routes/            → registra rotas
    ├─ controller/        → trata HTTP
    │   ├─ create_user.go
    │   └─ model/request/ → validação do body
    ├─ model/             → domínio do usuário
    │   └─ user.go
    └─ model/service/     → regras de negócio
        ├─ create_user.go
        └─ user_interface.go



## Commands
- Executar projeto
    `go run main.go`

- Atualizar e limpar arquivos
    `go mod tidy`