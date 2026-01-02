// Package api
package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName 	string `json:"lastName"`
	Biography string `json:"biography"`
}

type Response struct {
	Data any `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func sendJSON(w http.ResponseWriter, res Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(res)
	if err != nil {
		slog.Error("ocorreu um erro no json.Marshal da resposta", "error", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
	}

	w.WriteHeader(status)
	w.Write(data)
}

func CreateHandler(db map[uuid.UUID]User) http.Handler {
	handler := chi.NewMux()

	handler.Use(middleware.Recoverer)
	handler.Use(middleware.RequestID)
	handler.Use(middleware.Logger)

	handler.Get("/users/{uuid}", getUsers(db))

	return handler
}


func getUsers(db map[uuid.UUID]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uuidParam := chi.URLParam(r, "uuid")
		id, err := uuid.Parse(uuidParam)
		if err != nil {
			sendJSON(w, Response{Error: "uuid invalido"}, http.StatusBadRequest)
		}

		user, ok := db[id]
		if !ok {
			sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
		}

		sendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
