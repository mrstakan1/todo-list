#!/bin/sh
# Запускаем Go-API в фоне
api &
# Запускаем Nginx (в foreground, иначе контейнер завершится)
exec nginx -g 'daemon off;'
