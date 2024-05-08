FROM golang:1.22.2

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 8080

RUN go build ./cmd/app/main.go

CMD ["go", "run", "./cmd/app/main.go"]
