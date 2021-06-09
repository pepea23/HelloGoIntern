#!/bin/bash

set -o errexit

main() {
  echo "CREATING DATABASE $POSTGRES_DB_PROD AND $POSTGRES_DB_TEST"
  create_databases

  echo "CREATING UUID-OSSP EXTENSIONS" 
  create_extensions
}

create_databases() {
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -c "Create database $POSTGRES_DB_PROD"
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" -c "Create database $POSTGRES_DB_TEST"
}

create_extensions() {
  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$POSTGRES_DB_PROD" <<-EOSQL
     CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL

  psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname="$POSTGRES_DB_TEST" <<-EOSQL
     CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL
}

main "$@"