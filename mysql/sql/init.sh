#!/usr/bin/env bash

mysql -u user -ppassword --host=127.0.0.1 < "/docker-entrypoint-initdb.d/create-tb.sql"
