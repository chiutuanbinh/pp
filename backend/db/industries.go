package db

import (
	"context"
	"fmt"
	"log"
)

type Industry struct {
	ID   int
	Name string
}

func AddIndustry(name string, industryCategoryId int) int {
	var insertScript string = `INSERT INTO industries(name, industry_category_id) VALUES($1, $2) RETURNING id;`
	var industryId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name, industryCategoryId).Scan(&industryId)
	if err != nil {
		log.Fatal(err)
	}
	return industryId
}

func GetAllIndustries() []Industry {
	var query = fmt.Sprintf(`SELECT id, name FROM %[1]s`, industryTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	var res = make([]Industry, 0)
	for rows.Next() {
		i := Industry{}
		rows.Scan(&i.ID, &i.Name)
		res = append(res, i)
	}
	return res
}

func AddIndustryRelationType(name string, coefficient int) int {
	var insertScript string = `INSERT INTO industry_relation_types(name, coefficient) VALUES($1, $2) RETURNING id;`
	var industryRelationTypeId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name, coefficient).Scan(&industryRelationTypeId)
	if err != nil {
		log.Fatal(err)
	}
	return industryRelationTypeId
}

type IndustryRelationType struct {
	ID          int
	Name        string
	Coefficient float32
}

func GetAllIndustryRelationTypes() []IndustryRelationType {
	var query = fmt.Sprintf(`SELECT id, name, coefficient FROM %[1]s`, companyTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	var res = make([]IndustryRelationType, 0)
	for rows.Next() {
		c := IndustryRelationType{}
		rows.Scan(&c.ID, &c.Name, &c.Coefficient)
		res = append(res, c)
	}
	return res
}

type IndustryRelation struct {
	ID                     int
	FromCompanyID          int
	ToCompanyID            int
	IndustryRelationTypeID int
}

func GetAllIndustryRelations() []IndustryRelation {
	var query = fmt.Sprintf(`SELECT id, from_company_id, to_company_id, industry_relation_type_id FROM %[1]s`, industryRelationTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	var res = make([]IndustryRelation, 0)
	for rows.Next() {
		ir := IndustryRelation{}
		rows.Scan(&ir.ID, &ir.FromCompanyID, &ir.ToCompanyID, &ir.IndustryRelationTypeID)
		res = append(res, ir)
	}
	return res
}

func AddIndustryCategory(name string) int {
	var insertScript string = `INSERT INTO industry_categories(name) VALUES($1) RETURNING id;`
	var industryCategoryId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name).Scan(&industryCategoryId)
	if err != nil {
		log.Fatal(err)
	}
	return industryCategoryId
}
