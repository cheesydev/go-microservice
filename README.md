# go-microservice

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
