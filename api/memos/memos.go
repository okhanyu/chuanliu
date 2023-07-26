package memos

import (
	"github.com/gin-gonic/gin"
	"github.com/okhanyu/gohelper/gohelper_server"
	"log"
	service "rsshub/service/memos"
)

func List(c *gin.Context) {
	memosUserList, err := service.GetMemosUserList(c)
	if err != nil {
		log.Printf("service error is : %v", err)
		gohelper_server.FailWithError(c, err)
		c.Abort()
		return
	}
	gohelper_server.Success(c, ServiceObjectToApiObjectBatch(memosUserList))
}
