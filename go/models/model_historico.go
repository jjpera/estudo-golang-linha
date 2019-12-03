package models

import (
	"../enums"
	"time"
)

// Historico de operações
type Historico struct {
	ID        string                  `json:"id,omitempty"`
	Data      time.Time               `json:"data,omitempty"`
	Tipo      enums.TipoHistorico  `json:"tipo,omitempty"`
	Sistema   enums.Sistema        `json:"sistema,omitempty"`
	Descricao string                  `json:"descricao,omitempty"`
}
