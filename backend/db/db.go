package db

import (
	"context"
	"log"

	"mysite/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

const legalEntityTable = "legal_entities"
const personTable = "persons"
const companyTable = "companies"
const ownershipTable = "owns"
const projectTable = "projects"
const investmentTable = "investments"
const industryCategoryTable = "industry_categories"
const industryTable = "industries"
const operatingInTable = "operating_in"
const industryRelationTable = "industry_relations"
const industryRelationTypeTable = "industry_relation_types"
const financialStatementTypeTable = "financial_statement_types"
const financialStatementTable = "financial_statements"
const financialStatementLineTypeTable = "financial_statement_line_types"
const financialStatementLineTable = "financial_statement_lines"
const financialStatementLineSequenceTable = "financial_statement_line_sequences"
const exchangeTable = "exchanges"
const stockTable = "stocks"
const issueTable = "issues"
const stockPriceTable = "stock_prices"

func connect() {
	var err error
	pool, err = pgxpool.Connect(context.Background(), utils.Config.ConnStr)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	connect()
}
