package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Route entidade
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes lista
type Routes []Route

// NewRouter iniciando Router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// Index Hello World
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/linhas/",
		Index,
	},

	Route{
		"AlterarLinha",
		strings.ToUpper("Put"),
		"/linhas/linha/{codigo}",
		AlterarLinha,
	},

	Route{
		"BuscarLinha",
		strings.ToUpper("Get"),
		"/linhas/linha",
		BuscarLinha,
	},

	Route{
		"CadastrarLinha",
		strings.ToUpper("Post"),
		"/linhas/linha",
		CadastrarLinha,
	},

	Route{
		"ExcluirLinha",
		strings.ToUpper("Delete"),
		"/linhas/linha/{codigo}",
		ExcluirLinha,
	},
}
