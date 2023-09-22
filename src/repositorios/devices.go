package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type devices struct {
	db *sql.DB
}

func NewDevicesRepository(db *sql.DB) *devices {
	return &devices{db}
}

func (repositorio devices) Criar(device modelos.Device) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into devices (name, address, latitude, longitude, totalValue) values(?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, nil
	}
	defer statement.Close()

	resultado, erro := statement.Exec(device.Name, device.Address, device.Latitude, device.Longitude, device.TotalValue)
	if erro != nil {
		return 0, nil
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, nil
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio devices) Buscar(nameOrAddress string) ([]modelos.Device, error) {
	nameOrAddress = fmt.Sprintf("%%%s%%", nameOrAddress) // %nameOrAddress%

	linhas, erro := repositorio.db.Query(
		"select id, name, address, latitude, longitude, totalValue from devices where name like ? or address like ?",
		nameOrAddress, nameOrAddress,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var devices []modelos.Device

	for linhas.Next() {
		var device modelos.Device

		if erro = linhas.Scan(
			&device.ID,
			&device.Name,
			&device.Address,
			&device.Latitude,
			&device.Longitude,
			&device.TotalValue,
		); erro != nil {
			return nil, erro
		}

		devices = append(devices, device)
	}

	return devices, nil
}

func (repositorio devices) BuscarPorID(ID uint64) (modelos.Device, error) {
	linhas, erro := repositorio.db.Query(
		"select id, name, address, latitude, totalValue from devices where id = ?",
		ID,
	)

	if erro != nil {
		return modelos.Device{}, erro
	}

	defer linhas.Close()

	var device modelos.Device

	if linhas.Next() {
		if erro = linhas.Scan(
			&device.ID,
			&device.Name,
			&device.Address,
			&device.Latitude,
			&device.Longitude,
		); erro != nil {
			return modelos.Device{}, erro
		}
	}

	return device, nil
}

func (repositorio devices) AtualizarKwh(ID uint64, device modelos.Device) error {
	fmt.Println(device.TotalValue)

	statement, erro := repositorio.db.Prepare("update devices set totalValue = ? where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()
	if _, erro = statement.Exec(device.TotalValue, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio devices) Atualizar(ID uint64, device modelos.Device) error {
	statement, erro := repositorio.db.Prepare("update devices set name = ?, address = ?, latitude = ?, longitude = ?, totalValue where id = ?")

	if erro != nil {
		return erro
	}

	defer statement.Close()
	if _, erro = statement.Exec(device.Name, device.Address, device.Latitude, device.Longitude, device.TotalValue, ID); erro != nil {
		return erro
	}

	return nil
}

func (repositorio devices) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("delete from devices where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()
	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}
