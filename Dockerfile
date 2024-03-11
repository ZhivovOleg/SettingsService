# syntax=docker/dockerfile:1

################ GOLANG ################

FROM golang:alpine as go_builder

LABEL stage=backend_builder

ENV CGO_ENABLED 0
ENV GOOS linux

COPY ./api /build/api
COPY ./cmd /build/cmd
COPY ./docs /build/docs
COPY ./internal /build/internal
COPY ./api /build/api
COPY ./go.mod /build/go.mod
COPY ./go.sum /build/go.sum

WORKDIR /build

RUN mkdir result
RUN apk update --no-cache && apk add --no-cache tzdata
RUN go mod download
RUN go build -ldflags="-s -w" -o ./result ./cmd/SettingsService

################ VUE ##################

FROM node:lts-alpine as vue_builder

LABEL stage=frontend_builder

COPY ./web/src /build/src

RUN rm -rf /build/src/node_modules
RUN rm /build/src/package-lock.json

WORKDIR /build/src

RUN npm install && npm run build

################ RUN ##################

FROM alpine:latest as runner

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=go_builder build/result SettingsService/
COPY --from=vue_builder build/index.html SettingsService/web/index.html
COPY --from=vue_builder build/logo.svg SettingsService/web/logo.svg
COPY --from=vue_builder build/assets SettingsService/web/assets

ENV SettingsServicePort 9000
ENV SettingsServiceDbConnectionString postgres://pg:1@host.docker.internal:5432/servicesSettings_db
ENV SettingsServiceEnv dev

WORKDIR /SettingsService

EXPOSE 9000

ENTRYPOINT ["./SettingsService"]