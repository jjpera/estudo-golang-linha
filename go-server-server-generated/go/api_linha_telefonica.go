package swagger

import (
	"log"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/gorilla/mux"
)

var linhaService LinhaService

// CadastrarLinha inserir linha
func CadastrarLinha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var linha Linha
	_ = json.NewDecoder(r.Body).Decode(&linha)

	log.Printf("Insert Linha %s %s.", linha.Ddd, linha.Numero)

	var retornoLinha RetornoLinha
	retornoLinha = linhaService.Insert(linha)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(retornoLinha)
}

// BuscarLinha buscar linha
func BuscarLinha(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var ddd string = mux.Vars(r)["ddd"]
	var numero string = mux.Vars(r)["numero"]
	var pagina = stringToInt(mux.Vars(r)["pagina"])
	var qtdePagina = stringToInt(mux.Vars(r)["qtdePagina"])

	var retornoLinha RetornoLinha
	retornoLinha = linhaService.Get(ddd, numero, pagina, qtdePagina)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(retornoLinha)
}

// AlterarLinha alterar linha
func AlterarLinha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var retornoLinha RetornoLinha
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(retornoLinha)
}

// ExcluirLinha remover linha
func ExcluirLinha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var retornoLinha RetornoLinha
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(retornoLinha)
}

func stringToInt(snumber string) int {
	pagina, err := strconv.Atoi(snumber)
	if (err != nil) {
		return 0
	}
	return pagina
}
