package user

import "go-juejin/model"

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username" form:"username"`
	Offset   int    `json:"offset" form:"offset"`
	Limit    int    `json:"limit" form:"offset"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}
