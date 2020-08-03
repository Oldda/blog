package business

import(
	jwtgo "github.com/dgrijalva/jwt-go"
	"blog/model"
	"blog/util"
	"log"
	"time"
)

type UserBaseInfo struct{
	Nickname string `json:"nickname"`
	Avatar string `json:"avatar"`
	Userid int `json:"userid"`
	Authorization string `json:"Authorization"`
}

func LoginByUsername(username,password string)*UserBaseInfo{
	user := new(model.User)
	has,err := model.DB.Engine.Where("realname=? and password=?",username,password).Get(user)
	if has {
		token,err := GenerateToken(user.Id)
		if err == nil{
			return &UserBaseInfo{
				user.Nickname,
				user.Avatar,
				user.Id,
				token,
			}
		}
	}
	log.Println(err)
	return nil
}

func LoginByPhone(phone,captcha string)*UserBaseInfo{
	return nil
}

func GenerateToken(userid int)(string,error){
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