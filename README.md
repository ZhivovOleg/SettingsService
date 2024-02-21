## About

Self-learning Golang project.
<br />
Store and managing settings for microservices.

## Used packages
- [go-gin webserver](https://github.com/gin-gonic/gin)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [pgx - PostgreSQL Driver](https://github.com/jackc/pgx)
- [pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool)
- [zap logger](https://github.com/uber-go/zap)


## Develop & Debug & Test
 
Recommended IDE - VSCode
Environment for project in `./.vscode/launch.json`

For another IDE's don't forget set up env variable `"SettingsServiceEnv": "dev"`

1. [Install Go](https://go.dev/dl/)
0. Install golang `swag` utility:
    ```bash
    go install github.com/swaggo/swag/cmd/swag@latest
    ```
0. Install dependencies:
    ```bash
    go get .
    ```
0. Run test database docker image:
    ```bash
    docker compose up -d ./.test/database/docker-compose.yml
    ```

0. After changes regenerate swagger files:
    ```bash
    $HOME/go/bin/swag init
    ```

## Production 

build with remove the symbol and debug info:
```bash
go build -ldflags "-s -w"
```

## TODO

1. Add remove-method for exists json fields
0. 