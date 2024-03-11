current_os := $$(go env GOOS)
current_arch := $$(go env GOARCH)

ifdef v
version = $$v
endif


ifeq ($(MAKECMDGOALS),)
ifndef version
override version := 0.0.0
endif
endif

ifeq ($(MAKECMDGOALS),mac)
ifndef version
$(error ERROR: 'version' or 'v' flag must be defined)
endif
endif

ifeq ($(MAKECMDGOALS),win)
ifndef version
$(error ERROR: 'version' or 'v' flag must be defined)
endif
endif

ifeq ($(MAKECMDGOALS),lin)
ifndef version
$(error ERROR: 'version' or 'v' flag must be defined)
endif
endif

MAIN_PACKAGE_PATH := ./cmd/SettingsService
BINARY_PATH := ./bin
BINARY_NAME := SettingsService
TARGET_OS := linux
TARGET_ARCH := amd64
SUFFIX := 

.PHONY: default
default: TARGET_OS := ${current_os}
default: TARGET_ARCH := ${current_arch}
default: build

## help: print this help message
.PHONY: help
help:
	@echo
	@echo '    Usage:'
	@echo
	@echo '    make'
	@echo '    make v=<version number>'
	@echo '    make version=<version number>'
	@echo '    make [help, audit, swag, dep]'
	@echo '    make version=<version number> [lin, win, mac]'
	@echo
	@echo '  without params :: build project for current platform vith version 0.0.0'
	@echo '  v=<version number> equal version=<version number> :: build project for current platform with version <version number>'
	@echo
	@echo '    Command list:'
	@echo
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## audit: run quality control checks
.PHONY: audit
audit:
	@echo
	@echo VERIFY
	@echo
	go mod verify
	@echo
	@echo VET
	@echo
	go vet ./...
	@echo
	@echo STATIC CHECK
	@echo
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./... || true
	@echo
	@echo VULNERABILITIES
	@echo
	go run golang.org/x/vuln/cmd/govulncheck@latest ./... || true
	@echo
#	@echo TEST
#	@echo
#	go test -race -buildvcs -vet=off ./...

## lin: build project for linux_x64 
.PHONY: lin
lin: TARGET_OS := linux
lin: build

## mac: build project for macOS_arm64
.PHONY: mac
mac: TARGET_OS := darwin
mac: TARGET_ARCH := arm64
mac: build

## win: build project for win_x64 
.PHONY: win
win: TARGET_OS := windows
win: SUFFIX := .exe
win: build

## swag: generate API swagger documentation
.PHONY: swag
swag:
	go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/SettingsService/main.go -o docs

.PHONY: build
build: clean front swag 
	GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} go build -o ./bin/${TARGET_OS}-${TARGET_ARCH}/${version}/SettingsService${SUFFIX} -ldflags "-s -w -X main.Version=${version}" ${MAIN_PACKAGE_PATH}
	cp ./configs/appSettings.json ./bin/${TARGET_OS}-${TARGET_ARCH}/${version}/
	mkdir ./bin/${TARGET_OS}-${TARGET_ARCH}/${version}/web
	cp ./web/index.html ./bin/${TARGET_OS}-${TARGET_ARCH}/${version}/web/
	cp ./web/logo.svg ./bin/${TARGET_OS}-${TARGET_ARCH}/${version}/web/
	cp -R ./web/assets ./bin/${TARGET_OS}-${TARGET_ARCH}/${version}/web/

## front: build Vue SPA
.PHONY: front
front: clean-front
	(cd ./web/src/ && npm install && npm run build)

## dep: install dependencies for Vue(front) and Go(back) projects
.PHONY: dep
dep:
	@go mod .
	@(cd ./web/src/ && npm install)

.PHONY: clean-front
clean-front:
	@rm -rf ./web/assets
	@rm -rf ./web/index.html
	@rm -rf ./web/logo.svg

.PHONY: clean-back
clean-back:
	@rm -rf ./bin/${TARGET_OS}-${TARGET_ARCH}/${version}

## clean: remove all artifacts
.PHONY: clean
clean: clean-front clean-back

## image: build docker image
.PHONY: image
image: lin
	docker build --tag 'settingsservice' .

## 
.PHONY: docker
docker: image
	docker run -d -p 9000:9000 --name 'SettingsService' 'settingsservice'
