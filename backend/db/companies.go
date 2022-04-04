package db

import (
	"context"
	"fmt"
	"log"
)

type Company struct {
	ID            int
	LegalEntityId int
	Name          string
}

func AddCompany(name string) int {
	var legalEntityId int = addLegalEntity()
	var insertScript string = fmt.Sprintf(`INSERT INTO %[1]s(legal_entity_id, name) VALUES ($1, $2) RETURNING id;`, companyTable)
	var companyId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, legalEntityId, name).Scan(&companyId)
	if err != nil {
		log.Fatal(err)
	}
	return companyId
}

func GetAllCompany() []Company {
	var query = fmt.Sprintf(`SELECT id, legal_entity_id, name FROM %[1]s`, companyTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	var res = make([]Company, 0)
	for rows.Next() {
		c := Company{}
		rows.Scan(&c.ID, &c.LegalEntityId, &c.Name)
		res = append(res, c)
	}
	return res
}

func AddFinancialStatementType(name string) int {
	var insertScript string = `INSERT INTO financial_statement_types(name) VALUES($1) RETURN id;`
	var financialStatementTypeId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name).Scan(&financialStatementTypeId)
	if err != nil {
		log.Fatal(err)
	}
	return financialStatementTypeId
}
