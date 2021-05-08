package spyun

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	url2 "net/url"
	"sort"
	"strings"
)

type YunClient struct {
	AppId     string `json:"appid"`
	AppSecret string `json:"appsecret"`
	Timestamp string `json:"timestamp"`
	Sn        string `json:"sn"`
	PKey      string `json:"pkey"`
	Name      string `json:"name"`
	AutoCut   string `json:"auto_cut"`
	Voice     string `json:"voice"`
	Content   string `json:"content"`
	Times     string `json:"times"`
	Id        string `json:"id"`
	Date      string `json:"date"`
	Sign string `json:"sign"`
}

func (c *YunClient) SendPost(url string,yun YunClient) (*http.Response,error) {
	spYunMap,err:=c.ToMap(yun)
	if err!=nil{
		return nil,err
	}
	c.Sign=c.ToSign(spYunMap)
	data:=make(url2.Values)
	for k,v:=range spYunMap{
		if v!=""{
			data[k]=[]string{v}
		}
	}
	data["sign"]=[]string{c.Sign}
	res,err:=http.PostForm(url,data)
	if err!=nil{
		return nil,err
	}
	return res,nil
}
func (c *YunClient) SendGet(url string,yun YunClient) (*http.Response,error){
	spYunMap,err:=c.ToMap(yun)
	if err!=nil{
		return nil,err
	}
	c.Sign=c.ToSign(spYunMap)
	data:=make(url2.Values)
	for k,v:=range spYunMap{
		if v!=""{
			data[k]=[]string{v}
		}
	}
	res,err:=http.Get(url)
	if err!=nil{
		return nil,err
	}
	return res,nil
}

//签名
func (c *YunClient) ToSign(params map[string]string) string {
	//ASCII码从小到大排序
	var keys []string
	for k, v := range params {
		if v != "" && k != "appsecret" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	//缓冲区
	var buf bytes.Buffer
	for _, k := range keys {
		if len(params) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params[k])
			buf.WriteString(`&`)
		}
	}
	buf.WriteString("appsecret=" + c.AppSecret)
	//md5加密
	dataMd5 := md5.Sum(buf.Bytes())
	str := hex.EncodeToString(dataMd5[:]) //需转换成切片
	return strings.ToUpper(str)
}

//请求参数转map
func (c *YunClient) ToMap(yun YunClient) (map[string]string, error) {
	m := make(map[string]string)
	j, _ := json.Marshal(yun)
	err := json.Unmarshal(j, &m)
	return m, err
}
