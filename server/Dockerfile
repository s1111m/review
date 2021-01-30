FROM golang:1.15 as builder
WORKDIR /app
COPY . /app
RUN GO111MODULE=auto CGO_ENABLED=1 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app cmd/main.go
RUN pwd
RUN ls -lah

FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/app.env .

ENTRYPOINT ["/app/app"]
