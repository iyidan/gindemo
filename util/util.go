package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"sync"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iyidan/goutils/mise"
)

// Errno the type of errno
type Errno string

const (
	ErrnoSuccess Errno = "SUCC"
	ErrnoParam   Errno = "ERR_PARAM"
	ErrnoSign    Errno = "ERR_SIGN"
	ErrnoUser    Errno = "ERR_USER"
	ErrnoThird   Errno = "ERR_THIRD"
	ErrnoSystem  Errno = "ERR_SYSTEM"
	ErrnoTimeout Errno = "ERR_TIMEOUT"
)

var (
	apiDataPool = sync.Pool{
		New: func() interface{} {
			return &APIRespData{}
		},
	}
)

// APIRespData api's response data format
type APIRespData struct {
	Errno  Errno       `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func (ad *APIRespData) reset() {
	ad.Errno = ""
	ad.Errmsg = ""
	ad.Data = nil
}

// APISucc a success api reponse
func APISucc(c *gin.Context, data interface{}) {
	APIResp(c, data, ErrnoSuccess, "success", http.StatusOK)
}

// APIParamError a ErrnoParam api reponse
func APIParamError(c *gin.Context, errmsg string) {
	APIResp(c, nil, ErrnoParam, errmsg, http.StatusOK)
}

// APISystemError a ErrnoSystem api reponse
func APISystemError(c *gin.Context, errmsg string) {
	APIResp(c, nil, ErrnoSystem, errmsg, http.StatusInternalServerError)
}

// APIUserError a ErrnoUser api reponse
func APIUserError(c *gin.Context, errmsg string) {
	APIResp(c, nil, ErrnoUser, errmsg, http.StatusOK)
}

// APIStatusOk a statusok api reponse
func APIStatusOk(c *gin.Context, data interface{}, errno Errno, errmsg string) {
	APIResp(c, data, errno, errmsg, http.StatusOK)
}

// APIResp a api response
func APIResp(c *gin.Context, data interface{}, errno Errno, errmsg string, httpCode int) {
	//ret := &APIRespData{Errno: errno, Errmsg: errmsg, Data: data}
	//c.JSON(httpCode, ret)

	ret := apiDataPool.Get().(*APIRespData)
	ret.Errno = errno
	ret.Errmsg = errmsg
	ret.Data = data

	c.JSON(httpCode, ret)

	ret.reset()
	apiDataPool.Put(ret)
}

// GenSign sign implement
func GenSign(sk, method, path, date string, params map[string][]string, body []byte) string {
	p := GenSignParams(params, false)
	paramString := strings.Join(p, "&")
	bodyMD5 := ""
	if len(body) > 0 {
		h := md5.New()
		h.Write(body)
		bodyMD5 = hex.EncodeToString(h.Sum(nil))
	}
	stringToSign := method + "\n" +
		path + "\n" +
		bodyMD5 + "\n" +
		date + "\n" +
		paramString
	mac := hmac.New(sha1.New, []byte(sk))
	mac.Write([]byte(stringToSign))
	return hex.EncodeToString(mac.Sum(nil))
}

// GenSignParams generate sign params
// params support the below types
// 	map[string]string
// 	map[string]int
// 	map[string]float64 notice: the degree of value. you may first use mise.ToFixed() method
// 	map[string][]interface{} // interface{} can store string,int,float64 type
func GenSignParams(params map[string][]string, urlencode bool) []string {
	var p []string
	for k, a := range params {
		for _, v := range a {
			// empty string
			if len(v) == 0 {
				continue
			}
			if urlencode {
				v = url.QueryEscape(v)
			}
			p = append(p, fmt.Sprintf("%s=%s", k, v))
		}
	}
	sort.Strings(p)
	return p
}

// MapSliceMixed2MapSliceString change param-value to []string
func MapSliceMixed2MapSliceString(params map[string]interface{}) map[string][]string {

	if len(params) == 0 {
		return nil
	}
	data := make(map[string][]string)
	for k, a := range params {
		value, kind := mise.GetValueKind(a)
		switch kind {
		case reflect.Slice:
			for i := 0; i < value.Len(); i++ {
				sv := fmt.Sprintf("%v", value.Index(i))
				if _, ok := data[k]; ok {
					data[k] = append(data[k], sv)
				} else {
					data[k] = []string{sv}
				}
			}
		default:
			// empty string
			sv := fmt.Sprintf("%v", value)
			if _, ok := data[k]; ok {
				data[k] = append(data[k], sv)
			} else {
				data[k] = []string{sv}
			}
		}
	}

	return data
}
