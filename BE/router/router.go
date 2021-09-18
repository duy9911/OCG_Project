package router

import (
	"backend/controller"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/books", http.StatusSeeOther)

}
func HandleRequests() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/controlers", controller.GetProducts)
	log.Fatal(http.ListenAndServe(":9922", nil))
}
