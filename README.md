# wxspider
采集已关注的微信公众号新发布的文章，发送到指定服务器

示例：

https://wechatrank.com/

## 使用

1. 匹配`conf.toml`
2. 下载wxspider.exe并运行


## 实现过程

通过微信网页版相关接口，获得微信消息推送数据，匹配其中微信公众号文章的链接地址。
采集公众号文章地址相应的内容后，通过百度AI进行分类和标签，正文转换成markdown语法，发送到指定位置。

## 准备工作

1. 百度AI接口密令，需要文章分类和文章标签权限，请自行申请 https://ai.baidu.com
2. 采集服务器一台(一个64位windowns服务器)
3. 可登录微信网页版的微信号一个     不知道自己微信能不能登录网页版？自行测试 https://wx.qq.com
4. 数据发布服务器一台(具体参数请自行在服务器post打印)


## 配置conf.toml

> 请把`conf.toml`放在执行文件wxspider.exe位置

参考  conf.example.toml

	# This is a TOML document. Boom.

	[BaiDuAiConf]
	api_key = "kGMQC1R***********InjUL"
	secret_key = "H6Mk*******************0Y7DkH0p"

	[PostConfig]
	serve_url = "http://examples.com/"
	authorization_token = ""

	
## 其他
可以通过 http://localhost:3355 端口获取最新100条数据



## 警告：本项目没有质保，请勿商用