FROM golang:latest as builder

WORKDIR /app/ms-url-shortner

ADD . /app/ms-url-shortner/

RUN go mod download
RUN go mod vendor
RUN go build -o ms-url-shortner .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ms-url-shortner/ms-url-shortner /app/
COPY --from=builder /app/ms-url-shortner/config.yml /app/

RUN apk add libc6-compat

CMD ./ms-url-shortner

EXPOSE 4000

