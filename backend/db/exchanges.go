package db

import (
	"context"
	"fmt"
	"log"
)

type Exchange struct {
	ID   int
	Name string
	Code string
}

func AddExchange(name string, code string) int {
	var insertScript string = `INSERT INTO exchanges(name, code) VALUES($1,$2) RETURNING id;`
	var exchangeID int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name, code).Scan(&exchangeID)
	if err != nil {
		log.Println(err)
	}
	return exchangeID
}

func GetAllExchange() []Exchange {
	var query = fmt.Sprintf(`SELECT id, name, code FROM %[1]s`, exchangeTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Println(err)
	}
	var res = make([]Exchange, 0)
	for rows.Next() {
		e := Exchange{}
		rows.Scan(&e.ID, &e.Name, &e.Code)
		res = append(res, e)
	}
	return res
}

func GetExchange(id int) Exchange {
	var query = fmt.Sprintf(`SELECT id, name, code FROM %[1]s WHERE id=$1`, exchangeTable)
	e := Exchange{}
	err := pool.QueryRow(context.TODO(), query, id).Scan(&e.ID, &e.Name, &e.Code)
	if err != nil {
		log.Println(err)
	}
	return e
}

func GetExchangeByCode(code string) Exchange {
	var query = fmt.Sprintf(`SELECT id, name, code FROM %[1]s WHERE code=$1`, exchangeTable)
	e := Exchange{}
	err := pool.QueryRow(context.TODO(), query, code).Scan(&e.ID, &e.Name, &e.Code)
	if err != nil {
		log.Println(err)
	}
	return e
}

func (e Exchange) IsNull() bool {
	if e.ID == 0 {
		return true
	}
	return false
}
