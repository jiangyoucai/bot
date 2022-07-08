# stats

国家统计局统计用区划和城乡划分代码划抓取，http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/index.html

# how to use


```golang
func main() {
	fetch("http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/index.html")
}
```

output
```
Visiting http://xzqh.mca.gov.cn/defaultQuery?shengji=%B1%B1%BE%A9%CA%D0%A3%A8%BE%A9%A3%A9&diji=&xianji=

&{110000 北京市 0 0}
&{110101 东城区 110000 1}
&{110102 西城区 110000 1}
&{110105 朝阳区 110000 1}
&{110106 丰台区 110000 1}
&{110107 石景山区 110000 1}
&{110108 海淀区 110000 1}
&{110109 门头沟区 110000 1}
&{110111 房山区 110000 1}
&{110112 通州区 110000 1}
&{110113 顺义区 110000 1}
&{110114 昌平区 110000 1}
&{110115 大兴区 110000 1}
&{110116 怀柔区 110000 1}
&{110117 平谷区 110000 1}
&{110118 密云区 110000 1}
&{110119 延庆区 110000 1}
```

# data explain

```golang
type Place struct {
	Number string // 区划
	Name   string // 名称
	Origin string // 上级
	Level  int    // 等级
}
```