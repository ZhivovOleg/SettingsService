## About

Self-learning Golang project.
<br />
Store and managing settings for microservices.

## Used packages

### backend:
Golang
- [go-gin webserver](https://github.com/gin-gonic/gin)
- [gin-swagger](https://github.com/swaggo/gin-swagger)
- [pgx - PostgreSQL Driver](https://github.com/jackc/pgx)
- [pgxpool](https://pkg.go.dev/github.com/jackc/pgx/v4/pgxpool)
- [zap logger](https://github.com/uber-go/zap)

### frontend
Vue
- [Pinia state management](https://github.com/vuejs/pinia)
- [vue3-notification](https://github.com/kyvg/vue3-notification)
- [vanilla-jsoneditor](https://github.com/josdejong/svelte-jsoneditor)
- [tailwindcss](https://github.com/tailwindlabs/tailwindcss)
- [axios](https://github.com/axios/axios)

## Develop
 
Recommended IDE - VSCode.
<br />
Environment for project in `./.vscode/launch.json`.
<br />
For another IDE's don't forget set up env variables:
```bash
"SettingsServicePort": "9000",
"SettingsServiceDbConnectionString": "postgres://pg:1@localhost:5432/servicesSettings_db",
"SettingsServiceEnv": "dev"
```

### Requiremens
1. [Go](https://go.dev/dl/)
0. [Node](https://nodejs.org/en/download)

### Recommended extentions for VSCode
1. [Docker](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)
0. [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
0. [npm Intellisense](https://marketplace.visualstudio.com/items?itemName=christian-kohler.npm-intellisense)
0. [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)
0. [TypeScript Vue Plugin (Volar)](https://marketplace.visualstudio.com/items?itemName=Vue.vscode-typescript-vue-plugin)
0. [Volar](https://marketplace.visualstudio.com/items?itemName=Vue.volar) 
> ! disable Vetur

## Debug & Test

### Debug backend

1. Install golang `swag` utility:
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
0. Set breakpoints and press `F5`

### Debug frontend

1. Install node

0. Move to `./web/src`

0. Install dependencies `npm install`

0. `npm run dev`

### Run in Docker

```bash
make front
docker build --tag 'settingsservice' .
docker run -d -p 9000:9000 --name 'SettingsService' 'settingsservice'
```
or prepare image with tag `settingsservice`
```bash
make image
```
or run container with tag `settingsservice` and name `SettingsService`
```bash
make docker
```

## Production 

build with remove the symbol and debug info:
```bash
go build -o=SettingsService -ldflags "-s -w" ./cmd/SettingsService
(cd ./web/src/ && npm install && npm run build)
```
where:
- `-w` turns off DWARF debugging information
- `-s` turns off generation of the Go symbol table

or use `Makefile`:
```bash
make help
```

For non-docker environment don't forget set up env variables:
```bash
"SettingsServicePort": "9000",
"SettingsServiceDbConnectionString": "postgres://pg:1@localhost:5432/servicesSettings_db"
```

## TO DO

1. [healthchecks](https://github.com/tavsec/gin-healthcheck)
0. Add push mechanism for setting's updates
0. Add auth
0. Add unit-tests
0. Try another DB, for example mongo
0. fix config. How do it best? Now ENV 
0. Write help instructions

## History

1. Appy [project layout](https://github.com/golang-standards/project-layout/tree/master)
0. makefile
0. versioning
0. SPA
0. fix results of stat analyst (except swagger comments)
0. dockerize