package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/huynhminhtufu/go-blog-be/app/controllers"
	"github.com/huynhminhtufu/go-blog-be/app/lib"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	res := lib.Response{ResponseWriter: w}
	res.SendOK("Go server is running, this is default page, also a notfound page.")
}

// NewRouter ...
func NewRouter() *mux.Router {

	// Create main router
	mainRouter := mux.NewRouter().StrictSlash(true)
	mainRouter.KeepContext = true

	// Handle 404
	mainRouter.NotFoundHandler = http.HandlerFunc(notFound)

	/**
	 * meta-data
	 */
	mainRouter.Methods("GET").Path("/api/info").HandlerFunc(controllers.GetAPIInfo)

	/**
	 * /users
	 */
	// usersRouter.HandleFunc("/", l.Use(c.GetAllUsersHandler, m.SaySomething())).Methods("GET")

	var apiVersion string = "/api/v1"

	// User routes
	mainRouter.Methods("GET").Path(apiVersion + "/users").HandlerFunc(controllers.GetAllUsersHandler)
	mainRouter.Methods("POST").Path(apiVersion + "/users").HandlerFunc(controllers.CreateUserHandler)
	mainRouter.Methods("GET").Path(apiVersion + "/users/{id}").HandlerFunc(controllers.GetUserByIDHandler)
	mainRouter.Methods("PUT").Path(apiVersion + "/users/{id}").HandlerFunc(controllers.UpdateUserHandler)
	mainRouter.Methods("DELETE").Path(apiVersion + "/users/{id}").HandlerFunc(controllers.DeleteUserHandler)

	return mainRouter
}
