package user

import (
	"github.com/gin-gonic/gin"
	. "go-juejin/handler"
	"go-juejin/model"
	"go-juejin/pkg/auth"
	"go-juejin/pkg/errno"
	"go-juejin/pkg/token"
)

// Login generates the authentication token
// if the password was matched with the specified account.
func Login(c *gin.Context) {
	//Binding the data with the user struct
	var u model.UserModel
	if err := c.ShouldBind(&u); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Get the user information by the login username
	d, err := model.GetUser(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: d.Id, Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
