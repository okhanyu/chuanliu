package rss

import (
	"github.com/gin-gonic/gin"
	"github.com/okhanyu/gohelper/gohelper_server"
	"log"
	"rsshub/api/rss/model"
	service "rsshub/service/rss"
)

func ListRss(c *gin.Context) {
	var param model.GetListReq
	//err := c.ShouldBindWith(&param, binding.JSON)
	err := c.ShouldBindQuery(&param)
	if err != nil && c.Request.ContentLength != 0 {
		log.Printf("format param error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}

	rssList, err := service.GetList(c, ApiReqToServiceReq(param))
	if err != nil {
		log.Printf("service error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}
	gohelper_server.Success(c, ServiceObjectToApiObjectBatch(rssList))
}

func ListRssTags(c *gin.Context) {
	var param model.GetTagsReq
	//err := c.ShouldBindWith(&param, binding.JSON)
	err := c.ShouldBindQuery(&param)
	if err != nil && c.Request.ContentLength != 0 {
		log.Printf("format param error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}

	tagList, err := service.GetTags(c, ApiTagsReqToServiceTagsReq(param))
	if err != nil {
		log.Printf("service error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}
	gohelper_server.Success(c, tagList)
}

func WatchArticle(c *gin.Context) {
	var param model.WatchReq
	//err := c.ShouldBindWith(&param, binding.JSON)
	err := c.ShouldBindJSON(&param)
	if err != nil && c.Request.ContentLength != 0 {
		log.Printf("format param error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}

	err = service.Watch(c, param.Id)
	if err != nil {
		log.Printf("service error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}
	gohelper_server.Success(c, "success")
}

func ListRssUserRecent(c *gin.Context) {
	var param model.GetListReq
	//err := c.ShouldBindWith(&param, binding.JSON)
	err := c.ShouldBindQuery(&param)
	if err != nil && c.Request.ContentLength != 0 {
		log.Printf("format param error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}

	rssList, err := service.GetUserRecentArticleListByGroup(c, ApiReqToServiceReq(param))
	if err != nil {
		log.Printf("service error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}
	gohelper_server.Success(c, ServiceRecentObjectToApiRecentObjectBatch(rssList))
}
