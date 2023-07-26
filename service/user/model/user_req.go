package model

import (
	"rsshub/common/model/page"
)

type GetListReq struct {
	page.Page
	Where int `json:"where" form:"where"`
	Order int `json:"order" form:"order"`
}
