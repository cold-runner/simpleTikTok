package model

// StableUserFields 用于保存用户信息当中的稳定字段
type StableUserFields struct {
	Name            string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}
