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
 
Recommended IDE - VSCode.
<br />
Environment for project in `./.vscode/launch.json`.
<br />
For another IDE's don't forget set up env variable `"SettingsServiceEnv": "dev"`.

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
    swag init -g ./cmd/SettingsService/main.go -o ./docs
    ```

## Production 

build with remove the symbol and debug info:
```bash
go build -o=SettingsService -ldflags "-s -w" ./cmd/SettingsService
```
where:
- `-w` turns off DWARF debugging information
- `-s` turns off generation of the Go symbol table

or use `Makefile` run for your platform:
```bash
make version="1.0.0" build-mac
```
or 
```bash
make version="1.0.0" build-lin
```
or
```bash
make version="1.0.0" build-win
```

## TODO

1. dockerize 
0. Add push mechanism for setting's updates
0. Add auth
0. Add unit-tests
0. Try another DB, for example mongo
0. Write help instructions
0. 

## Closed fixes

1. Appy [project layout](https://github.com/golang-standards/project-layout/tree/master)
0. makefile
0. versioning
0. 