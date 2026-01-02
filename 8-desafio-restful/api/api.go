// Package api
package api

import (
	"8-desafio-restful/database"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

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

func CreateHandler() http.Handler {
	handler := chi.NewMux()

	handler.Use(middleware.Recoverer)
	handler.Use(middleware.RequestID)
	handler.Use(middleware.Logger)

	handler.Get("/users", getUsers)
	handler.Post("/users", postUser)
	handler.Get("/users/{uuid}", getUser)
	handler.Put("/users/{uuid}", putUser)
	handler.Delete("/users/{uuid}", deleteUser)

	return handler
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := database.FindAll()

	sendJSON(w, Response{Data: users}, http.StatusOK)
}

func postUser(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10000)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			sendJSON(w, Response{Error: "body too large"}, http.StatusRequestEntityTooLarge)
			return
		} else {
			sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		}
	}

	var user database.User
	if err := json.Unmarshal(data, &user); err != nil {
		sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		return
	}

	id, user := database.Insert(user)	

	sendJSON(w, Response{Data: makeResponse(id, user)}, http.StatusCreated )
}

func getUser(w http.ResponseWriter, r *http.Request) {
	uuidParam := chi.URLParam(r, "uuid")
	id, err := uuid.Parse(uuidParam)
	if err != nil {
		sendJSON(w, Response{Error: "uuid invalido"}, http.StatusBadRequest)
		return
	}

	user, ok := database.FindByID(id)
	if !ok {
		sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
		return
	}

	sendJSON(w, Response{Data: makeResponse(id, user)}, http.StatusOK)
}

func putUser(w http.ResponseWriter, r *http.Request) {
	uuidParam := chi.URLParam(r, "uuid")
	id, err := uuid.Parse(uuidParam)
	if err != nil {
		sendJSON(w, Response{Error: "uuid invalido"}, http.StatusBadRequest)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10000)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			sendJSON(w, Response{Error: "body too large"}, http.StatusRequestEntityTooLarge)
			return
		} else {
			sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
		}
	}

	var user database.User
	if err := json.Unmarshal(data, &user); err != nil {
		sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		return
	}

	user, ok := database.Update(id, user)
	if !ok {
		sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
		return
	}

	sendJSON(w, Response{Data: makeResponse(id, user)}, http.StatusOK)
}

func makeResponse(id uuid.UUID, user database.User) map[string]any {
	response := map[string]any{
		"id": id,
		"user": user,
	}

	return response
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	uuidParam := chi.URLParam(r, "uuid")
	id, err := uuid.Parse(uuidParam)
	if err != nil {
		sendJSON(w, Response{Error: "uuid invalido"}, http.StatusBadRequest)
		return
	}

	ok := database.Delete(id)
	if !ok {
		sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
		return
	}

	sendJSON(w, Response{Data: ok}, http.StatusOK)
}


