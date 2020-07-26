FROM golang:1-alpine as build

WORKDIR /app
COPY http_server_2020.go .
RUN go build http_server_2020.go

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/http_server_2020 /app/http_server_2020

EXPOSE 8081
ENTRYPOINT ["./http_server_2020"]