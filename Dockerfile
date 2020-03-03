FROM golang:1.13.8-alpine3.11 as builder

WORKDIR /build/
COPY . .
RUN \
  go mod download && \
  go mod verify
RUN CGO_ENABLED=0 GOOS=linux go build ./cmd/webserver

FROM scratch

WORKDIR /app/
EXPOSE 8080
COPY --from=builder /build/webserver .
CMD ["./webserver"]
