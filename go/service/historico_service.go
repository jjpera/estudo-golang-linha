package service

import (
	"../models"
	"log"
	"bytes"
    "net/http"
	"encoding/json"
)

// HistoricoService entidade
type HistoricoService struct {
}

// Historicourl Variavel com a URL do serviço
var Historicourl string

// SetHistoricourl seta o endereço do serviço
func  (m *HistoricoService) SetHistoricourl(url string) {
	Historicourl = url
}

// Insert adiciona historico
func (m *HistoricoService) Insert(historico models.Historico) error {
	log.Println(historico)
	log.Println(Historicourl)

	jsonValue, _ := json.Marshal(historico)
	response, err := http.Post(Historicourl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(response)
	return nil
}