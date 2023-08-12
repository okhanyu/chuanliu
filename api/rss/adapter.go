package rss

import (
	apiModel "rsshub/api/rss/model"
	serviceModel "rsshub/service/rss/model"
)

func ApiReqToServiceReq(req apiModel.GetListReq) serviceModel.GetListReq {
	return serviceModel.GetListReq{
		Page:     req.Page,
		Tag:      req.Tag,
		UserName: req.UserName,
		Where:    req.Where,
		Order:    req.Order,
	}
}
func ApiTagsReqToServiceTagsReq(req apiModel.GetTagsReq) serviceModel.GetTagsReq {
	return serviceModel.GetTagsReq{
		Page:  req.Page,
		Where: req.Where,
		Order: req.Order,
	}
}

func ServiceObjectToApiObjectBatch(res []serviceModel.GetRss) (result []apiModel.GetRss) {
	for _, re := range res {
		result = append(result, apiModel.GetRss{
			Id:          re.Id,
			Title:       re.Title,
			Link:        re.Link,
			Summary:     re.Summary,
			Tags:        re.Tags,
			UserName:    re.UserName,
			Watch:       re.Watch,
			Like:        re.Like,
			PubDate:     re.PubDate,
			PubDateShow: re.PubDate.Format("2006-01-02 15:04:05"),
			Cover:       re.Cover,
			UserId:      re.UserId,
			Avatar:      re.Avatar,
			//CreateTimeShow: time.Unix(re.CreateTime, 0).Format("2006-01-02 15:04:05"),
		})

	}
	return result
}

func ServiceRecentObjectToApiRecentObjectBatch(res []serviceModel.UserRecentArticle) (result []apiModel.
	UserRecentArticle) {
	for _, re := range res {
		result = append(result, apiModel.UserRecentArticle{
			Id:       re.Id,
			Title:    re.Title,
			Link:     re.Link,
			UserName: re.UserName,
			PubDate:  re.PubDate,
			//MinPubDateTime: re.MinPubDate,
			PubDateShow: re.PubDate.Format("2006-01-02 15:04:05"),
			//MinPubDateShow: re.MinPubDate.Format("2006-01-02 15:04:05"),
			UserId: re.UserId,
			//CreateTimeShow: time.Unix(re.CreateTime, 0).Format("2006-01-02 15:04:05"),
		})

	}
	return result
}
