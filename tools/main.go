package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func clean() {
	prefix := "./dat"
	dir, err := os.Open(prefix + "/raw/")
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		// fmt.Println(v.Name(), v.IsDir())
		// if v.Name() != "ACC_data.csv" {
		// 	continue
		// }
		outPath := prefix + "/clean/" + v.Name()
		fullPath := prefix + "/raw/" + v.Name()
		// fmt.Println(fullPath)
		rf, _ := os.Open(fullPath)
		wf, _ := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0644)
		defer rf.Close()
		defer wf.Close()
		reader := bufio.NewReader(rf)
		writer := bufio.NewWriter(wf)
		csvReader := csv.NewReader(reader)
		csvWriter := csv.NewWriter(writer)
		var elems []string
		var err error
		elems, err = csvReader.Read()
		if err != nil {
			fmt.Printf("%v %v", err, fullPath)
			continue
		}
		if elems[0] == "date" {
			elems, err = csvReader.Read()
		} else {
			csvWriter.Write(elems)
		}
		size := len(elems)
		var accumulator []string = make([]string, 0)
		// var flag bool = false
		// r, _ := regexp.Compile("\\s\\s*?\\s")
		for true {
			elems, err = csvReader.Read()
			if len(elems) == 0 {
				break
			}
			if len(elems) != size {
				if len(elems) == 1 {
					continue
				}
				// if r.MatchString(elems[0]) {
				// 	continue
				// }
				accumulator = append(accumulator, elems[len(elems)-1])
				// fmt.Println(accumulator)
				if len(accumulator) == size {

					csvWriter.Write(accumulator)
					accumulator = make([]string, 0)
				}
			} else {

				csvWriter.Write(elems)
			}

		}
		csvWriter.Flush()
		// break
	}
}

type Stock struct {
	ID         int `gorm:"primaryKey,index"`
	Code       string
	ExchangeId int
}

type StockPrice struct {
	ID           int `gorm:"primaryKey,index"`
	Date         int64
	OpeningPrice int32
	ClosingPrice int32
	Highest      int32
	Lowest       int32
	StockId      int
}

func insertDb() {
	DB, err := gorm.Open(postgres.Open("postgresql://postgres:123456@0.0.0.0:5432/test?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	prefix := "./dat/clean/"
	dir, err := os.Open(prefix)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := dir.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range files {
		fullPath := prefix + v.Name()
		rf, _ := os.Open(fullPath)
		defer rf.Close()
		csvReader := csv.NewReader(rf)
		flag := false
		for true {
			entry, _ := csvReader.Read()
			if !flag {
				fmt.Printf("%v %v\n", v.Name(), csvReader.FieldsPerRecord)
				flag = true
			}

			if len(entry) == 0 {
				break
			}
			var date time.Time
			var cPrice, oPrice, hPrice, lPrice float64
			var stock Stock = Stock{}
			if len(entry) == 14 {
				date, _ = time.Parse("02/01/2006", entry[0])
				cPrice, _ = strconv.ParseFloat(entry[2], 64)
				oPrice, _ = strconv.ParseFloat(entry[11], 64)
				hPrice, _ = strconv.ParseFloat(entry[12], 64)
				lPrice, _ = strconv.ParseFloat(entry[13], 64)
				DB.Where(&Stock{Code: strings.Split(v.Name(), "_")[0]}).First(&stock)
			} else if len(entry) == 12 {
				date, _ = time.Parse("02/01/2006", entry[0])
				cPrice, _ = strconv.ParseFloat(entry[2], 64)
				oPrice, _ = strconv.ParseFloat(entry[9], 64)
				hPrice, _ = strconv.ParseFloat(entry[10], 64)
				lPrice, _ = strconv.ParseFloat(entry[11], 64)
				DB.Where(&Stock{Code: strings.Split(v.Name(), "_")[0]}).First(&stock)
			} else if len(entry) == 13 {
				date, _ = time.Parse("02/01/2006", entry[0])
				cPrice, _ = strconv.ParseFloat(entry[2], 64)
				oPrice, _ = strconv.ParseFloat(entry[10], 64)
				hPrice, _ = strconv.ParseFloat(entry[11], 64)
				lPrice, _ = strconv.ParseFloat(entry[12], 64)
				DB.Where(&Stock{Code: strings.Split(v.Name(), "_")[0]}).First(&stock)
			}
			price := StockPrice{
				Date:         date.Unix(),
				ClosingPrice: int32(cPrice * 1000),
				OpeningPrice: int32(oPrice * 1000),
				Highest:      int32(hPrice * 1000),
				Lowest:       int32(lPrice * 1000),
				StockId:      stock.ID,
			}
			fmt.Println(price)
			DB.Create(&price)
			// fmt.Println(entry)
			// fmt.Printl	n(len(entry))
		}
	}

}

func t() {
	layout := "02/01/2006"
	t, err := time.Parse(layout, "05/05/2022")
	if err != nil {
		panic(err)
	}
	fmt.Println(t.Unix())
}
func main() {
	// extractAll()
	// extractTicker("ACV")
	// clean()
	insertDb()
	// t()
}
