#!/bin/sh

ls
go get -u github.com/labstack/echo/...
go get -u github.com/gorilla/handlers
go get -u github.com/gorilla/mux
go get -u google.golang.org/api/option
go get -u firebase.google.com/go
go get -u github.com/dgrijalva/jwt-go
go get -u github.com/go-sql-driver/mysql

git clone https://github.com/herumi/mcl.git
cp mcl/ffi/go/mcl/* go/src/key_manager/pkg/mcl/ -rf
cp mcl/ffi/go/mcl/* go/src/mcl_server/pkg/mcl/ -rf
ls