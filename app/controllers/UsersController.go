package controllers

import (
	"net/http"

	"github.com/huynhminhtufu/go-blog-be/app/lib"
	"github.com/huynhminhtufu/go-blog-be/app/models"
)

// GetAllUsersHandler ...
func GetAllUsersHandler(w http.ResponseWriter, req *http.Request) {
	res := lib.Response{ResponseWriter: w}
	user := new(models.User)
	users := user.FetchAll()
	res.SendOK(users)
}

// CreateUserHandler ...
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := lib.Request{ResponseWriter: w, Request: r}
	res := lib.Response{ResponseWriter: w}

	user := new(models.User)
	req.GetJSONBody(user)

	if err := user.Save(); err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendCreated(user)
}

// GetUserByIDHandler ...
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	req := lib.Request{ResponseWriter: w, Request: r}
	res := lib.Response{ResponseWriter: w}

	id, _ := req.GetVarID()
	user := models.User{
		ID: id,
	}

	if err := user.FetchByID(); err != nil {
		res.SendNotFound()
		return
	}

	res.SendOK(user)
}

// UpdateUserHandler ...
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := lib.Request{ResponseWriter: w, Request: r}
	res := lib.Response{ResponseWriter: w}

	id, _ := req.GetVarID()

	user := new(models.User)
	req.GetJSONBody(user)
	user.ID = id

	if err := user.Save(); err != nil {
		res.SendBadRequest(err.Error())
		return
	}

	res.SendOK(user)
}

// DeleteUserHandler ...
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	req := lib.Request{ResponseWriter: w, Request: r}
	res := lib.Response{ResponseWriter: w}

	id, _ := req.GetVarID()
	user := models.User{
		ID: id,
	}

	if err := user.Delete(); err != nil {
		res.SendNotFound()
		return
	}

	res.SendNoContent()
}
