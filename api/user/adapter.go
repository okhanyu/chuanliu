package user

import (
	apiModel "rsshub/api/user/model"
	serviceModel "rsshub/service/user/model"
)

func ApiReqToServiceReq(req apiModel.GetListReq) serviceModel.GetListReq {
	return serviceModel.GetListReq{
		Page:  req.Page,
		Where: req.Where,
		Order: req.Order,
	}
}

func ServiceStatisticsObjectToApiStatisticsObjectBatch(res []serviceModel.UserStatistics) (result []apiModel.UserStatistics) {
	for _, re := range res {
		result = append(result, apiModel.UserStatistics{
			Id:          re.Id,
			UserName:    re.UserName,
			RssLink:     re.RssLink,
			Avatar:      re.Avatar,
			SiteTitle:   re.UserName,
			SiteLink:    re.SiteLink,
			Watch:       re.Watch,
			Description: re.Description,
			CreateTime:  re.CreateTime,
			UpdateTime:  re.UpdateTime,
			Del:         re.Del,
			Total:       re.Total,
			//CreateTimeShow: re.CreateTime.Format("2006-01-02 15:04:05"),
			//CreateTimeShow: time.Unix(re.CreateTime, 0).Format("2006-01-02 15:04:05"),
		})

	}
	return result
}

func ServiceObjectToApiObjectBatch(res []serviceModel.User) (result []apiModel.User) {
	for _, re := range res {
		result = append(result, apiModel.User{
			Id:          re.Id,
			UserName:    re.UserName,
			RssLink:     re.RssLink,
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
