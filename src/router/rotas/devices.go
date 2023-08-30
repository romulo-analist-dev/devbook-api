package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasDevices = []Rota{
	{
		URI:                "/devices",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarDevice,
		RequerAutenticacao: false,
	},
	{
		URI:                "/devices",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDevices,
		RequerAutenticacao: false,
	},
	{
		URI:                "/devices/{deviceId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarDevice,
		RequerAutenticacao: false,
	},
	{
		URI:                "/devices/{deviceId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarDevice,
		RequerAutenticacao: false,
	},
	{
		URI:                "/devices/{deviceId}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarDevice,
		RequerAutenticacao: false,
	},
}
