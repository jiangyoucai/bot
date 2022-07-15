# spf

国家邮政局邮政普遍服务营业场所名录查询，https://zwfw.spb.gov.cn/gjj/pfyycsml

# how to use

```golang
func main() {
	for i := 1; i < 55394; i++ {
		path := fmt.Sprintf("https://zwfw.spb.gov.cn/gjj/pfyycsml/pfyycsmlDetail?uuid=%d", i)
		fetch(path)
	}
}
```

output

```
Visiting https://zwfw.spb.gov.cn/gjj/pfyycsml/pfyycsmlDetail?uuid=1
&{246122 安徽省 安庆市 怀宁县 秀山邮政所 怀宁县秀山乡栏坝街道 信件、印刷品、包裹、邮政汇兑}

Visiting https://zwfw.spb.gov.cn/gjj/pfyycsml/pfyycsmlDetail?uuid=2
&{246199 安徽省 安庆市 怀宁县 雷埠邮政所 怀宁县雷埠乡集镇区212线东侧 信件、印刷品、包裹、邮政汇兑}
```

# data explain

```golang
type ZIP struct {
	Number   string // 邮编
	Province string // 省
	City     string // 市
	County   string // 县
	Name     string // 名称
	Address  string // 地址
	Scope    string // 范围
}
```
