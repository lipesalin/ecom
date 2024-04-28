package user

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/lipesalin/ecom/types"

)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("Deve falhar caso os dados da request sejam inválidos", func(t *testing.T) {
		payloadUser := types.RegisterUserPayload{
			Name:     "felipe",
			Email:    "",
			Password: "123",
		}

		// parse JSON
		userJSON, errParseJSON := json.Marshal(payloadUser)

		if errParseJSON != nil {
			log.Fatal(errParseJSON)
		}

		// [Request]
		request, errRequest := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJSON))

		if errRequest != nil {
			t.Fatal(errRequest)
		}

		recorder := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(recorder, request)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("expected status code %d", http.StatusBadRequest, recorder.Code)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
