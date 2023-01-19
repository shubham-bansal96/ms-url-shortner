FROM golang:alpine AS builder

WORKDIR /app/ms-url-shortner

ADD . /app/ms-url-shortner/

LABEL org.opencontainers.image.source="https://github.com/shubham-bansal96/ms-url-shortner"

RUN go mod download
RUN go build -o ms-url-shortner .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ms-url-shortner/ms-url-shortner /app/
COPY --from=builder /app/ms-url-shortner/config.yml /app/

CMD ./ms-url-shortner

LABEL org.opencontainers.image.source="https://github.com/shubham-bansal96/ms-url-shortner"

EXPOSE 4000

