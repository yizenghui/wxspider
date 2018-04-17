package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	_ "github.com/CodyGuo/godaemon"
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"

	"github.com/Kevingong2013/wechat"
	wxs "github.com/yizenghui/wxspider"
)

var logger = logrus.WithFields(logrus.Fields{
	"module": "wxbot",
})

// Guard ...
type guard struct {
	bot *wechat.WeChat
}

func newGuard(bot *wechat.WeChat) *guard {
	return &guard{bot}
}

// AddFriend ...
func (g *guard) addFriend(username, content string) error {
	return g.verifyUser(username, content, 2)
}

// AcceptAddFriend ...
func (g *guard) acceptAddFriend(username, content string) error {
	return g.verifyUser(username, content, 3)
}

func (g *guard) verifyUser(username, content string, status int) error {

	url := fmt.Sprintf(`%s/webwxverifyuser?r=%s&%s`, g.bot.BaseURL, strconv.FormatInt(time.Now().Unix(), 10), g.bot.PassTicketKV())

	data := map[string]interface{}{
		`BaseRequest`:        g.bot.BaseRequest,
		`Opcode`:             status,
		`VerifyUserListSize`: 1,
		`VerifyUserList`: map[string]string{
			`Value`:            username,
			`VerifyUserTicket`: ``,
		},
		`VerifyContent`:  content,
		`SceneListCount`: 1,
		`SceneList`:      33,
		`skey`:           g.bot.BaseRequest.Skey,
	}

	bs, _ := json.Marshal(data)

	var resp wechat.Response

	err := g.bot.Excute(url, bytes.NewReader(bs), &resp)
	if err != nil {
		return err
	}
	if resp.IsSuccess() {
		return nil
	}
	return resp.Error()
}

func (g *guard) autoAcceptAddFirendRequest(msg wechat.EventMsgData) {
	if msg.MsgType == 37 {
		rInfo := msg.OriginalMsg[`RecommendInfo`].(map[string]interface{})
		err := g.addFriend(rInfo[`UserName`].(string),
			msg.OriginalMsg[`Ticket`].(string))
		if err != nil {
			logger.Error(err)
		}
		err = g.bot.SendTextMsg(`新添加了一个好友`, `filehelper`)
		if err != nil {
			logger.Error(err)
		}
	}
}

func pushLink(u string) {
	// fmt.Println("push url:", u)
	wxs.SpiderArticle(u)
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func regAndPostWxLink(content string) {
	if content != "" {
		reg := regexp.MustCompile(`http://mp.weixin.qq.com/s(?P<uri>([?a-zA-Z_=0-9&;\/]+))`)
		list := reg.FindAllString(content, -1)
		arr := make(map[string]string)

		for _, u := range list {
			arr[u] = "1"
		}
		for u := range arr {
			pushLink(u)
		}
	}
}

func postLocationArticle() {
	ticker := time.NewTicker(time.Second * 30)
	for _ = range ticker.C {
		go wxs.PublishArticle()
	}
}

func web() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, wxs.GetArticles())
	})
	e.Logger.Fatal(e.Start(":3355"))
}

func main() {
	go postLocationArticle()
	go web()
	options := wechat.DefaultConfigure()
	options.Debug = false
	bot, err := wechat.AwakenNewBot(options)
	if err != nil {
		panic(err)
	}
	g := newGuard(bot)

	bot.Handle(`/msg/solo`, func(evt wechat.Event) {

		data := evt.Data.(wechat.EventMsgData)
		go g.autoAcceptAddFirendRequest(data)

		// data := evt.Data.(wechat.EventMsgData)
		// if data.MsgType == 49 {
		// 	fmt.Println(`/msg/solo/` + data.OriginalMsg[`Url`].(string))
		// 	go pushLink(data.OriginalMsg[`Url`].(string))
		// }

		go regAndPostWxLink(data.Content)
		// fmt.Println(`/msg/solo/` + data.Content)
	})

	bot.Handle(`/msg/group`, func(evt wechat.Event) {
		data := evt.Data.(wechat.EventMsgData)
		fmt.Println(`/msg/group/` + data.Content)
	})

	bot.Handle(`/contact`, func(evt wechat.Event) {
		data := evt.Data.(wechat.EventContactData)
		fmt.Println(`/contact` + data.GGID)
	})

	bot.Handle(`/login`, func(arg2 wechat.Event) {
		isSuccess := arg2.Data.(int) == 1
		if isSuccess {
			fmt.Println(`login Success`)
		} else {
			fmt.Println(`login Failed`)
		}
	})

	// 8h 发一次消息
	bot.AddTimer(8 * time.Hour)
	bot.Handle(`/timer/8h`, func(arg2 wechat.Event) {
		data := arg2.Data.(wechat.EventTimerData)
		if bot.IsLogin {
			bot.SendTextMsg(fmt.Sprintf(`第%v次`, data.Count), `filehelper`)
		}
	})

	// 9:00 每天9点发一条消息
	bot.AddTiming(`9:00`)
	bot.Handle(`/timing/9:00`, func(arg2 wechat.Event) {
		// data := arg2.Data.(wechat.EventTimingtData)
		bot.SendTextMsg(`9:00 了`, `filehelper`)
	})
	bot.AddTiming(`17:00`)
	bot.Handle(`/timing/17:00`, func(arg2 wechat.Event) {
		// data := arg2.Data.(wechat.EventTimingtData)
		bot.SendTextMsg(`17:00 i m runing`, `filehelper`)
	})

	bot.Go()
}
