package main

import (
	"github.com/gorilla/mux"
	"github.com/jan4984/ev3play"
	"net/http"
	"github.com/gorilla/handlers"
	"os"
)

func main(){
	os.Unsetenv("http_proxy")
	r := mux.NewRouter()
	ev3play.RegisterMotorHandlers(r.PathPrefix("/motor").Subrouter())
	ev3play.RegisterSoundHandlers(r.PathPrefix("/sound").Subrouter())

	methods := [...]string{"GET","POST","CREATE","DELETE"}
	e := http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedMethods(methods[:]),
	)(r))
	panic(e)
}
