package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

// ValidateURLFormData 企业微信机器人填入Url时，需要做加密校验
type ValidateURLFormData struct {
	MsgSignature string `form:"msg_signature" binding:"required"`
	Timestamp    string `form:"timestamp" binding:"required"`
	Nonce        string `form:"nonce" binding:"required"`
	EchoStr      string `form:"echostr" binding:"required"`
}

func main() {
	r := gin.Default()

	msgCrypt := wxbizmsgcrypt.NewWXBizMsgCrypt(
		"oacrqp84m9Xy3ogt3nGnDUaiIRiR",
		"YqLDyEl6mNeYp3nUAv50bFgarPG93uBd52rd0B5ekRT",
		"",
		wxbizmsgcrypt.XmlType,
	)

	r.GET("/callback", func(c *gin.Context) {
		var formData ValidateURLFormData
		if c.ShouldBind(&formData) == nil {
			echoStr, cryptoErr := msgCrypt.VerifyURL(formData.MsgSignature, formData.Timestamp, formData.Nonce, formData.EchoStr)
			if cryptoErr != nil {
				c.String(http.StatusForbidden, cryptoErr.ErrMsg)
			} else {
				c.String(http.StatusOK, string(echoStr))
			}
		} else {
			c.String(http.StatusForbidden, "Auth Fail")
		}
	})
	r.Run("0.0.0.0:80")
}
