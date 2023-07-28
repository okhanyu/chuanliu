package rss

import (
	daoModel "rsshub/dao/rss/model"
	serviceModel "rsshub/service/rss/model"
	"strings"
)

func ServiceReqToDaoReq(req serviceModel.GetListReq) daoModel.GetRssListReq {
	return daoModel.GetRssListReq{
		Page:     req.Page,
		Tag:      req.Tag,
		UserName: req.UserName,
		Where:    req.Where,
		Order:    req.Order,
	}
}

func ServiceTagsReqToDaoTagsReq(req serviceModel.GetTagsReq) daoModel.GetTagsReq {
	return daoModel.GetTagsReq{
		Page:  req.Page,
		Where: req.Where,
		Order: req.Order,
	}
}

func DaoObjectToServiceObjectBatch(req []daoModel.GetRss) (result []serviceModel.GetRss) {
	for _, re := range req {
		var tagsHandle []string
		if re.Tags != "" {
			tagsHandle = strings.Split(re.Tags, ",")
		}
		if re.Summary == "" {
			re.Summary = re.Content
		}
		result = append(result, serviceModel.GetRss{
			Id:         re.Id,
			Title:      re.Title,
			Tags:       tagsHandle,
			Link:       re.Link,
			UserName:   re.UserName,
			Del:        re.Del,
			Watch:      re.Watch,
			Like:       re.Like,
			Summary:    re.Summary,
			CreateTime: re.CreateTime,
			//CreateTimeShow: re.CreateTime.Format("2006-01-02 15:04:05"),
			PubDate: re.PubDateTime,
			UserId:  re.UserId,
			Cover:   re.Cover,
			Avatar:  re.Avatar,
			//CreateTimeShow: time.Unix(re.CreateTime, 0).Format("2006-01-02 15:04:05"),

		})

	}
	return result
}

func DaoRecentObjectToServiceRecentObjectBatch(req []daoModel.UserRecentArticle) (result []serviceModel.UserRecentArticle) {
	for _, re := range req {
		result = append(result, serviceModel.UserRecentArticle{
			Id:       re.Id,
			Title:    re.Title,
			Link:     re.Link,
			UserName: re.UserName,
			UserId:   re.UserId,
			//CreateTimeShow: re.CreateTime.Format("2006-01-02 15:04:05"),
			PubDate:    re.PubDateTime,
			MinPubDate: re.MinPubDateTime,
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
