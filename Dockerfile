FROM golang:1.26-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o main ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main .

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 CMD wget -qO- http://localhost:8000/ping || exit 1

EXPOSE 8000

CMD ["./main"]