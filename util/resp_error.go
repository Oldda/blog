package util

type resp map[string]interface{}

type ApiResp struct{
	Status bool `json:"status"`
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

func NewApiResp(codeStr string,data... interface{})(int,ApiResp){
	if result,ok := res[codeStr];ok {
		code,_ := result["http_code"].(int)
		body,_ := result["api_body"].(ApiResp)
		if data != nil{
			body.Data = data
		}
		return code,body
	}
	code,_ := res["RESP_NOT_FOUND"]["http_code"].(int)
	body,_ := res["RESP_NOT_FOUND"]["api_body"].(ApiResp)
	return code,body
}

var	res = map[string]resp{
	//公共错误
	"SUCCESS" : resp{
		"http_code" : 200,
		"api_body" : ApiResp{
			Status:true,
			Code:200,
			Msg:"请求成功",
		},
	},
	"RESP_NOT_FOUND" : resp{
		"http_code" : 302,
		"api_body" : ApiResp{
			Status:false,
			Code:1000,
			Msg:"未知错误",
		},
	},
	"MUST_PARAMS_LOST" : resp{
		"http_code" : 302,
		"api_body" : ApiResp{
			Status:false,
			Code:1001,
			Msg:"缺少必要参数",
		},
	},
	//用户相关
	"LOGIN_METHOD_NOT_ALLOWD" : resp{
		"http_code" : 302,
		"api_body" : ApiResp{
			Status:false,
			Code:2000,
			Msg:"不支持的登录方式",
		},
	}, 
	"UNAME_OR_PWD_ERROR" : resp{
		"http_code" : 302,
		"api_body" : ApiResp{
			Status:false,
			Code:2001,
			Msg:"用户名或密码错误",
		},
	},
	"AUTHORIZATION_TOKEN_ERROR" : resp{
		"http_code" : 302,
		"api_body" : ApiResp{
			Status:false,
			Code:2002,
			Msg:"登录token不可用",
		},
	},
	"AUTHORIZATION_TOKEN_EXPIRED" : resp{
		"http_code" : 302,
		"api_body" : ApiResp{
			Status:false,
			Code:2003,
			Msg:"登录token已过期",
		},
	},
	"AUTHORIZATION_TOKEN_LOST" : resp{
		"http_code" : 302,
		"api_body" : ApiResp{
			Status:false,
			Code:2004,
			Msg:"登录token丢失",
		},
	},
}
