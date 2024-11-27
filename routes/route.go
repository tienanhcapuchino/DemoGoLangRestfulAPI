package routes

import (
	"net/http"
	store "rest-api-demo/interfaces/users"
	"rest-api-demo/types"
	"rest-api-demo/utils"

	"github.com/gorilla/mux"
)

type Handler struct {
	store store.UserStore
}

func NewHandler(store store.UserStore) *Handler {
	return &Handler{
		store: store,
	}
}

func NewHandlerNoParam() *Handler{
	return &Handler{}
}


func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/get-all-users", h.getAllUsers).Methods("GET")
}

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users [4]types.UserModel
	users[0] = types.UserModel{Name: "user1", Email: "user1@example.com", Age: 25}
	users[1] = types.UserModel{Name: "Bob", Email: "bob@example.com", Age: 30}
	users[2] = types.UserModel{Name: "Bob1", Email: "bob1@example.com", Age: 30}
	users[3] = types.UserModel{Name: "Bob2", Email: "bob2@example.com", Age: 30}

	utils.WriteJSON(w, http.StatusOK, users)
}