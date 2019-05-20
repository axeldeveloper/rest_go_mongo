package handler

/*
import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"rest.gorn.mongo/api/app/model"
)

func (u *model.Employee) Create(mgosession *mgo.Session, database, collection string) error {
	err := mgosession.DB(database).C(collection).Insert(u)
	if err != nil {
		return err
	}
	return err
}
func (u *model.Employee) Read(mgosession *mgo.Session, database, collection string, selector bson.M) (*[]Employee, error) {
	var results []User
	err := mgosession.DB(database).C(collection).Find(selector).All(&results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}
func (u *model.Employee) Update(mgosession *mgo.Session, database, collection string, selector, change bson.M) error {
	err := mgosession.DB(database).C(collection).Update(selector, bson.M{"$set": change})
	if err != nil {
		return err
	}
	return nil
}
func (u *model.Employee) Delete(mgosession *mgo.Session, database, collection string, selector bson.M) (*mgo.ChangeInfo, error) {
	info, err := mgosession.DB(database).C(collection).RemoveAll(selector)
	if err != nil {
		return nil, err
	}
	return info, nil
}

*/
