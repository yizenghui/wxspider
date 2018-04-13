// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wxspider

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
	// "golang.org/x/net/html/charset"
)

func Test_XTags(t *testing.T) {
	 tags := AiTags{}
	tt := `{"log_id": 7971186989667946038}`
	rbyte := []byte(tt)

	json.Unmarshal(rbyte, &tags)
	t.Fatal(tags)
}

func Test_GetTags(t *testing.T) {

	tt := "iphone手机出现“白苹果”原因及解决办法，用苹果手机的可以看下"
	cc := `如果下面的方法还是没有解决你的问题建议来我们门店看下成都市锦江区红星路三段99号银石广场24层01室。在通电的情况下掉进清水，这种情况一不需要拆机处理。尽快断电。用力甩干，但别把机器甩掉，主意要把屏幕内的水甩出来。如果屏幕残留有水滴，干后会有痕迹。^H3 放在台灯，射灯等轻微热源下让水分慢慢散去。`

	a := &Article{
		Title: tt,
		Cont:  cc,
	}
	a3 := &Article{
		Title: "你若记得，他便不悔",
		Cont: `也许，爱上他就在这一瞬间。

		1986年，我们高中毕业前夕，空军飞行学院在湖州招收飞行学员。这时我给他写了一封信，信中有这么一句话：“是一名真正的男子汉，你就去蓝天翱翔吧！”
		
		王伟说：“我的人生第一是飞行，第二是我们之间的爱情，它们对我来讲，就像飞机的两翼缺一不可。我一定要飞出来，否则决不回来见你！”“感情上，爱你胜过爱飞行，而理智上要我爱飞行比爱你更甚。”

		在此后的近三年的时间里，我们没有再见面。毕业后的王伟并没有立即回到分别了两年半的我的身边，而是直接去了海军航空兵的训练基地报到。他写信告诉我：
		
		“此生唯有飞行和你左右我。你能给我最大的幸福和快乐，飞行也能。我爱飞行事业，同样爱你。感情上，爱你胜过爱飞行，而理智上要我爱飞行比爱你更甚——这就是我目前（也许是一生）的主要矛盾，以及它们的辩证关系。”收到这样的来信，我觉得我爱的王伟成熟了。我在给王伟的信中写到：

		“在我的心中，一个热爱飞行事业、刚毅、坚强、奋斗的王伟，他吃得起苦、受得起累，他自信地向前走，这种对目标的追求和未来美好的渴望，使我热爱他。如此一个王伟，使我不在乎他对我关心有多少，不在乎他跟我在一起的时间有多长，不在乎别人一对对卿卿我我的甜蜜感受，不在乎他不能在我身边的孤独和寂寞，在乎的是他对未来的追求和生命的热爱。”当年江南雨巷中徘徊的少年终于成长为搏击蓝天的雄鹰了。

		向往蓝天的王伟是一个浪漫的人，我们的家中墙上挂着的是他亲手画的油画。一次，王伟画了一幅中国海军航空兵驾驶着未来的新型战机从战舰上起飞的油画挂在了家中的墙上，并对前来家中做客的团政委讲：“画上那名飞行员就是我自己！”我知道，这是他的追求与梦想！看着英姿勃勃的丈夫，我在想：部队真是锻炼人，当年在江南雨巷中徘徊的那个少年终于成长为搏击蓝天的雄鹰了。我觉得好幸福，在空旷的机场跑道旁，我和王伟手挽着手，漫步在一望无际的天地，享受落日的霞光……`,
	}
	a2 := &Article{
		Title: "If you remember, he will not regret",
		Cont: `Maybe, in love with him at this moment.



		In 1986, on the eve of our graduation from high school, the Air Force Flight Academy enrolled cadets in Huzhou. At this time, I wrote him a letter, which contained a sentence: "a real man, you can go to the blue sky."
		
		
		
		Wang Wei said, "the first thing in my life is flying, and the second is love between us. They are just like the two wings of a plane. I must fly out, or I will never come back to see you. "Emotionally, I love you more than flying, and intellectually, I love flying more than I love you."
		
		
		
		We haven't seen each other for nearly three years. Wang Wei didn't immediately return to my side after two and a half years, but went directly to the training base of naval air force to report. He wrote to me:
		
		
		
		"This life is only flying and you are about me. You can give me the greatest happiness and joy, and I can fly too. I love flying, and I love you. Emotionally, love you better than flying, and intellectually I love flying more than love you - this is the main contradiction of my present (perhaps life), and their dialectical relationship. " After receiving such a letter, I think Wang Wei, who I love, is mature. In my letter to Wang Wei, I wrote:
		
		
		
		"In my heart, Wang Wei, who loves flying, resolute, strong, and struggled, he eats hard and gets tired, he walks forward confidently, the pursuit of the goal and the good desire for the future, which makes me love him. Such a Wang Wei makes me not care about how much he cares for me, how long he is with me, the sweet feelings of a couple of people, the loneliness and loneliness that he can't be around me, the love of his future and his life. " The eagle hovering boy finally grows into fighting the blue sky in the Jiangnan rain lane.
		
		
		
		Wang Wei, who yearns for the blue sky, is a romantic man. He painted oil paintings on his own walls in our house. One time, Wang Wei painted a Chinese Navy airman driving a new aircraft that took off from the warship on the wall of the house, and said to the party political commissar who came to his home, "the pilot is myself!" I know, this is his pursuit and dream! Looking at the heroic and beautiful husband, I was thinking: the army was really a workout. The young man who was wandering in the rain lane in the south of the Yangtze finally grew into an eagle to fight the blue sky. I feel so happy that Wang Wei and I are walking hand in hand in the open field beside the open airport runway, enjoying the sunset. `,
	}
	// url := fmt.Sprintf(`https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=%v`, GetToken())
	// t.Fatal(url)
	// t.Fatal(a2.AiGetTags())
	t.Fatal(a.AiGetTags())
	t.Fatal(a2.AiGetTags())
	t.Fatal(a3.AiGetTags())
}

func Test_Cp(t *testing.T) {
	song := make(map[string]interface{})

	tt := "go语言标准json库Marshal不支持gbk格式string?"
	cc := `整个go的语言基础就是utf8，包括源码，运行时都是utf8，除此之外，还有很多编码，你觉得gbk常用是因为你是中国人，日本和韩国的还不一定呢。json那么实现，就简化整个json的解码实现，至于其它编码，再做个转码的就好了。`

	// gbkTitle, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(tt)), simplifiedchinese.GBK.NewEncoder()))

	// gbkContent, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(cc)), simplifiedchinese.GBK.NewEncoder()))

	song["title"] = tt
	song["content"] = cc
	// src := "编码转换内容内容"
	// enc := mahonia.NewEncoder("GBK")
	// // output := enc.ConvertString(tt)

	// song["title"] = enc.ConvertString(tt)
	// song["content"] = enc.ConvertString(cc)

	bytesData, err := json.Marshal(song)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	gbkBytesData, err := UTF8ToGBK(bytesData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	reader := bytes.NewReader(gbkBytesData)
	url := `https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=24.01a7fba39af897d7e5c3141b28962bd4.2592000.1526178157.282335-11067381`
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	// request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		t.Fatal(err.Error())
		return
	}

	// reader2, err := charset.NewReader(resp.Body, strings.ToLower(resp.Header.Get("Content-Type")))

	// respBytes, err := ioutil.ReadAll(reader2)

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err.Error())
		return
	}
	//byte数组直接转成string，优化内存
	str := string(respBytes)

	// t.Fatal(str)
	udata := ConvertStrEncode(str, "gbk", "utf-8")
	t.Fatal(udata)
}

func UrlEncode(instr string) string {
	return url.QueryEscape(instr)
}
func Test_Cp2(t *testing.T) {

	tt := "iphone手机出现“白苹果”原因及解决办法，用苹果手机的可以看下"
	cc := `如果下面的方法还是没有解决你的问题建议来我们门店看下成都市锦江区红星路三段99号银石广场24层01室。在通电的情况下掉进清水，这种情况一不需要拆机处理。尽快断电。用力甩干，但别把机器甩掉，主意要把屏幕内的水甩出来。如果屏幕残留有水滴，干后会有痕迹。^H3 放在台灯，射灯等轻微热源下让水分慢慢散去。`

	gbkTitle := UrlEncode(ConvertStrEncode(tt, "utf-8", "gbk"))

	gbkContent := UrlEncode(ConvertStrEncode(cc, "utf-8", "gbk"))

	post_arg := map[string]interface{}{
		"title":   gbkTitle,
		"content": gbkContent,
	}
	post_json, err := json.Marshal(post_arg)
	if nil != err {
		panic(err)
	}
	real_uri := `https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=24.01a7fba39af897d7e5c3141b28962bd4.2592000.1526178157.282335-11067381`

	req, err := http.NewRequest("POST", real_uri, bytes.NewReader(post_json))
	client := http.Client{}

	resp, _ := client.Do(req)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		panic(err)
	}

	udata := ConvertStrEncode(string(data), "gbk", "utf-8")

	if "null" == udata {
		panic(errors.New("貌似返回空"))
	}

	map_result := make(map[string]interface{})
	json.Unmarshal([]byte(udata), &map_result)

	panic(map_result)
	error_msg, ok := map_result["error_msg"]
	if ok {
		panic(error_msg)
	}

	// map_result["tags"].([]interface{})

	fmt.Println(map_result["tags"].([]interface{}))
}
