package utils

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const JWT_SIGNATURE_KEY string = "1GN1T3CH"

type ApiKey struct {
	code string
	jwt.StandardClaims
}

type Token struct {
	Token string `json:"token"`
}

func CreateToken(code string) (*Token, error) {
	var err error

	ttl := 24 * time.Hour
	claims := ApiKey{
		code,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(ttl).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(JWT_SIGNATURE_KEY))
	if err != nil {
		return nil, err
	}

	return &Token{signedToken}, nil

}

var Bundle *i18n.Bundle

var authSecret []byte = []byte("")

func SetAuthSecret(val string) {
	authSecret = []byte(val)
}

var notAuth []string = []string{} //List of endpoints that doesn't require auth

func SetNoAuth(paths []string) {
	notAuth = append(notAuth, paths...)
}

type favContextKey string

func GetValuesCtx(ctx context.Context) (jwt.MapClaims, bool) {
	val, ok := ctx.Value(favContextKey("values")).(jwt.MapClaims)
	return val, ok
}

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestPath := r.URL.Path //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if strings.Contains(requestPath, value) {
				next.ServeHTTP(w, r)
				return
			}
		}

		//response := make(map[string]interface{})

		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			response := Message(false, "Missing auth token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-Type", "application/json")
			Respond(w, response)
			return
		}

		splitted := strings.Split(authorizationHeader, " ")
		if len(splitted) != 2 {
			response := Message(false, "Invalid/Malformed auth token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Header().Add("Content-Type", "application/json")
			Respond(w, response)
			return
		}

		tokenString := splitted[1] //Grab the token part, what we are truly interested in

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// authSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return authSecret, nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					response := Message(false, "Token is not valid")
					w.WriteHeader(http.StatusUnauthorized)
					w.Header().Add("Content-Type", "application/json")
					Respond(w, response)
					return
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					response := Message(false, "Token Expired")
					w.WriteHeader(http.StatusUnauthorized)
					w.Header().Add("Content-Type", "application/json")
					Respond(w, response)
					return
				} else {
					response := Message(false, "Malformed authentication token")
					w.WriteHeader(http.StatusUnauthorized)
					w.Header().Add("Content-Type", "application/json")
					Respond(w, response)
					return
				}
			} else {
				response := Message(false, "Malformed authentication token")
				w.WriteHeader(http.StatusUnauthorized)
				w.Header().Add("Content-Type", "application/json")
				Respond(w, response)
				return
			}
		}

		ctx := context.WithValue(r.Context(), favContextKey("values"), claims)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

func GetDbNameCtx(ctx context.Context) (string, bool) {
	values, isvalid := GetValuesCtx(ctx)
	if isvalid {
		dbname, ok := values["dbname"].(string)
		if ok {
			return dbname, true
		}
	}
	return "", false
}

func GetUserIdCtx(ctx context.Context) (float64, bool) {
	values, isvalid := GetValuesCtx(ctx)
	if isvalid {
		value, ok := values["user_id"].(float64)
		if ok {
			return value, true
		}

	}
	return 0, false
}

func GetDbHostCtx(ctx context.Context) (string, bool) {
	values, isvalid := GetValuesCtx(ctx)
	if isvalid {
		dbhost, ok := values["dbhost"].(string)
		if ok {
			return dbhost, true
		}
	}
	return "", false
}

func GetGrpIdCtx(ctx context.Context) (float64, bool) {
	values, isvalid := GetValuesCtx(ctx)
	if isvalid {
		grpId, ok := values["grp_id"].(float64)
		if ok {
			return grpId, true
		}
	}
	return 0, false
}

func GetJasperHostCtx(ctx context.Context) (string, bool) {
	values, isvalid := GetValuesCtx(ctx)
	if isvalid {
		jasperhost, ok := values["jasperhost"].(string)
		if ok {
			return jasperhost, true
		}
	}
	return "", false
}
