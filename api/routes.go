package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func HandlePostUser(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body User
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			SendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		id := uuid.New()

		slog.Info("user id", "id", id)

		db[id.String()] = body
		SendJSON(w, Response{Data: body}, http.StatusCreated)
	}
}

func HandleGetUsers(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var result []User
		for _, value := range db {
			result = append(result, value)
		}

		SendJSON(w, Response{Data: result}, http.StatusOK)
	}
}

func HandleGetUserById(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, ok := db[id]
		if !ok {
			http.Error(w, "not found user", http.StatusNotFound)
			return
		}

		SendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func HandleDeleteUserById(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, ok := db[id]
		if !ok {
			http.Error(w, "not found user", http.StatusNotFound)
			return
		}

		delete(db, id)

		SendJSON(w, Response{Data: user}, http.StatusOK)
	}
}

func HandleUpdateUserById(db map[string]User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		user, ok := db[id]
		if !ok {
			http.Error(w, "not found user", http.StatusNotFound)
			return
		}

		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			SendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if u.FirstName != "" {
			user.FirstName = u.FirstName
		}

		if u.LastName != "" {
			user.LastName = u.LastName
		}

		if u.Biography != "" {
			user.Biography = u.Biography
		}

		db[id] = user

		SendJSON(w, Response{Data: user}, http.StatusOK)
	}
}
