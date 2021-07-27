package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

/*
	统一登录生成token
*/
func GetJwtToken(iat, seconds int64, secretKey, username, id, staffId, userId string) (string, error) {
	claims := make(jwt.MapClaims)

	//Exp1 := time.Now().Add(time.Second * time.Duration(seconds)).Unix()
	//timeLayout := "2006-01-02 15:04:05" //转化所需模板
	//时间戳转化为日期
	//datetime := time.Unix(Exp1, 0).Format(timeLayout)

	fmt.Println("11111", iat, seconds, secretKey, username, id, staffId, userId)
	claims["expire"] = iat+seconds // 过期时间
	claims["iat"] = iat         //生成时间
	claims["id"] = userId
	claims["userName"] = username
	claims["No"] = staffId
	claims["company_id"] = id
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
