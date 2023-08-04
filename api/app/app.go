package app

import (
	"fmt"
	"log"
	"net/http"

	"rest_go_mongo/api/app/handler"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *mgo.Session
}

// App initialize with predefined configuration
func (a *App) Initialize() {
	//session, err := mgo.Dial("mongodb://127.0.0.1:27017/mestre")
	session, err := mgo.Dial("mongodb://localhost:27017/mestre")
	//session, err := mgo.DialWithInfo(info)
	log.Println("Porta: ", err)

	if err != nil {
		//panic(err)
		log.Println("Não foi possível conectar ao mongo: ", err.Error())
		fmt.Println("", err)
		return
	}
	//defer session.Close()
	//error check on every access
	//session.SetSafe(&mgo.Safe{})
	session.SetMode(mgo.Monotonic, true)
	a.DB = session
	a.Router = mux.NewRouter()
	a.setRouters()
}

// Set all required routers
//	a.Get  ("/employees/{id:[0-9]+}", a.GetEmployeeByID)
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/employees", a.GetAllEmployees)
	a.Get("/employees/{id}", a.GetEmployeeByID)
	a.Get("/employees/nome/{name}", a.GetEmployeeByName)
	a.Post("/employees", a.CreateEmployee)
	a.Put("/employees/{id}", a.UpdateEmployee)
	a.Delete("/employees/{id}", a.DeleteEmployee)
	a.Put("/employees/{id}/disable", a.DisableEmployee)

	a.Put("/employees/{id}/enable", a.EnableEmployee)
}

// GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

/******************************************
// Handlers to manage Employee Data
*******************************************/
func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r) // ok
}
func (a *App) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployeeByID(a.DB, w, r) // ok
}
func (a *App) GetEmployeeByName(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployeeByName(a.DB, w, r) // ok
}
func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r) // ok
}
func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.UpdateEmployee(a.DB, w, r)
}
func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DeleteEmployee(a.DB, w, r)
}
func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.DisableEmployee(a.DB, w, r)
}

func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	handler.EnableEmployee(a.DB, w, r)
}

// Run the app on it's router
// Executando o aplicativo no roteador
func (a *App) Run(host string) {

	log.Println("Porta: ", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
