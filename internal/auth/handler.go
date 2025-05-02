package auth

import (
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/req"
	"go/adv-demo/pkg/resp"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(body)
		data := LoginResponse{
			Token: "123",
		}
		resp.ResponseJSON(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(body)
		data := RegisterResponse{
			Message: "Вы успешно зарегистрировались",
		}
		resp.ResponseJSON(w, data, http.StatusOK)
	}
}
