package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	. "go-juejin/handler"
	"go-juejin/model"
	"go-juejin/pkg/errno"
	"go-juejin/util"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.ShouldBind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	//Validate the data.
	//if err:= u.Validate();err!=nil{
	//	SendResponse(c,errno.ErrValidation,nil)
	//	return
	//}
	if err := r.checkParam(); err != nil {
		SendResponse(c, err, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		log.Errorf(err, "Create user err %s", u)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}

func (r *CreateRequest) checkParam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty")
	}
	return nil
}
