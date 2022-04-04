package db

import (
	"context"
	"log"
)

func AddExchange(name string) int {
	var insertScript string = `INSERT INTO exchanges(name) VALUES($1) RETURN id;`
	var exchangeID int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name).Scan(&exchangeID)
	if err != nil {
		log.Fatal(err)
	}
	return exchangeID
}
