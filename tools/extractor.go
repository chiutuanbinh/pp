package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func extractTable(page *rod.Page, writer *csv.Writer) {
	elem := page.Timeout(5 * time.Second).MustElement(`#GirdTable2, #GirdTable`).CancelTimeout()
	s, _ := elem.Attribute("id")

	t, _ := elem.Text()
	lines := strings.Split(t, "\n")
	if *s == "GirdTable2" {
		lines = lines[7:]
	} else {
		lines = lines[9:]
	}

	for _, line := range lines {
		fields := strings.Split(line, "\t")
		for i := range fields {
			fields[i] = strings.Trim(fields[i], " ")
		}
		writer.Write(fields)
	}
	writer.Flush()
}

func extractTicker(ticker string, browser *rod.Browser) {
	fileName := fmt.Sprintf("dat/%v_data.csv", ticker)
	_, err := os.Stat(fileName)
	if err == nil {
		return
	}
	url := fmt.Sprintf("https://s.cafef.vn/Lich-su-giao-dich-%v-1.chn", ticker)
	page := browser.MustPage(url).Timeout(15 * time.Second).MustWaitLoad().CancelTimeout()
	defer page.Close()
	f, err := os.Create(fileName)
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()
	writer := bufio.NewWriter(f)
	csvWriter := csv.NewWriter(writer)
	// header := [11]string{
	// 	"date",
	// 	"recalibrate",
	// 	"close",
	// 	"change",
	// 	"KLKL",
	// 	"GTKL",
	// 	"KLTT",
	// 	"GTTT",
	// 	"Open",
	// 	"Highest",
	// 	"Lowest",
	// }
	// csvWriter.Write(header[:])
	extractTable(page, csvWriter)
	for true {
		elem := page.Timeout(3 * time.Second).MustElement(`#ctl00_ContentPlaceHolder1_ctl03_panelAjax> div > div > div > table > tbody > tr > td:last-child `)
		icon, _ := elem.Text()
		// fmt.Println(icon)
		if icon != ">" {
			break
		}

		elem.MustClick()
		wait := page.Timeout(3 * time.Second).WaitEvent(proto.NetworkDataReceived{})
		wait()
		time.Sleep(200 * time.Millisecond)
		// fmt.Print(page.MustElement(`#ctl00_ContentPlaceHolder1_ctl03_divHO > div > div > table > tbody > tr > td:nth-last-child(2) > a`).Text())
		extractTable(page, csvWriter)
	}

	// page.MustWaitLoad().MustScreenshot(fmt.Sprintf("%v.png", ticker))
}

func extractTickerLatest(ticker string) {
	url := fmt.Sprintf("https://s.cafef.vn/Lich-su-giao-dich-%v-1.chn", ticker)
	page := rod.New().NoDefaultDevice().MustConnect().MustPage(url)
	err := rod.Try(func() {
		t, _ := page.Timeout(time.Second * 5).MustElement(`#ctl00_ContentPlaceHolder1_ctl03_rptData2_ctl01_itemTR`).CancelTimeout().Text()
		fmt.Println(t)
	})
	if err != nil {
		// fmt.Println(err)
	}

	// page.MustWaitLoad().MustScreenshot(fmt.Sprintf("%v.png", ticker))
}

func extractAll() {
	file := "company_stock.csv"
	f, _ := os.Open(file)
	defer f.Close()
	csv := csv.NewReader(f)
	csv.Read()
	companyList := make([]string, 0)
	for i := 0; i < 3200; i++ {
		l, _ := csv.Read()
		// fmt.Println(len(l))
		if len(l) == 0 {
			break
		}
		companyList = append(companyList, l[1])
		// err := rod.Try(func() {
		// 	extractTicker()
		// })
		// if err != nil {
		// 	fmt.Println(err)
		// 	continue
		// }
	}
	var browser *rod.Browser = rod.New().NoDefaultDevice().MustConnect()
	sem := make(chan int, 4)
	for i := range companyList {
		sem <- 1 // will block if there is MAX ints in sem
		go func(company string) {
			fmt.Println(company)
			err := rod.Try(func() {
				extractTicker(company, browser)
			})
			if err != nil {
				fmt.Println(err)
			}
			<-sem // removes an int from sem, allowing another to proceed
		}(companyList[i])
	}
}
