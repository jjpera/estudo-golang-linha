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

// Insert adiciona historico
func (m *HistoricoService) Insert(historico models.Historico) error {
	log.Println(historico)

	jsonValue, _ := json.Marshal(historico)
	response, err := http.Post("http://localhost:8081/historico/historico", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(response)
	return nil
}