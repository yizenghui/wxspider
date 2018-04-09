package wxspider

import (
	"github.com/imroc/req"
)

//AiData 要获取描述的数据
type AiData struct {
	Title string
	Desc  string
}

type AiTags struct {
	LogID string `json:"log_id"`
	Items []AiTag
}

type AiTag struct {
	Score string `json:"score"`
	Tag   string `json:"tag"`
}

func (a Article) AiGetTags() (tags AiTags, err error) {

	url := `https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=24.ef1d835acc4ff52eeaa9e1806ad5442d.2592000.1525877150.282335-10531269`
	// url := fmt.Sprintf(`https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?access_token=%v`, GetToken())
	resp, err := req.Post(url, req.Param{
		"title":   a.Title,
		"content": a.Cont,
	})
	if err != nil {
		return tags, err
	}
	// tags := new(AiTags)
	if err := resp.ToJSON(tags); err != nil {
		return tags, err
	}
	return tags, nil
}

// type AuthResponse struct {
// 	AccessToken      string `json:"access_token"`  //要获取的Access Token
// 	ExpireIn         string `json:"expire_in"`     //Access Token的有效期(秒为单位，一般为1个月)；
// 	RefreshToken     string `json:"refresh_token"` //以下参数忽略，暂时不用
// 	Scope            string `json:"scope"`
// 	SessionKey       string `json:"session_key"`
// 	SessionSecret    string `json:"session_secret"`
// 	ERROR            string `json:"error"`             //错误码；关于错误码的详细信息请参考鉴权认证错误码(http://ai.baidu.com/docs#/Auth/top)
// 	ErrorDescription string `json:"error_description"` //错误描述信息，帮助理解和解决发生的错误。
// }
