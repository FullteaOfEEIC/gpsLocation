package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"github.com/stretchr/gomniauth"
	//"github.com/stretchr/gomniauth/providers/facebook"
	//"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
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
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		//facebook.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/facebook"),
		//github.New("クライアントID", "秘密の値", "http://localhost:8080/auth/callback/github"),
		google.New("26512943495-mv1onnih74b1br9r6eh9scpjeq4po9kq.apps.googleusercontent.com", "2ej7Mg7fzmhaDGwzTOOtu75m", "http://localhost:8080/auth/callback/google"),
	)
	http.Handle("/register", MustAuth(&templateHandler{filename: "register.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
