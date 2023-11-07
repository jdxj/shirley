package v1

import "github.com/gogf/gf/v2/frame/g"

type GetTokenReq struct {
	g.Meta `path:"/token" method:"get"`

	Name string `v:"required"`
	Pass string `v:"required"`
}

type GetTokenRes struct {
	Token string `json:"token"`
}
