package model

type WxUser struct {
	ID       int    `db:"id"`
	Staff_ID string `db:staffID`
	User_ID  string	`db:userId`
	Nickname string	`db:nickname`
	Gender   int	`db:gender`
	Language string	`db:language`
	City     string	`db:city`
	Province string	`db:province`
	Country  string	`db:country`
	OpenId   string	`db:openId`
	UnionId  string	`db:unionId`
	IsValid  bool	`db:isValid`
}
