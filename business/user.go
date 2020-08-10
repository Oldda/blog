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

//用户名密码登录
func LoginByUsername(username,password string)*UserBaseInfo{
	user := new(model.User)
	has,err := model.DB.Engine.Where("email=? and password=?",username,password).Get(user)
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

//手机号验证码登录
func LoginByPhone(phone,captcha string)*UserBaseInfo{
	var err error
	cpt,err := util.RdbGetString(phone+"_sms")
	if err == nil && cpt == captcha{
		user := new(model.User)
		has,err := model.DB.Engine.Where("phone = ?",phone).Get(user)
		if has && err == nil {
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
	}
	return nil
}

//生成token
func GenerateToken(userid int)(string,error){
	// 构造SignKey: 签名和解签名需要使用一个值
	j := util.NewJWT()
	// 构造用户claims信息(负荷)
	claims := util.CustomClaims{
		userid,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 7200), // 签名过期时间2个小时
			Issuer:    "oldda.cn",                    // 签名颁发者
		},
	}

	// 根据claims生成token对象
	return j.CreateToken(claims)
}

//添加管理员
func CreateUser(nickname,realname,password,phone,email,avatar string)(int,string){
	user := new(model.User)
	return user.CreateUser(nickname,realname,password,phone,email,avatar)
}

//管理员列表
func UserList(page,limit int)model.ListData{
	user := new(model.User)
	return user.UserList(page,limit)
}

func UserUpdate(id int,nickname,realname,password,phone,email,avatar string)error{
	user := new(model.User)
	return user.UserUpdate(id,nickname,realname,password,phone,email,avatar)
}

func UserDelete(id int)error{
	user := new(model.User)
	return user.UserDelete(id)
}

func UserGet(id int)(*model.User,error){
	user := new(model.User)
	return user.UserGet(id)
}