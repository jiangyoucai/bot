package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type Area struct {
	Number     string // 行政区划
	Name       string // 名称
	Population string // 人口
	Acreage    string // 面积
	Station    string // 驻地
	Area       string // 区号
	Zip        string // 邮编
	Level      string // 等级
}

func main() {
	fetch("北京市（京）", "", "", &Area{})
	fetch("天津市（津）", "", "", &Area{})
	fetch("河北省（冀）", "", "", &Area{
		Number:     "130000",
		Name:       "河北省",
		Station:    "石家庄市",
		Population: "7763",
		Acreage:    "190000",
		Level:      "0",
	})
	fetch("山西省（晋）", "", "", &Area{
		Number:     "140000",
		Name:       "山西省",
		Station:    "太原市",
		Population: "3540",
		Acreage:    "160000",
		Level:      "0",
	})
	fetch("内蒙古自治区（内蒙古）", "", "", &Area{
		Number:     "150000",
		Name:       "内蒙古自治区",
		Station:    "呼和浩特市",
		Population: "2442",
		Acreage:    "1180000",
		Level:      "0",
	})
	fetch("辽宁省（辽）", "", "", &Area{
		Number:     "210000",
		Name:       "辽宁省",
		Station:    "沈阳市",
		Population: "4190",
		Acreage:    "150000",
		Level:      "0",
	})
	fetch("吉林省（吉）", "", "", &Area{
		Number:     "220000",
		Name:       "辽宁省",
		Station:    "吉林市",
		Population: "2602",
		Acreage:    "190000",
		Level:      "0",
	})
	fetch("黑龙江省（黑）", "", "", &Area{
		Number:     "230000",
		Name:       "黑龙江省",
		Station:    "哈尔滨市",
		Population: "3555",
		Acreage:    "460000",
		Level:      "0",
	})
	fetch("上海市（沪）", "", "", &Area{})
	fetch("江苏省（苏）", "", "", &Area{
		Number:     "320000",
		Name:       "江苏省",
		Station:    "南京市",
		Population: "7858",
		Acreage:    "110000",
		Level:      "0",
	})
	fetch("浙江省（浙）", "", "", &Area{
		Number:     "330000",
		Name:       "浙江省",
		Station:    "杭州市",
		Population: "5039",
		Acreage:    "100000",
		Level:      "0",
	})
	fetch("安徽省（皖）", "", "", &Area{
		Number:     "340000",
		Name:       "安徽省",
		Station:    "合肥市",
		Population: "7119",
		Acreage:    "140000",
		Level:      "0",
	})
	fetch("福建省（闽）", "", "", &Area{
		Number:     "350000",
		Name:       "福建省",
		Station:    "福州市",
		Population: "3896",
		Acreage:    "120000",
		Level:      "0",
	})
	fetch("江西省（赣）", "", "", &Area{
		Number:     "360000",
		Name:       "江西省",
		Station:    "南昌市",
		Population: "5039",
		Acreage:    "170000",
		Level:      "0",
	})
	fetch("山东省（鲁）", "", "", &Area{
		Number:     "370000",
		Name:       "山东省",
		Station:    "济南市",
		Population: "10139",
		Acreage:    "160000",
		Level:      "0",
	})
	fetch("河南省（豫）", "", "", &Area{
		Number:     "410000",
		Name:       "河南省",
		Station:    "郑州市",
		Population: "11486",
		Acreage:    "170000",
		Level:      "0",
	})
	fetch("湖北省（鄂）", "", "", &Area{
		Number:     "420000",
		Name:       "湖北省",
		Station:    "武汉市",
		Population: "6178",
		Acreage:    "190000",
		Level:      "0",
	})
	fetch("湖南省（湘）", "", "", &Area{
		Number:     "430000",
		Name:       "湖南省",
		Station:    "长沙市",
		Population: "7320",
		Acreage:    "210000",
		Level:      "0",
	})
	fetch("广东省（粤）", "", "", &Area{
		Number:     "440000",
		Name:       "广东省",
		Station:    "广州市",
		Population: "9663",
		Acreage:    "180000",
		Level:      "0",
	})
	fetch("广西壮族自治区（桂）", "", "", &Area{
		Number:     "450000",
		Name:       "广西壮族自治区",
		Station:    "南宁市",
		Population: "5695",
		Acreage:    "240000",
		Level:      "0",
	})
	fetch("海南省（琼）", "", "", &Area{
		Number:     "460000",
		Name:       "海南省",
		Station:    "海口市",
		Population: "937",
		Acreage:    "34000",
		Level:      "0",
	})
	fetch("重庆市（渝）", "", "", &Area{})
	fetch("四川省（川、蜀）", "", "", &Area{
		Number:     "510000",
		Name:       "四川省",
		Station:    "成都市",
		Population: "9100",
		Acreage:    "490000",
		Level:      "0",
	})
	fetch("贵州省（黔、贵）", "", "", &Area{
		Number:     "520000",
		Name:       "贵州省",
		Station:    "贵阳市",
		Population: "4571",
		Acreage:    "180000",
		Level:      "0",
	})
	fetch("云南省（滇、云）", "", "", &Area{
		Number:     "530000",
		Name:       "云南省",
		Station:    "昆明市",
		Population: "4792",
		Acreage:    "390000",
		Level:      "0",
	})
	fetch("西藏自治区（藏）", "", "", &Area{
		Number:     "540000",
		Name:       "西藏自治区",
		Station:    "拉萨市",
		Population: "335",
		Acreage:    "1230000",
		Level:      "0",
	})
	fetch("陕西省（陕、秦）", "", "", &Area{
		Number:     "610000",
		Name:       "陕西省",
		Station:    "西安市",
		Population: "4052",
		Acreage:    "210000",
		Level:      "0",
	})
	fetch("甘肃省（甘、陇）", "", "", &Area{
		Number:     "620000",
		Name:       "甘肃省",
		Station:    "兰州市",
		Population: "2783",
		Acreage:    "430000",
		Level:      "0",
	})
	fetch("青海省（青）", "", "", &Area{
		Number:     "630000",
		Name:       "青海省",
		Station:    "西宁市",
		Population: "589",
		Acreage:    "720000",
		Level:      "0",
	})
	fetch("宁夏回族自治区（宁）", "", "", &Area{
		Number:     "640000",
		Name:       "宁夏回族自治区",
		Station:    "银川市",
		Population: "686",
		Acreage:    "66000",
		Level:      "0",
	})
	fetch("新疆维吾尔自治区（新）", "", "", &Area{
		Number:     "650000",
		Name:       "新疆维吾尔自治区",
		Station:    "乌鲁木齐市",
		Population: "2286",
		Acreage:    "1660000",
		Level:      "0",
	})
	fetch("香港特别行政区（港）", "", "", &Area{})
	fetch("澳门特别行政区（澳）", "", "", &Area{})
}

func fetch(province, city, country string, args *Area) {
	data := []*Area{args}

	c := colly.NewCollector()
	c.OnHTML("tr[class]", func(e *colly.HTMLElement) {

		item := Area{Level: "1"}

		e.ForEach("td", func(index int, i *colly.HTMLElement) {
			switch index {
			case 0:
				item.Name = strings.ReplaceAll(i.Text, "+", "")
			case 1:
				item.Station = strings.TrimSpace(i.Text)
			case 2:
				item.Population = strings.TrimSpace(i.Text)
			case 3:
				item.Acreage = strings.TrimSpace(i.Text)
			case 4:
				item.Number = strings.TrimSpace(i.Text)
			case 5:
				item.Area = strings.TrimSpace(i.Text)
			case 6:
				item.Zip = strings.TrimSpace(i.Text)
			}
		})
		data = append(data, &item)
	})

	// Find and visit all links
	c.OnHTML("tr[type]", func(e *colly.HTMLElement) {
		item := Area{Level: "2"}

		e.ForEach("td", func(index int, i *colly.HTMLElement) {
			switch index {
			case 0:
				item.Name = strings.ReplaceAll(i.Text, "+", "")
			case 1:
				item.Station = strings.TrimSpace(i.Text)
			case 2:
				item.Population = strings.TrimSpace(i.Text)
			case 3:
				item.Acreage = strings.TrimSpace(i.Text)
			case 4:
				item.Number = strings.TrimSpace(i.Text)
			case 5:
				item.Area = strings.TrimSpace(i.Text)
			case 6:
				item.Zip = strings.TrimSpace(i.Text)
			}
		})
		data = append(data, &item)
	})

	c.OnScraped(func(r *colly.Response) {
		result := []interface{}{}
		for _, v := range data {
			if v.Number != "" {
				if v.Acreage == "" || v.Population == "" {
					v.Population = "0"
					v.Acreage = "0"
				}
				fmt.Println(v)

				result = append(result, v)
			}
		}

		// todo
		toStorage(result)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	path := fmt.Sprintf("http://xzqh.mca.gov.cn/defaultQuery?shengji=%s&diji=%s&xianji=%s", toGBK(province), toGBK(city), toGBK(country))

	c.Visit(path)
}

func toGBK(s string) string {
	tmp, err := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(s)) //utf-8 转 gbk
	if err != nil {
		panic(err)
	}
	return url.QueryEscape(string(tmp))
}

func toStorage(data []interface{}) {
	// query := "INSERT INTO area(number, name, station, population, acreage, area, zip, level) VALUES (:number, :name, :station, :population, :acreage, :area, :zip, :level);"
	// result, err := db.NewWrite().ListNamed(query, data...)
	// fmt.Println(result, err)
}
