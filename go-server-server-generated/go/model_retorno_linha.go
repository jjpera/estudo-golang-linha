package swagger

type RetornoLinha struct {
	Codigo      int64   `json:"codigo,omitempty"`
	Descricao   string  `json:"descricao,omitempty"`
	Registros   int64   `json:"registros,omitempty"`
	Pagina      int64   `json:"pagina,omitempty"`
	QtdePagina  int64   `json:"qtdePagina,omitempty"`
	ListaLinhas []Linha `json:"listaLinhas,omitempty"`
}
