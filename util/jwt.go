package util

import(
	"github.com/dgrijalva/jwt-go"
)

// 定义一个jwt对象
type JWT struct {
	// 声明签名信息
	SigningKey []byte
}

// 初始化jwt对象
func NewJWT() *JWT {
	return &JWT{
		[]byte("oldda.cn"),
	}
}

// 自定义有效载荷(这里采用自定义的Name和Email作为有效载荷的一部分)
type CustomClaims struct {
	UserId int
	// StandardClaims结构体实现了Claims接口(Valid()函数)
	jwt.StandardClaims
}


// 调用jwt-go库生成token
// 指定编码的算法为jwt.SigningMethodHS256
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#Token
	// 返回一个token的结构体指针
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}


// token解码
func (j *JWT) ParserToken(tokenString string) (*CustomClaims, string) {
	// https://gowalker.org/github.com/dgrijalva/jwt-go#ParseWithClaims
	// 输入用户自定义的Claims结构体对象,token,以及自定义函数来解析token字符串为jwt的Token结构体指针
	// Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	// func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		// https://gowalker.org/github.com/dgrijalva/jwt-go#ValidationError
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, "AUTHORIZATION_TOKEN_EXPIRED"
			} else {
				return nil, "AUTHORIZATION_TOKEN_ERROR"
			}
		}
	}

	// 将token中的claims信息解析出来并断言成用户自定义的有效载荷结构
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, ""
	}
	return nil, "AUTHORIZATION_TOKEN_EXPIRED"
}