package db

import (
	"context"
	"fmt"
	"log"
)

type Industry struct {
	ID                 int
	Name               string
	IndustryCategoryId int
	Description        string
}

func (i *Industry) AddToDb() {
	i.ID = AddIndustry(i.Name, i.IndustryCategoryId, i.Description)
}

func AddIndustry(name string, industryCategoryId int, description string) int {
	var insertScript string = `INSERT INTO industries(name, industry_category_id, description) VALUES($1, $2, $3) RETURNING id;`
	var industryId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name, industryCategoryId).Scan(&industryId)
	if err != nil {
		log.Println(err)
	}
	return industryId
}

func GetAllIndustries() []Industry {
	var query = fmt.Sprintf(`SELECT id, name FROM %[1]s`, industryTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Println(err)
	}
	var res = make([]Industry, 0)
	for rows.Next() {
		i := Industry{}
		rows.Scan(&i.ID, &i.Name)
		res = append(res, i)
	}
	return res
}
func GetIndustry(id int) Industry {
	var query = fmt.Sprintf(`SELECT id, name FROM %[1]s WHERE id=$1`, industryTable)
	i := Industry{}
	err := pool.QueryRow(context.TODO(), query, id).Scan(&i.ID, &i.Name)
	if err != nil {
		log.Println(err)
	}
	return i
}

func AddIndustryRelationType(name string, coefficient int) int {
	var insertScript string = `INSERT INTO industry_relation_types(name, coefficient) VALUES($1, $2) RETURNING id;`
	var industryRelationTypeId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name, coefficient).Scan(&industryRelationTypeId)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
	}
	return industryCategoryId
}
