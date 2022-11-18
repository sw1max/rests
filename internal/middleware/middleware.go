package middleware

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"rests.com/internal/authdb"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, req)
	})

  // UserContextKey is the key in a request's context used to check if the request
  // has an authenticated user. The middleware will set the value of this key to
  // the username, if the user war properly authenticated with a password.
  const UserContextKey = "user"

  // BasicAuth is middleware that verifies the request has appropriate basic auth
  // set up with a user:password pair verified by authdb.
  func BasicAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
      user, pass, ok := req.BasicAuth()
      if ok && authdb.VerifyUserPass(user, pass) {
        newctx := context.WithValue(req.Context(), UserContextKey, user)
        next.ServeHTTP(w, req.WithContext(newctx))
      } else {
        w.Header().Set("WWW-Authenticate", `Basic realm="api"`)
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
      }
    })
  }
}
