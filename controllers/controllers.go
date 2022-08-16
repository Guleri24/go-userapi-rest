package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Welcome(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Write([]byte("Welcome to ModelInfo"))
}
