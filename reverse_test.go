package reverse

import "testing"


func TestReverseString(t *testing.T) {
	str := "Test"
	if res := reverseString(str); res != "tseT" {
		t.Error("Expected tseT, got ", res)
	}
}