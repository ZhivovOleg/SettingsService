ifndef version
ifneq ($(MAKECMDGOALS),help)
ifneq ($(MAKECMDGOALS),audit)
ifneq ($(MAKECMDGOALS),swag)
$(error ERROR: 'version' flag must be defined)
endif
endif
endif
endif

MAIN_PACKAGE_PATH := ./cmd/SettingsService
BINARY_PATH := ./bin
BINARY_NAME := SettingsService
TARGET_OS := linux
TARGET_ARCH := amd64
SUFFIX := 

## help: print this help message
.PHONY: help
help:
	@echo
	@echo '    Usage:'
	@echo
	@echo '    make [help, audit, swag]'
	@echo '    make version=<version number> [lin, win, mac]'
	@echo
	@echo '    Command list:'
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

.PHONY: all
all: lin

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
build: swag
	GOOS=${TARGET_OS} GOARCH=${TARGET_ARCH} go build -o ./bin/${TARGET_OS}-${TARGET_ARCH}/${version}/SettingsService${SUFFIX} -ldflags "-s -w -X main.Version=${version}" ${MAIN_PACKAGE_PATH}