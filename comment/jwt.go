package comment

import (
	"errors"
	"mini/functions/info"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//生成token
func JwtMarsh(uid uint, isadmin bool) string {
	saltInit()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":     uid,
		"admin":   isadmin,
		"intime":  time.Now(),                         //注册时间
		"outtime": time.Now().Add(7 * 24 * time.Hour), //过期时间
	})
	tokenString, _ := token.SignedString([]byte(salt))
	return tokenString
}

//解析token
func JwtUnMarsh(tokenString string) map[string]interface{} {
	saltInit()
	//解析传入的token
	//第二个参数是一个回调函数，作用是判断生成token所用的签名算法是否和传入token的签名算法是否一致。
	//算法匹配就返回密钥，用来解析token.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("加密算法匹配错误")
		}
		return []byte(salt), nil
	})
	//err不为空，说明token已过期
	if err != nil {
		return nil
	}

	//将获取的token中的Claims强转为MapClaims
	claims, ok := token.Claims.(jwt.MapClaims)
	//判断token是否有效
	if !(ok && token.Valid) {
		return nil
	}
	return claims
}

var salt string

func saltInit() {
	if len(salt) == 0 {
		salt = info.Allinfo.Jwt.Salt
	}
}

func TokenId(c echo.Context) (id uint, admin bool) {
	token, err := c.Cookie("token")
	if err != nil {
		return
	}
	tokenmap := JwtUnMarsh(token.Value)
	//判断过期
	if pass, ok := tokenmap["outtime"].(string); ok {
		var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
		st, _ := time.Parse(timeLayoutStr, pass)  //string转time
		if time.Now().After(st) {
			//过期
			return
		}
	}
	uid := tokenmap["uid"].(float64)
	admin = tokenmap["admin"].(bool)
	return uint(uid), admin
}
