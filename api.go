package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type APIServer struct {
	listerAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listerAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlerFunc(s.handleGetAccount))

	log.Println("JSON API server runnning in port:", s.listerAddr)

	http.ListenAndServe(s.listerAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		fmt.Println("get")
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		fmt.Println("post")
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "DELETE" {
		fmt.Println("delete")
		return s.handleDeleteAccount(w, r)
	}

	fmt.Println("METHOD NYA GA BENAR...")
	return nil
}
func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	//account := NewAccount("Rey", "Dan")
	id := mux.Vars(r)["id"]
	fmt.Println(id)

	return WriteJSON(w, http.StatusOK, &Account{})
}
func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	//w.Header().Set("Content-Type", "application/json") // mengganti nilai header yang sudah ada dengan nilai yang baru
	w.Header().Add("Content-Type", "application/json") // menambah header kalau memakai add.

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			_ = WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
