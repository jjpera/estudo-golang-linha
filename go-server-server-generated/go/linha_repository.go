package swagger

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LinhaRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "linhas"
)

func (m *LinhaRepository) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}
func (m *LinhaRepository) FindAll() ([]Linha, error) {
	var linhas []Linha
	err := db.C(COLLECTION).Find(bson.M{}).All(&linhas)
	return linhas, err
}
func (m *LinhaRepository) FindById(id string) (Linha, error) {
	var linha Linha
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&linha)
	return linha, err
}
func (m *LinhaRepository) Insert(linha Linha) error {
	err := db.C(COLLECTION).Insert(&linha)
	return err
}
func (m *LinhaRepository) Delete(linha Linha) error {
	err := db.C(COLLECTION).Remove(&linha)
	return err
}
func (m *LinhaRepository) Update(linha Linha) error {
	err := db.C(COLLECTION).UpdateId(linha.Id, &linha)
	return err
}
