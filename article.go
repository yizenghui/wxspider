package wxspider

import (
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
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
		a.Save()
		log.Println("spider", a.ID, a.Title, a.URL)
	}
	return nil
}

//PublishArticle 发布本地文章
func PublishArticle() error {
	var a Article
	rows := a.GetPlanPublushArticle()
	for _, row := range rows {
		e := PostArticle(row)
		if e == nil {
			row.PublishAt = time.Now().Unix()
			row.Save()
			log.Println("post", row.ID, row.Title, row.URL, row.PublishAt)
		}
	}
	return nil
}

//PostArticle 采集文章并保存到本地
func PostArticle(article Article) error {

	client := http.Client{}
	data := make(url.Values)
	data["title"] = []string{article.Title}
	data["app_id"] = []string{article.AppID}
	data["app_name"] = []string{article.AppName}
	data["app_cover"] = []string{article.OriHead}
	data["url"] = []string{article.URL}
	data["intro"] = []string{article.Intro}
	data["cover"] = []string{article.Cover}
	data["author"] = []string{article.Author}

	i64, err := strconv.ParseInt(article.PubAt, 10, 64)
	if err != nil {
		// fmt.Println(err)
		return errors.New("时间转化失败")
	}
	pubAt := time.Unix(i64, 0).Format("2006-01-02 15:04:05")

	data["pub_at"] = []string{pubAt}
	data["category"] = []string{`测试`}
	// data["tags"] = []string{`php`, `golang`}

	// tags category

	resp, err := client.PostForm("http://wxapi.readfollow.com/api/v1/article", data)
	resp.Body.Close()
	if err != nil {
		// log.Println(" %s  ", err.Error)
		return err
	}

	// formPost, err := goquery.NewDocumentFromReader(resp.Body)

	// resp.Body.Close()

	// postMsg, err := formPost.Html()
	// // panic(err)
	// log.Println(" %s  ", postMsg)
	return nil
}
