FROM golang:1.23.4-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o server ./cmd/server

FROM scratch
WORKDIR /app
ENV JWT_SECRET=devsecret
COPY --from=build /app/server .
CMD ["./server"]