package memos

import (
	apiModel "rsshub/api/memos/model"
	serviceModel "rsshub/service/memos/model"
)

func ServiceObjectToApiObjectBatch(res []serviceModel.MemosUser) (result []apiModel.MemosUser) {
	for _, re := range res {
		result = append(result, apiModel.MemosUser{
			Id:          re.Id,
			UserName:    re.UserName,
			MemosLink:   re.MemosLink,
			Avatar:      re.Avatar,
			SiteTitle:   re.UserName,
			SiteLink:    re.SiteLink,
			Description: re.Description,
			CreateTime:  re.CreateTime,
			UpdateTime:  re.UpdateTime,
			//CreateTimeShow: re.CreateTime.Format("2006-01-02 15:04:05"),
			//CreateTimeShow: time.Unix(re.CreateTime, 0).Format("2006-01-02 15:04:05"),
		})

	}
	return result
}
