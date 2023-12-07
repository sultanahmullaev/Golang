FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go get -v -u github.com/golang-migrate/migrate/v4

COPY . ./

RUN go build -o main ./cmd/api/

EXPOSE 4000
EXPOSE 5432

CMD ["migrate -path ./migrations -database postgres://sports:pa55word@db:5432/sports?sslmode=disable up"]
CMD ["./main"]