FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go fmt ./...
RUN go build ./cmd/kruder

EXPOSE 5000
CMD ["./kruder"]

