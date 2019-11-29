package swagger

import (
	"time"
)

type Historico struct {
	Id        string         `json:"id,omitempty"`
	Data      time.Time      `json:"data,omitempty"`
	Tipo      *TipoHistorico `json:"tipo,omitempty"`
	Sistema   *Sistema       `json:"sistema,omitempty"`
	Descricao string         `json:"descricao,omitempty"`
}
