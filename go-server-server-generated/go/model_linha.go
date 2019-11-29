package swagger

import (
	"gopkg.in/mgo.v2/bson"
)

// Linha Telefonica
type Linha struct {
	ID     bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Ddd    string        `json:"ddd,omitempty"`
	Numero string        `json:"numero,omitempty"`
}
