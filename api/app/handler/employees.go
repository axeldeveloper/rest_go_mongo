package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"rest.gorn.mongo/api/app/model"
)

func GetAllEmployees(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	employees := []model.Employee{}
	c := session.DB("mestre").C("employees")
	_ = c.Find(nil).All(&employees)

	respondJSON(w, http.StatusOK, employees)
}

func CreateEmployee(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	employee := model.Employee{}

	Jerr := json.NewDecoder(r.Body).Decode(&employee)

	//fmt.Fprintf(w, "%s %s is %d years old!", employee.Name, employee.City, employee.Age)

	if Jerr != nil {
		respondError(w, http.StatusBadRequest, Jerr.Error())
		return
	}
	defer r.Body.Close()
	c := session.DB("mestre").C("employees")
	err := c.Insert(&employee)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, employee)
}

func GetEmployee(session *mgo.Session, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	fmt.Println("Var Nome ", name)

	employee := getEmployeeOr404(session, name, w, r)
	if employee == nil {
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

/*
func UpdateEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&employee); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func DeleteEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	if err := db.Delete(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}

func DisableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	employee.Disable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}

func EnableEmployee(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	employee := getEmployeeOr404(db, name, w, r)
	if employee == nil {
		return
	}
	employee.Enable()
	if err := db.Save(&employee).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, employee)
}
*/

// getEmployeeOr404 gets a employee instance if exists, or respond the 404 error otherwise
func getEmployeeOr404(session *mgo.Session, name string, w http.ResponseWriter, r *http.Request) *model.Employee {

	employee := model.Employee{}
	c := session.DB("mestre").C("employees")
	//err := c.First(&employee, model.Employee{Name: name} )
	conditions := bson.M{"name": name}
	err := c.Find(conditions).One(&employee)
	fmt.Println("employee ", err)
	if err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		//fmt.Println("Erro 404 ", err)
		log.Fatal(err)
		//return nil
	}

	return &employee
}

/*

func GetByID(id string) []ToDoItem {
	var result ToDoItem
	var res []ToDoItem
	_ = c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&result)
	res = append(res, result)
	return res
}


*/
