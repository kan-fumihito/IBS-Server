#!/bin/sh
docker-compose build
docker-compose up -d
docker-compose exec golang go build -tags bn256 mcl_server/cmd/server.go
docker-compose exec golang go build -tags bn256 key_maneger/cmd/masterkey.go

docker-compose exec mysql bash -c "chmod 775 /docker-entrypoint-initdb.d/init.sh"
docker-compose exec mysql bash -c "/docker-entrypoint-initdb.d/init.sh"
