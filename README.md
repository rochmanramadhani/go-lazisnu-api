### Golang [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) REST API LAZISNU

LAZISNU backend service with clean architecture.

- HOST : http://localhost:8801

### Pre Requisite

- Go version 1.19

### Full list what has been used:

* [echo](https://github.com/labstack/echo) - Web framework
* [go-redis](https://github.com/go-redis/redis) - Type-safe Redis client for Golang
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [jwt-go](https://github.com/golang-jwt/jwt) - JSON Web Tokens (JWT)
* [uuid](https://github.com/google/uuid) - UUID
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [swag](https://github.com/swaggo/swag) - Swagger
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gomock](https://github.com/golang/mock) - Mocking framework
* [CompileDaemon](https://github.com/githubnemo/CompileDaemon) - Compile daemon for Go
* [Docker](https://www.docker.com/) - Docker

### Getting Started

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/rochmanramadhani/go-lazisnu-api.git
    ```

2. Navigate to the project directory:

   ```bash
   cd go-lazisnu-api
   ```

3. Install the dependencies:

   ```bash
    go mod tidy
    ```

4. Install Tools:

   ```bash
    # swaggo
    go install github.com/swaggo/swag/cmd/swag@latest
   
    # sqlc
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
   
    # golang-migrate
    go install github.com/golang-migrate/migrate@latest
   
    # mockgen
    go install github.com/golang/mock/mockgen@latest
    ```

5. Run the application:
   ```base
   make docker-local
   ```

   ```bash
   make app-run
   ```

### Architecture

This project built in clean architecture that contains some layer :

1. Driver
2. Factory
3. Delivery
4. Repository
5. Usecase
6. Model

### Packages

This project have some existing driver :

1. Http (rest, ws, web)
2. Database (postgres, mysql)
3. Elasticsearch
4. Firebase
5. Sentry
6. Websocket
7. Cron

### Author

LAZISNU Team
