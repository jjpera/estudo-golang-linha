package swagger

import (
	"time"
)

// Historico de operações
type Historico struct {
	ID        string         `json:"id,omitempty"`
	Data      time.Time      `json:"data,omitempty"`
	Tipo      TipoHistorico `json:"tipo,omitempty"`
	Sistema   Sistema       `json:"sistema,omitempty"`
	Descricao string         `json:"descricao,omitempty"`
}
