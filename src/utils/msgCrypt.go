package utils

import (
	"github.com/jeferwang/WeWorkRobot/src/conf"
	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
)

var msgCrypt *wxbizmsgcrypt.WXBizMsgCrypt

func GetMsgCrypt() *wxbizmsgcrypt.WXBizMsgCrypt {
	if msgCrypt == nil {
		config := conf.GetWxWorkConfig()
		msgCrypt = wxbizmsgcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, "", wxbizmsgcrypt.XmlType)
	}
	return msgCrypt
}
