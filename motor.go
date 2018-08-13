package ev3play

import (
	"net/http"
	"github.com/gorilla/mux"
	"sync"
	"github.com/ev3go/ev3dev"
)

var lock sync.Mutex
var motors = make(map[string]*ev3dev.TachoMotor)

func RegisterMotorHandlers(router *mux.Router){
	router.HandleFunc("/{port}/{driverName}", create).Methods("CREATE")
	router.HandleFunc("/{port}", deletee).Methods("DELETE")
	router.HandleFunc("/{port}/read/{property}", read).Methods("GET")
	router.HandleFunc("/{port}/write/{property}/{value}", write).Methods("POST")
}

func create(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	port := vars["port"]
	lock.Lock()
	defer lock.Unlock()
	if _,ok:=motors[port];ok{
		http.Error(w, "already created", 400)
		return
	}

	motor,err := ev3dev.TachoMotorFor(port, vars["driverName"])
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	motors[port] = motor
}

func read(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	port := vars["port"]
	lock.Lock()
	defer lock.Unlock()

	if m,ok:=motors[port];!ok{
		http.Error(w, "not created", 400)
		return
	}else{
		v,e := attributeOf(m.Path(), m.String(), vars["property"])
		if e != nil{
			http.Error(w, e.Error(), 500)
			return
		}
		w.Write([]byte(v))
	}
}

func write(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	port := vars["port"]
	lock.Lock()
	defer lock.Unlock()

	if m,ok:=motors[port];!ok{
		http.Error(w, "not created", 400)
		return
	}else{
		e := setAttributeOf(m.Path(), m.String(), vars["property"], vars["value"])
		if e != nil {
			http.Error(w, e.Error(), 500)
			return
		}
	}
}

func deletee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	port := vars["port"]
	lock.Lock()
	defer lock.Unlock()

	if _,ok:=motors[port];!ok{
		http.Error(w, "not created", 400)
		return
	}

	delete(motors, port)
}