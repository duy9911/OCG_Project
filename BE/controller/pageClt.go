package controller

import (
	"backend/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOnePage(w http.ResponseWriter, r *http.Request) {

	fmt.Println("start")
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	page, err := model.OnePageCollections(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(page)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //set header type
	fmt.Fprintf(w, "%s\n", uj)
	w.WriteHeader(http.StatusOK) // 200
}

//
func GetAllPages(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	pages, err := model.AllPageCollections()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	uj, err := json.Marshal(pages)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
