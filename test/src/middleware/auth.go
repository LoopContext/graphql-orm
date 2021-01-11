package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/loopcontext/go-graphql-orm/test/src/utils"
	"github.com/rs/zerolog/log"

	"github.com/dgrijalva/jwt-go"
)

var (
	// TokenAuthHeaderName is a string in the header. Default value is "Bearer"
	TokenAuthHeaderName = "Bearer"

	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"
	TokenLookup = "query:token,cookie:jwt,header:Authorization"

	// ErrInvalidTokenOrClaims when HTTP status 403 is given
	ErrInvalidTokenOrClaims = errors.New("token or claims invalid")

	// ErrForbidden when HTTP status 403 is given
	ErrForbidden = errors.New("you don't have permission to access this resource")

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = errors.New("auth header is empty")

	// ErrEmptyAPIKeyHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAPIKeyHeader = errors.New("api key header is empty")

	// ErrMissingExpField missing exp field in token
	ErrMissingExpField = errors.New("missing exp field")

	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
	ErrInvalidAuthHeader = errors.New("auth header is invalid")

	// ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyQueryToken = errors.New("query token is empty")

	// ErrEmptyCookieToken can be thrown if authing with a cookie, the token cokie is empty
	ErrEmptyCookieToken = errors.New("cookie token is empty")

	// ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrEmptyParamToken = errors.New("parameter token is empty")

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")
)

// w.Header().Set("Access-Control-Allow-Origin", "*")

func authError(res http.ResponseWriter, err error) {
	res.WriteHeader(http.StatusUnauthorized)
	errEncode := json.NewEncoder(res).Encode(map[string]string{"message": "[Auth] error: " + err.Error()})
	if errEncode != nil {
		log.Err(fmt.Errorf("[Wrapped] errEncode: %w\n[Original]: %q", errEncode, err))
		return
	}
}

// if needed:
// func apiKeyFromHeader(req *http.Request, key string) (string, error) {
// 	apiKey := req.Header.Get(key)
// 	if apiKey == "" {
// 		return "", ErrEmptyAPIKeyHeader
// 	}
// 	return apiKey, nil
// }

func tokenFromQuery(req *http.Request, key string) (string, error) {
	token := req.URL.Query().Get(key)
	if token == "" {
		return "", ErrEmptyQueryToken
	}
	return token, nil
}

func tokenFromCookie(req *http.Request, key string) (string, error) {
	cookie, err := req.Cookie(key)
	if cookie.String() == "" {
		return "", ErrEmptyCookieToken
	}
	return cookie.Value, err
}

func jwtFromHeader(req *http.Request, key string) (string, error) {
	authHeader := req.Header.Get(key)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == TokenAuthHeaderName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

// ParseToken parse jwt token from gin context
func ParseToken(req *http.Request) (t *jwt.Token, err error) {
	var token string
	methods := strings.Split(TokenLookup, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			token, err = jwtFromHeader(req, v)
		case "query":
			token, err = tokenFromQuery(req, v)
		case "cookie":
			token, err = tokenFromCookie(req, v)
		}
	}
	if err != nil {
		return nil, err
	}
	SigningAlgorithm := utils.MustGet("AUTH_JWT_SIGNING_ALGORITHM")
	Key := []byte(utils.MustGet("AUTH_JWT_SECRET"))
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(SigningAlgorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		return Key, nil
	})
}

// AuthJWT auth middleware struct
type AuthJWT struct {
	Path string
}

// Middleware auth func, called each request
func (a *AuthJWT) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		log.Debug().Msgf("[JWTAuth.Middleware] Applied to path: %s - current path: %s", a.Path, req.RequestURI)
		if !strings.HasPrefix(req.RequestURI, a.Path) {
			next.ServeHTTP(res, req)
			return
		}
		t, err := ParseToken(req)
		if err != nil {
			authError(res, err)
		} else {
			if claims, ok := t.Claims.(jwt.MapClaims); t.Valid && ok {
				if claims["exp"] != nil {
					// You might use these fields to filter out token requests
					// issuer := claims["iss"].(string)
					// userid := claims["jti"].(string)
					// email := claims["email"].(string)
					// if claims["aud"] != nil {
					// audiences := claims["aud"].(interface{})
					// log.Debug().Msgf("audiences: %s", audiences)
					// }
					// if claims["alg"] != nil {
					// algo := claims["alg"].(string)
					// log.Debug().Msgf("algo: %s", algo)
					// }
					next.ServeHTTP(res, req)
				} else {
					authError(res, ErrMissingExpField)
				}
			} else {
				authError(res, ErrInvalidTokenOrClaims)
			}
		}
	})
}
