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
	statementRecord, erro := repositorio.db.Prepare(
		"insert into records (value, device_id) values(?, ?);",
	)
	if erro != nil {
		fmt.Println("erro aqui 0")
		return 0, nil
	}

	statementDevice, erro := repositorio.db.Prepare(
		"update devices set totalValue = ? where id = ?",
	)

	if erro != nil {
		fmt.Println("erro aqui 1")
		return 0, nil
	}
	defer statementRecord.Close()
	defer statementDevice.Close()

	resultadoRecord, erro := statementRecord.Exec(record.Value, record.Device.ID)
	if erro != nil {
		fmt.Println("erro aqui 2")
		return 0, nil
	}
	statementDevice.Exec(record.Device.TotalValue, record.Device.ID)

	ultimoIDInserido, erro := resultadoRecord.LastInsertId()
	if erro != nil {
		fmt.Println("erro aqui 3")
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio records) Buscar() ([]modelos.Record, error) {
	linhas, erro := repositorio.db.Query(
		"select D.*, R.id, R.value, R.createdAt from devices D inner join (SELECT device_id, MAX(createdAt) as maxCreatedAt from records group by device_id) as latestRecords ON D.id = latestRecords.device_id INNER JOIN records AS R ON D.id = R.device_id AND latestRecords.maxCreatedAt = R.createdAt;",
	)

	if erro != nil {
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
			&device.TotalValue,
			&record.ID,
			&record.Value,
			&record.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		record.Device = &device
		records = append(records, record)
	}

	return records, nil
}

func (repositorio records) BuscarPorID(ID uint64) (modelos.Record, error) {
	linha, erro := repositorio.db.Query(
		"SELECT D.*, R.id, R.value, R.createdAt FROM devices D INNER JOIN records R ON D.id = R.device_id WHERE R.device_id = ? ORDER BY R.id DESC LIMIT 1;",
		ID,
	)

	if erro != nil {
		return modelos.Record{}, erro
	}

	defer linha.Close()

	var record modelos.Record

	if linha.Next() {
		var device modelos.Device

		if erro = linha.Scan(&device.ID,
			&device.Name,
			&device.Address,
			&device.Latitude,
			&device.Longitude,
			&device.TotalValue,
			&record.ID,
			&record.Value,
			&record.CreatedAt); erro != nil {
			return modelos.Record{}, erro
		}
		record.Device = &device
	}

	return record, nil
}
