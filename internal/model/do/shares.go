// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Shares is the golang structure of table shares for DAO operations like Where/Data.
type Shares struct {
	g.Meta      `orm:"table:shares, do:true"`
	Id          interface{} //
	Number      interface{} //
	Protocol    interface{} //
	Uid         interface{} //
	Address     interface{} //
	Port        interface{} //
	Security    interface{} //
	Encryption  interface{} //
	PublicKey   interface{} //
	HeaderType  interface{} //
	Fingerprint interface{} //
	Network     interface{} //
	Flow        interface{} //
	Sni         interface{} //
	ShortIds    interface{} //
	Remarks     interface{} //
	UpdatedAt   *gtime.Time //
}
