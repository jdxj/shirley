package shares

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/jdxj/shirley/api/shares/v1"
	"github.com/jdxj/shirley/internal/dao"
	"github.com/jdxj/shirley/internal/model/do"
)

func (c *ControllerV1) UpdateShare(ctx context.Context, req *v1.UpdateShareReq) (res *v1.UpdateShareRes, err error) {
	data := &do.Shares{
		Number:      req.Number,
		Protocol:    req.Protocol,
		Uid:         req.Uid,
		Address:     req.Address,
		Port:        req.Port,
		Security:    req.Security,
		Encryption:  req.Encryption,
		PublicKey:   req.PublicKey,
		HeaderType:  req.HeaderType,
		Fingerprint: req.Fingerprint,
		Network:     req.Network,
		Flow:        req.Flow,
		Sni:         req.Sni,
		ShortIds:    req.ShortIds,
		Remarks:     req.Remarks,
	}
	err = dao.Shares.UpdateShare(ctx, req.Id, data)
	if err != nil {
		g.Log().Errorf(ctx, "UpdateShare err: %s", err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	return &v1.UpdateShareRes{}, nil
}
