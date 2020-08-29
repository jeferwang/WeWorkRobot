package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeferwang/WeWorkRobot/src/entity"
	"github.com/jeferwang/WeWorkRobot/src/utils"
)

func Callback(c *gin.Context) {
	var formData entity.ValidateURLFormData
	if c.ShouldBind(&formData) == nil {
		echoStr, cryptoErr := utils.GetMsgCrypt().VerifyURL(formData.MsgSignature, formData.Timestamp, formData.Nonce, formData.EchoStr)
		if cryptoErr != nil {
			c.String(http.StatusForbidden, cryptoErr.ErrMsg)
		} else {
			c.String(http.StatusOK, string(echoStr))
		}
	} else {
		c.String(http.StatusForbidden, "Auth Fail")
	}
}
