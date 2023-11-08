// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package sub

import (
	"context"
	
	"github.com/jdxj/shirley/api/sub/v1"
)

type ISubV1 interface {
	GetSub(ctx context.Context, req *v1.GetSubReq) (res *v1.GetSubRes, err error)
}


