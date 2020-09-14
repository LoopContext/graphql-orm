FROM golang:alpine as base

RUN apk --update upgrade && apk add --no-cache bash git openssh curl build-base
# removing apk cache
RUN rm -rf /var/cache/apk/*

WORKDIR /graphql-server
COPY . /graphql-server/

RUN go get -v
RUN go get -v -u github.com/cosmtrek/air

CMD ["./scripts/run-dev.sh"]
