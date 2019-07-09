## 关于

 全局替换 api.service.com ,这个字符串是你根目录的名称，如 app_name
 
     D:/goWork/src/app_name

导入 gz2017.sql
修改 conf/ 下数据库配置

> 查看效果
1. go run main.go [下载缺少的包]
2. 访问 http://localhost:8008/health
3. 访问 http://localhost:8008/user/getUserInfo

## 开发
> - 包管理工具使用[glide](http://glidedocs.readthedocs.io/zh/latest/)
> - 安装glide: 
```bash
    #使用国人改进版本
    go get github.com/xkeyideal/glide
    go install github.com/xkeyideal/glide
```
> - 解决墙外的包下载： 将mirrors.yaml放到 用户目录 ~/.glide目录
> - 下载或更新包：glide install || glide update


## 接口验证
> - 请求URL必备2个参数timestamp和api_sign，例子：
> - POST /oss/uploadBase64Image 完整请求URL： 
> - http://vendor.api.7shangzuo.com.7vipseat.com/oss/uploadBase64Image?timestamp=1519645493&utm_medium=admin&api_sign=0452ce29789f86cda5ee3e3e4919ace3
> - 请求body: application/x-www-form-urlencoded  image=base64字符串
> - 返回：
```javascript
    {
        "status": "0000",
        "msg": "success",
        "data": {
            "uri": "images/20180225/4b2d7438aeba1ecd53196cc6938e2ddb.jpeg"
        }
    }
```

 