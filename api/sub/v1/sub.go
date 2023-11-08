package v1

import "github.com/gogf/gf/v2/frame/g"

type GetSubReq struct {
	g.Meta `path:"/sub" method:"get"`

	Path string `v:"required"`
}

type GetSubRes struct {
}
