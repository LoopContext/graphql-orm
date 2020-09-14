package gen

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/apollotracing"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"gopkg.in/gormigrate.v1"
)

// GetHTTPServeMux HTTP Mux
func GetHTTPServeMux(r ResolverRoot, db *DB, migrations []*gormigrate.Migration) *mux.Router {

	mux := mux.NewRouter()

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
		mux.HandleFunc(os.Getenv("API_VERSION")+"/migrate", func(res http.ResponseWriter, req *http.Request) {
			err := db.Migrate(migrations)
			if err != nil {
				http.Error(res, err.Error(), 400)
			}
			fmt.Fprintf(res, "OK")
		})
		mux.HandleFunc(os.Getenv("API_VERSION")+"/automigrate", func(res http.ResponseWriter, req *http.Request) {
			err := db.AutoMigrate()
			if err != nil {
				http.Error(res, err.Error(), 400)
			}
			fmt.Fprintf(res, "OK")
		})
	}
	gqlBasePath := os.Getenv("API_GRAPHQL_BASE_RESOURCE")
	if gqlBasePath == "" {
		gqlBasePath = "/graphql"
	}
	mux.HandleFunc(os.Getenv("API_VERSION")+gqlBasePath, func(res http.ResponseWriter, req *http.Request) {
		ctx := initContextWithJWTClaims(req)
		ctx = context.WithValue(ctx, KeyLoaders, loaders)
		ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)
		req = req.WithContext(ctx)
		gqlHandler.ServeHTTP(res, req)
	})

	if os.Getenv("EXPOSE_PLAYGROUND_ENDPOINT") == "true" {
		playgroundHandler := playground.Handler("GraphQL playground", os.Getenv("API_VERSION")+gqlBasePath)
		mux.HandleFunc(os.Getenv("API_VERSION")+gqlBasePath+"/playground", func(res http.ResponseWriter, req *http.Request) {
			ctx := initContextWithJWTClaims(req)
			ctx = context.WithValue(ctx, KeyLoaders, loaders)
			ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)
			req = req.WithContext(ctx)
			if req.Method == "GET" {
				playgroundHandler(res, req)
			}
		})
	}
	return mux
}

// GetHTTPVercel func for be used with Vercel deployments
func GetHTTPVercel(r ResolverRoot, db *DB, migrations []*gormigrate.Migration, res http.ResponseWriter, req *http.Request) {
	if os.Getenv("DEBUG") == "true" {
		log.Debug().Msgf("Path base: %s", path.Base(req.URL.Path))
	}
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
		if path.Base(req.URL.Path) == "migrate" {
			err := db.Migrate(migrations)
			if err != nil {
				http.Error(res, err.Error(), 400)
			}
			fmt.Fprintf(res, "OK")
		}
		if path.Base(req.URL.Path) == "automigrate" {
			err := db.AutoMigrate()
			if err != nil {
				http.Error(res, err.Error(), 400)
			}
			fmt.Fprintf(res, "OK")
		}
	}
	gqlBasePath := os.Getenv("API_GRAPHQL_BASE_RESOURCE")
	if gqlBasePath == "" {
		gqlBasePath = "graphql"
	}
	if path.Base(req.URL.Path) == gqlBasePath {
		ctx := initContextWithJWTClaims(req)
		ctx = context.WithValue(ctx, KeyLoaders, loaders)
		ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)
		req = req.WithContext(ctx)
		gqlHandler.ServeHTTP(res, req)
	}

	if os.Getenv("EXPOSE_PLAYGROUND_ENDPOINT") == "true" && path.Base(req.URL.Path) == "playground" {
		playgroundHandler := playground.Handler("GraphQL playground", gqlBasePath)
		ctx := initContextWithJWTClaims(req)
		ctx = context.WithValue(ctx, KeyLoaders, loaders)
		ctx = context.WithValue(ctx, KeyExecutableSchema, executableSchema)
		req = req.WithContext(ctx)
		if req.Method == "GET" {
			playgroundHandler(res, req)
		}
	}
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

// JWTClaims JWT Claims
type JWTClaims struct {
	jwtgo.StandardClaims
	Email  string    `json:"email"`
	Avatar string    `json:"avatar,omitempty"`
	Roles  []JWTRole `json:"roles,omitempty"`
	Scope  *string   `json:"scope,omitempty"`
}

// JWTRole defines a role
type JWTRole struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
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
