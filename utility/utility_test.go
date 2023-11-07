package utility

import "testing"

func TestNewToken(t *testing.T) {
	key := []byte("abc")
	sign, err := NewToken(key)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("sign: %s\n", sign)

	mc, err := ParseToken(key, sign)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	t.Logf("mc: %+v\n", mc)
}
