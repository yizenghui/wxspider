package wxspider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

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

	data["title"] = a.Title
	data["content"] = a.Cont

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

	// log.Fatal(udata)
	// t.Fatal(str)
	rbyte := []byte(udata)

	json.Unmarshal(rbyte, &tags)

	// rp := resp.Response()
	// // reader, err := charset.NewReader(rp.Body, strings.ToLower(rp.Header.Get("Content-Type")))
	// // bs, _ := ioutil.ReadAll(reader)
	// bs, _ := ioutil.ReadAll(rp.Body)
	// htmlStr := string(bs)
	// log.Fatal(htmlStr)
	// log.Fatal(htmlStr)
	// // reader, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))
	// // log.Fatel(reader)
	// // "title":   gbk_title,
	// // "content": a.Cont,
	// if err != nil {
	// 	return tags, err
	// }

	// // log.Fatal(resp.Response)
	// // tags := new(AiTags)
	// if err := resp.ToJSON(tags); err != nil {
	// 	return tags, err
	// }
	return tags, nil
}

//httpClient 默认http.Client
var httpClient *http.Client

func init() {
	client := *http.DefaultClient
	client.Timeout = time.Second * 5
	httpClient = &client
}

// HTTPGetJSON 通过传入url和结构，提取出页面中的值
func HTTPGetJSON(url string, response interface{}) error {
	httpResp, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return DecodeJSONHttpResponse(httpResp.Body, response)
}

//HTTPPostJSON  通过传入url和内容，提交内容后，提取出页面中的值
func HTTPPostJSON(url string, body []byte, response interface{}) error {
	httpResp, err := httpClient.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return fmt.Errorf("http.Status: %s", httpResp.Status)
	}
	return DecodeJSONHttpResponse(httpResp.Body, response)
}

//DecodeJSONHttpResponse 解决json
func DecodeJSONHttpResponse(r io.Reader, v interface{}) error {
	body, err := ioutil.ReadAll(r)

	log.Fatal(string(body))
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
