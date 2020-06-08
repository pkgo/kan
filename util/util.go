package util

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func FlexBind(c *gin.Context, body interface{}) {
	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	_ = json.Unmarshal(bodyBytes, &body)
	_ = c.Request.Body.Close() //  must close
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
}
