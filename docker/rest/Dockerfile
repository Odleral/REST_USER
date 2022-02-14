# Build enviroment
#-----------------
FROM golang:1.17-alpine as build-env
WORKDIR /REST

RUN apk update && apk add --no-cach gcc musl-dev git bash

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app

CMD ["/REST/bin/app"]

# Deployment enviroment
#----------------------
FROM alpine
RUN apk update && apk add --no-cache bash postgresql-client

COPY --from=build-env /REST/bin/app /REST/

EXPOSE 80
CMD ["/REST/app"]