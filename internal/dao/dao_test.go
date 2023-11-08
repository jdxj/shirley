package dao

import (
	"context"
	"fmt"
	"testing"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/jdxj/shirley/internal/model/do"
)

func TestSharesDao_GetShares(t *testing.T) {
	t.Parallel()

	res, err := Shares.GetShares(context.Background())
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	for _, v := range res {
		fmt.Printf("%+v\n", v)
	}
}

func TestSharesDao_InsertShare(t *testing.T) {
	t.Parallel()

	var (
		ctx  = gctx.New()
		data = do.Shares{
			Protocol:    "def",
			Uid:         "def",
			Address:     "def",
			Port:        123,
			Security:    "def",
			Encryption:  "def",
			PublicKey:   "def",
			HeaderType:  "def",
			Fingerprint: "def",
			Network:     "def",
			Flow:        "def",
			Sni:         "def",
			ShortIds:    "def",
			Remarks:     "def",
		}
	)
	err := Shares.InsertShare(ctx, &data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", data)
}

func TestSharesDao_DeleteShare(t *testing.T) {
	t.Parallel()

	err := Shares.DeleteShare(context.Background(), 1)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestSharesDao_UpdateShare(t *testing.T) {
	t.Parallel()

	var (
		id   uint64 = 2
		data        = &do.Shares{
			Protocol: "abc",
			Uid:      "def",
		}
	)
	err := Shares.UpdateShare(context.Background(), id, data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}
