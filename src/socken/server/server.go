package server

import (
	"fmt"
	"net/http"
	"net/url"
)

var StaticServer http.Handler = http.FileServer(http.Dir("./assets"))

var NewPlayer = func(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	names := r.PostForm["name"]
	if len(names) != 1 || len(names[0]) < 1 {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	name := url.QueryEscape(names[0])

	http.Redirect(w, r, fmt.Sprintf("/player/%s", name), http.StatusFound)
}

var Player = func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello: %q", r.URL)
}
