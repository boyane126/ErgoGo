package user

import "ErgoGo/internal/model"

type User struct {
	model.BaseModel

	Phone         string `json:"phone"`
	Avatar        string `json:"avatar"`
	Nickname      string `json:"nickname"`
	BusinessId    string `json:"business_id"`
	WechatOpenid  string `json:"wechat_openid"`
	WechatUnionid string `json:"wechat_unionid"`

	model.CommonTimestampsField
}
