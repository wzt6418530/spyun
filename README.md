# spyun
商鹏云打印机go语言版

## 打印示例
``` golang
package main

import (
	"fmt"
	"spyun"
	"time"
	"strconv"
)

func main() {
	nowTime:=strconv.Itoa(int(time.Now().Unix()))
	client:=new(spyun.YunClient)
	client.AppId="你的appid"
	client.AppSecret="你的appsecret"
	client.Sn="打印机编号"
	client.Content=`
<L1><C>测试打印</C></L1>
<CUT>
`
	client.Timestamp=nowTime
	res,err:=client.SendPost(spyun.PrinterPrint,*client)
	fmt.Println(res)
	fmt.Println(err)
}
```





