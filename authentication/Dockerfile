FROM golang:1.23-alpine3.21

WORKDIR /app/authentication

COPY go.mod go.sum ./

RUN go mod download

COPY ./authentication/ .

COPY ../.env .

RUN go build -o /app/authentication/auth /app/authentication/main.go

EXPOSE 8090

CMD ["/app/authentication/auth"]

