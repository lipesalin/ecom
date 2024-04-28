package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"github.com/lipesalin/ecom/service/auth"
	"github.com/lipesalin/ecom/types"
	"github.com/lipesalin/ecom/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Teste")
}

func (h *Handler) handleRegister(w http.ResponseWriter, request *http.Request) {
	// JSON da request
	var payloadUser types.RegisterUserPayload
	errParseJSON := utils.ParseJSON(request, &payloadUser)

	if errParseJSON != nil {
		utils.WriteError(w, http.StatusBadRequest, errParseJSON)
	}

	// validar request
	errValidator := utils.Validate.Struct(payloadUser)
	if errValidator != nil {
		errorsValidation := errValidator.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, errorsValidation)
		return
	}

	// verifica se o usuário existe
	_, errFindUser := h.store.GetUserByEmail(payloadUser.Email)

	if errFindUser != nil {
		utils.WriteError(w, http.StatusBadRequest, errFindUser)
		return
	}

	// Passa o hash na senha
	hashedPassword, errHashPassword := auth.HashPassword(payloadUser.Password)

	if errHashPassword != nil {
		utils.WriteError(w, http.StatusInternalServerError, errHashPassword)
	}

	// caso não exista, criar novo usuário
	errCreateUser := h.store.CreateUser(types.User{
		Name:     payloadUser.Name,
		Email:    payloadUser.Email,
		Password: hashedPassword,
	})

	if errCreateUser != nil {
		utils.WriteError(w, http.StatusInternalServerError, errCreateUser)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
