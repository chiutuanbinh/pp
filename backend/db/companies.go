package db

import (
	"context"
	"fmt"
	"log"
)

type Company struct {
	ID            int    `json:"id"`
	LegalEntityId int    `json:"legal_entity_id,omitempty"`
	Name          string `json:"name"`
}

func (c *Company) AddToDb() {
	c.ID = AddCompany(c.Name, c.LegalEntityId)
	log.Println(c)
}

func AddCompany(name string, legalEntityId int) int {
	if legalEntityId == 0 {
		legalEntityId = addLegalEntity()
	}
	var insertScript string = fmt.Sprintf(`INSERT INTO %[1]s(legal_entity_id, name) VALUES ($1, $2) RETURNING id;`, companyTable)
	var companyId int = 0
	err := pool.QueryRow(context.TODO(), insertScript, legalEntityId, name).Scan(&companyId)
	if err != nil {
		log.Println(err)
	}
	return companyId
}

func GetAllCompanies() []Company {
	var query = fmt.Sprintf(`SELECT id, legal_entity_id, name FROM %[1]s`, companyTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Println(err)
	}
	var res = make([]Company, 0)
	for rows.Next() {
		c := Company{}
		rows.Scan(&c.ID, &c.LegalEntityId, &c.Name)
		res = append(res, c)
	}
	return res
}

func GetCompany(id int) Company {
	var query = fmt.Sprintf(`
		SELECT c.id, c.legal_entity_id, c.name , 
		FROM %[1]s  as c
		WHERE id=$1 `, companyTable)
	c := Company{}
	err := pool.QueryRow(context.TODO(), query, id).Scan(&c.ID, &c.LegalEntityId, &c.Name)
	if err != nil {
		log.Println(err)
	}
	return c
}

func AddFinancialStatementType(name string) int {
	var insertScript string = `INSERT INTO financial_statement_types(name) VALUES($1) RETURNING id;`
	var financialStatementTypeId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name).Scan(&financialStatementTypeId)
	if err != nil {
		log.Println(err)
	}
	return financialStatementTypeId
}
