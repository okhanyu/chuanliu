package user

import (
	"github.com/gin-gonic/gin"
	dao "rsshub/dao/memos"
	"rsshub/service/memos/model"
)

func GetMemosUserList(_ *gin.Context) ([]model.MemosUser, error) {
	userDaoList, err := dao.GetMemosUserList()
	if err != nil {
		return nil, err
	}
	return DaoObjectToServiceObjectBatch(userDaoList), nil
}
