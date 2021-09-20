package controller

import (
	"backend/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetOneCollection(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	page, err := model.OneCollectionCategories(id)
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
