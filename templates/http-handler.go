package templates

// HTTPHandler ...
var HTTPHandler = `package gen

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	jwtgo "github.com/dgrijalva/jwt-go"
	"gopkg.in/gormigrate.v1"
)

// GetHTTPServeMux ...
func GetHTTPServeMux(r ResolverRoot, db *DB, migrations []*gormigrate.Migration) *http.ServeMux {
	mux := http.NewServeMux()

	executableSchema := NewExecutableSchema(Config{Resolvers: r})
	gqlHandler := handler.New(executableSchema)
	gqlHandler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	gqlHandler.AddTransport(transport.Options{})
	gqlHandler.AddTransport(transport.GET{})
	gqlHandler.AddTransport(transport.POST{})
	gqlHandler.AddTransport(transport.MultipartForm{})
	gqlHandler.Use(extension.FixedComplexityLimit(300))
	if os.Getenv("DEBUG") == "true" {
		gqlHandler.Use(extension.Introspection{})
	}
	gqlHandler.Use(apollotracing.Tracer{})
	gqlHandler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	loaders := GetLoaders(db)

	if os.Getenv("EXPOSE_MIGRATION_ENDPOINT") == "true" {
		mux.HandleFunc("/migrate", func(res http.ResponseWriter, req *http.Request) {
			err := db.Migrate(migrations)
			if err != nil {
				http.Error(res, err.Error(), 400)
			}
			fmt.Fprintf(res, "OK")
		})
		mux.HandleFunc("/automigrate", func(res http.ResponseWriter, req *http.Request) {
			err := db.AutoMigrate()
			if err != nil {
				http.Error(res, err.Error(), 400)
			}
			fmt.Fprintf(res, "OK")
		})
	}
	mux.HandleFunc("/graphql", func(res http.ResponseWriter, req *http.Request) {
		ctx := initContextWithJWTClaims(req)
		ctx = context.WithValue(ctx, KeyLoaders, loaders)
		ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)
		req = req.WithContext(ctx)
		gqlHandler.ServeHTTP(res, req)
	})

	if os.Getenv("EXPOSE_PLAYGROUND_ENDPOINT") == "true" {
		playgroundHandler := playground.Handler("GraphQL playground", "/graphql")
		mux.HandleFunc("/graphql/playground", func(res http.ResponseWriter, req *http.Request) {
			ctx := initContextWithJWTClaims(req)
			ctx = context.WithValue(ctx, KeyLoaders, loaders)
			ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)
			req = req.WithContext(ctx)
			if req.Method == "GET" {
				playgroundHandler(res, req)
			}
		})
	}
	handler := mux

	return handler
}

func initContextWithJWTClaims(req *http.Request) context.Context {
	claims, _ := getJWTClaims(req)
	var principalID *string
	if claims != nil {
		principalID = &(*claims).Subject
	}
	ctx := context.WithValue(req.Context(), KeyJWTClaims, claims)
	if principalID != nil {
		ctx = context.WithValue(ctx, KeyPrincipalID, principalID)
	}
	return ctx
}

// GetPrincipalIDFromContext ...
func GetPrincipalIDFromContext(ctx context.Context) *string {
	v, _ := ctx.Value(KeyPrincipalID).(*string)
	return v
}

// GetJWTClaimsFromContext ...
func GetJWTClaimsFromContext(ctx context.Context) *JWTClaims {
	val, _ := ctx.Value(KeyJWTClaims).(*JWTClaims)
	return val
}

// JWTClaims ...
type JWTClaims struct {
	jwtgo.StandardClaims
	Scope *string
}

func getJWTClaims(req *http.Request) (*JWTClaims, error) {
	var p *JWTClaims

	tokenStr := strings.Replace(req.Header.Get("authorization"), "Bearer ", "", 1)
	if tokenStr == "" {
		return p, nil
	}

	p = &JWTClaims{}
	jwtgo.ParseWithClaims(tokenStr, p, nil)
	return p, nil
}

// Scopes ...
func (c *JWTClaims) Scopes() []string {
	s := c.Scope
	if s != nil && len(*s) > 0 {
		return strings.Split(*s, " ")
	}
	return []string{}
}

// HasScope ...
func (c *JWTClaims) HasScope(scope string) bool {
	for _, s := range c.Scopes() {
		if s == scope {
			return true
		}
	}
	return false
}
`
