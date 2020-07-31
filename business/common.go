package business

import(
	jwtgo "github.com/dgrijalva/jwt-go"
	"blog/util"
	"time"
)

type Common struct{

}

// token生成器
// md 为上面定义好的middleware中间件
func (com *Common)GenerateToken(userid int)(string,error){
	// 构造SignKey: 签名和解签名需要使用一个值
	j := util.NewJWT()
	// 构造用户claims信息(负荷)
	claims := util.CustomClaims{
		userid,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "oldda.cn",                    // 签名颁发者
		},
	}

	// 根据claims生成token对象
	return j.CreateToken(claims)
}