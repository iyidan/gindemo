package middleware

import (
	"bytes"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/iyidan/gindemo/conf"
	"github.com/iyidan/gindemo/log"
	"github.com/iyidan/gindemo/util"
)

// APISign api signature
func APISign() gin.HandlerFunc {
	isprod := conf.IsOnProd()
	maxBodySize := conf.Int64("maxBodySize")
	return func(c *gin.Context) {
		// nosign
		_, ok := c.GetQuery("nosign")
		if !isprod && ok {
			c.Next()
			return
		}

		errno, errmsg := SignRequest(maxBodySize, c)
		if errno != util.ErrnoSuccess {
			util.APIResp(c, nil, errno, errmsg, http.StatusOK)
			c.Abort()
			return
		}

		c.Next()
	}
}

// SignRequest api request check signature
func SignRequest(maxBodySize int64, c *gin.Context) (errno util.Errno, errmsg string) {
	datehdr := c.GetHeader("Date")
	if len(datehdr) == 0 {
		return util.ErrnoSign, "sign-header: Date empty"
	}
	authhdr := c.GetHeader("Authorization")
	if len(authhdr) == 0 {
		return util.ErrnoSign, "sign-header: Authorization empty"
	}
	auths := strings.Split(authhdr, " ")
	if len(auths) != 2 || len(auths[0]) == 0 || len(auths[1]) == 0 {
		return util.ErrnoSign, "sign-header: Authorization format error"
	}
	sk := getSkByAk(auths[0])
	if len(sk) == 0 {
		return util.ErrnoSign, "ak error"
	}

	// query values
	values := c.Request.URL.Query()
	var rawBody []byte
	var err error

	// form values
	ct := c.GetHeader("Content-Type")
	// RFC 2616, section 7.2.1 - empty type
	//   SHOULD be treated as application/octet-stream
	if ct == "" {
		ct = "application/octet-stream"
	}
	ct, _, err = mime.ParseMediaType(ct)
	if ct == "application/x-www-form-urlencoded" || ct == "multipart/form-data" {
		err = c.Request.ParseForm()
		if err != nil {
			return util.ErrnoSystem, "parse form error:" + err.Error()
		}
		err = c.Request.ParseMultipartForm(maxBodySize)
		if err != nil && err != http.ErrNotMultipart {
			return util.ErrnoSystem, "parse form error:" + err.Error()
		}
		for k, a := range c.Request.PostForm {
			// replace existing key in queryString
			for _, v := range a {
				values.Set(k, v)
			}
		}
		if c.Request.MultipartForm != nil && c.Request.MultipartForm.File != nil {
			for k, a := range c.Request.MultipartForm.Value {
				// replace existing key in queryString
				for _, v := range a {
					values.Set(k, v)
				}
			}
		}
	} else {
		// raw body
		rawBody, err = c.GetRawData()
		if err != nil {
			return util.ErrnoSystem, "read body error:" + err.Error()
		}
		// reset body
		c.Request.Body.Close()
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(rawBody))
	}

	genSign := util.GenSign(sk, c.Request.Method, c.Request.URL.Path, datehdr, values, rawBody)
	if genSign != auths[1] {
		log.Errorf("sign_request_error: sk=%s,method=%s,path=%s,datehdr=%s,authhdr=%s,params=%#v,body=%s,genSign=%s",
			sk, c.Request.Method, c.Request.URL.Path, datehdr, authhdr, values, rawBody, genSign)
		return util.ErrnoSign, "sign error"
	}
	return util.ErrnoSuccess, ""
}

func getSkByAk(ak string) string {
	apps := conf.MapStringString("client.appList")
	for k, v := range apps {
		if ak == k {
			return v
		}
	}
	return ""
}
