package models

// Retorno gerar
type Retorno struct {
	Codigo     int64  `json:"codigo,omitempty"`
	Descricao  string `json:"descricao,omitempty"`
	Registros  int32  `json:"registros,omitempty"`
	Pagina     int32  `json:"pagina,omitempty"`
	QtdePagina int32  `json:"qtdePagina,omitempty"`
}
