package auth

import (
	"encoding/json"
	"fmt"
	"go/adv-demo/configs"
	"go/adv-demo/pkg/resp"
	"net/http"
	"net/mail"
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
		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			resp.ResponseJSON(w, err.Error(), http.StatusBadRequest)
			return
		}
		if payload.Email == "" {
			resp.ResponseJSON(w, "Email required", http.StatusUnauthorized)
			return
		}
		//match, _ := regexp.MatchString(`[A-Za-z0-9\._%+\-]+@[A-Za-z0-9\.\-]+\.[A-Za-z]{2,}`, payload.Email)
		mailAddress, err := mail.ParseAddress(payload.Email)
		fmt.Println(mailAddress.Address)
		if err != nil {
			resp.ResponseJSON(w, "Wrong email", http.StatusUnauthorized)
			return
		}
		if payload.Password == "" {
			resp.ResponseJSON(w, "Password required", http.StatusUnauthorized)
			return
		}
		fmt.Println(payload)
		data := LoginResponse{
			Token: "123",
		}
		resp.ResponseJSON(w, data, http.StatusOK)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("register")
	}
}
