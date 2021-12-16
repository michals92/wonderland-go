package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/michals92/wonderland-go/errors"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(response http.ResponseWriter, request *http.Request)) {

	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) AUTH_GET(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, JwtVerify(http.HandlerFunc(f))).Methods("GET")
}

func (*muxRouter) AUTH_POST(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, JwtVerify(http.HandlerFunc(f))).Methods("POST")
}

func (*muxRouter) POST(uri string, f func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Http server running on port %v \n", port)
	http.ListenAndServe(":"+port, muxDispatcher)
}

func JwtVerify(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("Authorization") //Grab the token from the header

		header = strings.TrimSpace(header)

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(errors.ServiceError{Message: "missing jwt header"})
			return
		}

		jwt, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("ACCESS_SECRET")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(errors.ServiceError{Message: "error parsing jwt"})
			return
		}

		//TODO: - get parameters from JWT token
		ctx := context.WithValue(r.Context(), "userId", jwt)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
