package util

import(
	"time"
    "math/rand"
)

//生成随机字符串
func GetRandStr(n int) (randStr string) {
    chars := "ABCDEFGHIJKMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz23456789"
    charsLen := len(chars)
    if n > 10 {
        n = 10
    }

    rand.Seed(time.Now().UnixNano())
    for i := 0; i < n; i++ {
        randIndex := rand.Intn(charsLen)
        randStr += chars[randIndex : randIndex+1]
    }
    return randStr
}