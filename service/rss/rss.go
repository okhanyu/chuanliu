package rss

import (
	"github.com/gin-gonic/gin"
	dao "rsshub/dao/rss"
	"rsshub/service/rss/model"
	"strings"
)

//var LikeIp map[string]map[int]int
//
//func init() {
//	if LikeIp == nil {
//		LikeIp = make(map[string]map[int]int)
//	}
//}

func GetList(c *gin.Context, param model.GetListReq) ([]model.GetRss, error) {
	rssDaoList, err := dao.GetRssList(c, ServiceReqToDaoReq(param))
	if err != nil {
		return nil, err
	}
	return DaoObjectToServiceObjectBatch(rssDaoList), nil
}

func GetTags(c *gin.Context, param model.GetTagsReq) ([]string, error) {
	tagList, err := dao.GetTags(c, ServiceTagsReqToDaoTagsReq(param))
	if err != nil {
		return nil, err
	}
	handleMap := make(map[string]string)
	var newTags []string
	for _, s := range tagList {
		items := strings.Split(s, ",")
		for _, item := range items {
			item = strings.ReplaceAll(item, "\n", "")
			item = strings.TrimSpace(item)
			if _, ok := handleMap[item]; !ok && strings.TrimSpace(item) != "" {
				handleMap[item] = handleMap[item]
				//tagList = append(tagList[:index], tagList[index+1:]...)
				newTags = append(newTags, item)
			}
		}
	}
	return newTags, nil
}

func Watch(c *gin.Context, param int) error {
	err := dao.Watch(c, param)
	if err != nil {
		return err
	}
	return nil
}

func Like(c *gin.Context, param int) error {
	err := dao.Like(c, param)
	if err != nil {
		return err
	}
	return nil
}

func GetUserRecentArticleListByGroup(c *gin.Context, param model.GetListReq) ([]model.UserRecentArticle, error) {

	rssDaoList, err := dao.GetUserRecentArticleListByGroup(c, ServiceReqToDaoReq(param))
	if err != nil {
		return nil, err
	}
	return DaoRecentObjectToServiceRecentObjectBatch(rssDaoList), nil

}
