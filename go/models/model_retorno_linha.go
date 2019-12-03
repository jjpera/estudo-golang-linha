package models

// RetornoLinha Entidade para retorno das linhas encontradas
type RetornoLinha struct {
	Codigo      int64   `json:"codigo,omitempty"`
	Descricao   string  `json:"descricao,omitempty"`
	Registros   int   `json:"registros,omitempty"`
	Pagina      int   `json:"pagina,omitempty"`
	QtdePagina  int   `json:"qtdePagina,omitempty"`
	ListaLinhas []Linha `json:"listaLinhas,omitempty"`
}
