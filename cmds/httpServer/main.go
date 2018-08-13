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

	e := http.ListenAndServe(":8080", handlers.CORS()(r))
	panic(e)
}
