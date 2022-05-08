package db

import (
	"context"
	"fmt"
	"log"
)

type Stock struct {
	ID         int
	Code       string
	ExchangeId int
}

func (s *Stock) AddToDb() {
	s.ID = AddStock(s.Code, s.ExchangeId)
}

func AddStock(code string, exchangeID int) int {
	var insertScript string = `INSERT INTO stocks(code, exchange_id) VALUES($1, $2) RETURNING id;`
	var stockID int = -1
	err := pool.QueryRow(context.TODO(), insertScript, code, exchangeID).Scan(&stockID)
	if err != nil {
		log.Println(err)
	}
	return stockID
}

func AddIssue(stockId int, companyId int, timestamp int64, amount int, initialPrice int) int {
	var insertScript string = `INSERT INTO issues(stock_id, company_id, date, amount, initial_price) VALUES($1, $2, $3, $4, $5) RETURNING id;`
	var issueId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, stockId, companyId, timestamp, amount, initialPrice).Scan(&issueId)
	if err != nil {
		log.Println(err)
	}
	return issueId
}

func AddStockPrice(timestamp int64, openingPrice int, closingPrice int, highest int, lowest int, stockId int) int {
	var insertScript string = `INSERT INTO stock_prices(date, opening_price, closing_price, highest, lowest, stock_id) VALUES($1, $2, $3, $4, $5, $6) RETURNING id;`
	var priceId int = -1
	err := pool.QueryRow(context.TODO(), insertScript, timestamp, openingPrice, closingPrice, highest, lowest, stockId).Scan(&priceId)
	if err != nil {
		log.Println(err)
	}
	return priceId
}

func GetAllStocks() []Stock {
	var query = fmt.Sprintf(`SELECT id, code, exchange_id FROM %[1]s`, stockTable)
	rows, err := pool.Query(context.TODO(), query)
	if err != nil {
		log.Println(err)
	}
	var res = make([]Stock, 0)
	for rows.Next() {
		s := Stock{}
		rows.Scan(&s.ID, &s.Code, &s.ExchangeId)
		res = append(res, s)
	}
	return res
}

func GetStock(id int) Stock {
	var query = fmt.Sprintf(`SELECT id, code, exchange_id FROM %[1]s WHERE id=$1`, stockTable)
	s := Stock{}
	err := pool.QueryRow(context.TODO(), query, id).Scan(&s.ID, &s.Code, &s.ExchangeId)
	if err != nil {
		log.Println(err)
	}
	return s
}
