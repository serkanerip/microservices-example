FROM golang:alpine as builder

RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api -ldflags="-w -s" cmd/server.go

FROM scratch

WORKDIR /app
COPY --from=builder /app/api /usr/bin/

ENTRYPOINT ["api"]