FROM golang:1.24.2-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o server ./cmd/server
EXPOSE 8080
ENTRYPOINT ["./server"]
