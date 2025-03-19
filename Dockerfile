FROM golang:1.20

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o main ./api/api.go

EXPOSE 1323
CMD ["./main"]