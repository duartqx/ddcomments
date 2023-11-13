package controllers

import (
	"encoding/json"
	"net/http"

	h "github.com/duartqx/ddcomments/application/http"
	s "github.com/duartqx/ddcomments/application/services"
	u "github.com/duartqx/ddcomments/domains/entities/user"
)

type UserController struct {
	userService *s.UserService
}

func NewUserController(userService *s.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var response *h.HttpResponse

	switch r.Method {
	case http.MethodPost:
		response = uc.post(r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if body, err := json.Marshal(response.Body); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(response.Status)
		w.Write(body)
	}
}

func (uc UserController) post(r *http.Request) *h.HttpResponse {

	var user *u.UserEntity

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return &h.HttpResponse{
			Body:   map[string]string{"error": "Json Decode Error"},
			Status: http.StatusBadRequest,
		}
	}

	if err := uc.userService.Create(user); err != nil {
		return &h.HttpResponse{
			Body:   err.Error(),
			Status: http.StatusBadRequest,
		}
	}

	return &h.HttpResponse{
		Body:   user.Clean(),
		Status: http.StatusCreated,
	}
}
