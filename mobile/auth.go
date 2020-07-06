package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/stretchr/objx"
	"os"
	"encoding/json"
	"github.com/stretchr/gomniauth"
	//"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	//"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
)

type Credential struct {
  Service string `json:service`
  ClientID string `json:clientID`
  Secret string `json:secret`
  Redirect string `json:redirect`
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		w.Header().Set("Location", "login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("認証プロバイダの取得に失敗しました:", provider, "-", err)
		}
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln("GetBeginAuthURL呼び出し中にエラーが発生しました", provider, "-", err)
		}
		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)

	case "callback":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("認証プロバイダの取得に失敗しました", err)
		}

		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln("認証を完了できませんでした", err)
		}
		user, err := provider.GetUser(creds)
		if err != nil {
			log.Fatalln("ユーザーの取得に失敗しました", err)
		}

		authCookieValue := objx.New(map[string]interface{}{
			"name": user.Name,
		}).MustBase64()
		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/",
		})
		w.Header()["Location"] = []string{"/register"}
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "不正なURLです")
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}


func credential(filename string){
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	file, err := os.Open(filename)
  if err != nil {
    log.Fatalln(err)
  }
  defer file.Close()
	var cred []*Credential
  decoder := json.NewDecoder(file)

  err =decoder.Decode(&cred)
  if err != nil {
    log.Fatalln(err)
  }


	clientID := cred[0].ClientID
	secret := cred[0].Secret
	redirect := cred[0].Redirect

	gomniauth.WithProviders(
		github.New(clientID, secret, redirect),
	)
}
