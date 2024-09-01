# Etapa de compilacion 
FROM golang:1.23-alpine AS builder
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=0

COPY ./*go* .
RUN go mod tidy
RUN go build -o server

# Imagen para produccion
FROM alpine:latest AS prod

COPY --from=builder /app/server /usr/local/bin/server
EXPOSE 5000

ENTRYPOINT ["/usr/local/bin/server"]
