package main

import (
	_ "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/unrolled/render"
)

var (
	ren = render.New(render.Options{
		IsDevelopment: true,
	})
)

func GetHandlers() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", rootViewHandler)
	router.HandleFunc("/7d8830b69114244387aa33c18fc4b0cf", addViewHandler)
	router.HandleFunc("/add", addPostHandler)
	router.HandleFunc("/read/{title}", readHandler)
	router.HandleFunc("/{[a-z]+}", notFoundHandler)
	return router
}

func rootViewHandler(w http.ResponseWriter, r *http.Request) {
	editons := GetAllMagazines()
	ren.HTML(w, http.StatusOK, "all", editons)
}

func addViewHandler(w http.ResponseWriter, r *http.Request) {
	ren.HTML(w, http.StatusOK, "new", nil)
}

func addPostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	PanicIf(err)

	edition := new(Edition)
	decoder := schema.NewDecoder()

	err = decoder.Decode(edition, r.PostForm)
	PanicIf(err)
	ProcessEdition(edition)

	SaveToDb(edition)

	http.Redirect(w, r, "/read/"+edition.Title, http.StatusFound)
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	edition := GetEditionByTitle(title)
	if edition.Title != "" {
		ren.HTML(w, http.StatusOK, "index", edition)
	} else {
		notFoundHandler(w, r)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/404.html", http.StatusTemporaryRedirect)
}
