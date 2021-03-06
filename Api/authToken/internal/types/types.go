// Code generated by goctl. DO NOT EDIT.
package types

type UserRequest struct {
	Mobile string `form:"mobile"`
	Passwd string `form:"passwd"`
	Code   string `form:"code,optional"`
}

type UserResponse struct {
	Id        string `json:"id"`
	Expire    int64  `json:"exp"`
	Iat       int64  `json:"iat"`
	UserName  string `json:"user_name"`
	CompanyId string `json:"company_id"`
	No        string `json:"No"`
}

type TokenRequest struct {
	Code   string `form:"code"`
}

type TokenResponse struct {
	Token  string `json:"token"`
	Expire int64  `json:"ExpiresAt"`
}

