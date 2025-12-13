FROM golang:1.25.5-alpine AS base

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

FROM base AS build

RUN apk add --no-cache git make gcc musl-dev

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/main.go