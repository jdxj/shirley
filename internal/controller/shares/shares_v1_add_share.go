package shares

import (
	"context"

	"github.com/jdxj/shirley/api/shares/v1"
	"github.com/jdxj/shirley/internal/dao"
	"github.com/jdxj/shirley/internal/model/do"
)

func (c *ControllerV1) AddShare(ctx context.Context, req *v1.AddShareReq) (res *v1.AddShareRes, err error) {
	data := &do.Shares{
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
	err = dao.Shares.InsertShare(ctx, data)
	if err != nil {
		return
	}
	return &v1.AddShareRes{
		Id:     data.Id,
		Number: data.Number,
	}, nil
}
