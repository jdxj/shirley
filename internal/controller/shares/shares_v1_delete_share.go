package shares

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/jdxj/shirley/api/shares/v1"
	"github.com/jdxj/shirley/internal/dao"
)

func (c *ControllerV1) DeleteShare(ctx context.Context, req *v1.DeleteShareReq) (res *v1.DeleteShareRes, err error) {
	err = dao.Shares.DeleteShare(ctx, req.Id)
	if err != nil {
		g.Log().Errorf(ctx, "DeleteShare err: %s", err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	return &v1.DeleteShareRes{}, nil
}
