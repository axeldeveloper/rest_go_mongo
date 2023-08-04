package handler

import (
	"encoding/json"
	"net/http"
	"rest_go_mongo/api/app/model"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
)

const (
	DBNAME     = "mestre"
	COLLECTION = "employees"
)

/***************************************************
	CRUDO MONGODB GOLANG
*****************************************************/
func GetAllEmployees(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	employees := []model.Employee{}
	_ = session.DB(DBNAME).C(COLLECTION).Find(nil).All(&employees)
	respondJSON(w, http.StatusOK, employees)
}
func GetEmployeeByID(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	employee, erro := FindById(session, id)
	if erro != nil {
		//respondError(w, http.StatusNotAcceptable, "Registro não encontado")
		respondError(w, http.StatusNotFound, erro.Error())
		//log.Fatal("Erro: ", erro)
		return
	}
	//if employee == nil {return}
	respondJSON(w, http.StatusOK, employee)
}
func GetEmployeeByName(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	employee, erro := Find_By(session, name)
	if erro != nil {
		//respondError(w, http.StatusNotAcceptable, "Registro não encontado")
		respondError(w, http.StatusNotFound, erro.Error())
		//log.Fatal(err)
		return
	}
	//if employee == nil {return}
	respondJSON(w, http.StatusOK, employee)
}
func CreateEmployee(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	employee := model.Employee{}
	Jerr := json.NewDecoder(r.Body).Decode(&employee)
	if Jerr != nil {
		respondError(w, http.StatusBadRequest, Jerr.Error())
		return
	}
	defer r.Body.Close()
	err := session.DB(DBNAME).C(COLLECTION).Insert(&employee)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}
func UpdateEmployee(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	employee, erro := FindId_ByID(session, id)
	//if employee == nil {}
	if erro != nil {
		respondError(w, http.StatusNotAcceptable, "Registro não encontado")
		return //log.Fatal("Erro: ", erro)
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	data := bson.M{"$set": bson.M{
		"name":   employee.Name,
		"city":   employee.City,
		"age":    employee.Age,
		"status": employee.Status}}

	err := session.DB(DBNAME).C(COLLECTION).Update(selector, data)
	if err != nil {
		respondError(w, http.StatusMethodNotAllowed, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}
func DeleteEmployee(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	employee, erro := FindId_ByID(session, id)

	if erro != nil {
		respondError(w, http.StatusNotFound, erro.Error())
		return //log.Fatal("Erro: ", erro)
	}

	//if employee == nil {return}
	err := session.DB(DBNAME).C(COLLECTION).Remove(&employee)

	if err != nil {
		respondError(w, http.StatusNotFound, erro.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}
func DisableEmployee(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	employee, erro := FindId_ByID(session, id)
	if erro != nil {
		respondError(w, http.StatusNotFound, erro.Error())
		return //log.Fatal("Erro: ", erro)
	}
	employee.Disable()
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	//data 	 := bson.M{"$set": bson.M{"status": false}}
	_, err := session.DB(DBNAME).C(COLLECTION).Upsert(selector, &employee)
	if err != nil {
		respondError(w, http.StatusNotFound, erro.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}
func EnableEmployee(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	employee, erro := FindId_ByID(session, id)

	if erro != nil {
		respondError(w, http.StatusNotFound, erro.Error())
		return
	}

	employee.Enable()
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	//data 	 := bson.M{"$set": bson.M{"status": false}}
	_, err := session.DB(DBNAME).C(COLLECTION).Upsert(selector, &employee)
	if err != nil {
		respondError(w, http.StatusNotFound, erro.Error())
		return
	}

	respondJSON(w, http.StatusOK, employee)
}

/***************************************************
	FILTROS MONGODB GOLANG
*****************************************************/
// Find a movie by its id
func FindById(session *mgo.Session, id string) (*model.Employee, error) {
	var employee = model.Employee{}
	selector := bson.M{"_id": bson.ObjectIdHex(id)}
	err := session.DB(DBNAME).C(COLLECTION).Find(selector).One(&employee)
	//err := session.DB(DBNAME).C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&employee)
	return &employee, err
}
func FindId_ByID(session *mgo.Session, Id string) (*model.Employee, error) {
	employee := model.Employee{}
	c := session.DB(DBNAME).C(COLLECTION)
	err := c.FindId(bson.ObjectIdHex(Id)).One(&employee)
	return &employee, err
}
func Find_By(session *mgo.Session, name string) (*model.Employee, error) {
	employee := model.Employee{}
	selector := bson.M{"name": name}
	err := session.DB(DBNAME).C(COLLECTION).Find(selector).One(&employee)
	return &employee, err
}
func GetByID(session *mgo.Session, id string) []model.Employee {
	var result model.Employee
	var res []model.Employee
	_ = session.DB(DBNAME).C(COLLECTION).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	res = append(res, result)
	return res
}
