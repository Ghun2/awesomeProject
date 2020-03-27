# Dockerfile
FROM golang:1.14.1-alpine3.11

WORKDIR /app
COPY . .

RUN apk update && \
  apk add git && \
  go get github.com/cespare/reflex

EXPOSE 7777
CMD ["reflex", "-c", "reflex.conf"]