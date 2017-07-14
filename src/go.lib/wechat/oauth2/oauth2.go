package oauth2

import (
	"go.lib/wechat/config"
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)
const (
	APIURI 		= "https://api.weixin.qq.com/%s"
)

type WechatCodeData struct {
	Errcode			int
	Errmsg			string
	Access_token	string
	Expires_in		int
	Refresh_token	string
	Openid			string
	Scope			string
}

type WechatUserData struct {
	Errcode		int
	Errmsg		string
	Openid		string		// 用户的唯一标识
	Nickname	string		// 用户昵称
	Sex			int			// 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Province 	string		// 用户个人资料填写的省份
	City		string		// 普通用户个人资料填写的城市
	Country		string		// 国家，如中国为CN
	Headimgurl	string		// 用户头像
	Unionid		string		// 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
}

func GetAutoUserInfo(code string) (wechatUserInfo WechatUserData, err error) {
	u, _ := url.Parse(fmt.Sprintf(APIURI, "sns/oauth2/access_token"))
	q := u.Query()
	q.Set("appid", 		config.APPID)
	q.Set("secret", 	config.APPSECRET)
	q.Set("code", 		code)
	q.Set("grant_type", "authorization_code")

	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return wechatUserInfo, err
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return wechatUserInfo, err
	}

	codeInfo := WechatCodeData{}
	err = json.Unmarshal(result, &codeInfo)
	if err != nil {
		return wechatUserInfo, err
	}
	if codeInfo.Errcode != 0 {
		return wechatUserInfo, fmt.Errorf("%s", codeInfo.Errmsg)
	}

	uInfo, _ := url.Parse(fmt.Sprintf(APIURI, "sns/userinfo"))
	qInfo := uInfo.Query()
	qInfo.Set("access_token",	codeInfo.Access_token)
	qInfo.Set("openid", 		codeInfo.Openid)
	qInfo.Set("lang", 			"zh_CN")

	uInfo.RawQuery = qInfo.Encode()

	resInfo, err := http.Get(uInfo.String())
	if err != nil {
		return wechatUserInfo, err
	}

	resultInfo, err := ioutil.ReadAll(resInfo.Body)

	resInfo.Body.Close()
	if err != nil {
		return wechatUserInfo, err
	}
	err = json.Unmarshal(resultInfo, &wechatUserInfo)
	if err != nil {
		return wechatUserInfo, err
	}
	if wechatUserInfo.Errcode != 0 {
		return wechatUserInfo, fmt.Errorf("error value is %d", wechatUserInfo)
	}
	return wechatUserInfo, nil
}