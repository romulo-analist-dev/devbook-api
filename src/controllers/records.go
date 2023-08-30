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
)

func BuscarRecords(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		fmt.Println("db error")
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRecordsRepository(db)
	fmt.Println("É aqui?")
	records, erro := repositorio.Buscar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		fmt.Println("repositorio error")
		return
	}

	respostas.JSON(w, http.StatusOK, records)
}

func CreateRecord(message string) {
	data := strings.Split(message, ",")

	deviceID, erro := strconv.ParseUint(data[0], 10, 64)
	value, erro := strconv.ParseFloat(data[1], 64)

	if erro != nil {
		return
	}

	device := &modelos.Device{
		ID: deviceID,
	}

	record := modelos.Record{
		Value:  value,
		Device: device,
	}

	db, erro := banco.Conectar()
	if erro != nil {
		fmt.Println("erro ao conectar no banco")
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRecordsRepository(db)
	record.ID, erro = repositorio.Criar(record)
	if erro != nil {
		fmt.Println("erro ao executar o método Criar")
		return
	}
}

func CriarRecord(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var record modelos.Record
	if erro = json.Unmarshal(corpoRequest, &record); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NewRecordsRepository(db)
	record.ID, erro = repositorio.Criar(record)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, record)
}
