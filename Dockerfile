# syntax=docker/dockerfile:1
LABEL org.opencontainers.image.source https://github.com/clarkbains/waypoint-proxy
FROM golang:1.16-alpine AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download -x
COPY ./proto ./proto
COPY ./routers ./routers
COPY *.go ./
RUN go build -o /wp
CMD ["/wp"]

##
## Deploy
##
FROM alpine:latest
WORKDIR /
COPY root.pem ./
COPY --from=build /wp /wp
EXPOSE 8080
ENTRYPOINT ["/wp"]
