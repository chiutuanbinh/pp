package db

import (
	"context"
	"log"
)

type Person struct {
	ID            int
	Name          string
	Age           int
	LegalEntityId int
	ImageUrl      string
}

func AddPerson(name string, age int, image_url string) int {
	var legalEntityId int = addLegalEntity()
	var insertScript string = `INSERT INTO persons(legal_entity_id, name, age, image_url) VALUES ($1, $2, $3, $4) RETURNING id;`
	var personId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, legalEntityId, name, age, image_url).Scan(&personId)
	if err != nil {
		log.Fatal(err)
	}
	return personId
}
