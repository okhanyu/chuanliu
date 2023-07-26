package user

import (
	daoModel "rsshub/dao/memos/model"
	serviceModel "rsshub/service/memos/model"
)

func DaoObjectToServiceObjectBatch(req []daoModel.MemosUser) (result []serviceModel.MemosUser) {
	for _, re := range req {
		result = append(result, serviceModel.MemosUser{
			Id:          re.Id,
			UserName:    re.UserName,
			MemosLink:   re.MemosLink,
			Avatar:      re.Avatar,
			SiteTitle:   re.UserName,
			SiteLink:    re.SiteLink,
			Description: re.Description,
			CreateTime:  re.CreateTime,
			UpdateTime:  re.UpdateTime,
			Del:         re.Del,
			//CreateTimeShow: re.CreateTime.Format("2006-01-02 15:04:05"),
			//CreateTimeShow: time.Unix(re.CreateTime, 0).Format("2006-01-02 15:04:05"),
		})
	}
	return result
}
