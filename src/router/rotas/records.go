package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasRecords = []Rota{
	{
		URI:                "/records",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarRecords,
		RequerAutenticacao: false,
	},
}
