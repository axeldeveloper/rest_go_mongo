package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/gorilla/mux"
	"rest_go_mongo/api/app/handler"
)

const (
	hosts      = "dsxxxxx.XXXX.com:61375"
	database   = "mestre"
	username   = "XXXXX"
	password   = "XXXXXXx@193"
	collection = "employes"
)

// App has router and db instances
type App struct {
	Router *mux.Router
	DB     *mgo.Session 
}

// App initialize with predefined configuration
func (a *App) Initialize() {

		//fmt.Println("app config:", a)

		/*
		info := &mgo.DialInfo{
			Addrs:    []string{hosts},
			Timeout:  60 * time.Second,
			Database: database,
			Username: username,
			Password: password,
		}
		*/
		session, err := mgo.Dial("mongodb://localhost:27017/mestre")
	
	
		//session, err := mgo.DialWithInfo(info)
	if err != nil {
		//panic(err)
		log.Println("Não foi possível conectar ao mongo: ", err.Error())
		fmt.Println("", err)
		return
	}

	session.SetMode(mgo.Monotonic, true)
	a.DB = session
	a.Router = mux.NewRouter()
	a.setRouters()

}

func (a *App) Initialize2() {

	a.Router = mux.NewRouter()
	a.setRouters()

}


// Set all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Get("/employees", a.GetAllEmployees)
	a.Post("/employees", a.CreateEmployee)
	a.Get("/employees/{name}", a.GetEmployee)
	a.Put("/employees/{title}", a.UpdateEmployee)
	a.Delete("/employees/{title}", a.DeleteEmployee)
	a.Put("/employees/{title}/disable", a.DisableEmployee)
	a.Put("/employees/{title}/enable", a.EnableEmployee)
}

/// r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler).Name("article")

// Wrap the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Wrap the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// Wrap the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Wrap the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

/******************************************
// Handlers to manage Employee Data
*******************************************/

func (a *App) GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	handler.GetAllEmployees(a.DB, w, r)
}

func (a *App) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	handler.CreateEmployee(a.DB, w, r)
}

func (a *App) GetEmployee(w http.ResponseWriter, r *http.Request) {
	handler.GetEmployee(a.DB, w, r)
}

func (a *App) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	///handler.UpdateEmployee(a.DB, w, r)
}

func (a *App) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	//handler.DeleteEmployee(a.DB, w, r)
}

func (a *App) DisableEmployee(w http.ResponseWriter, r *http.Request) {
	//handler.DisableEmployee(a.DB, w, r)
}

func (a *App) EnableEmployee(w http.ResponseWriter, r *http.Request) {
	//handler.EnableEmployee(a.DB, w, r)
}


// Run the app on it's router
// Executando o aplicativo no roteador
func (a *App) Run(host string) {
	
	log.Println("Porta: ", host ) 
	log.Fatal(http.ListenAndServe(host, a.Router))
}
