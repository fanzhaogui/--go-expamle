package lib

import (
	"sort"
	"api.service.com/util"
)

// 生成api_sign
func MakeSign(params map[string]string, screctKey string) string {

	if len(params) == 0 {
		return ""
	}

	// 获取所有键名，用于排序
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}

	// 按键正向排序
	sort.Strings(keys)

	// 将map拼接成字符串
	var signData string
	for _, v := range keys {
		signData += v + "=" + params[v] + "&"
	}

	// 末尾添加密钥
	signData += screctKey
	//println(signData)
	return util.MD5(signData)
}

type ApiRes struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// 初始化api返回数据
func InitApiRes() *ApiRes {
	apiRes := new(ApiRes)
	apiRes.Status = Success
	apiRes.Msg = Msg(Success)
	apiRes.Data = make(map[string]interface{})

	return apiRes
}
