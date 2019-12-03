package service

import (
	"../enums"
	"../models"
	"../repository"
	"log"
)

// LinhaService serviço para linha
type LinhaService struct {
}

var linhaRepository repository.LinhaRepository
var historicoService HistoricoService

// Insert método para buscar as linhas telefonicas
func (m *LinhaService) Insert(linha models.Linha) (retornoLinha models.RetornoLinha) {
	log.Println("Insert linha")
	err := linhaRepository.Insert(linha)
	if err != nil {
		retornoLinha.Codigo = 500
		retornoLinha.Descricao = err.Error()
		log.Println("Insert error: $s", err.Error())
	} else {
		retornoLinha.Codigo = 0
		retornoLinha.Pagina = 0
		retornoLinha.QtdePagina = 1
		retornoLinha.Registros = 1
		retornoLinha.ListaLinhas = []models.Linha{linha}
		log.Println("Insert success: $s", linha)

		historico := models.Historico{
			Tipo: enums.POST, 
			Sistema: enums.LINHA,
			Descricao: linha.Ddd + linha.Numero,
		}

		historicoService.Insert(historico)
	}

	return retornoLinha
}

// Get método para buscar as linhas telefonicas
func (m *LinhaService) Get(ddd string, numero string, pagina int, qtdePagina int) (retornoLinha models.RetornoLinha) {
	var linhas []models.Linha

	linhas, err := linhaRepository.Find(ddd, numero, pagina, qtdePagina)
	if err != nil {
		retornoLinha.Codigo = 500
		retornoLinha.Descricao = err.Error()

		log.Println("Get error: $s", err.Error())
	} else {
		count, error := linhaRepository.Count(ddd, numero)
		if error != nil {
			log.Println("Erro em buscar total: $s", error.Error())
		}

		retornoLinha.Codigo = 0
		retornoLinha.Pagina = pagina
		retornoLinha.QtdePagina = qtdePagina
		retornoLinha.Registros = count
		retornoLinha.ListaLinhas = linhas

		historico := models.Historico{
			Tipo: enums.GET, 
			Sistema: enums.LINHA,
			Descricao: "DDD: " + ddd + " Numero: " + numero,
		}

		historicoService.Insert(historico)
	}

	return retornoLinha
}

// Update atualiza linha
func (m *LinhaService) Update(Codigo string, linha models.Linha) (retornoLinha models.RetornoLinha) {
	linhaExists, err := linhaRepository.FindByID(Codigo)
	if linhaExists == (models.Linha{}) {
		retornoLinha.Codigo = 500
		retornoLinha.Descricao = "Error: linha não encontrada"
	} else if err != nil {
		retornoLinha.Codigo = 500
		retornoLinha.Descricao = err.Error()
	} else {
		linha.ID = linhaExists.ID
		error := linhaRepository.Update(linha)
		if error != nil {
			retornoLinha.Codigo = 500
			retornoLinha.Descricao = err.Error()
		} else {
			retornoLinha.Codigo = 0
			retornoLinha.Pagina = 0
			retornoLinha.QtdePagina = 1
			retornoLinha.Registros = 1
			retornoLinha.ListaLinhas = []models.Linha{linha}
			log.Println("Update success: $s", linha)
			
			historico := models.Historico{
				Tipo: enums.PUT, 
				Sistema: enums.LINHA,
				Descricao: linha.Ddd + linha.Numero,
			}

			historicoService.Insert(historico)
		}
	}

	return retornoLinha
}

// Delete remove linha
func (m *LinhaService) Delete(Codigo string) (retornoLinha models.RetornoLinha) {
	linhaExists, err := linhaRepository.FindByID(Codigo)
	if linhaExists == (models.Linha{}) {
		retornoLinha.Codigo = 500
		retornoLinha.Descricao = "Error: linha não encontrada"
	} else if err != nil {
		retornoLinha.Codigo = 500
		retornoLinha.Descricao = err.Error()
	} else {
		error := linhaRepository.Delete(linhaExists)
		if error != nil {
			retornoLinha.Codigo = 500
			retornoLinha.Descricao = err.Error()
		} else {
			retornoLinha.Codigo = 0
			retornoLinha.Pagina = 0
			retornoLinha.QtdePagina = 1
			retornoLinha.Registros = 1
			retornoLinha.ListaLinhas = []models.Linha{linhaExists}
			log.Println("Delete success: $s", linhaExists)
			
			historico := models.Historico{
				Tipo: enums.DELETE, 
				Sistema: enums.LINHA,
				Descricao: linhaExists.Ddd + linhaExists.Numero,
			}

			historicoService.Insert(historico)
		}
	}

	return retornoLinha
}