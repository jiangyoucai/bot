package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type ZIP struct {
	Number   string
	Province string
	City     string
	County   string
	Name     string
	Address  string
	Scope    string
}

func main() {
	for i := 1; i < 55394; i++ {
		path := fmt.Sprintf("https://zwfw.spb.gov.cn/gjj/pfyycsml/pfyycsmlDetail?uuid=%d", i)
		fetch(path)
	}
}

func fetch(path string) {
	c := colly.NewCollector(colly.MaxDepth(1))
	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		if e.Attr("class") != "infocente" {
			return
		}

		var zip ZIP
		e.ForEach("td", func(index int, i *colly.HTMLElement) {
			switch index {
			case 1:
				zip.Province = strings.TrimSpace(i.Text)
			case 3:
				zip.City = strings.TrimSpace(i.Text)
			case 5:
				zip.County = strings.TrimSpace(i.Text)
			case 7:
				zip.Name = strings.TrimSpace(i.Text)
			case 9:
				zip.Address = strings.TrimSpace(i.Text)
			case 11:
				zip.Scope = strings.TrimSpace(i.Text)
			case 13:
				zip.Number = strings.TrimSpace(i.Text)
			}
		})

		if zip.Number != "" {
			fmt.Println(zip)
			// toStorage(&zip)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(path)

}

// func toStorage(data *ZIP) {
// 	query := "INSERT INTO zip (number, name, province, city, county, address, scope) VALUES (:number, :name, :province, :city, :county, :address, :scope);"
// 	result, err := db.NewWrite().ItemNamed(query, data)
// 	fmt.Println(result, err)
// }
