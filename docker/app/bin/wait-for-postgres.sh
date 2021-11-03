#!/usr/bin/env bash

host="$1"
shift
cmd="$@"

until psql postgresql://"${DB_USER}:${DB_PASS}@$host/${DB_NAME}" -c "select 1"; do
  >&2 echo "Postgresql is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgresql is up - executing command"
exec "$cmd"