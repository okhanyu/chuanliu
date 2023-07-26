package user

import (
	"github.com/gin-gonic/gin"
	"github.com/okhanyu/gohelper/gohelper_server"
	"log"
	"rsshub/api/user/model"
	service "rsshub/service/user"
)

func ListByGroup(c *gin.Context) {
	var param model.GetListReq
	//err := c.ShouldBindWith(&param, binding.JSON)
	err := c.ShouldBindQuery(&param)
	if err != nil && c.Request.ContentLength != 0 {
		log.Printf("format param error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}
	userList, err := service.GetListByGroup(c, ApiReqToServiceReq(param))
	if err != nil {
		log.Printf("service error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}
	gohelper_server.Success(c, ServiceStatisticsObjectToApiStatisticsObjectBatch(userList))
}

func List(c *gin.Context) {
	userList, err := service.GetList(c)
	if err != nil {
		log.Printf("service error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}
	gohelper_server.Success(c, ServiceObjectToApiObjectBatch(userList))
}
