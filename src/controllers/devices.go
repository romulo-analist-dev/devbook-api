package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	emitter "github.com/emitter-io/go"
	"github.com/gorilla/mux"
)

func CriarDevice(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var device modelos.Device
	if erro = json.Unmarshal(corpoRequest, &device); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewDevicesRepository(db)
	device.ID, erro = repositorio.Criar(device)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, device)
}

func BuscarDevices(w http.ResponseWriter, r *http.Request) {
	nameOrAddress := strings.ToLower(r.URL.Query().Get("device"))

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewDevicesRepository(db)
	devices, erro := repositorio.Buscar(nameOrAddress)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	respostas.JSON(w, http.StatusOK, devices)
}

func BuscarDevice(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	deviceID, erro := strconv.ParseUint(parametros["deviceId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewDevicesRepository(db)
	device, erro := repositorio.BuscarPorID(deviceID)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusOK, device)
}

func AtualizarDevice(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	deviceID, erro := strconv.ParseUint(parametros["deviceId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var device modelos.Device
	fmt.Println(corpoRequisicao)
	if erro = json.Unmarshal(corpoRequisicao, &device); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewDevicesRepository(db)
	if erro = repositorio.Atualizar(deviceID, device); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}

func AtualizarDeviceTotalKwh(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	_, erro := strconv.ParseUint(parametros["deviceId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var data TotalKwh
	if erro = json.Unmarshal(corpoRequisicao, &data); erro != nil {
		fmt.Println("bad request: ", erro)
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}
	fmt.Println("data: ")
	opts := emitter.NewClientOptions()
	opts.AddBroker("tcp://broker.hivemq.com:1883")

	topico := "alguma-coisa/demo/config"
	message := "1," + fmt.Sprintf("%f", data.TotalValue)

	client := emitter.NewClient(opts)
	wait(client.Connect())
	client.Publish(topico, message, 0)

	respostas.JSON(w, http.StatusNoContent, nil)
}

type TotalKwh struct {
	TotalValue float64 `json:"totalValue"`
}

func wait(t emitter.Token) {
	t.Wait()
	if t.Error() != nil {
		panic(t.Error())
	}
}

func DeletarDevice(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	deviceID, erro := strconv.ParseUint(parametros["deviceId"], 10, 64)
	if erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewDevicesRepository(db)
	if erro = repositorio.Deletar(deviceID); erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusNoContent, nil)
}
