package json

import (
	"testing"
)

func TestDeepCopyMap(t *testing.T) {
	src := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{"c": 2},
	}
	dst := DeepCopyMap(src)
	dst["d"] = 3
	dstB := dst["b"].(map[string]interface{})
	dstB["c"] = 4
	t.Log(src)
	t.Log(dst)
}
