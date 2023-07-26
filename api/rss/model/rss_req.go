package model

import (
	"rsshub/common/model/page"
)

type GetListReq struct {
	page.Page
	UserName string `json:"user_name" form:"user_name"`
	Tag      string `json:"tag" form:"tag"`
	Link     string `json:"link" form:"link"`
	Where    int    `json:"where" form:"where"`
	Order    int    `json:"order" form:"order"`
}

type WatchReq struct {
	Id int `json:"id" form:"id"`
}

type GetTagsReq struct {
	page.Page
	Where int `json:"where" form:"where"`
	Order int `json:"order" form:"order"`
}
