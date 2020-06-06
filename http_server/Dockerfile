FROM golang:1-alpine as build

WORKDIR /app
COPY http_server.go .
RUN go build http_server.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/http_server /app/http_server

EXPOSE 8180
ENTRYPOINT ["./http_server"]