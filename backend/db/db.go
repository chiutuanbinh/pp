package db

import (
	"context"
	"log"

	"mysite/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

func connect() {
	var err error
	pool, err = pgxpool.Connect(context.Background(), utils.Config.ConnStr)
	if err != nil {
		log.Fatalln(err)
	}
}

func addLegalEntity() int {
	var insertScript string = "INSERT INTO legal_entities DEFAULT VALUES  RETURNING id;"
	var legalEntityId int = -1
	err := pool.QueryRow(context.TODO(), insertScript).Scan(&legalEntityId)
	if err != nil {
		log.Fatal(err)
	}
	return legalEntityId
}

func AddCompany(name string) int {
	var legalEntityId int = addLegalEntity()
	var insertScript string = `INSERT INTO companies(legal_entity_id, name) VALUES ($1, $2) RETURNING id;`
	var companyId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, legalEntityId, name).Scan(&companyId)
	if err != nil {
		log.Fatal(err)
	}
	return companyId
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

func AddPersonOwnerships(legalentityId int, companyId int) {

}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	connect()
}
