package repository

import (
	"../models"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// LinhaRepository entidade do repositorio
type LinhaRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

// COLLECTION tabela do banco mongo
const (
	COLLECTION = "linhas"
)

// Connect conecta no banco
func (m *LinhaRepository) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(m.Database)
}

// Insert Adiciona linha telefonica
func (m *LinhaRepository) Insert(linha models.Linha) error {
	err := db.C(COLLECTION).Insert(&linha)
	return err
}

// Find Busca Linha Telefonica
func (m *LinhaRepository) Find(Ddd string, Numero string, Pagina int, QtdePagina int) ([]models.Linha, error) {
	filter := bson.M{}
	if Ddd != "" {
		filter["ddd"] = Ddd
	}

	if Numero != "" {
		filter["numero"] = Numero
	}

	var linhas []models.Linha
	err := db.C(COLLECTION).Find(filter).Sort("ddd", "numero").Skip(Pagina * QtdePagina).Limit(QtdePagina).All(&linhas)
	return linhas, err
}

// Count busca quantidade de linhas existentes
func (m *LinhaRepository) Count(Ddd string, Numero string) (int, error) {
	filter := bson.M{}
	if Ddd != "" {
		filter["ddd"] = Ddd
	}

	if Numero != "" {
		filter["numero"] = Numero
	}

	return db.C(COLLECTION).Find(filter).Count()
}

// FindByID busca pelo Id 
func (m *LinhaRepository) FindByID(id string) (models.Linha, error) {
	var linha models.Linha
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&linha)
	return linha, err
}

// Delete Remove atualiza linha telefonica
func (m *LinhaRepository) Delete(linha models.Linha) error {
	err := db.C(COLLECTION).Remove(&linha)
	return err
}

// Update atualiza linha telefonica
func (m *LinhaRepository) Update(linha models.Linha) error {
	err := db.C(COLLECTION).UpdateId(linha.ID, &linha)
	return err
}
