package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/tealeg/xlsx"
)

func getTime() string {
	t := time.Now()
	return t.Format("2006-01-02-15-04-05")
}
func main() {
	xCurTime := getTime()
	//fmt.Println(os.Args[1])
	fName := "toyScrape" + xCurTime + ".csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	upc := []string{}
	xlFileName := os.Args[1]
	xlFile, err := xlsx.OpenFile(xlFileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, sheet := range xlFile.Sheets {
		ctr := 1
		for {
			if ctr == len(sheet.Rows) {
				break
			}
			row := sheet.Rows[ctr]
			if row.Cells[0].String() == "" {
				break
			}
			upc = append(upc, row.Cells[0].Value)
			//upc = append(upc, "653569825586")
			hldUPC := row.Cells[0].String()
			url := "https://www.upcitemdb.com/upc/" + hldUPC
			//fmt.Println(url)
			c := colly.NewCollector()
			toyName := ""
			c.OnHTML(".num li:nth-of-type(1)", func(e *colly.HTMLElement) {
				toyName = e.Text
				//fmt.Println(toyName)
			})
			manufacturer := ""
			c.OnHTML(".detail-list tr", func(e *colly.HTMLElement) {
				holdText := e.ChildText("td")
				if holdText[0:6] == "Brand:" {
					manufacturer = strings.TrimSpace(holdText[7:])
					//fmt.Println("manufacturer>>>> ", strings.TrimSpace(manufacturer))
				}

			})

			c.Visit(url)
			time.Sleep(2 * time.Second)
			fmt.Println(hldUPC, manufacturer, toyName)
			writer.Write([]string{hldUPC, manufacturer, toyName})

			/*fmt.Println(hldUPC, manufacturer, toyName)
			row.Cells[3].Value = toyName
			row.Cells[4].Value = manufacturer
			xlNewFileName := "2" + xlFileName
			err = xlFile.Save(xlNewFileName)
			if err != nil {
				fmt.Printf(err.Error())
			}
			*/
			ctr++
		}
	}
	// Write CSV header

	// Instantiate default collector

	// Limit the number of threads started by colly to two
	// when visiting links which domains' matches "*httpbin.*" glob

	// Start scraping in four threads on https://httpbin.org/delay/2
	/*for i := 0; i < 4; i++ {
		c.Visit(fmt.Sprintf("%s?n=%d", url, i))
	}
	// Start scraping on https://httpbin.org/delay/2
	c.Visit(url)
	// Wait until threads are finished
	c.Wait()
	*/
}
