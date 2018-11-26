package wxspider

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/yizenghui/sda/wechat"
)

func init() {
	DB().AutoMigrate(&Article{})

}

//SpiderArticle 采集文章并保存到本地
func SpiderArticle(urlStr string) error {
	var a Article
	article, err := wechat.Find(urlStr)
	if err == nil {

		if article.URL == "" {
			return errors.New("不支持该链接！")
		}

		a.GetArticleByURL(article.URL)
		a.AppID = article.AppID
		a.AppName = article.AppName
		a.RoundHead = article.RoundHead
		a.OriHead = article.OriHead
		a.URL = article.URL
		a.Title = article.Title
		a.Intro = article.Intro
		a.Cover = article.Cover
		a.Author = article.Author
		a.PubAt = article.PubAt
		a.Cont = article.Content
		a.Body = article.MdContent
		a.Copyright = article.Copyright
		a.WxID = article.WxID
		a.WxIntro = article.WxIntro
		a.Video = article.Video
		a.Audio = article.Audio
		a.Category = `其它`

		// data["tags"] = []string{`php`, `golang`}
		tags, err := a.AiGetTags()
		if err == nil {
			var tarr []string
			for _, t := range tags.Items {
				tarr = append(tarr, t.Tag)
			}
			if len(tarr) > 0 {
				tagStr := strings.Join(tarr, ",")
				a.Tags = tagStr
			}
		}
		// 增加夜读标签
		if strings.Contains(a.Title, `夜读`) && a.Audio != `` {
			if a.Tags != `` {
				a.Tags = fmt.Sprintf(`%v,夜读`, a.Tags)
			} else {
				a.Tags = `夜读`
			}
		}

		if len(article.Images) > 0 {
			var imgarr []string
			for _, img := range article.Images {
				if CheckImage(img) {
					imgarr = append(imgarr, img)
				}
			}
			if len(imgarr) > 0 {
				a.Images = strings.Join(imgarr, ";")
				// log.Println(`img`, a.Images)
			}
		}

		// // 增加音频标签
		// if a.Audio != `` {
		// 	if a.Tags != `` {
		// 		a.Tags = fmt.Sprintf(`%v,音频`, a.Tags)
		// 	} else {
		// 		a.Tags = `音频`
		// 	}
		// }
		// // 视频标签
		// if a.Video != `` {
		// 	if a.Tags != `` {
		// 		a.Tags = fmt.Sprintf(`%v,视频`, a.Tags)
		// 	} else {
		// 		a.Tags = `视频`
		// 	}
		// }

		categories, err := a.AiGetCategories()
		if err == nil {
			var carr1 []string
			for _, t := range categories.Item.TopCategory {
				carr1 = append(carr1, t.Tag)
			}
			if len(carr1) > 0 {
				categoryStr1 := strings.Join(carr1, ",")
				a.Category = categoryStr1 // 一级分类
			}

			var carr []string
			for _, t := range categories.Item.SecondCatrgory { //二级分类
				carr = append(carr, t.Tag)
			}
			if len(carr) > 0 {
				categoryStr := strings.Join(carr, ",")
				a.Categories = categoryStr
			}
		}

		a.Save()
		log.Println("spider", a.ID, a.Title, a.URL, a.Category, a.Categories, a.Tags)
	}
	return nil
}

//PublishArticle 发布本地文章
func PublishArticle() error {
	var a Article
	rows := a.GetPlanPublushArticle()
	for _, row := range rows {
		time.Sleep(time.Second * 2)
		pistID, e := PostArticle(row)
		if e == nil {
			row.PublishAt = time.Now().Unix()
			row.PostID = pistID
			row.Save()
			log.Println("post", row.ID, row.Title, row.URL, row.PublishAt)
		} else {
			row.PublishAt = -1
			row.Save()
			log.Println("post err", row.ID, row.Title, row.URL)
		}
	}
	return nil
}

//GetArticles 获取文章列表
func GetArticles() []Article {
	var a Article
	rows := a.GetArticles()
	return rows
}

//PostMessage 发布数据返回消息 成功返回 id 失败返回json
type PostMessage struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

//PostArticle 发布文章
func PostArticle(article Article) (int64, error) {

	client := http.Client{}
	// 获取系统配置(发送地址和授权令牌)
	cf := GetConf()

	data := make(url.Values)
	data["title"] = []string{article.Title}
	data["app_id"] = []string{article.AppID}
	data["app_name"] = []string{article.AppName}
	data["app_cover"] = []string{article.OriHead}
	data["url"] = []string{article.URL}
	data["intro"] = []string{article.Intro}
	data["body"] = []string{article.Cont}
	data["mdbody"] = []string{article.Body}
	data["cover"] = []string{article.Cover}
	data["author"] = []string{article.Author}
	data["wxid"] = []string{article.WxID}
	data["wxintro"] = []string{article.WxIntro}
	data["tags"] = []string{article.Tags}
	data["images"] = []string{article.Images}
	data["category"] = []string{article.Category}
	data["categories"] = []string{article.Categories}
	data["copyright"] = []string{article.Copyright}
	data["video"] = []string{article.Video}
	data["audio"] = []string{article.Audio}
	// post 时把授权密码放在数据包里面
	// data["authorization_token"] = []string{cf.PostConfig.AuthorizationToken}

	i64, err := strconv.ParseInt(article.PubAt, 10, 64)
	if err != nil {
		// fmt.Println(err)
		return 0, errors.New("时间转化失败")
	}
	pubAt := time.Unix(i64, 0).Format("2006-01-02 15:04:05")
	data["pub_at"] = []string{pubAt}

	// 把数据包post到配置的位置
	resp, err := client.PostForm(cf.PostConfig.ServeURL, data)
	// resp, err := http.NewRequest("POST", cf.PostConfig.ServeURL, strings.NewReader(data.Encode()))

	// log.Println(cf.PostConfig.ServeURL)
	resp.Header.Add("accept", `application/json`)
	resp.Header.Add("Authorization", fmt.Sprintf(`Bearer %v`, cf.PostConfig.AuthorizationToken))

	if err != nil {
		// log.Println(" %v  ", err.Error)
		return 0, err
	}

	// formPost, err := goquery.NewDocumentFromReader(resp.Body)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()

	//byte数组直接转成string，优化内存
	// postMsg := string(respBytes)
	// postMsg, err := formPost.Html()
	// // // panic(err)
	// log.Println(" %s  ", postMsg)
	if err != nil {
		// log.Println(" %s  ", err.Error)
		return 0, err
	}

	var pm PostMessage
	err = json.Unmarshal(respBytes, &pm)
	if err != nil {
		log.Println(string(respBytes))
		return 0, errors.New(`服务器 返回数据出错`)
	}
	if pm.Message != `` {
		return 0, errors.New(pm.Message)
	}
	fmt.Println(pm)
	return pm.ID, nil
	// i64, err = strconv.ParseInt(postMsg, 10, 64)
	// if err != nil {
	// 	log.Println(postMsg)
	// 	// log.Println(" %s  ", err.Error)
	// 	return 0, err
	// }
	// fmt.Println("posted id", i64)
	// return i64, nil

	// if b := strings.Contains(postMsg, `mp.weixin.qq.com`); b == true {
	// 	return nil
	// }
	// return errors.New(`发布失败了`)
}
