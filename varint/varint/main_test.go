package varint

import (
	"testing"
	"varint/core"
)

func TestMain(t *testing.T) {
	t.Log(core.EncodeUInt64(292))
	t.Log(core.EncodeUInt64(1292))
	t.Log(core.EncodeUInt64(12292))
}
