# go-microservice

[![GitHub Actions](https://github.com/cheesydev/go-microservice/workflows/Test/badge.svg?branch=master)](https://github.com/cheesydev/go-microservice/actions?query=branch%3Amaster)

Simple Go HTTP server.

Run it locally.
```
go test ./..
go build ./cmd/webserver
./webserver
```

Build image, run and publish.
```
docker build -t rafaelportela/go-microservice:0.1 .
docker run --rm -p 8080:8080 rafaelportela/go-microservice:0.1
docker push rafaelportela/go-microservice:0.1
```
