package middlewares

import (
	"fmt"
	"net/http"

	"github.com/huynhminhtufu/go-blog-be/app/lib"
)

// DemoMiddleware ...
func DemoMiddleWare() lib.Middleware {
	return func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("before wuhuuu middleware")
			h.ServeHTTP(w, r)
			fmt.Println("after wuhuuu middleware")
		}
	}
}
