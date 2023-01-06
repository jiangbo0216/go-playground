package fmt_test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	smile := fmt.Sprintf(`%q`, '😀')
	if smile != "'😀'" {
		t.Errorf(`smile = %q, want '😀'`, smile)
	}
}
