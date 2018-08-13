package main

import (
	"github.com/gorilla/mux"
	"github.com/jan4984/ev3play"
	"net/http"
	"github.com/gorilla/handlers"
)

func main(){
	r := mux.NewRouter()
	ev3play.RegisterMotorHandlers(r.PathPrefix("/motor").Subrouter())

	methods := [...]string{"GET","POST","CREATE","DELETE"}
	e := http.ListenAndServe(":8080", handlers.CORS(
		handlers.AllowedMethods(methods[:]),
	)(r))
	panic(e)
}
