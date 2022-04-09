package db

import (
	"context"
	"fmt"
	"log"
)

type Exchange struct {
	ID   int
	Name string
}

func AddExchange(name string) int {
	var insertScript string = `INSERT INTO exchanges(name) VALUES($1) RETURNING id;`
	var exchangeID int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name).Scan(&exchangeID)
	if err != nil {
		log.Fatal(err)
	}
	return exchangeID
}

func GetAllExchange() []Exchange {
	var query = fmt.Sprintf(`SELECT id, name FROM %[1]s`, exchangeTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	var res = make([]Exchange, 0)
	for rows.Next() {
		e := Exchange{}
		rows.Scan(&e.ID, &e.Name)
		res = append(res, e)
	}
	return res
}
