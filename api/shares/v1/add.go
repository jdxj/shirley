package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddShareReq struct {
	g.Meta `path:"/shares" method:"post"`

	Protocol    string `v:"required"`
	Uid         string `v:"required"`
	Address     string `v:"required"`
	Port        int    `v:"required"`
	Security    string `v:"required"`
	Encryption  string `v:"required"`
	PublicKey   string `v:"required"`
	HeaderType  string `v:"required"`
	Fingerprint string `v:"required"`
	Network     string `v:"required"`
	Flow        string `v:"required"`
	Sni         string `v:"required"`
	ShortIds    string `v:"required"`
	Remarks     string `v:"required"`
}

type AddShareRes struct {
	Id     any `json:"id"`
	Number any `json:"number"`
}
