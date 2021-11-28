#!/usr/bin/env bash
mysql -u docker -pdocker main < "/docker-entrypoint-initdb.d/001-create-tables.sql"
mysql -u docker -pdocker main < "/docker-entrypoint-initdb.d/002-company-insert.sql"
mysql -u docker -pdocker main < "/docker-entrypoint-initdb.d/003-user-insert.sql"