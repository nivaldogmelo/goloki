FROM golang:1.20.5-alpine as build

RUN apk update && apk add gcc g++

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build

FROM alpine

COPY --from=build /app/goloki /

CMD ["/goloki", "gen", "1"]
