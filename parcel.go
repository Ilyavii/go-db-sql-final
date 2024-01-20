package main

import (
	"database/sql"
	"fmt"
)

type ParcelStore struct {
	db *sql.DB
}

func NewParcelStore(db *sql.DB) ParcelStore {
	return ParcelStore{db: db}
}

func (s ParcelStore) Add(p Parcel) (int, error) {

	// реализуйте добавление строки в таблицу parcel, используйте данные из переменной p
	// название продукта и цена передаются через параметры
	res, err := db.Exec("INSERT INTO parcel (number, client, status, address, created_at) VALUES (:number, :client, :status, :address, :created_at)",
		sql.Named("number", p.Number),
		sql.Named("client", p.Client),
		sql.Named("status", p.Status),
		sql.Named("address", p.Address),
		sql.Named("created_at", p.CreatedAt))
	if err != nil {
		fmt.Println(err)
		return
	}
	id := int(res.LastInsertId())
	// верните идентификатор последней добавленной записи
	return id, nil
}

func (s ParcelStore) Get(number int) (Parcel, error) {
	// реализуйте чтение строки по заданному number
	// здесь из таблицы должна вернуться только одна строка

	// заполните объект Parcel данными из таблицы
	p := Parcel{}

	return p, nil
}

func (s ParcelStore) GetByClient(client int) ([]Parcel, error) {
	// реализуйте чтение строк из таблицы parcel по заданному client
	// здесь из таблицы может вернуться несколько строк

	// заполните срез Parcel данными из таблицы
	var res []Parcel

	return res, nil
}

func (s ParcelStore) SetStatus(number int, status string) error {
	// реализуйте обновление статуса в таблице parcel

	_, err := s.db.Exec("UPDATE parcel SET status = :status  WHERE number = :number",
		sql.Named("status", status),
		sql.Named("number", number))

	return err
}

func (s ParcelStore) SetAddress(number int, address string) error {
	_, err := s.db.Exec("UPDATE parcel SET address = :address  WHERE number = :number",
		sql.Named("address", address),
		sql.Named("number", number))

	return err
}

func (s ParcelStore) Delete(number int) error {
	// реализуйте удаление строки из таблицы parcel
	// удалять строку можно только если значение статуса registered

	return nil
}
