package swagger

// RetornoHistorico retorno do serviço
type RetornoHistorico struct {
	Codigo          int64       `json:"codigo,omitempty"`
	Descricao       string      `json:"descricao,omitempty"`
	Registros       int32       `json:"registros,omitempty"`
	Pagina          int32       `json:"pagina,omitempty"`
	QtdePagina      int32       `json:"qtdePagina,omitempty"`
	ListaHistoricos []Historico `json:"listaHistoricos,omitempty"`
}
