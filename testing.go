package simply

import (
	"strings"
	"testing"
)

func AssertContains(t *testing.T, str, substr string) {
	if !strings.Contains(str, substr) {
		t.Errorf(`"%s" is not in "%s"`, substr, str)
	}
}
