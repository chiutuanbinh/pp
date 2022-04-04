package db

import (
	"context"
	"log"
)

func addLegalEntity() int {
	var insertScript string = "INSERT INTO legal_entities DEFAULT VALUES  RETURNING id;"
	var legalEntityId int = -1
	err := pool.QueryRow(context.TODO(), insertScript).Scan(&legalEntityId)
	if err != nil {
		log.Fatal(err)
	}
	return legalEntityId
}
