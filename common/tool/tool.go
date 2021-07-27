package tool

import (
	"crypto/md5"
	"fmt"
	"github.com/patrickmn/go-cache"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

/*
  工具类
*/

var (
	//key的过期时间是一小时(自定义时间)，并且每10s清除缓存中的过期key
	c = cache.New(3600*time.Second, 10*time.Second)
)

//获取本机IP
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func SplitEerr(err string) []string {
	arr := strings.Split(err, "!")
	return arr
}

func SetCache(Session_token, name string) {
	c.Set(name, Session_token, cache.DefaultExpiration)
}

func GetCache(name string) string {
	value, found := c.Get(name)
	if found {
		return value.(string)
	} else {
		return ""
	}
}

// 获取字符串中的中文
func Chinese(str string) string {
	r := []rune(str)
	strSlice := []string{}
	cnstr := ""
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			cnstr = cnstr + string(r[i])
			strSlice = append(strSlice, cnstr)

		}
	}
	if 0 == len(strSlice) {
		//无中文，需要跳过，后面再找规律
	}
	return cnstr
}

//万能密码 ，md5加密
func UniversalEncryption() string {
	month := int(time.Now().Month())
	day := time.Now().Day()
	str := "GBM_CY_OA" + strconv.Itoa(month) + strconv.Itoa(day)
	w := md5.New()
	io.WriteString(w, str) //将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil))
	return md5str2
}
