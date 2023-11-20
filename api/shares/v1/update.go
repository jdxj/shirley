package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateShareReq struct {
	g.Meta `path:"/shares/:id" method:"patch"`

	Id          uint64 `v:"required"`
	Number      any    `v:"min:0|max:255"`
	Protocol    any
	Uid         any
	Address     any
	Port        any `v:"min:1025|max:65535"`
	Security    any
	Encryption  any
	PublicKey   any
	HeaderType  any
	Fingerprint any
	Network     any
	Flow        any
	Sni         any
	ShortIds    any
	Remarks     any
}

type UpdateShareRes struct {
}
