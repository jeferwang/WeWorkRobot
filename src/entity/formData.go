package entity

// ValidateURLFormData 企业微信机器人填入Url时，需要做加密校验
type ValidateURLFormData struct {
	MsgSignature string `form:"msg_signature" binding:"required"`
	Timestamp    string `form:"timestamp" binding:"required"`
	Nonce        string `form:"nonce" binding:"required"`
	EchoStr      string `form:"echostr" binding:"required"`
}
