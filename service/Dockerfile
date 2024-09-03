FROM golang:1.21-alpine

RUN apk add curl

COPY ./ /go/src/app

WORKDIR /go/src/app/src

RUN go mod download && go build -o /app/app

EXPOSE 8000

CMD ["/app/app"]
