package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/jdxj/shirley/internal/model/entity"
)

type GetSharesReq struct {
	g.Meta `path:"/shares" method:"get"`
}

type GetSharesRes struct {
	Shares []*entity.Shares `json:"shares"`
}
