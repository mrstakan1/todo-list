######################## 1. Сборка фронтенда ########################
FROM node:20-alpine AS frontend
WORKDIR /app
COPY frontend/package*.json ./
RUN npm ci
COPY frontend/ .
RUN npm run build             # → /app/dist

######################## 2. Сборка бэкенда ##########################
FROM golang:1.22-alpine AS backend
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/ .
RUN CGO_ENABLED=0 go build -o /api ./cmd/server

######################## 3. Финальный образ ########################
FROM alpine:3.20
# минимально необходимые пакеты
RUN apk add --no-cache nginx

# копируем статические файлы фронта
COPY --from=frontend /app/dist /usr/share/nginx/html
# копируем Go-бинарь
COPY --from=backend /api /usr/local/bin/api

# конфиг nginx (по умолчанию подходит, только root уже указан)
RUN sed -i 's/80 default_server;/80;/' /etc/nginx/conf.d/default.conf \
 && sed -i 's/index  index.html index.htm;/index  index.html;/' /etc/nginx/conf.d/default.conf

# стартовый скрипт
COPY docker/start.sh /start.sh
RUN chmod +x /start.sh

ENV DB_DSN="postgres://todo:todo@db:5432/todo?sslmode=disable" \
    JWT_SECRET="devsecret" \
    API_PORT=8080

EXPOSE 80
ENTRYPOINT ["/bin/sh", "/start.sh"]