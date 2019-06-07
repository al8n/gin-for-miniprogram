package bash_profile

// 基础配置
const (
	DBConnect = "mongodb://localhost:27017" 	// 更换成你自己的uri地址
	DBName = "mini_program"						// 更换成你自己的数据库名字


	WxUserCollection = "wx_users"								// 更换成你自己的表单名字
	WxSite = "https://api.weixin.qq.com/sns/jscode2session?"
	WxAppId = "wx6f9cc63ed9ded89bc" 								// 使用自己的AppId 进行替换
	WxSecret = "421169661fc3ef99110fe20c75a64ad1v" 				// 使用自己的Secret 进行替换
	WxHttpTail = "&grant_type=authorization_code"
)