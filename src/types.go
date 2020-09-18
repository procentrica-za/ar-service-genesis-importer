package main

import (
	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
}
type Config struct {
	CRUDHost        string
	CRUDPort        string
	ListenServePort string
}

type toAssetRegisterResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
