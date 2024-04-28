package user

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/lipesalin/ecom/types"
	"github.com/lipesalin/ecom/utils"

)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, request *http.Request) {
	// JSON da request
	var payloadUser types.RegisterUserPayload
	err := utils.ParseJSON(request, payloadUser)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// verifica se o usuário existe
	// caso não exista, criar novo usuário
}
