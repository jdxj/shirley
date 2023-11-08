package v1

import "github.com/gogf/gf/v2/frame/g"

type DeleteShareReq struct {
	g.Meta `path:"/shares/:id" method:"delete"`

	Id uint64 `v:"required"`
}

type DeleteShareRes struct {
}
