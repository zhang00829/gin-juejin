package user

import (
	"github.com/gin-gonic/gin"
	. "go-juejin/handler"
	"go-juejin/model"
	"go-juejin/pkg/errno"
	"strconv"
)

//Delete delete an user by the user identifier
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	SendResponse(c, nil, nil)
}
