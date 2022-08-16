package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CreateUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
}

func DeleteUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func UpdateUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {}

func GetUser(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Write([]byte("Welcome to ModelInfo"))

}

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rw.Write([]byte("Welcome to ModelInfo"))

}
