package model

import "rsshub/common/model/page"

type GetUserReq struct {
	Id      int
	RssLink string
	Where   int
	Order   int
}

type GetListReq struct {
	page.Page
	Where int
	Order int
}
