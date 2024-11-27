package routes

import (
	"net/http"
	store "rest-api-demo/interfaces/users"
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
	utils.WriteJSON(w, http.StatusOK, map[string]string{"login": "success"})
}