FROM golang:1.25.5-alpine AS base

RUN apk add --no-cache git ca-certificates

WORKDIR /app

COPY go.mod go.sum ./

#ENV GOPROXY=direct

RUN go mod download

FROM base AS dev

CMD ["go", "run", "./cmd/ritmotrack/main.go"]

FROM base AS build

RUN apk add --no-cache git make gcc musl-dev
# Копируем весь проект
COPY . .

# Собираем статический бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app/main.go