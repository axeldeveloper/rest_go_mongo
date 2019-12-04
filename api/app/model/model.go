package model

import (
	"github.com/globalsign/mgo/bson"
)

type Employee struct {
	_id    bson.ObjectId `bson:"_id,omitempty"`
	Name   string        `json:"name"`
	City   string        `json:"city"`
	Age    int           `json:"age"`
	Status bool          `json:"status"`
}

func (e *Employee) Disable() {
	e.Status = false
}

func (p *Employee) Enable() {
	p.Status = true
}