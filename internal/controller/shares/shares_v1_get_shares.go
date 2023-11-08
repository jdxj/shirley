package shares

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/jdxj/shirley/api/shares/v1"
	"github.com/jdxj/shirley/internal/dao"
)

func (c *ControllerV1) GetShares(ctx context.Context, req *v1.GetSharesReq) (res *v1.GetSharesRes, err error) {
	shares, err := dao.Shares.GetShares(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "GetShares err: %s", err)
		return nil, err
	}
	return &v1.GetSharesRes{Shares: shares}, nil
}
