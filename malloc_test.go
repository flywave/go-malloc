package malloc

import "testing"

func TestMallocTrim(t *testing.T) {
	Free(Malloc(1 << 12))
	if !MallocTrim(0) {
		t.Errorf("MallocTrim() want = %v, got = %v", true, false)
	}
}
