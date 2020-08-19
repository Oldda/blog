package util

import(
	"time"
    "math/rand"
    "encoding/json"
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

/**
    json和go数据类型的转换
**/
// json转map
func JSONToMap(jsonStr string) (map[string]interface{},error) {

    var tempMap map[string]interface{}

    err := json.Unmarshal([]byte(jsonStr), &tempMap)

    return tempMap,err
}