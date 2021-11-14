#!/usr/bin/env bash
set -e

host="$1"
shift
cmd="$@"

until psql postgresql://"${POSTGRES_USER}:${POSTGRES_PASSWORD}@$host/${POSTGRES_DB}" -c "select 1"; do
  >&2 echo "Postgresql is unavailable - sleeping"
  sleep 5
done

>&2 echo "Postgresql is up - executing command"
exec "$cmd"