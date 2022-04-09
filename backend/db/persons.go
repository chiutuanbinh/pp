package db

import (
	"context"
	"fmt"
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

func GetAllPersons() []Person {
	var query = fmt.Sprintf(`SELECT id,name, age, legal_entity_id, image_url FROM %[1]s`, personTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	var res = make([]Person, 0)
	for rows.Next() {
		p := Person{}
		rows.Scan(&p.ID, &p.Name, &p.Age, &p.LegalEntityId, &p.ImageUrl)
		res = append(res, p)
	}
	return res
}
