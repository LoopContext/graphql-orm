package templates

// DockerfileDev development dockerfile
var DockerfileDev = `FROM golang:alpine as base

RUN apk update && apk upgrade && apk add --no-cache bash git openssh curl

WORKDIR /graphql-server

COPY . /graphql-server/
RUN go mod download

RUN GO111MODULE=off go get github.com/oxequa/realize

CMD ["./scripts/run-dev.sh"]
`

// DockerfileProd production dockerfile
var DockerfileProd = `# Multistaged build production golang service
FROM golang:alpine as base

FROM base AS ci

RUN apk update && apk upgrade && apk add --no-cache git
RUN mkdir /build
ADD . /build/
WORKDIR /build

# Build prod
FROM ci AS build-env

RUN ls
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix \
    cgo -ldflags '-extldflags "-static"' -o server .

FROM alpine AS prod
RUN apk --no-cache add ca-certificates

COPY --from=build-env build/server ./graphql-server/

# Set all the ENV variables here
CMD ["./graphql-server/server", "start"]
`

// Dockerfile ...
var Dockerfile = `FROM golang as builder

ENV GO111MODULE=on
WORKDIR /go/src/{{.Config.Package}}

COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /tmp/app *.go

FROM loopcontext/wait-for as wait-for

FROM alpine:3.5

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=wait-for /usr/local/bin/wait-for /usr/local/bin/wait-for
COPY --from=builder /tmp/app /usr/local/bin/app

# https://serverfault.com/questions/772227/chmod-not-working-correctly-in-docker
RUN chmod +x /usr/local/bin/app

ENTRYPOINT []
CMD [ "/bin/sh", "-c", "wait-for ${DATABASE_URL} && app start"]
`

// RunDevSh ...
var RunDevSh = `#!/bin/sh
app="graphql-server"
printf "\nStart running: $app\n"
# Set all ENV vars for the server to run
export $(grep -v '^#' .env | xargs)
time /$GOPATH/bin/realize start --no-config --run *.go start --cors
# This should unset all the ENV vars, just in case.
# unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/''' | xargs)
printf "\nStopped running: $app\n\n"
`

// RunSh ...
var RunSh = `#!/bin/sh
buildPath="build"
app="./"
program="$buildPath/$app"

printf "\nStart app: $app\n"
# Set all ENV vars for the program to run
export $(grep -v '^#' .env | xargs)
time ./$program
# This should unset all the ENV vars, just in case.
unset $(grep -v '^#' .env | sed -E 's/(.*)=.*/\1/' | xargs)
printf "\nStopped app: $app\n\n"
`

// DotenvExample example .env file
var DotenvExample = `DATABASE_URL=DATABASE_URL=postgres://test:test@host.docker.internal/test?sslmode=disable
EXPOSE_MIGRATION_ENDPOINT=false
TABLE_NAME_PREFIX=
EVENT_TRANSPORT_URL=
EVENT_TRANSPORT_SOURCE=
`

// DockerComposeYml file
var DockerComposeYml = `version: "3.8"
services:
  dev:
    command: ./scripts/run-dev.sh
    build:
      context: .
      dockerfile: docker/dev.dockerfile
    env_file:
      - .env.dev
    ports:
      - 8081:8081
      - 5002:5002
    volumes:
      - .:./
  dev-linux:
    network_mode: "host"
    command: ./scripts/run-dev.sh
    build:
      context: .
      dockerfile: docker/dev.dockerfile
    env_file:
      - .env.dev
    ports:
      - 8081:8081
      - 5002:5002
    volumes:
      - .:./
  prod:
    build:
      context: .
      dockerfile: docker/prod.dockerfile
    env_file:
      - .env
    ports:
      - 80:80
  docker:
    build:
      context: .
      dockerfile: Dockerfile
    env_file:
      - .env
    ports:
      - 80:80
`
