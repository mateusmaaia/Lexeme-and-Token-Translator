FROM golang:1.12 AS build-env
EXPOSE 8090

LABEL maintainer "mateusinocenciomaia@gmail.com"

# Just to download dependencies and caching them
RUN mkdir /lexeme-and-token-translator
WORKDIR /lexeme-and-token-translator
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/lexeme-and-token-translator

FROM alpine:latest

ENV PATH = ""

WORKDIR /app
COPY --from=build-env /go/bin/lexeme-and-token-translator .

CMD ["/app/lexeme-and-token-translator ${PATH}"]