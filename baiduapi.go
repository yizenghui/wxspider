package wxspider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

//AiData 要获取描述的数据
type AiData struct {
	Title string
	Desc  string
}

type AiTags struct {
	LogID     int64   `json:"log_id"`
	Items     []AiTag `json:"items"`
	ErrorMSG  string  `json:"error_msg"`
	ErrorCode int     `json:"error_code"`
}

type AiTag struct {
	Score float64 `json:"score"`
	Tag   string  `json:"tag"`
}

type AiTagParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AiCategories struct {
	LogID     int64          `json:"log_id"`
	Item      AiCategoryItem `json:"item"`
	ErrorMSG  string         `json:"error_msg"`
	ErrorCode int            `json:"error_code"`
}
type AiCategoryItem struct {
	TopCategory    []AiTag `json:"lv1_tag_list"`
	SecondCatrgory []AiTag `json:"lv2_tag_list"`
}

//UTF8ToGBK utf8 2 gbk
func UTF8ToGBK(utf8 []byte) (gbk []byte, err error) {
	reader := transform.NewReader(
		bytes.NewReader(utf8), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 字符串转编码
func ConvertStrEncode(inStr, inCharset, outCharset string) string {
	if outCharset == "" {
		outCharset = inCharset
	}

	inCharset = strings.ToLower(inCharset)
	outCharset = strings.ToLower(outCharset)

	if inCharset == outCharset {
		return inStr
	}

	if inCharset == "gbk" || inCharset == "gb2312" {
		inCharset = "gb18030"
	}

	// 输入字符串解码为utf-8
	var destr string
	if inCharset != "utf8" && inCharset != "utf-8" {
		destr = mahonia.NewDecoder(inCharset).ConvertString(inStr)
	} else {
		destr = inStr
	}

	if outCharset == "utf8" || outCharset == "utf-8" {
		return destr
	}
	// 转换为 outCharset
	return mahonia.NewEncoder(outCharset).ConvertString(destr)
}

//AiGetTags 文章通过ai.baidu.com 获取标签
func (a Article) AiGetTags() (tags AiTags, err error) {

	data := make(map[string]interface{})

	if a.Title == "" {
		return
	}
	if a.Cont == "" {
		return
	}
	data["title"] = a.Title
	data["content"] = a.Cont

	log.Println("tags: title,content", a.Title, a.Cont)

	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	gbkBytesData, err := UTF8ToGBK(bytesData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(gbkBytesData)

	url := `https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=24.01a7fba39af897d7e5c3141b28962bd4.2592000.1526178157.282335-11067381`
	// url := fmt.Sprintf(`https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=%v`, GetToken())

	request, err := http.NewRequest("POST", url, reader)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}

	// reader2, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))

	// respBytes, err := ioutil.ReadAll(reader2)

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	//byte数组直接转成string，优化内存
	str := string(respBytes)

	udata := ConvertStrEncode(str, "gbk", "utf-8")

	rbyte := []byte(udata)

	json.Unmarshal(rbyte, &tags)

	return tags, nil
}

//AiGetCategories 文章通过ai.baidu.com 获取标签
func (a Article) AiGetCategories() (tags AiCategories, err error) {

	data := make(map[string]interface{})

	data["title"] = a.Title
	data["content"] = a.Cont

	if a.Title == "" {
		return
	}
	if a.Cont == "" {
		return
	}
	log.Println("categories: title,content", a.Title, a.Cont)

	bytesData, err := json.Marshal(data)
	if err != nil {
		return
	}
	gbkBytesData, err := UTF8ToGBK(bytesData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(gbkBytesData)

	url := `https://aip.baidubce.com/rpc/2.0/nlp/v1/topic?access_token=24.01a7fba39af897d7e5c3141b28962bd4.2592000.1526178157.282335-11067381`
	// url := fmt.Sprintf(`https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=%v`, GetToken())

	request, err := http.NewRequest("POST", url, reader)

	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	//byte数组直接转成string，优化内存
	str := string(respBytes)

	udata := ConvertStrEncode(str, "gbk", "utf-8")

	// log.Fatal(udata)
	// t.Fatal(str)
	rbyte := []byte(udata)

	json.Unmarshal(rbyte, &tags)

	return tags, nil
}
