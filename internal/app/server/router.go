package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go_test_learning/internal/app/model"
	"net/http"
	"strconv"
)

func NewRouter(s *APIServer) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user/list", getUserList(s))
	router.HandleFunc("/user/create", userCreate(s)).Methods("POST")
	router.HandleFunc("/user/update", userUpdate(s)).Methods("POST")
	return router
}

func userCreate(s *APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user.ID = primitive.ObjectID{}

		if err := user.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.Store.User().Create(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func userUpdate(s *APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := &model.User{}
		if err := json.NewDecoder(r.Body).Decode(user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := user.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.Store.User().Update(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

func getUserList(s *APIServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var page int64 = 1
		var perPage int64 = 10

		if pg, ok := r.URL.Query()["page"]; ok {
			p, err := strconv.ParseInt(pg[0], 10, 64)
			if err == nil {
				if p > 0 {
					page = p
				}
			}
		}

		if pp, ok := r.URL.Query()["perPage"]; ok {
			p, err := strconv.ParseInt(pp[0], 10, 64)
			if err == nil {
				if p > 0 {
					perPage = p
				}
			}
		}

		skip := (page - 1) * perPage

		users, err := s.Store.User().SelectUsers(skip, perPage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(users)
	}
}
