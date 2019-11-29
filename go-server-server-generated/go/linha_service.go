package swagger

import (
	"log"
)

type LinhaService struct {
}

var linhaRepository LinhaRepository

func (m *LinhaService) Insert(linha Linha) (retornoLinha RetornoLinha) {
	log.Printf("Insert linha")
	err := linhaRepository.Insert(linha)
	if err != nil {
		retornoLinha.Codigo = 500
		retornoLinha.Descricao = err.Error()
		log.Printf("Insert error: $s", err.Error())
	} else {
		retornoLinha.Codigo = 0
		retornoLinha.Pagina = 0
		retornoLinha.QtdePagina = 1
		retornoLinha.Registros = 1
		retornoLinha.ListaLinhas = []Linha{linha}
		log.Printf("Insert success: $s", linha)
	}

	return retornoLinha
}

func (m *LinhaService) Get(ddd string, numero string, pagina int64, qtdePagina int64) (retornoLinha RetornoLinha) {
	log.Printf("Get linha")
	var linhas []Linha
	linhas, err := linhaRepository.FindAll()
	if err != nil {
		retornoLinha.Codigo = 500
		retornoLinha.Descricao = err.Error()
		log.Printf("Get error: $s", err.Error())
	} else {
		retornoLinha.Codigo = 0
		retornoLinha.Pagina = 1
		retornoLinha.QtdePagina = 1
		retornoLinha.Registros = 1
		retornoLinha.ListaLinhas = linhas
		log.Printf("Get success")
	}

	return retornoLinha
}
