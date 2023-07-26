package user

import (
	"github.com/gin-gonic/gin"
	dao "rsshub/dao/user"
	"rsshub/service/user/model"
)

func GetListByGroup(_ *gin.Context, param model.GetListReq) ([]model.UserStatistics, error) {
	userDaoList, err := dao.GetUserListByGroup(ServiceListReqToDaoListReq(param))
	if err != nil {
		return nil, err
	}
	return DaoStatisticsObjectToServiceStatisticsObjectBatch(userDaoList), nil
}

func GetList(_ *gin.Context) ([]model.User, error) {
	userDaoList, err := dao.GetUserList()
	if err != nil {
		return nil, err
	}
	return DaoObjectToServiceObjectBatch(userDaoList), nil
}
