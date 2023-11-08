package sub

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"github.com/jdxj/shirley/api/sub/v1"
	"github.com/jdxj/shirley/internal/config"
	"github.com/jdxj/shirley/internal/dao"
)

func (c *ControllerV1) GetSub(ctx context.Context, req *v1.GetSubReq) (res *v1.GetSubRes, err error) {
	if req.Path != config.Shirley.Path {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized)
	}
	shares, err := dao.Shares.GetShares(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "GetShares err: %s", err)
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	var (
		r   = ghttp.RequestFromCtx(ctx)
		buf = bytes.NewBuffer(nil)
	)
	for _, s := range shares {
		u := &url.URL{
			Scheme:   s.Protocol,
			User:     url.User(s.Uid),
			Host:     fmt.Sprintf("%s:%d", s.Address, s.Port),
			Fragment: s.Remarks,
		}
		query := u.Query()
		query.Set("security", s.Security)
		query.Set("encryption", s.Encryption)
		query.Set("pbk", s.PublicKey)
		query.Set("headerType", s.HeaderType)
		query.Set("fp", s.Fingerprint)
		query.Set("type", s.Network)
		query.Set("flow", s.Flow)
		query.Set("sni", s.Sni)
		query.Set("sid", s.ShortIds)
		u.RawQuery = query.Encode()

		buf.WriteString(u.String())
		buf.WriteByte('\n')
	}

	data := base64.StdEncoding.WithPadding(base64.NoPadding).
		EncodeToString(buf.Bytes())
	r.Response.Write(data)
	return nil, nil
}
