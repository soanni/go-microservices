#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	CREATE USER shippy WITH ENCRYPTED PASSWORD 'shippy';
	CREATE DATABASE users;
	GRANT ALL PRIVILEGES ON DATABASE users TO shippy;
EOSQL