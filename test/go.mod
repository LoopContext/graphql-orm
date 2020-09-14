module github.com/loopcontext/go-graphql-orm/test

go 1.14

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/akrylysov/algnhsa v0.12.1
	github.com/cloudevents/sdk-go v1.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/gorilla/mux v1.6.2
	github.com/graph-gophers/dataloader v5.0.0+incompatible
	github.com/jinzhu/gorm v1.9.14
	github.com/loopcontext/cloudevents-aws-transport v1.0.8
	github.com/loopcontext/go-graphql-orm v0.0.0-20201126122147-e0ff8f68c13b
	github.com/loopcontext/graphql-orm v1.0.8
	github.com/mitchellh/mapstructure v1.3.2
	github.com/rs/cors v1.7.0
	github.com/rs/zerolog v1.20.0
	github.com/urfave/cli v1.22.4
	github.com/vektah/gqlparser/v2 v2.0.1
	gopkg.in/gormigrate.v1 v1.6.0
)

replace github.com/loopcontext/go-graphql-orm v0.0.0-20201126122147-e0ff8f68c13b => ../
