FROM golang:1.20-alpine

WORKDIR /go/src/app
COPY . .

RUN apk upgrade --update && apk --no-cache add git

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]