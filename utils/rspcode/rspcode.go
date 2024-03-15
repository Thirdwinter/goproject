package rspcode

const (
	SUCCESS = 200
	ERROR   = 500
	// code =1000... 100开头表示用户模块错误
	ERROR_USERNAME_USED  = 1001 // 用户名以被使用
	ERROR_PASSWORD_WRONG = 1002 // 用户密码错误
	ERROR_USER_NOT_EXIST = 1003 // 用户不存在

	ERROR_USER_NO_RIGHT     = 1008
	ERROR_PASSWORD_NO_EXIST = 1009

	// code =2000... 200开头表示文章模块错误

	ERROR_ART_NOT_EXIST = 2001

	// code =3000... 300开头表示分类模块错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002

	// code = 4000... 表示token相关
	ERROR_TOKEN_EXIST      = 4001 // 用户携带token不存在
	ERROR_TOKEN_LONGTIME   = 4002 // 用户携带token超时
	ERROR_TOKEN_WRONG      = 4003 // 用户携带token错误
	ERROR_TOKEN_TYPE_WRONG = 4004 // 用户token格式错误
	RTOKEN_SUCCESS         = 4005 // 用户token续期成功
	ERROR_TOKEN_CREATE     = 4006 // 生成token错误
	ERROR_TOKEN_NR         = 4007 // token续期失败
)

var (
	codemsg = map[int]string{
		SUCCESS:                 "OK!",
		ERROR:                   "FAIL!",
		ERROR_USERNAME_USED:     "用户名重复!",
		ERROR_PASSWORD_WRONG:    "密码错误!",
		ERROR_PASSWORD_NO_EXIST: "请输入密码",
		ERROR_USER_NOT_EXIST:    "用户不存在!",
		//ERROR_USER_NO_RIGHT:     "用户不正确",
		ERROR_TOKEN_EXIST:       "TOKEN不存在!",
		ERROR_TOKEN_LONGTIME:    "TOKEN超时!",
		ERROR_TOKEN_WRONG:       "TOKEN错误!",
		ERROR_TOKEN_TYPE_WRONG:  "TOKEN格式错误",
		ERROR_CATENAME_USED:     "分类已存在",
		ERROR_CATE_NOT_EXIST:    "分类不存在",
		ERROR_ART_NOT_EXIST:     "文章不存在",
		ERROR_USER_NO_RIGHT:     "用户权限不足",
		RTOKEN_SUCCESS:          "用户token续期成功",
		ERROR_TOKEN_NR:          "token无法续签",
	}
)

func GetMsg(code int) string {
	return codemsg[code]
}
