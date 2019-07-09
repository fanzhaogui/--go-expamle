package lib

import (
	"api.service.com/db"
	"api.service.com/model"
	"api.service.com/service/wechat/token"
	"time"
)

const (
	// VendorAccessTokenRedisKey redis key
	VendorAccessTokenRedisKey = "api_service:access_token"
)


//GetToken redis
func GetToken(refresh bool) (string, error) {
	redis := db.GetRedis()
	sugar := GetLogger().Sugar()
	accessToken := redis.Get(VendorAccessTokenRedisKey).Val()

	if accessToken == "" || refresh {
		token, expire, err := token.AccessToken()

		sugar.Infof("token expire:%s", expire)

		if err != nil {
			sugar.Errorf("GetToken fail err:%v", err)
			return "", err
		}

		err = redis.Set(VendorAccessTokenRedisKey, token, time.Second*3600).Err()
		if err != nil {
			sugar.Errorf("restore token fail err:%v", err)
			return "", err
		}
		return token, nil
	}
	return accessToken, nil
}

/*
 * 获取配置
 */
func GetConfig(key string) (string, error) {
	redis := db.GetRedis()
	sugar := GetLogger().Sugar()
	cacheKey := "act:conf:" + key
	value := redis.Get(cacheKey).Val()

	if value == "" {
		var config model.Config
		db.GetDB().Where("option_key=?", key).First(&config)
		if config.Id > 0 {
			err := redis.Set(cacheKey, config.OptionValue, time.Second*3600).Err()
			if err != nil {
				sugar.Errorf("getconfig redis set error:%v", err)
				return "", nil
			}
			value = config.OptionValue
		}
	}
	return value, nil
}
