## Build
FROM docker.karimi.dev/library/golang:1.19-bulseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /simple-store

## Deploy
FROM docker.karimi.dev/gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /simple-store /simple-store

# TODO: Fix ownership and make it nonroot
# USER nonroot:nonroot

ENTRYPOINT ["/simple-store"]
