module github.com/loopcontext/go-graphql-orm/test

go 1.16

require (
	github.com/99designs/gqlgen v0.13.0
	github.com/akrylysov/algnhsa v0.12.1
	github.com/cloudevents/sdk-go v1.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/loopcontext/cloudevents-aws-transport v1.0.9
	github.com/loopcontext/go-graphql-orm v0.0.0-20210302110458-c32b8f56ab03
	github.com/mitchellh/mapstructure v1.4.1
	github.com/rs/cors v1.7.0
	github.com/rs/zerolog v1.20.0
	github.com/urfave/cli v1.22.5
	github.com/vektah/gqlparser/v2 v2.1.0
	gopkg.in/gormigrate.v1 v1.6.0
)

replace github.com/loopcontext/go-graphql-orm v0.0.0-20210302221454-2d745f04c960 => ../
