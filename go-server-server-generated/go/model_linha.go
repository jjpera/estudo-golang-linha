package swagger

import (
	"gopkg.in/mgo.v2/bson"
)

type Linha struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Ddd    string        `json:"ddd,omitempty"`
	Numero string        `json:"numero,omitempty"`
}
