// Package token 微信 access_token
package token

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"api.service.com/service/wechat"
)

const tokenAPI = "/cgi-bin/token"

// 获取 access_token 成功返回数据
type response struct {
	wechat.Response
	AccessToken string `json:"access_token"`
	ExpireIn    int    `json:"expires_in"`
}

// AccessToken 通过微信服务器获取 access_token 以及其有效期
func AccessToken() (string, int, error) {
	url, err := url.Parse(wechat.BaseURL + tokenAPI)
	if err != nil {
		return "", 0, err
	}

	query := url.Query()

	query.Set("appid", wechat.AppID)
	query.Set("secret", wechat.AppSecret)
	query.Set("grant_type", "client_credential")

	url.RawQuery = query.Encode()

	res, err := http.Get(url.String())
	if err != nil {
		return "", 0, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", 0, errors.New(wechat.WeChatServerError)
	}

	var data response
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", 0, err
	}

	if data.Errcode != 0 {
		return "", 0, errors.New(data.Errmsg)
	}

	return data.AccessToken, data.ExpireIn, nil
}
