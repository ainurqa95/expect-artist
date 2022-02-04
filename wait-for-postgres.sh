#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
username="$3"
db="$4"
shift
cmd="$@"

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -U "$username" "$db" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command KEK"
exec $cmd