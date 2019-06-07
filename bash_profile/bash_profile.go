package bash_profile

// 基础配置
const (
	DBConnect = "mongodb://localhost:27017" 	// 更换成你自己的uri地址
	DBName = "mini_program"						// 更换成你自己的数据库名字
	UserCollection = "users"					// 更换成你自己的表单名字

	WxSite = "https://api.weixin.qq.com/sns/jscode2session?"
	WxAppId = "wx6f9cc63euiased89b" 							// 使用自己的AppId 进行替换
	WxSecret = "421169661fiiif99110fe20c75a64ad1" 				// 使用自己的Secret 进行替换
	WxHttpTail = "&grant_type=authorization_code"
)