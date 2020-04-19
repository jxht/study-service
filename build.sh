#!/bin/sh

set -e
set -x

DIR=$(cd `dirname $0` && pwd -P)
cd $DIR
go version
go env

case $1 in
  "run" )
    go run .
    ;;
  "add-pre-commit" )
    cp ./bin/pre-commit .git/hooks/pre-commit
    chmod +x .git/hooks/pre-commit
    ;;
  * )
    ;;
esac