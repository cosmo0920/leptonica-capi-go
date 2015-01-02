package leptonica_test

import (
	lept "./leptonica"
	"testing"
)

func TestVersion(t *testing.T) {
	result := lept.Version()
	if result == "" {
		t.Errorf("result = %v cannot empty string", result)
	}
}
