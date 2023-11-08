// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package shares

import (
	"context"
	
	"github.com/jdxj/shirley/api/shares/v1"
)

type ISharesV1 interface {
	AddShare(ctx context.Context, req *v1.AddShareReq) (res *v1.AddShareRes, err error)
	DeleteShare(ctx context.Context, req *v1.DeleteShareReq) (res *v1.DeleteShareRes, err error)
	GetShares(ctx context.Context, req *v1.GetSharesReq) (res *v1.GetSharesRes, err error)
	UpdateShare(ctx context.Context, req *v1.UpdateShareReq) (res *v1.UpdateShareRes, err error)
}


