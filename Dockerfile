FROM golang:alpine AS builder

ENV GO111MODULE=on
WORKDIR /app
ADD ./ /app

RUN apk update --no-cache
RUN apk add git
RUN go build -o golang-test  .

FROM alpine
WORKDIR /app1
COPY --from=builder /app /app1

EXPOSE 8000
ENTRYPOINT ["/app1/golang-test"]
