package token

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/jdxj/shirley/api/token/v1"
	"github.com/jdxj/shirley/internal/config"
	"github.com/jdxj/shirley/utility"
)

func (c *ControllerV1) GetToken(ctx context.Context, req *v1.GetTokenReq) (res *v1.GetTokenRes, err error) {
	var (
		name = config.Shirley.Name
		pass = config.Shirley.Pass
		key  = config.Shirley.Key
	)
	if req.Name != name || req.Pass != pass {
		return nil, gerror.NewCode(gcode.CodeValidationFailed)
	}

	token, err := utility.NewToken(key)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	return &v1.GetTokenRes{Token: token}, nil
}
