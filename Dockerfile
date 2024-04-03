FROM golang:1.22-alpine AS BUILD
RUN apk add --no-cache ca-certificates gcc g++ && update-ca-certificates
WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=1 go build -C cmd/http -v -o /usr/local/bin/app

FROM alpine
RUN apk add --no-cache sqlite
COPY --from=BUILD /usr/local/bin/app /app
COPY --from=BUILD /etc/ssl/certs /etc/ssl/certs
EXPOSE 8080
ENTRYPOINT [ "/app" ]
