package db

import (
	"context"
	"log"
)

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
