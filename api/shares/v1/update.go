package v1

import "github.com/gogf/gf/v2/frame/g"

type UpdateShareReq struct {
	g.Meta `path:"/shares/:id" method:"patch"`

	Id          uint64 `v:"required"`
	Protocol    any
	Uid         any
	Address     any
	Port        any
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
