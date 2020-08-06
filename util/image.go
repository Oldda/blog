package util

import(
	"github.com/qiniu/api.v7/v7/auth/qbox"
    "github.com/qiniu/api.v7/v7/storage"
)

//AK:NTFEIKTPfcU8BJDi0PNGy0C7sksUCt3RczbNfbTH
//SK:zX9NwDeywmiyM1gHyGqIR6j72KdSopyEMKyN1aVX
//空间名称：oldda-blog
//访问域名：pic.oldda.cn

func GetQNUploadToken()string{
	bucket:="oldda-blog"
	putPolicy := storage.PutPolicy{
	        Scope: bucket,
	}
	ak := "NTFEIKTPfcU8BJDi0PNGy0C7sksUCt3RczbNfbTH"
	sk := "zX9NwDeywmiyM1gHyGqIR6j72KdSopyEMKyN1aVX"
	mac := qbox.NewMac(ak, sk)
	return putPolicy.UploadToken(mac)
}