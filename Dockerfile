# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS build
WORKDIR /app
COPY root.pem ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./proto ./proto
COPY ./routers ./routers
COPY *.go ./
RUN RUN go build -o /app
CMD ["/app"]

##
## Deploy
##
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY root.pem ./
COPY --from=build /app /app
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]
