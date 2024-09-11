package api

import "net/http"

func PostSigninHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am signed in!"))

}

func PostRefreshHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I am refreshed!"))
}
