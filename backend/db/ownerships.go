package db

import (
	"context"
	"fmt"
	"log"
)

func AddOwnerships(legalEntityId int, companyId int) int {
	var insertScript string = `INSERT INTO owns(legal_entity_id, company_id) VALUES($1,$2) RETURNING id;`
	ownsId := -1
	err := pool.QueryRow(context.TODO(), insertScript, legalEntityId, companyId).Scan(&ownsId)

	if err != nil {
		log.Fatal(err)
	}

	return ownsId
}

func AddPersonOwnerships(personId int, companyId int) int {
	var queryScript string = `SELECT legal_entity_id FROM persons WHERE id=$1;`
	var legal_entity_id = -1
	err := pool.QueryRow(context.TODO(), queryScript, personId).Scan(&legal_entity_id)
	if err != nil {
		log.Fatal(err)
	}
	return AddOwnerships(legal_entity_id, companyId)
}

func AddCompanyOwnerships(owningCompanyId int, ownedCompanyId int) int {
	var queryScript string = fmt.Sprintf(`SELECT legal_entity_id FROM %[1]s WHERE id=$1;`, companyTable)
	var legal_entity_id = -1
	err := pool.QueryRow(context.TODO(), queryScript, owningCompanyId).Scan(&legal_entity_id)
	if err != nil {
		log.Fatal(err)
	}
	return AddOwnerships(legal_entity_id, ownedCompanyId)
}
