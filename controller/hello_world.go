package controller

import "net/http"

func NewController() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helo world"))
	}
}
