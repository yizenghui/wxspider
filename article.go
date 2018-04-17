package wxspider

import (
	"errors"
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
		e := PostArticle(row)
		if e == nil {
			row.PublishAt = time.Now().Unix()
			row.Save()
			// time.Sleep(time.Second)
			log.Println("post", row.ID, row.Title, row.URL, row.PublishAt)
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
	data["tags"] = []string{article.Tags}
	data["category"] = []string{article.Category}
	data["categories"] = []string{article.Categories}

	i64, err := strconv.ParseInt(article.PubAt, 10, 64)
	if err != nil {
		// fmt.Println(err)
		return errors.New("时间转化失败")
	}
	pubAt := time.Unix(i64, 0).Format("2006-01-02 15:04:05")
	data["pub_at"] = []string{pubAt}

	// tags category

	// resp, err := client.PostForm("http://wxapi.readfollow.com/api/v1/article", data)
	resp, err := client.PostForm("https://wx.readfollow.com/api/links/", data)
	resp.Body.Close()
	if err != nil {
		// log.Println(" %s  ", err.Error)
		return err
	}

	// formPost, err := goquery.NewDocumentFromReader(resp.Body)

	// resp.Body.Close()

	// postMsg, err := formPost.Html()
	// // // panic(err)
	// log.Println(" %s  ", postMsg)
	return nil
}
