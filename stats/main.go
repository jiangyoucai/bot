package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Place struct {
	Number string // 区划
	Name   string // 名称
	Origin string // 上级
	Level  int    // 等级
}

func main() {
	path := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/11.html"
	fetch(path)
}

func fetch(path string) {
	data := []interface{}{}
	c := colly.NewCollector(colly.MaxDepth(4))
	c.OnHTML("tr[class]", func(e *colly.HTMLElement) {
		switch e.Attr("class") {
		case "provincetr":
			e.ForEach("a[href]", func(index int, i *colly.HTMLElement) {
				item := Place{Level: 0}

				item.Name = i.Text
				item.Number = fmt.Sprintf("%s0000", i.Attr("href")[:2])

				if item.Number != "" {
					item.Origin = "0"
					data = append(data, &item)
				}

				path := i.Request.AbsoluteURL(i.Attr("href"))
				i.Request.Visit(path)

			})
		case "citytr":
			item := Place{Level: 1}
			e.ForEach("a[href]", func(index int, i *colly.HTMLElement) {
				switch index {
				case 0:
					item.Number = i.Text[:6]
				case 1:
					item.Name = i.Text

					path := i.Request.AbsoluteURL(i.Attr("href"))
					i.Request.Visit(path)
				}
			})
			if item.Number != "" {
				item.Origin = fmt.Sprintf("%s00", item.Number[:4])
				data = append(data, &item)
			}
		case "countytr":
			item := Place{Level: 2}
			e.ForEach("a[href]", func(index int, i *colly.HTMLElement) {
				switch index {
				case 0:
					item.Number = i.Text[:6]
				case 1:
					item.Name = i.Text

					path := i.Request.AbsoluteURL(i.Attr("href"))
					i.Request.Visit(path)
				}
			})
			if item.Number != "" {
				item.Origin = fmt.Sprintf("%s00", item.Number[:4])
				data = append(data, &item)
			}
		case "towntr":
			item := Place{Level: 3}
			e.ForEach("a[href]", func(index int, i *colly.HTMLElement) {
				switch index {
				case 0:
					item.Number = i.Text
				case 1:
					item.Name = i.Text

					path := i.Request.AbsoluteURL(i.Attr("href"))
					i.Request.Visit(path)
				}
			})
			if item.Number != "" {
				item.Origin = fmt.Sprintf("%s", item.Number[:6])
				data = append(data, &item)
			}
		case "villagetr":
			item := Place{Level: 4}
			e.ForEach("td", func(index int, i *colly.HTMLElement) {
				switch index {
				case 0:
					item.Number = i.Text
				case 2:

					item.Name = i.Text
				}
			})
			if item.Number != "" {
				item.Origin = fmt.Sprintf("%s000", item.Number[:9])
				data = append(data, &item)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(path)

	// todo
	// toStorage(data)
}

// func toStorage(data []interface{}) {
// 	query := "INSERT INTO village (number, name, origin, level) VALUES (:number, :name, :origin, :level);"
// 	result, err := db.NewWrite().ListNamed(query, data...)
// 	fmt.Println(result, err)
// }
