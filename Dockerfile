FROM golang:1.21.3-alpine

# RUN apt updatge && apt update -y && \ apt install -y git \ make openssh-client

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ./bin/main.exe ./app/main.go

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]