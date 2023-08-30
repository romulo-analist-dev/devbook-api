package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type records struct {
	db *sql.DB
}

func NewRecordsRepository(db *sql.DB) *records {
	return &records{db}
}

func (repositorio records) Criar(record modelos.Record) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into records (value, device_id) values(?, ?)",
	)
	if erro != nil {
		fmt.Println("erro ao preparar a query")
		return 0, nil
	}
	defer statement.Close()

	resultado, erro := statement.Exec(record.Value, record.Device.ID)
	if erro != nil {
		fmt.Println("erro ao executar a query: " + erro.Error())
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		fmt.Println("erro ao obter LastInsertId")
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio records) Buscar() ([]modelos.Record, error) {
	linhas, erro := repositorio.db.Query(
		"select D.*, R.id, R.value, R.createdAt from devices D inner join (SELECT device_id, MAX(createdAt) as maxCreatedAt from records group by device_id) as latestRecords ON D.id = latestRecords.device_id INNER JOIN records AS R ON D.id = R.device_id AND latestRecords.maxCreatedAt = R.createdAt;",
	)

	if erro != nil {
		fmt.Println("erro ao preparar a query")
		return nil, erro
	}

	defer linhas.Close()

	var records []modelos.Record

	for linhas.Next() {
		var record modelos.Record
		var device modelos.Device

		if erro = linhas.Scan(
			&device.ID,
			&device.Name,
			&device.Address,
			&device.Latitude,
			&device.Longitude,
			&record.ID,
			&record.Value,
			&record.CreatedAt,
		); erro != nil {
			fmt.Println("erro ao executar a query: " + erro.Error())
			return nil, erro
		}
		record.Device = &device
		records = append(records, record)
	}

	return records, nil
}

func (repositorio records) BuscarPorID(ID uint64) (modelos.Record, error) {
	linhas, erro := repositorio.db.Query(
		"select id, value, createdAt, D.* from records inner join devices D on device_id = D.id where id = ?",
		ID,
	)

	if erro != nil {
		return modelos.Record{}, erro
	}

	defer linhas.Close()

	var record modelos.Record

	if linhas.Next() {
		if erro = linhas.Scan(
			&record.ID,
			&record.Value,
			&record.CreatedAt,
			&record.Device,
		); erro != nil {
			return modelos.Record{}, erro
		}
	}

	return record, nil
}
