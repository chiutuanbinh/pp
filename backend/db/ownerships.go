package db

import (
	"context"
	"fmt"
	"log"
)

type Ownership struct {
	ID            int
	LegalEntityId int
	StockId       int
	Amount        int
}

func AddOwnerships(legalEntityId int, stockId int, amount int) int {
	var insertScript string = `INSERT INTO owns(legal_entity_id, stock_id, amount) VALUES($1,$2,$3) RETURNING id;`
	ownsId := -1
	err := pool.QueryRow(context.TODO(), insertScript, legalEntityId, stockId, amount).Scan(&ownsId)

	if err != nil {
		log.Println(err)
	}

	return ownsId
}

func AddPersonOwnerships(personId int, stockId int, amount int) int {
	var queryScript string = `SELECT legal_entity_id FROM persons WHERE id=$1;`
	var legal_entity_id = -1
	err := pool.QueryRow(context.TODO(), queryScript, personId).Scan(&legal_entity_id)
	if err != nil {
		log.Println(err)
	}
	return AddOwnerships(legal_entity_id, stockId, amount)
}

func AddCompanyOwnerships(owningCompanyId int, stockId int, amount int) int {
	var queryScript string = fmt.Sprintf(`SELECT legal_entity_id FROM %[1]s WHERE id=$1;`, companyTable)
	var legal_entity_id = -1
	err := pool.QueryRow(context.TODO(), queryScript, owningCompanyId).Scan(&legal_entity_id)
	if err != nil {
		log.Println(err)
	}
	return AddOwnerships(legal_entity_id, stockId, amount)
}
