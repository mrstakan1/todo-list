######################  build frontend  ######################
FROM node:20-alpine AS frontend
WORKDIR /app
COPY frontend/package*.json ./
RUN npm ci
COPY frontend .
RUN npm run build                       # → /app/dist
######################  build backend   ######################
FROM golang:1.24.2-alpine AS backend
WORKDIR /app
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend .
RUN CGO_ENABLED=0 go build -o /api ./cmd/server
######################  final image     ######################
FROM alpine:3.20
ENV DB_USER=todo \
    DB_PASS=todo \
    DB_NAME=todo \
    PGDATA=/var/lib/postgresql/data \
    API_PORT=8080 \
    JWT_SECRET=devsecret

# пакеты: nginx + postgresql + tini
RUN apk add --no-cache nginx postgresql16 postgresql16-client tini

# инициализация БД
RUN install -d -o postgres -g postgres "$PGDATA" && \
    su postgres -c "initdb -D $PGDATA" && \
    su postgres -c "pg_ctl -D $PGDATA -o '-c listen_addresses=localhost' -w start" && \
    su postgres -c "psql -c \"CREATE USER ${DB_USER} WITH PASSWORD '${DB_PASS}';\"" && \
    su postgres -c "createdb -O ${DB_USER} ${DB_NAME}" && \
    su postgres -c "pg_ctl -D $PGDATA -m fast stop"

# копируем артефакты
COPY --from=frontend /app/dist      /usr/share/nginx/html
COPY --from=backend  /api           /usr/local/bin/api

# старт-скрипт
COPY docker/start.sh                /start.sh
RUN chmod +x /start.sh

EXPOSE 80
ENTRYPOINT ["/bin/sh", "/start.sh"]
