package middleware

import (
	"fmt"
	"time"

	"goproject/global"
	"goproject/utils/rspcode"

	_ "github.com/ThirdWinter/Go/mylog"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//var (
//	code int
//)

type MyClaims struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

func Jwt() []byte {
	return []byte(global.Config.System.JwtKey)
}

// 生成token
func SetToken(username string, role int) (string, string, int) {
	expireTime := time.Now().Add(24 * time.Hour)
	r_expireTime := time.Now().Add(48 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	// 创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	// 生成token
	a_Token, aerr := token.SignedString(Jwt())
	// 生成token错误处理逻辑
	if aerr != nil {
		return "", "", rspcode.ERROR_TOKEN_CREATE
	}
	// rToken 不需要存储任何自定义数据
	// rtoken生成错误时返回空
	r_Token, rerr := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: r_expireTime.Unix(),
		Issuer:    "ginblog",
	}).SignedString(Jwt())
	if rerr != nil {
		return a_Token, "", rspcode.SUCCESS
	}
	return a_Token, r_Token, rspcode.SUCCESS
}

// 验证token
func CheckToken(token string) (jwt.Claims, int) {

	claims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Jwt(), nil
	})
	if claims == nil {
		return nil, rspcode.ERROR_TOKEN_TYPE_WRONG
	}
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token 已过期
				return nil, rspcode.ERROR_TOKEN_LONGTIME
			}
			// 其他验证错误
			return nil, rspcode.ERROR_TOKEN_WRONG
		}
	}

	// 处理正常情况
	return claims.Claims, rspcode.SUCCESS // token可以解析,但是有可能过期
}

// 刷新token
// 如果at正常,直接放行;如果ak是过期错误且携带rk,校验rk,rk正确后返回新ak和rk,
func RefreshToken(aToken string, rToken string) (newToken string, newRtoken string, code int) {
	// 1. 判断rToken格式正确,没有错误
	if _, err := jwt.Parse(rToken, func(token *jwt.Token) (interface{}, error) {
		return Jwt(), nil
	}); err != nil {
		return "", "", rspcode.ERROR_TOKEN_NR
	}

	// 2. 从旧的aToken中解析出claims数据
	var claims MyClaims
	_, err := jwt.ParseWithClaims(aToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return Jwt(), nil
	})
	v, _ := err.(*jwt.ValidationError)

	// 当atoken是过期错误,且rtoken没有过期就创建一个新的access token
	if v.Errors == jwt.ValidationErrorExpired {
		newAToken, newRToken, errCode := SetToken(claims.Username, claims.Role)
		if errCode == rspcode.SUCCESS {
			return newAToken, newRToken, rspcode.RTOKEN_SUCCESS
		}
		return "", "", rspcode.ERROR_TOKEN_NR
	}

	return "", "", rspcode.ERROR_TOKEN_NR
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		atoken := c.GetHeader("Authorization")
		rtoken := c.GetHeader("RefreshAuthorization")
		if len(atoken) == 0 {
			c.JSON(200, gin.H{
				"msg":  "没有携带 atoken",
				"code": rspcode.ERROR_TOKEN_EXIST, //未授权
			})
			c.Abort()
			return
		}
		//校验正确
		info, tCode := CheckToken(atoken)
		a, ok := info.(*MyClaims)
		fmt.Println("a.role:",a.Role)
		if !ok {
			fmt.Println("no")
			return
		}
		fmt.Printf("info: %#v", a.Username)
		switch tCode {
		case 200:
			{
				//校验成功
				c.Set("role", a.Role)
				c.Set("username", a.Username)
				c.Next()
				return
			}
		case 4002:
			{
				if len(rtoken) == 0 {
					c.JSON(200, gin.H{
						"code": tCode,
						"msg":  rspcode.GetMsg(tCode),
					})
					c.Abort()
					return
				} else if len(rtoken) != 0 {
					//刷新
					newAToken, newRToken, code := RefreshToken(atoken, rtoken)
					if code != rspcode.RTOKEN_SUCCESS {
						c.JSON(200, gin.H{
							"msg":  "刷新错误,请重新登录",
							"code": code,
						})
						c.Abort()
						return
					}
					c.JSON(200, gin.H{
						"msg":    rspcode.GetMsg(code),
						"atoken": newAToken,
						"rtoken": newRToken,
						"code":   code,
					})
					c.Abort()
					return
				}
			}
		default:
			c.JSON(200, gin.H{
				"code": tCode,
				"msg":  rspcode.GetMsg(tCode),
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

//"atoken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiZXhwIjoxNzA5NDUxOTk0LCJpc3MiOiJnaW5ibG9nIn0.g6ln3LIv2-gQcNmDQFys0mv5w1Cj5tPEV5j0lkTrGZM",
//"msg": "OK!",
//"rtoken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk1MjM5ODQsImlzcyI6ImdpbmJsb2cifQ.Kj_LbCkVNXEL086J7WhMRhWToge295iWz_VfiRV1puM",
//"status": 200
