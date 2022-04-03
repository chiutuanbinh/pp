package db

import (
	"context"
	"fmt"
	"log"

	"mysite/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

var companyTable string = "companies"

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

type Company struct {
	ID            int
	LegalEntityId int
	Name          string
}

type Person struct {
	ID            int
	Name          string
	Age           int
	LegalEntityId int
	ImageUrl      string
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

func AddIndustryCategory(name string) int {
	var insertScript string = `INSERT INTO industry_categories(name) VALUES($1) RETURNING id;`
	var industryCategoryId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name).Scan(&industryCategoryId)
	if err != nil {
		log.Fatal(err)
	}
	return industryCategoryId
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

func AddIndustryRelationType(name string, coefficient int) int {
	var insertScript string = `INSERT INTO industry_relation_types(name, coefficient) VALUES($1, $2) RETURNING id;`
	var industryRelationTypeId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name, coefficient).Scan(&industryRelationTypeId)
	if err != nil {
		log.Fatal(err)
	}
	return industryRelationTypeId
}

func AddExchange(name string) int {
	var insertScript string = `INSERT INTO exchanges(name) VALUES($1) RETURN id;`
	var exchangeID int = -1
	err := pool.QueryRow(context.TODO(), insertScript, name).Scan(&exchangeID)
	if err != nil {
		log.Fatal(err)
	}
	return exchangeID
}

func AddStock(code string, exchangeID string) int {
	var insertScript string = `INSERT INTO stocks(code, exchange_id) VALUES($1, $2) RETURN id;`
	var stockID int = -1
	err := pool.QueryRow(context.TODO(), insertScript, code, exchangeID).Scan(&stockID)
	if err != nil {
		log.Fatal(err)
	}
	return stockID
}

func AddIssue(stockId int, companyId int, timestamp int64, amount int, initialPrice int) int {
	var insertScript string = `INSERT INTO issues(stock_id, company_id, date, amount, initial_price) VALUES($1, $2, $3, $4, $5) RETURN id;`
	var issueId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, stockId, companyId, timestamp, amount, initialPrice).Scan(&issueId)
	if err != nil {
		log.Fatal(err)
	}
	return issueId
}

func AddStockPrice(timestamp int64, openingPrice int, closingPrice int, highest int, lowest int, stockId int) int {
	var insertScript string = `INSERT INTO stock_prices(date, opening_price, closing_price, highest, lowest, stock_id) VALUES($1, $2, $3, $4, $5, $6) RETURN id;`
	var priceId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, timestamp, openingPrice, closingPrice, highest, lowest, stockId).Scan(&priceId)
	if err != nil {
		log.Fatal(err)
	}
	return priceId
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

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	connect()
}
