# Multistaged build production golang service
FROM golang:alpine as base

FROM base AS ci

# To add sqlite3 support, add build-base to the package list
RUN apk update && apk upgrade && apk add --no-cache git
RUN mkdir /build
ADD . /build/
WORKDIR /build

# Build prod
FROM ci AS build-env

RUN go mod download

# To add sqlite3 support, change to CGO_ENABLED=1
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix \
    cgo -ldflags '-extldflags "-static"' -o server .

FROM alpine AS prod
RUN apk --no-cache add ca-certificates

COPY --from=build-env build/server ./graphql-server/

# Set all the ENV variables here
CMD ["./graphql-server/server", "start"]
