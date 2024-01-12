#!/bin/bash -e

dbs=("sqlite" "mysql")

if [[ $(pwd) == *"go-flyway/tests"* ]]; then
  cd ..
fi

if [ -d tests ]
then
  cd tests
  go get -u -t ./...
  go mod download
  go mod tidy
  cd ..
fi
if [ -d tests ]
then
  cd tests
  go test ./...
fi

