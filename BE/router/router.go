package router

import (
	"backend/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/api/producst", controller.GetProducts)

	r.HandleFunc("/api/home/pages/collections", controller.GetAllPages)
	r.HandleFunc("/api/home/page/{id:[0-9]+}/collections", controller.GetOnePage)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9911", r))
}
