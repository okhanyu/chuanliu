package user

import (
	daoModel "rsshub/dao/user/model"
	serviceModel "rsshub/service/user/model"
)

func ServiceListReqToDaoListReq(req serviceModel.GetListReq) daoModel.GetListReq {
	return daoModel.GetListReq{
		Page:  req.Page,
		Where: req.Where,
		Order: req.Order,
	}
}
func DaoObjectToServiceObjectBatch(req []daoModel.User) (result []serviceModel.User) {
	for _, re := range req {
		result = append(result, serviceModel.User{
			Id:          re.Id,
			UserName:    re.UserName,
			RssLink:     re.RssLink,
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

func DaoStatisticsObjectToServiceStatisticsObjectBatch(req []daoModel.UserStatistics) (result []serviceModel.UserStatistics) {
	for _, re := range req {
		result = append(result, serviceModel.UserStatistics{
			Id:          re.Id,
			UserName:    re.UserName,
			RssLink:     re.RssLink,
			Avatar:      re.Avatar,
			SiteTitle:   re.UserName,
			SiteLink:    re.SiteLink,
			Watch:       re.Watch,
			Like:        re.Like,
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

//for _, rss_model := range rssList {
//var coordinates res.Coordinates
//err = json.Unmarshal([]byte(say.Location), &coordinates)
//if err != nil {
//return nil, err
//}
//var contentResModel res.ContentResModel
//err = json.Unmarshal([]byte(say.Content), &contentResModel)
//if err != nil {
//return nil, err
//}
//var tags []string
//err = json.Unmarshal([]byte(say.Tags), &tags)
//if err != nil {
//return nil, err
//}
//resSays = append(resSays, res.ResModel{
//Content:      contentResModel,
//Location:     coordinates,
//LocationShow: say.LocationShow,
//UserName:     say.UserName,
//CreateTime:   say.CreateTime,
//Praises:      say.Praises,
//Comments:     say.Comments,
//Uuid:         say.Uuid,
//Id:           say.Id,
//Tags:         tags,
//})
//}
