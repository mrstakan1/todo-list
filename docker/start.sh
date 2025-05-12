#!/usr/bin/env sh
set -eu

echo "▶️  Starting PostgreSQL …"
# — слушаем только localhost внутри контейнера
su postgres -c "postgres -D \"$PGDATA\" -c listen_addresses='localhost'" &
PID_PG=$!

# ─────────── Wait for Postgres ───────────
echo -n "⏳  Waiting for DB to be ready "
for i in $(seq 1 30); do
    if pg_isready -h localhost -U "$DB_USER" >/dev/null 2>&1; then
        echo " ✓"
        break
    fi
    printf '.'
    sleep 1
done

echo "▶️  Starting Go API on :$API_PORT …"
/usr/local/bin/api &
PID_API=$!

echo "▶️  Starting Nginx …"
#  «daemon off» — чтобы nginx работал на переднем плане
nginx -g "daemon off;" &
PID_NGINX=$!

# ─────────── Graceful shutdown ───────────
term_handler() {
  echo "\n🛑  Caught SIGTERM — stopping services …"
  kill $PID_NGINX $PID_API $PID_PG 2>/dev/null || true
  wait
  echo "✅  Services stopped — exiting."
  exit 0
}

trap term_handler TERM INT

# Ждём любой дочерний процесс
wait