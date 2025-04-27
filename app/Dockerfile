FROM golang:1.24.1 AS go-build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .


EXPOSE 3000
CMD ["./main"]




