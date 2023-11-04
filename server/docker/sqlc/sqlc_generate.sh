#/bin/bash
docker pull kjconroy/sqlc
docker run --rm -v $(pwd)/../..:/src -w /src kjconroy/sqlc generate
