package swagger

import (
	"log"
	"bytes"
    "net/http"
	"encoding/json"
)

// HistoricoService entidade
type HistoricoService struct {
}

// Insert adiciona historico
func (m *HistoricoService) Insert(historico Historico) error {
	jsonValue, _ := json.Marshal(historico)
	response, err := http.Post("http:/localhost:8081/historico/historico", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	log.Println(response)
	return nil
}