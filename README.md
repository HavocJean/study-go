# Study Golang (Basic MVC)

## User API


### Create User
Esta é a rota principal do projeto em Go, usando Gin. Ela recebe os dados de um usuário, valida, converte para o domínio e envia ao service responsável pela regra de negócio.
- **POST** `/users` 
- Body:
    ```json
    {
        "email": "user@example.com",
        "password": "P@ssw0rd!",
        "name": "John Doe",
        "age": 25
    }
    ```

### Processo (Create User)
1. A requisição POST `/users` chega em:
   - `internal/routes/routes.go`
2. O controller recebe e valida o JSON:
   - `internal/controller/create_user.go` (usa `request.UserRequest`)
3. O JSON é convertido para domínio (UserDomain):
   - `internal/model/user.go`
4. O service aplica regras de negócio:
   - `internal/model/service/create_user.go`
5. O service chama o repositório para persistir os dados:
   - `internal/model/repository/create_user_repository.go`
6. O repositório converte domínio em entidade e insere na coleção do MongoDB:
   - `internal/model/repository/entity/user_entity.go`
   - `internal/model/repository/entity/converter/convert.go`
7. A resposta é construída e retornada:
   - `internal/view/domain_to_response.go`

> Request → Controller (valida) → Domain → Service (regras) → Repository (Mongo) → Response


## Structure
```
internal/
    ├─ config/
    │   ├─ database/mongodb/
    │   │   └─ mongodb_connection.go
    │   ├─ rest_error/
    │   │   └─ rest_error.go
    │   └─ validation/
    │       └─ validate_user.go
    ├─ routes/
    │   └─ routes.go
    ├─ controller/
    │   ├─ create_user.go
    │   └─ model/
    │       └─ request/
    │           └─ user_request.go
    ├─ model/
    │   ├─ user.go
    │   └─ repository/
    │       ├─ user_repository.go
    │       ├─ create_user_repository.go
    │       └─ entity/
    │           ├─ user_entity.go
    │           └─ converter/
    │               └─ convert.go
    ├─ model/service/
    │   ├─ create_user.go
    │   └─ user_interface.go
    └─ view/
        └─ domain_to_response.go

root files:
- `main.go`
- `init_dependecies.go`
- `Dockerfile`
- `docker-compose.yml`
- `.env`
```


## How to run
- Local (dev):
  ```bash
  go run main.go
  ```
- Local Docker Compose:
  ```bash
  docker compose up --build
  ```


## Commands
- Atualizar e limpar dependências:
    `go mod tidy`