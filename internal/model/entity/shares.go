// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Shares is the golang structure for table shares.
type Shares struct {
	Id          uint64      `json:"id"          ` //
	Number      uint        `json:"number"      ` //
	Protocol    string      `json:"protocol"    ` //
	Uid         string      `json:"uid"         ` //
	Address     string      `json:"address"     ` //
	Port        uint        `json:"port"        ` //
	Security    string      `json:"security"    ` //
	Encryption  string      `json:"encryption"  ` //
	PublicKey   string      `json:"publicKey"   ` //
	HeaderType  string      `json:"headerType"  ` //
	Fingerprint string      `json:"fingerprint" ` //
	Network     string      `json:"network"     ` //
	Flow        string      `json:"flow"        ` //
	Sni         string      `json:"sni"         ` //
	ShortIds    string      `json:"shortIds"    ` //
	Remarks     string      `json:"remarks"     ` //
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` //
}
