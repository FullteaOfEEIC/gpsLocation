package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates",
				t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	credential("credential.json")
	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/register", &templateHandler{filename: "register.html"})
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.Handle("/show", &templateHandler{filename: "show.html"})
	http.HandleFunc("/auth", loginHandler)
	http.HandleFunc("/register_core", gpsRegisterHandler)
	http.HandleFunc("/show_core", gpsShowHandler)

	crt := "/ssl/dblec.crt"
	key := "/ssl/dblec.key"
	if err := http.ListenAndServeTLS(":8080", crt, key, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
